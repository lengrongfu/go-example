package controller

import (
	"encoding/json"
	"fmt"
	"github.com/lengrongfu/go-example/httptest_demo2/dto"
	"github.com/lengrongfu/go-example/httptest_demo2/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	TEST_LOGIN_USERNAME = "lrf"
	TEST_LOGIN_PASSWORD = "123456"
)

type testlogin struct {
	m map[string]*model.UserInfo
}

func (t *testlogin) FindUserInfo(userName,pwd string) (*model.UserInfo,error)  {
	if userName == TEST_LOGIN_USERNAME && pwd == TEST_LOGIN_PASSWORD {
		return TEST_USER_INFO,nil
	}
	return nil,fmt.Errorf("invalid username:%s or password:%s",userName,pwd)
}


func (t *testlogin) GetUserInfoByCache(token string)(*model.UserInfo,error)  {
	return nil,nil
}

func (t *testlogin) SetUserInfoToCache(token string,info *model.UserInfo) error  {
		t.m[token] = info
		return nil
}

func TestEnv_Login(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.PostForm = make(url.Values)
	request.PostForm.Set(UserName,TEST_LOGIN_USERNAME)
	request.PostForm.Set(Password,TEST_LOGIN_PASSWORD)

	recorder := httptest.NewRecorder()

	et := &testlogin{
		m:make(map[string]*model.UserInfo),
	}
	env := NewEnv(et,et)

	handler := http.HandlerFunc(env.Login)
	handler.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Fatal(fmt.Errorf("expecte http.code is 200 , got is %d ",recorder.Code))
	}

	var dRes dto.RespDTO
	err = json.Unmarshal([]byte(recorder.Body.String()), &dRes)
	if err != nil || &dRes == nil{
		t.Fatal(err)
	}
	expectResp := "{\"Code\":200,\"MSg\":\"Success\",\"Data\":\"70b4c0d3b2431ec6f65b445c260614e4\"}"
	if dRes.Code != 200 || dRes.MSg != "Success" || dRes.Data == ""{
		t.Fatal(fmt.Errorf("expect return value is %s, got is %s",expectResp,recorder.Body.String()))
	}
}
