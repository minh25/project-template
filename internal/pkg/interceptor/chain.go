package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

// ChainUnaryInterceptors ...
func ChainUnaryInterceptors(interceptors []grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Copy from grpc source

		// the struct ensures the variables are allocated together, rather than separately, since we
		// know they should be garbage collected together. This saves 1 allocation and decreases
		// time/call by about 10% on the microbenchmark.
		var state struct {
			i    int
			next grpc.UnaryHandler
		}
		state.next = func(ctx context.Context, req interface{}) (interface{}, error) {
			if state.i == len(interceptors)-1 {
				return interceptors[state.i](ctx, req, info, handler)
			}
			state.i++
			return interceptors[state.i-1](ctx, req, info, state.next)
		}
		return state.next(ctx, req)
	}
}

// IgnoreHealthCheckInterceptor ...
func IgnoreHealthCheckInterceptor(
	interceptor grpc.UnaryServerInterceptor,
) grpc.UnaryServerInterceptor {
	const healthLive = "/health.v1.HealthCheckService/Liveness"
	const healthReady = "/health.v1.HealthCheckService/Readiness"

	return func(
		ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if info.FullMethod == healthLive || info.FullMethod == healthReady {
			return handler(ctx, req)
		}
		return interceptor(ctx, req, info, handler)
	}
}

// ChainUnaryInterceptorIgnoreHealthCheck ...
func ChainUnaryInterceptorIgnoreHealthCheck(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		IgnoreHealthCheckInterceptor(
			ChainUnaryInterceptors(interceptors),
		),
	)
}
