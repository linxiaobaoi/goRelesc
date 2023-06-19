package main

import (
	"github.com/linxiaobaoi/goRelesc.git/config"
	"github.com/linxiaobaoi/goRelesc.git/router"
)

func main() {
	config.InitMySQL()
	config.InitRedis()
	router.AppRouter()

}
