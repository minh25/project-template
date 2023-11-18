package server

import "removebyyourname/pb"

type server struct {
	pb.UnimplementedServiceServer
}

func New() *server {
	return &server{}
}
