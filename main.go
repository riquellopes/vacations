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

func client() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error to connect %v", err)
	}

	defer dial.Close()

	client := protons.NewHenriqueVacationsClient(dial)

	asking := &protons.HenriqueMessage{Msg: "When do you start your vacation?"}
	r, err := client.WhenDoYouStartYourVacation(context.Background(), asking)

	if err != nil {
		log.Fatalf("Error to replay %v", err)
	}

	log.Printf("gRPC say: %s", r.GetMsg())
}

func main() {
	go server()

	client()
}
