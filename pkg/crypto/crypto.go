package crypto

import (
	"context"
	"fmt"

	pb "github.com/palak92/ziggy/api"
	"github.com/palak92/ziggy/pkg/coinmarketcap"
	"github.com/palak92/ziggy/pkg/db"
)

type Server struct {
	pb.UnimplementedCryptoServer
	DB *db.CoinsDB
}

func (c *Server) ListCoins(ctx context.Context, req *pb.ListCoinsRequest) (*pb.ListCoinsResponse, error) {
	coins := []*pb.Coin{}
	coinsData, err := c.DB.AllCoins()
	if err != nil {
		return nil, fmt.Errorf("error while getting coins from database:%w", err)
	}
	for _, c := range coinsData {
		coin := pb.Coin{
			Id:         fmt.Sprint(c.UnvID),
			Name:       c.Name,
			Symbol:     c.Symbol,
			Tracked:    c.Tracked,
			Price:      c.Price,
			LastSynced: c.LastSynced,
		}

		coins = append(coins, &coin)
	}
	res := &pb.ListCoinsResponse{
		Coins: coins,
	}
	return res, nil
}

func (c *Server) GetCoinHistory(ctx context.Context, req *pb.GetCoinHistoryRequest) (*pb.GetCoinHistoryResponse, error) {
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

func (c *Server) TrackCoin(ctx context.Context, req *pb.TrackCoinRequest) (*pb.TrackCoinResponse, error) {
	res := &pb.TrackCoinResponse{}
	return res, nil
}
