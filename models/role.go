package models

import "time"

type (
	RoleList []*RoleModel
)

//角色表
type RoleModel struct {
	Id 		int 	`orm:"pk;auto" `
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`

	Name 	string
	PId 	int
}

func (m *RoleModel) TableName() string {
	return TNRole()
}

func NewRoleModel() *RoleModel {
	return new(RoleModel)
}

func NewRoleList() RoleList {
	return make(RoleList, 0)
}

func (m *RoleModel) Add() (int64,error) {
	return GetOrm().Insert(m)
}

func (m *RoleModel) UpdateById() (int64,error) {
	return GetOrm().Update(m)
}

func (m *RoleModel) DeleteById(id int) (int64,error) {
	return GetOrm().QueryTable(TNRole()).Filter("id",id).Delete()
}

func (m *RoleModel) GetById(id int) (error) {
	return GetOrm().QueryTable(TNRole()).Filter("id",id).One(m)
}


func (l *RoleList) All(offset, limit int) (int64,error) {
	return GetOrm().QueryTable(TNRole()).Offset(offset).Limit(limit).All(l)
}