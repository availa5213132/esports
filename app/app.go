package app

import (
	model "esports/server/app/models"
	"esports/server/app/router"
)

// Start 启动器方法
func Start() {
	model.NewMysql()
	model.NewRdb()
	model.NewMongoDB()
	router.New()
}
