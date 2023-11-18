package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalid = status.Error(codes.InvalidArgument, "invalid")
