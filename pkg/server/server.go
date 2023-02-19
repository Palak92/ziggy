package server

// import (
// 	"context"

// 	pb "github.com/palak92/ziggy/api"
// )

// type Server struct{}

// func (s *Server) GetPrice(ctx context.Context, req *pb.GetPriceRequest) (*pb.GetPriceResponse, error) {
// 	// Connect to a cryptocurrency price source and retrieve the price for the requested symbol
// 	// Replace <YOUR_API_KEY> with your API key
// 	price, err := getPrice(req.Symbol, "<YOUR_API_KEY>")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create a response message and return it
// 	res := &pb.GetPriceResponse{
// 		Price: price,
// 	}
// 	return res, nil
// }
