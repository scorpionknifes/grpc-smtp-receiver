package example-receiver

func main(){
	// Start gRPC server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

type Server struct{}

func (s *Server) SendEmail(ctx context.Context, email *Email) (*Status, error) {
	log.Printf("Received Email from smtp server: %s", email.data)
	return &Status{Status:true}
}