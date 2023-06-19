package main

import (
	"github.com/linxiaobaoi/goRelesc/config"
	"github.com/linxiaobaoi/goRelesc/router"
)

func main() {
	config.InitMySQL()
	config.InitRedis()
	router.AppRouter()

}
