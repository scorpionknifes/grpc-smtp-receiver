package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Start gRPC server

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen on port 3001: %v", err)
	}

	log.Println("gRPC Receiver for SMTP server listening on port 3001")
	grpcServer := grpc.NewServer()

	s := Server{}
	RegisterSMTPServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 3001: %v", err)
	}

}

// Server gRPC server
type Server struct{}

// SendEmail send emails
func (s *Server) SendEmail(ctx context.Context, email *Email) (*Status, error) {
	log.Printf("Received Email from smtp server: %s", email.Data)
	return &Status{Status: true}, nil
}
