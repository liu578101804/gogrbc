package models

import (
	"time"
)

//用户表
type UserModel struct {
	Id       int		`orm:"pk;auto" `
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`

	Account string 	`orm:"unique"`
	Password string
}

func (m *UserModel) TableName() string {
	return TNUser()
}

func NewUserModel() *UserModel {
	return new(UserModel)
}

func (m *UserModel) Add() (int64,error) {
	return GetOrm().Insert(m)
}

func (m *UserModel) FindByAccount(account string) error {
	return GetOrm().QueryTable(TNUser()).Filter("account",account).One(m)
}