package models

import "github.com/astaxie/beego/orm"

func init()  {

	orm.RegisterModel(
		new(UserModel),
		new(NodeModel),
		new(RoleModel),
		new(RoleNodeModel),
		new(RoleUserModel),
	)
}

func TNUser() string {
	return "md_user"
}

func TNNode() string {
	return "md_node"
}

func TNRole() string {
	return "md_role"
}

func TNRoleNode() string {
	return "md_role_node"
}

func TNRoleUser() string {
	return "md_role_user"
}

func GetOrm() orm.Ormer {
	o := orm.NewOrm()
	o.Using("default")
	return o
}