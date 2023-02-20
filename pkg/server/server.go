package server

import (
	"context"

	pb "github.com/palak92/ziggy/bazel-bin/api/crypto_go_proto_/github.com/palak92/ziggy/api"
)

type Crypto struct {
	pb.UnimplementedCryptoServer
}

func (c *Crypto) ListCoins(ctx context.Context, req *pb.ListCoinsRequest) (*pb.ListCoinsResponse, error) {
	coins := []*pb.Coin{}
	res := &pb.ListCoinsResponse{
		Coins: coins,
	}
	return res, nil
}

func (c *Crypto) GetCoinHistory(ctx context.Context, req *pb.GetCoinHistoryRequest) (*pb.GetCoinHistoryResponse, error) {
	res := &pb.GetCoinHistoryResponse{}
	return res, nil
}

func (c *Crypto) TrackCoin(ctx context.Context, req *pb.TrackCoinRequest) (*pb.TrackCoinResponse, error) {
	res := &pb.TrackCoinResponse{}
	return res, nil
}
