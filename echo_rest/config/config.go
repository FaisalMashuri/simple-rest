package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_Type     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}

func GetDBType() string {
	conf := GetConfig()
	return conf.DB_Type
}

func GetMYSQLConnectionString() string {
	conf := GetConfig()

	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME)

	return dataBase
}
