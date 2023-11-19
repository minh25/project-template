package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"removebyyourname/config"
	"removebyyourname/internal/pkg/interceptor"
	"removebyyourname/internal/server"
	"removebyyourname/pb"
	"sync"
	"syscall"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startServer(conf config.Config, tracer opentracing.Tracer, zapLogger *zap.Logger, db *sql.DB) {
	grpcServer := grpc.NewServer(
		interceptor.ChainUnaryInterceptorIgnoreHealthCheck(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(interceptor.RecoveryHandlerFunc)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_ctxtags.UnaryServerInterceptor(),
			interceptor.UnaryServerJaegerImplementation(tracer),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_validator.StreamServerInterceptor(),
		),
	)

	pb.RegisterServiceServer(grpcServer, server.New(db))
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	startHTTPAndGRPCServers(conf, grpcServer)
}

func startHTTPAndGRPCServers(conf config.Config, grpcServer *grpc.Server) {
	fmt.Println("GRPC:", conf.Server.GRPC.ListenString())
	fmt.Println("HTTP:", conf.Server.HTTP.ListenString())

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
	)

	ctx := context.Background()
	grpcHost := conf.Server.GRPC.String()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(constant.MaxResponseLength)),
	}

	registerGRPCGateway(ctx, mux, grpcHost, opts)

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())
	httpMux.Handle("/", mux)

	httpServer := &http.Server{
		Addr:    conf.Server.HTTP.ListenString(),
		Handler: httpMux,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Println("Shutdown HTTP server successfully")
	}()

	go func() {
		defer wg.Done()

		listener, err := net.Listen("tcp", conf.Server.GRPC.ListenString())
		if err != nil {
			panic(err)
		}

		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
		fmt.Println("Shutdown gRPC server successfully")
	}()

	//--------------------------------
	// Graceful Shutdown
	//--------------------------------
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func registerGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	_ = pb.RegisterServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
