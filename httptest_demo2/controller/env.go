package controller

import "github.com/lengrongfu/go-example/httptest_demo2/model"

type Env struct {
	redis model.KvStore
	sql model.DbStore
}

func NewEnv(redis model.KvStore,db model.DbStore) *Env {
	return &Env{
		redis:redis,
		sql:db,
	}
}