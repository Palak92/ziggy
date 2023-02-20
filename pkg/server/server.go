package server

import (
	"context"

	pb "github.com/bazelbuild/rules_go/bazel-bin/api/crypto_go_proto_/github.com/palak92/ziggy/api"
)

type Crypto struct {
	pb.UnimplementedCryptoServer
}

func (c *Crypto) GetCoinInformation(ctx context.Context, req *pb.GetCoinInformationRequest) (*pb.GetCoinInformationResponse, error) {
	coins := []*pb.Coin{}
	res := &pb.GetCoinInformationResponse{
		Coins: coins,
	}
	return res, nil
}
