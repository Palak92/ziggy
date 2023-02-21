package server

import (
	"context"
	"fmt"

	pb "github.com/palak92/ziggy/api"
	"github.com/palak92/ziggy/pkg/coinmarketcap"
)

type Crypto struct {
	pb.UnimplementedCryptoServer
}

func (c *Crypto) ListCoins(ctx context.Context, req *pb.ListCoinsRequest) (*pb.ListCoinsResponse, error) {
	coins := []*pb.Coin{}
	list, err := coinmarketcap.ListCoins()
	if err != nil {
		return nil, fmt.Errorf("error while getting list from coinmarketcap:%w", err)
	}
	for _, crypto := range list.CryptoMarket {
		coin := pb.Coin{
			Id:         string(crypto.ID),
			Name:       crypto.Name,
			Symbol:     crypto.Symbol,
			Tracked:    false,
			Price:      fmt.Sprint(crypto.Quote["USD"].Price),
			LastSynced: crypto.LastUpdated,
		}

		coins = append(coins, &coin)
	}
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
