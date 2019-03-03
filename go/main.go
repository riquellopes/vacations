package main

import (
	"context"
	"log"
	"net"

	"github.com/riquellopes/vacations/go/protons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server -
type Server struct {
}

// WhenDoYouStartYourVacation -
func (s *Server) WhenDoYouStartYourVacation(c context.Context, req *protons.HenriqueMessage) (*protons.HenriqueMessage, error) {
	log.Printf("Asking from gRPC: %s", req.GetMsg())
	return &protons.HenriqueMessage{Msg: "01/03/2019"}, nil
}

func server() {
	server, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}

	defer server.Close()

	grpcServer := grpc.NewServer()
	protons.RegisterHenriqueVacationsServer(grpcServer, &Server{})

	reflection.Register(grpcServer)
	grpcServer.Serve(server)
}

func main() {
	go server()
}
