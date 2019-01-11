package context


import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/buaazp/fasthttprouter"
)


type key int

const (
	mysqlConnKey key = 0
	clickHouseConnKey key = 1
	redisClientsKey key = 2
	routerKey key = 3
)

//
func NewWrapperContext(parent context.Context,
	mysql *sql.DB,
	clickHouse *sql.DB,
	redisClients []*redis.Client,
	router *fasthttprouter.Router) context.Context {

	return &wrapperContext{parent,
		mysql,
		clickHouse,
		redisClients,
		router}
}

//
func NewWrapperContextDef(
	mysql *sql.DB,
	clickHouse *sql.DB,
	redisClients []*redis.Client,
	router *fasthttprouter.Router) context.Context {

	return &wrapperContext{
		context.Background(),
		mysql,
		clickHouse,
		redisClients,
		router}
}

//
type wrapperContext struct {
	context.Context
	mysqlConn *sql.DB
	clickHouseConn *sql.DB
	redisClients []*redis.Client
	router *fasthttprouter.Router
}


//
func (w*wrapperContext) Value(key interface{}) interface{} {
	switch key {
	case mysqlConnKey: return w.mysqlConn
	case clickHouseConnKey: return w.clickHouseConn
	case redisClientsKey: return w.redisClients
	case routerKey: return w.router
	}
	return w.Context.Value(key)
}

//
func MysqlConn(ctx context.Context) (*sql.DB, bool) {
	conn, ok := ctx.Value(mysqlConnKey).(*sql.DB)
	return conn, ok
}

//
func ClickHouseConn(ctx context.Context) (*sql.DB, bool) {
	conn, ok := ctx.Value(clickHouseConnKey).(*sql.DB)
	return conn, ok
}

//
func FastHttpRouter(ctx context.Context) (*fasthttprouter.Router, bool) {
	router, ok := ctx.Value(routerKey).(*fasthttprouter.Router)
	return router, ok
}

//
func RedisClients(ctx context.Context) ([]*redis.Client, bool) {
	clients, ok := ctx.Value(redisClientsKey).([]*redis.Client)
	return clients, ok
}