package controller

import (
	"context"
	"github.com/lengrongfu/go-example/httptest_demo2/dto"
	"log"
	"net/http"
)

const (
	TOKEN = "token"
	USER_INFO = "user_info"
)

//Auth 认证
func (env *Env)Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		token := r.Header.Get(TOKEN)
		userInfo, e := env.redis.GetUserInfoByCache(token)
		if e != nil {
			log.Println(e.Error())
			w.Write(dto.Error(dto.NO_LOGIN.MSg,nil))
			return
		}
		value := context.WithValue(r.Context(), USER_INFO, *userInfo)
		r = r.WithContext(value)
		next.ServeHTTP(w,r)
	})
}