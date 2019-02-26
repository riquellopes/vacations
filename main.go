package main

import (
	"context"
	"log"
	"net"

	"github.com/riquellopes/vacations/protons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server -
type Server struct {
}

// WhenDoYouStartYourVacation -
func (s *Server) WhenDoYouStartYourVacation(c context.Context, req *protons.HenriqueMessage) (*protons.HenriqueMessage, error) {
	return &protons.HenriqueMessage{Msg: "01/03/2019"}, nil
}

func main() {
	server, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Printf("Error to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	protons.RegisterHenriqueVacationsServer(grpcServer, &Server{})

	reflection.Register(grpcServer)
	grpcServer.Serve(server)
}
