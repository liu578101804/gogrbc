package models

import "time"

type (
	NodeList []*NodeModel
)


//节点表
type NodeModel struct {
	Id		int		`orm:"pk;auto" `
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`

	Name 		string
	Controller 	string
	Function 	string
	Path 		string
}

func (m *NodeModel) TableName() string {
	return TNNode()
}

func NewNodeModel() *NodeModel {
	return new(NodeModel)
}
func NewNodeList() NodeList  {
	return make(NodeList, 1)
}

func (m *NodeModel) Add() (int64,error) {
	return GetOrm().Insert(m)
}

func (m *NodeModel) UpdateById() (int64,error) {
	return GetOrm().Update(m)
}

func (m *NodeModel) DeleteById(id int) (int64,error) {
	return GetOrm().QueryTable(TNNode()).Filter("id",id).Delete()
}

func (m *NodeModel) GetById(id int) (error) {
	return GetOrm().QueryTable(TNNode()).Filter("id",id).One(m)
}

func (m *NodeModel) AllInIds(nodeIds []int) (list NodeList,err error) {
	_,err = GetOrm().QueryTable(TNNode()).Filter("id__in", nodeIds).All(&list)
	return
}

func (l *NodeList) All(offset, limit int) (int64,error) {
	return GetOrm().QueryTable(TNNode()).Offset(offset).Limit(limit).All(l)
}