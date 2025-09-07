package models

import "pledge-backend/db"

func InitTable() {
	// 多签
	db.Mysql.AutoMigrate(&MultiSign{})
	// 代币信息
	db.Mysql.AutoMigrate(&TokenInfo{})
	db.Mysql.AutoMigrate(&TokenList{})
	// pool data
	db.Mysql.AutoMigrate(&PoolData{})
	// pool base
	db.Mysql.AutoMigrate(&PoolBases{})
}
