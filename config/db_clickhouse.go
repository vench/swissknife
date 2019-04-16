package config

import (
	"database/sql"
	"fmt"
	_ "github.com/kshvakov/clickhouse"
	"log"
	"os"
)

var (
	clickHouseConnect          *sql.DB
	clickHouseConnectOpen      = false
)

//
func ClickHouseOneConnect() *sql.DB {
	if !clickHouseConnectOpen {
		conn := NewClickHouseConnect(config.ClickHouse)
		clickHouseConnectOpen = true
		clickHouseConnect = conn
	}
	return clickHouseConnect
}

//
func ClickHouseOneClose() {
	if clickHouseConnectOpen {
		clickHouseConnect.Close()
	}
}

//
func NewClickHouseConnect(conf *Database) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"tcp://%s:%d?debug=%d&database=%s&read_timeout=30&write_timeout=30&password=%s&username=%s",
		conf.Ip,
		conf.Port,
		conf.Debug,
		conf.Database,
		conf.Password,
		conf.Username,
	)

	conn, err := sql.Open("clickhouse", dataSourceName)
	if err != nil {
		log.Fatalf(`%s`, err.Error())
		os.Exit(1)
	}

	if conf.MaxOpenConn > 0 {
		conn.SetMaxOpenConns(conf.MaxOpenConn)
	} else {
		conn.SetMaxOpenConns(chDefaultMaxOpenConnections)
	}

	return conn
}