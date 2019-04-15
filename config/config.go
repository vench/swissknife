package config

import (
	"github.com/jinzhu/configor"
	"os"
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	scontext "github.com/vench/swissknife/context"
	"context"
)

const (
	chDefaultMaxOpenConnections = 256
	envConfigFilePath = `APP_CONFIG_FILE_PATH`
)

var (
	configPath = "./cmd/config.yml"
	isLoad     = false
)

//
type database struct {
	Ip       	 string `default:"localhost"`
	Database 	 string `default:"default"`
	Username 	 string `default:""`
	Password 	 string `default:""`
	Port     	 uint   `default:"9000"`
	Debug    	 uint8  `default:"0"`
	MaxOpenConn  int    `default:"0"`
}

//
var config =  struct {
	ClickHouse 	*database
	Mysql		*database
	Web struct {
		Ip   string `default:"0.0.0.0"`
		Port uint   `default:"8087"`
	}
	Redis struct {
		List []string `default:"[localhost:6379]"`
	}
}{}

//
func ConfigLoad() {
	if !isLoad {
		isLoad = true
		ConfigLoadStruct(&config)
	}
}

//
func ConfigLoadStruct(confStr interface{}) {
	if path, ok := os.LookupEnv(envConfigFilePath); ok {
		configPath = path
	}

	err := configor.Load(confStr, configPath)
	if err != nil {
		log.Fatal("ConfigLoad: ", err)
		os.Exit(1)
	}
}

//
func ConfigWebAddr() string {
	ConfigLoad()
	return fmt.Sprintf("%s:%d", config.Web.Ip, config.Web.Port)
}

//
func CreateSwissKnifeContext() context.Context {
	ConfigLoad()

	router := fasthttprouter.New()
	ctx := scontext.NewWrapperContextDef(
		NewMysqlConnect(config.Mysql),
		NewClickHouseConnect(config.ClickHouse),
		GetListClientRedis(),
		router	)

	return ctx
}

//
func ConfigPrint() {
	out, err := json.Marshal(config)
	if err != nil {
		panic (err)
	}
	fmt.Println(string(out))
}