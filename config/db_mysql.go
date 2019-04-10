package config


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlConnect          *sql.DB
	mysqlConnectOpen      = false
)

//
func MysqlOneConnect() *sql.DB {
	if !mysqlConnectOpen {
		conn := NewMysqlConnect(config.Mysql)
		mysqlConnectOpen = true
		mysqlConnect = conn
	}
	return mysqlConnect
}

//
func MysqlOneClose() {
	if mysqlConnectOpen {
		mysqlConnect.Close()
	}
}


//
func NewMysqlConnect(conf *database) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.Username,
		conf.Password,
		conf.Ip,
		conf.Port,
		conf.Database)

	d, e := sql.Open("mysql", connStr)
	if e != nil {
		return nil
	}

	if conf.MaxOpenConn > 0 {
		d.SetMaxOpenConns(conf.MaxOpenConn)
	}

	stmt, e := d.Prepare("SET NAMES utf8")
	if e != nil {
		return nil
	}
	defer stmt.Close()

	_, e = stmt.Exec()
	if e != nil {
		return nil
	}
	return d

}