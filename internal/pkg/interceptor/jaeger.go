package interceptor

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.tekoapis.com/tekbone/app/common/golang/logger"
	"go.tekoapis.com/tekbone/app/common/golang/utils"
	"google.golang.org/grpc"
)

func UnaryServerJaegerImplementation(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		unsSpan := tracer.StartSpan(info.FullMethod)
		defer unsSpan.Finish()

		spanCtx, ok := unsSpan.Context().(jaeger.SpanContext)
		if ok {
			logger.Infof("TraceId: %s\n", spanCtx.TraceID().String())
			ctx = context.WithValue(ctx, utils.TraceIdCtxKey, spanCtx.TraceID().String())
		}

		return handler(ctx, req)
	}
}
