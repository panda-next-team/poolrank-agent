package model

import "time"

type Quote struct {
	Id                int32 `xorm:"pk autoincr"`
	EnTag             string
	CirculatingSupply float64
	NumMarketPairs    int32
	CmcRank           int32
	EurPrice          float64
	Volume24H         float64 `xorm:"'volume_24h'"`
	PercentChange1H   float64 `xorm:"'percent_change_1h'"`
	PercentChange24H  float64 `xorm:"'percent_change_24h'"`
	PercentChange7D   float64 `xorm:"'percent_change_7d'"`
	MarketCap         float64
	CreatedAtTs       int64
	CreatedAt         time.Time `xorm:"created"`
	UpdatedAtTs       int64
	UpdatedAt         time.Time `xorm:"updated"`
}
