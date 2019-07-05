package main

import (
	"fmt"
	"net/http"
)
import (
	"encoding/json"
)

type RespDTO struct {
	Code int
	MSg string
	Data []byte
}

func Login(w http.ResponseWriter, r *http.Request) {
	dto := RespDTO{
		Code:200,
		MSg:"Login Success",
	}
	bytes, _ := json.Marshal(dto)
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/login", Login)
	fmt.Println(http.ListenAndServe(":8088", nil))
}