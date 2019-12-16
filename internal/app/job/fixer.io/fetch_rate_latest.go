package fixer_io

import (
	"github.com/LordotU/go-fixerio"
	"github.com/go-xorm/xorm"
	"github.com/panda-next-team/poolrank-agent/internal/app/model"
	"github.com/panda-next-team/poolrank-agent/internal/pkg"
	"time"
)

type FetchRateLatestConfig struct {
	Mysql   pkg.MysqlConfig
	Fixer   FixerIOConfig
	Symbols []string
}

type FetchRateLatestJob struct {
	Engine *xorm.Engine
	Fixer  *gofixerio.FixerIO
}

func (job *FetchRateLatestJob) Run(symbols []string) error {
	latestRates, err := job.Fixer.GetLatest(symbols)
	if err != nil {
		return err
	}

	var forErr error

	for currency, currencyRate := range latestRates.Rates {
		entity := new(model.Rate)
		entity.Base = latestRates.Base
		entity.Currency = currency

		has, err := job.Engine.Exist(entity)
		if err != nil {
			forErr = err
			break
		}

		if has {
			_, forErr = job.Engine.Update(&model.Rate{CurrencyRate: currencyRate, UpdatedAtTs: time.Now().Unix()}, entity)
		} else {
			entity.CurrencyRate = currencyRate
			entity.CreatedAtTs = time.Now().Unix()
			entity.UpdatedAtTs = time.Now().Unix()
			_, forErr = job.Engine.Insert(entity)
		}

		if forErr != nil {
			break
		}
	}

	return forErr
}
