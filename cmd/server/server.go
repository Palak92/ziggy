package main

import (
	"log"
	"net"

	pb "github.com/palak92/ziggy/api"
	"github.com/palak92/ziggy/pkg/crypto"
	"github.com/palak92/ziggy/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbUser     = "root"
	dbPassword = "password"
	dbAddress  = "127.0.0.1:3306"
)

func main() {
	// start database
	db, err := db.New(dbUser, dbPassword, dbAddress)
	if err != nil {
		log.Fatalf("failed to connect to database : %v", err)
	}

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterCryptoServer(s, &crypto.Server{DB: db})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
