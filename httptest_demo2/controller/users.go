package controller

import (
	"github.com/lengrongfu/go-example/httptest_demo2/dto"
	"github.com/lengrongfu/go-example/httptest_demo2/model"
	"net/http"
)

func (e *Env) UserInfo (w http.ResponseWriter, r *http.Request)  {
	userInfo,ok := r.Context().Value(USER_INFO).(model.UserInfo)
	if !ok {
		w.Write(dto.Error(dto.NO_LOGIN.MSg,dto.NO_LOGIN.Data))
		return
	}
	w.Write(dto.Ok(userInfo))
}
