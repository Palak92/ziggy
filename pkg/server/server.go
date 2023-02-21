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
			Id:         fmt.Sprint(crypto.ID),
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
	qs, err := coinmarketcap.GetHistoricalQuotes("21971", 10, "asc")
	if err != nil {
		return nil, fmt.Errorf("error while getting list from coinmarketcap:%w", err)
	}

	dps := []*pb.DataPoint{}

	for _, q := range qs.HistoricalQuote["21971"].Quotes {
		dp := pb.DataPoint{
			Name:     qs.HistoricalQuote["21971"].Name,
			Id:       fmt.Sprint(qs.HistoricalQuote["21971"].ID),
			Price:    fmt.Sprint(q.Quote["USD"].Price),
			SyncedOn: q.Timestamp,
		}

		dps = append(dps, &dp)
	}
	return &pb.GetCoinHistoryResponse{
		DataPoints: dps,
	}, nil
}

func (c *Crypto) TrackCoin(ctx context.Context, req *pb.TrackCoinRequest) (*pb.TrackCoinResponse, error) {
	res := &pb.TrackCoinResponse{}
	return res, nil
}
