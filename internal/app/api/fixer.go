package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/panda-next-team/poolrank-agent/internal/app/model"
	pb "github.com/panda-next-team/poolrank-proto/agent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FixerService struct {
	Engine *xorm.Engine
}

func (s *FixerService) GetRate(ctx context.Context, in *pb.GetRateRequest) (*pb.GetRateResponse, error) {
	if in.Base == "" {
		st := status.New(codes.InvalidArgument, "Invalid argument base")
		return nil, st.Err()
	}

	if in.Currency == "" {
		st := status.New(codes.InvalidArgument, "Invalid argument currency")
		return nil, st.Err()
	}

	rate := new(model.Rate)
	has, err := s.Engine.Where("base = ?", in.Base).And("currency = ?", in.Currency).Get(rate)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if !has {
		st := status.New(codes.NotFound, "Not found entity.")
		return nil, st.Err()
	}

	return &pb.GetRateResponse{Rate: loadRate(rate)}, nil
}

func (s *FixerService) GetRates(ctx context.Context, in *pb.GetRatesRequest) (*pb.GetRatesResponse, error) {
	if in.Base == "" {
		st := status.New(codes.InvalidArgument, "Invalid argument base")
		return nil, st.Err()
	}

	rates := make([]*model.Rate, 0)
	err := s.Engine.Where("base = ?", in.Base).Find(&rates)
	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	pbRates := make([]*pb.Rate, len(rates))
	for i, rate := range rates {
		pbRates[i] = loadRate(rate)
	}
	return &pb.GetRatesResponse{Rates: pbRates}, nil
}

func loadRate(model *model.Rate) *pb.Rate {
	entity := new(pb.Rate)
	entity.Id = model.Id
	entity.Base = model.Base
	entity.Currency = model.Currency
	entity.CurrencyRate = model.CurrencyRate
	entity.CratedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}
