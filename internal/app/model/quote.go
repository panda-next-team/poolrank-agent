package model

import "time"

type Quote struct {
	Id                int32     `xorm:"pk autoincr"`
	EnTag             string    `xorm:"en_tag"`
	CirculatingSupply float64   `xorm:"circulating_supply"`
	NumMarketPairs    int32     `xorm:"num_market_pairs"`
	CmcRank           int32     `xorm:"cmc_rank"`
	EurPrice          float64   `xorm:"eur_price"`
	Volume24H         float64   `xorm:"volume_24h"`
	PercentChange1H   float64   `xorm:"percent_change_1h"`
	PercentChange24H  float64   `xorm:"percent_change_24h"`
	PercentChange7D   float64   `xorm:"percent_change_7d"`
	MarketCap         float64   `xorm:"market_cap"`
	CreatedAtTs       int64     `xorm:"created_at_ts"`
	CreatedAt         time.Time `xorm:"created"`
	UpdatedAtTs       int64     `xorm:"updated_at_ts"`
	UpdatedAt         time.Time `xorm:"updated"`
}
