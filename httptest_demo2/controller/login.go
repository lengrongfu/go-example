package controller

import (
	"github.com/lengrongfu/go-example/httptest_demo2/dto"
	"github.com/lengrongfu/go-example/httptest_demo2/util"
	"net/http"
	"strings"
)

const (
	UserName = "userName"
	Password = "password"
)

const (
	ErrUserNameOrPwdNotNull = "UserName Or Password is not null"
	NotLogin = "User not login"
)

func (e *Env) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	userName := r.FormValue(UserName)
	password := r.FormValue(Password)
	if strings.TrimSpace(userName) == "" || strings.TrimSpace(password) == "" {
		respDTO := dto.Error(ErrUserNameOrPwdNotNull, nil)
		w.Write(respDTO)
		return
	}
	info, err := e.sql.FindUserInfo(userName, password)
	if err != nil {
		respDTO := dto.Error(err.Error(), nil)
		w.Write(respDTO)
		return
	}
	token := util.Token()
	err = e.redis.SetUserInfoToCache(token, info)
	if err != nil {
		respDTO := dto.Error(err.Error(), nil)
		w.Write(respDTO)
		return
	}
	ok := dto.Ok(token)
	w.Write(ok)
}

