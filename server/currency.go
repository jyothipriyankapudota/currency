package server

import (
	"context"

	protos "github.com/jyothipriyankapudota/currency/protos/currency"

	hclog "github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l, protos.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "distination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
