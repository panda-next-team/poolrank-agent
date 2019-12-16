package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/panda-next-team/poolrank-agent/internal/app/model"
	pb "github.com/panda-next-team/poolrank-proto/agent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoinMarketCapService struct {
	Engine *xorm.Engine
}

func (s *CoinMarketCapService) GetQuote(ctx context.Context, in *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	if in.EnTag == "" {
		st := status.New(codes.InvalidArgument, "Invalid argument en_tag")
		return nil, st.Err()
	}

	quote := new(model.Quote)
	has, err := s.Engine.Where("en_tag = ?", in.EnTag).Get(quote)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if !has {
		st := status.New(codes.NotFound, "Not found entity.")
		return nil, st.Err()
	}

	return &pb.GetQuoteResponse{Quote: loadQuote(quote)}, nil
}

func loadQuote(quote *model.Quote) *pb.Quote {
	entity := new(pb.Quote)
	entity.Id = quote.Id
	entity.EnTag = quote.EnTag
	entity.CirculatingSupply = quote.CirculatingSupply
	entity.MarketCap = quote.MarketCap
	entity.NumMarketPairs = quote.NumMarketPairs
	entity.CmcRank = quote.CmcRank
	entity.Volume_24H = quote.Volume24H
	entity.PercentChange_7D = quote.PercentChange7D
	entity.PercentChange_1H = quote.PercentChange1H
	entity.PercentChange_24H = quote.PercentChange24H
	entity.MarketCap = quote.MarketCap
	entity.CreatedAtTs = quote.CreatedAtTs
	entity.CreatedAt = quote.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = quote.UpdatedAtTs
	entity.UpdatedAt = quote.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}
