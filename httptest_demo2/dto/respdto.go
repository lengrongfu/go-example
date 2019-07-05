package dto

import (
	"encoding/json"
)

var (
	NO_LOGIN = RespDTO{Code:-1,MSg:"用户未登陆"}
)

type RespDTO struct {
	Code int
	MSg string
	Data interface{}
}

func Ok(data interface{}) []byte {
	okRsp := RespDTO{Code:200,MSg:"Success",Data:data}
	bytes, e := json.Marshal(&okRsp)
	if e != nil {
		return Error(e.Error(),nil)
		return nil
	}
	return bytes
}

func Error(msg string,data interface{}) []byte {
	errRsp := RespDTO{Code:500,MSg:msg,Data:data}
	bytes, _ := json.Marshal(errRsp)
	return bytes
}
