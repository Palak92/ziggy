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

func (s *Server) ListCoins(ctx context.Context, req *pb.ListCoinsRequest) (*pb.ListCoinsResponse, error) {
	cns := req.CoinNames

	var cd []*db.Coin
	if len(cns) == 0 {
		coinsData, err := s.allCoinsFromDB()
		if err != nil {
			return nil, fmt.Errorf("error while getting all coins from db: %w", err)
		}
		cd = coinsData
	} else {
		coinsData, err := s.coinsByNames(cns)
		if err != nil {
			return nil, fmt.Errorf("error while getting coins from db of names%v: %w", cns, err)
		}
		cd = coinsData
	}

	coins := []*pb.Coin{}
	for _, c := range cd {
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

func (c *Server) allCoinsFromDB() ([]*db.Coin, error) {
	coinsData, err := c.DB.AllCoins()
	if err != nil {
		return nil, fmt.Errorf("error while getting coins from database:%w", err)
	}
	return coinsData, nil
}

func (s *Server) coinsByNames(names []string) ([]*db.Coin, error) {
	var coins []*db.Coin
	for _, n := range names {
		c, err := s.DB.CoinsByName(n)
		if err != nil {
			return nil, fmt.Errorf("error while getting coin of name %q from database:%w", n, err)
		}
		coins = append(coins, c)
	}

	return coins, nil
}
func (s *Server) GetCoinHistory(ctx context.Context, req *pb.GetCoinHistoryRequest) (*pb.GetCoinHistoryResponse, error) {
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
