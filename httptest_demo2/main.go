package main

import (
	"fmt"
	"github.com/lengrongfu/go-example/httptest_demo2/config"
	"github.com/lengrongfu/go-example/httptest_demo2/controller"
	"log"
	"math/rand"
	"net/http"
	"time"
)


func main() {
	log.Println("Server start ... ")
	rand.Seed(time.Now().UnixNano())
	redisCli, err := config.NewRedis("localhost:6379")
	if err != nil {
		log.Panic(err)
	}
	dbStore, err := config.NewDB("root:123456@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8mb4,utf8")
	if err != nil {
		log.Panic(err)
	}
	env := controller.NewEnv(redisCli,dbStore)

	http.Handle("/api/login",http.HandlerFunc(env.Login))
	http.Handle("/api/userInfo",env.Auth(http.HandlerFunc(env.UserInfo)))
	fmt.Println(http.ListenAndServe(":8088", nil))
}