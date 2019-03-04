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
func (s *Server) WhenDoYouStartYourVacation(c context.Context, req *protons.VacationRequest) (*protons.VacationResponse, error) {
	log.Printf("Asking from gRPC: %s", req.GetName())
	reply := "Will be 03/07/2019."
	return &protons.VacationResponse{Reply: &reply}, nil
}

func server() {
	server, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}

	defer server.Close()

	grpcServer := grpc.NewServer()
	protons.RegisterVacationServiceServer(grpcServer, &Server{})

	reflection.Register(grpcServer)
	grpcServer.Serve(server)
}

func main() {
	server()
}
