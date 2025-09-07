package tasks

import (
	"pledge-backend/db"
	"pledge-backend/schedule/common"
	"pledge-backend/schedule/services"
	"time"

	"github.com/jasonlvhit/gocron"
)

func Task() {

	// get environment variables
	common.GetEnv()

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//init task
	services.NewPool().UpdateAllPoolInfo()
	services.NewTokenPrice().UpdateContractPrice()
	services.NewTokenSymbol().UpdateContractSymbol()
	services.NewTokenLogo().UpdateTokenLogo()
	services.NewBalanceMonitor().Monitor()
	// services.NewTokenPrice().SavePlgrPrice()
	services.NewTokenPrice().SavePlgrPriceTestNet()

	//run pool task
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)

	// 更新全部 poolbases/pooldata/token_info (存在就更新，不存在就插入Token，ChainId，CreatedAt，UpdatedAt)
	_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(services.NewPool().UpdateAllPoolInfo)
	// 访问BscPledgeOracle(间接访问 chainlink) ：  token_info.price
	_ = s.Every(1).Minute().From(gocron.NextTick()).Do(services.NewTokenPrice().UpdateContractPrice)
	// 更新名称 token_info.symbol （根据合约地址）
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenSymbol().UpdateContractSymbol)
	// 更新图标 Logo（根据合约地址）
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenLogo().UpdateTokenLogo)
	//
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewBalanceMonitor().Monitor)
	//
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewTokenPrice().SavePlgrPriceTestNet)
	<-s.Start() // Start all the pending jobs

}
