package model

import "time"

type Rate struct {
	Id           int32 `xorm:"pk autoincr"`
	Base         string
	Currency     string
	CurrencyRate float64
	CreatedAtTs  int64
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAtTs  int64
	UpdatedAt    time.Time `xorm:"updated"`
}
