package interceptor

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

// RecoveryHandlerFunc ...
func RecoveryHandlerFunc(p interface{}) error {
	fmt.Println("stacktrace from panic:\n" + string(debug.Stack()))
	return status.Errorf(codes.Internal, "panic: %s", p)
}
