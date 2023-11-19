package server

import (
	"database/sql"
	"removebyyourname/pb"
)

type server struct {
	pb.UnimplementedServiceServer
}

func New(_ *sql.DB) *server {
	return &server{}
}
