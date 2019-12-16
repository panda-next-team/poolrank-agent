package coinmarketcap_com

import (
	"fmt"
	"github.com/go-xorm/xorm"
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
	"github.com/panda-next-team/poolrank-agent/internal/app/model"
	"github.com/panda-next-team/poolrank-agent/internal/pkg"
	"time"
)

const (
	DefaultCurrency = "EUR"
)


type FetchQuoteLatestConfig struct {
	Mysql      pkg.MysqlConfig
	CoinMarket CoinMarketCapConfig
	Symbols    []string
}

type FetchQuoteLatestJob struct {
	Engine    *xorm.Engine
	CmcClient *cmc.Client
}

func (job *FetchQuoteLatestJob) Run(symbols []string) error {
	fmt.Println(symbols)
	var forErr error
	for _, symbol := range symbols {
		quotes, err := job.CmcClient.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
			Symbol:  symbol,
			Convert: DefaultCurrency,
		})
		if err != nil {
			forErr = err
			break
		}


		for _, quote := range quotes {
			entity := new(model.Quote)
			entity.EnTag = symbol
			has, err := job.Engine.Exist(entity)
			if err != nil {
				forErr = err
				goto exit
			}

			newEntity := new(model.Quote)
			newEntity.CirculatingSupply = quote.CirculatingSupply
			newEntity.CmcRank = int32(quote.CMCRank)
			newEntity.NumMarketPairs = int32(quote.NumMarketPairs)
			newEntity.EurPrice = quote.Quote[DefaultCurrency].Price
			newEntity.Volume24H = quote.Quote[DefaultCurrency].Volume24H
			newEntity.PercentChange1H = quote.Quote[DefaultCurrency].PercentChange1H
			newEntity.PercentChange24H = quote.Quote[DefaultCurrency].PercentChange24H
			newEntity.PercentChange7D = quote.Quote[DefaultCurrency].PercentChange7D
			newEntity.MarketCap = quote.Quote[DefaultCurrency].MarketCap

			if has {
				newEntity.UpdatedAtTs = time.Now().Unix()
				_, forErr = job.Engine.Update(newEntity, entity)
			} else {
				newEntity.EnTag = entity.EnTag
				newEntity.CreatedAtTs = time.Now().Unix()
				newEntity.UpdatedAtTs = time.Now().Unix()
				_, forErr = job.Engine.Insert(newEntity)
			}

			if forErr != nil {
				goto exit
			}

		}
	}
exit:
	return forErr
}
