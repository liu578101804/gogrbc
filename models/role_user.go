package models

import "time"

type (
	RoleUserList []*RoleUserModel
)

//角色与用户关联表
type RoleUserModel struct {
	Id 			int		`orm:"pk;auto" `
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`

	RoleId 		int
	UserId 		int
}

func (m *RoleUserModel) TableName() string {
	return TNRoleUser()
}

func NewRoleUserModel() *RoleUserModel {
	return new(RoleUserModel)
}

func NewRoleUserList() RoleUserList {
	return make(RoleUserList, 0)
}

func (m *RoleUserModel) Add() (int64,error) {
	return GetOrm().Insert(m)
}

func (m *RoleUserModel) DeleteByUserId(userId int) (int64,error) {
	return GetOrm().QueryTable(TNRoleUser()).Filter("user_id", userId).Delete()
}


func (l *RoleUserList) Add() (int64,error) {
	return GetOrm().InsertMulti(1, l)
}

func (l *RoleUserList) GetByUserId(userId int) (int64,error) {
	return GetOrm().QueryTable(TNRoleUser()).Filter("user_id", userId).All(l)
}

