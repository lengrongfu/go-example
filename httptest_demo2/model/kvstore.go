package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisCli struct {
	cli *redis.Client
}

func NewRedisCli(cli *redis.Client) *RedisCli  {
	return &RedisCli{
		cli:cli,
	}
}

type KvStore interface {
	GetUserInfoByCache(token string)(*UserInfo,error)
	SetUserInfoToCache(token string,info *UserInfo) error
}


//GetUserInfoByCache from redis get userInfo by token
func (rc *RedisCli) GetUserInfoByCache(token string)(*UserInfo,error) {
	if rc.cli == nil {
		return nil,fmt.Errorf("redis is disconnet")
	}
	result, e := rc.cli.Get(token).Result()
	if e != nil {
		return nil,e
	}
	var userInfo UserInfo
	e = json.Unmarshal([]byte(result), &userInfo)
	if e != nil {
		return nil,e
	}
	return &userInfo,nil
}

func (rc *RedisCli) SetUserInfoToCache(token string,info *UserInfo) error  {
	_, e := rc.cli.Set(token, *info, time.Minute * 30).Result()
	return e
}