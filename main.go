package main

import (
	"Spider/bootstrap"
	"Spider/cmd"
	btsConfig "Spider/config"
	"Spider/pkg/config"
	"flag"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 初始化配置
	initEnv()

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化 DB
	bootstrap.SetupDB()

	err := cmd.Execute()
	if err != nil {
		return
	}
}

func initEnv() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=example 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
}
