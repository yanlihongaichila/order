package main

import (
	"github.com/yanlihongaichila/framework/app"
	"github.com/yanlihongaichila/framework/nacos"
	"log"
	"rpc/model"
	"strconv"
)

func main() {
	//获取配置
	viper, err := nacos.InitViper("rpc", "/config", "nacos")
	if err != nil {
		return
	}
	atoi, err := strconv.Atoi(viper["port"])
	if err != nil {
		log.Println(err)
		return
	}

	err = app.Init(viper["address"], viper["group"], viper["name"], uint64(atoi), "mysql", "redis")
	if err != nil {
		log.Println(err)
		return
	}

	//创建表
	err = model.Migrate()
	if err != nil {
		log.Println(err)
		return
	}

}
