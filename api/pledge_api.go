package main

import (
	"fmt"
	"pledge-backend/api/middlewares"
	"pledge-backend/api/models"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/api/models/ws"
	"pledge-backend/api/routes"
	"pledge-backend/api/static"
	"pledge-backend/api/validate"
	"pledge-backend/config"
	"pledge-backend/db"

	"github.com/gin-gonic/gin"
)

func main() {

	//init mysql
	db.InitMysql()

	//init redis
	db.InitRedis()
	models.InitTable()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// websocket server
	go ws.StartServer()

	// get plgr price from kucoin-exchange
	// 实时获取 kucoin 交易所 代币的的价格 （plgr 是改项目的代币，目前已经没有了）
	// 并缓存到redis
	go kucoin.GetExchangePrice()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	staticPath := static.GetCurrentAbPathByCaller()
	fmt.Println("staticPath", staticPath)
	app.Static("/storage/", staticPath)
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":" + config.Config.Env.Port)

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
