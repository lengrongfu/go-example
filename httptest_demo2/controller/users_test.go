package controller

import (
	"encoding/json"
	"fmt"
	"github.com/lengrongfu/go-example/httptest_demo2/dto"
	"github.com/lengrongfu/go-example/httptest_demo2/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnv_UserInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/userInfo", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(TOKEN,TEST_TOKEN)

	rr := httptest.NewRecorder()

	rt := new(redisTest)
	rt.m = make(map[string]*model.UserInfo)
	env := &Env{redis:rt}
	err = env.redis.SetUserInfoToCache(TEST_TOKEN, TEST_USER_INFO)
	if err != nil {
		t.Fatal(err)
	}
	handler := env.Auth(http.HandlerFunc(env.UserInfo))
	handler.ServeHTTP(rr, req)
	var dResp dto.RespDTO
	err = json.Unmarshal([]byte(rr.Body.String()), &dResp)
	if err != nil {
		t.Fatal(err)
	}
	//如果返回未登陆
	if dResp.Code != 200 || dResp.MSg != "Success" || dResp.Data == nil {
		t.Fatal(fmt.Errorf("expect return value is %s, got is %s","",rr.Body.String()))
	}
}
