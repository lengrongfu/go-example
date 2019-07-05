package model

import (
	"database/sql"
)

type DbCli struct {
	cli *sql.DB
}

func NewDbCli(cli *sql.DB) *DbCli {
	return &DbCli{
		cli:cli,
	}
}

type DbStore interface {
	FindUserInfo(userName,pwd string) (*UserInfo,error)
}

func (b DbCli) FindUserInfo(userName,pwd string) (*UserInfo,error)	  {
	var userInfo UserInfo
	err := b.cli.QueryRow("select * from user_info where name = ? and pwd = ?", userName, pwd).
		Scan(&userInfo.Id, &userInfo.Name, &userInfo.Pwd, &userInfo.Age, &userInfo.Gender, &userInfo.Phone, &userInfo.Addr, &userInfo.IsDelete, &userInfo.InsertTime, &userInfo.UpdateTime)
	if err != nil {
		return  nil,err
	}
	return &userInfo,nil
}