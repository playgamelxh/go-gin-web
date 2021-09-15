package main

import (
	"flag"
	"fmt"
	"go-gin-web/conf"
	"go-gin-web/router"
	api_router "go-gin-web/app/api/router"
	default_router "go-gin-web/app/default/router"
)

// 环境变量
var env string

// 配置文件
var webConfig *conf.YamlConfig

func main()  {

	// 解析环境变量
	flag.StringVar(&env, "env", "dev", "环境")
	flag.Parse()

	// 获取环境配置
	webConfig = conf.GetConfig(env)

	//加载路由
	router.Include(api_router.Routers, default_router.Routers)

	// 1.创建路由
	r := router.Init()

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(fmt.Sprintf("%s:%d", webConfig.Host, webConfig.Port))
}
