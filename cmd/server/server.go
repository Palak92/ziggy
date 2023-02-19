package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bazelbuild/rules_go/bazel-bin/api/crypto_go_proto_/github.com/palak92/ziggy/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) GetPrice(ctx context.Context, req *pb.GetPriceRequest) (*pb.GetPriceResponse, error) {

	res := &pb.GetPriceResponse{
		Price: 1,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCryptoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
