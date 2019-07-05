package controller

import (
	"fmt"
	"github.com/lengrongfu/go-example/httptest_demo2/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	TEST_USER_INFO = &model.UserInfo{
		Name:"冷荣富",
		Pwd:"123456",
		Age:18,
		Gender:"男",
		Phone:"17700000000",
	}
	TEST_TOKEN = "12abc3456defgh789ijk"
	TEST_NOLOGIN = "{\"Code\":500,\"MSg\":\"用户未登陆\",\"Data\":null}"
)

type redisTest struct {
	m map[string]*model.UserInfo
}


func (r *redisTest) GetUserInfoByCache(token string)(*model.UserInfo,error)  {
	if token  == "" {
		return nil,fmt.Errorf("token is nil")
	}
	if token != TEST_TOKEN {
		return nil,fmt.Errorf("invalid token %s",token)
	}
	info,ok := r.m[token]
	if !ok {
		return nil,fmt.Errorf("not exist userinfo %s",token)
	}
	return info,nil
}

func (r *redisTest) SetUserInfoToCache(token string,info *model.UserInfo) error  {
	r.m[token] = info
	return nil
}

// 测试middleware
/**
  测试中间层，同时测试了缓存层
 */
func TestEnv_Auth(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/userInfo", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(TOKEN,TEST_TOKEN)
	//测试Auth认证是否通过
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		info,ok  := r.Context().Value(USER_INFO).(model.UserInfo)
		if !ok {
			t.Fatal("return value error")
		}
		if !reflect.DeepEqual(info,*TEST_USER_INFO) {
			t.Fatalf("expect return value is %+v, got is %+v",info,TEST_USER_INFO)
		}
	})

	rr := httptest.NewRecorder()

	rt := new(redisTest)
	rt.m = make(map[string]*model.UserInfo)
	env := &Env{redis:rt}
	err = env.redis.SetUserInfoToCache(TEST_TOKEN, TEST_USER_INFO)
	if err != nil {
		t.Fatal(err)
	}
	handler := env.Auth(testHandler)
	handler.ServeHTTP(rr, req)
	//如果返回未登陆
	if rr.Body.String() != "" {
		t.Fatalf("expect return value is %s, got is %s","",rr.Body.String())
	}
}

func BenchmarkEnv_Auth(b *testing.B) {

}