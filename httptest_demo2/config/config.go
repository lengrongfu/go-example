package config

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lengrongfu/go-example/httptest_demo2/model"
)



func NewRedis(addr string) (*model.RedisCli,error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, e := redisCli.Ping().Result()
	if e != nil {
		panic(e)
	}
	cli := model.NewRedisCli(redisCli)
	return cli,nil
}

func NewDB(add string) (*model.DbCli,error) {
	db, e := sql.Open("mysql", add)
	if e != nil {
		panic(e)
	}
	e = db.Ping()
	if e != nil  {
		panic(e)
	}
	dbStore := model.NewDbCli(db)
	return dbStore,nil
}