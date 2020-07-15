package main

import (
	"log"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var receivers map[string]*grpc.ClientConn
var conf Config

func main() {

	//Read config
	viper.SetConfigFile("./config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into config file, %v", err)
	}

	receivers = make(map[string]*grpc.ClientConn)
	for _, server := range conf.Server {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(server.GRPCServer, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %s", err)
		}
		defer conn.Close()
		receivers[server.EmailDomain] = conn
	}
	log.Println(receivers)

	// Start SMTP server
	s := smtp.NewServer(&SMTPHandlers{})

	s.Addr = ":" + conf.SMTPPort
	s.ReadTimeout = conf.Timeout * time.Second
	s.WriteTimeout = conf.Timeout * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Printf("Starting server at: %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
