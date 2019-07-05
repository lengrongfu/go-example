package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserInfo struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Pwd        string    `json:"pwd"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Phone      string    `json:"phone"`
	Addr       string    `json:"addr"`
	IsDelete   bool      `json:"is_delete"`
	InsertTime *time.Time `json:"insert_time"`
	UpdateTime *time.Time `json:"update_time"`
}

func (u UserInfo) String() string {
	return fmt.Sprintf("UserInfo[ Id:%d,Name:%s,Pwd:%s,Age:%d,Gender:%s,"+
		"Phone:%s,Addr:%s,IsDelete:%t,InsertTime:%s,UpdateTime:%s", u.Id, u.Name, u.Pwd,
		u.Age, u.Gender, u.Phone, u.Addr, u.IsDelete, u.InsertTime, u.UpdateTime)
}

func (u UserInfo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
