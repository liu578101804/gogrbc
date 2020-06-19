package models

import "time"

type (
	RoleNodeList []*RoleNodeModel
)

//角色与权限关联表
type RoleNodeModel struct {
	Id 			int		`orm:"pk;auto" `
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`

	RoleId 		int
	NodeId 		int

	Path 		string 	`orm:"-"`
}

func (m *RoleNodeModel) TableName() string {
	return TNRoleNode()
}

func NewRoleNodeModel() *RoleNodeModel {
	return new(RoleNodeModel)
}

func NewRoleNodeList() RoleNodeList {
	return make(RoleNodeList,0)
}

func (l *RoleNodeList) Add() (int64,error) {
	return GetOrm().InsertMulti(1, l)
}

func (l *RoleNodeList) GetByRoleId(roleId int) (int64,error) {
	return GetOrm().QueryTable(TNRoleNode()).Filter("role_id", roleId).All(l)
}


func (m *RoleNodeModel) DeleteByRoleId(roleId int) (int64,error) {
	return GetOrm().QueryTable(TNRoleNode()).Filter("role_id", roleId).Delete()
}
