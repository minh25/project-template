package server

import (
	"context"
	"removebyyourname/pb"
)

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Code:    200,
		Message: "OK",
		Data: &pb.GetUserResponse_User{
			Id:    1,
			Name:  "removebyyourname",
			Email: "nothing",
		},
	}, nil
}
