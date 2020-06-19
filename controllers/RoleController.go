package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liu578101804/gorbac/models"
	"strings"
	"strconv"
	"fmt"
	"github.com/liu578101804/gorbac/cache"
)

type RoleController struct {
	BaseController
}

// @Title 添加角色
// @Description 添加角色
// @Param	name		formData	string	true  "名字"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /add [post]
func (c *RoleController) Add() {

	name 	:= c.GetString("name","")

	m := models.NewRoleModel()
	m.Name = name

	if _,err := m.Add();err != nil{
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}
	c.JsonResult(0, "ok")
}

// @Title 更新角色下面的权限
// @Description 更新角色下面的权限
// @Param	role_id		formData	string	true  "角色id"
// @Param	nodes		formData	string	true  "权限id，用,（英文的逗号）分割"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /updateNode [post]
func (c *RoleController) UpdateNode() {

	nodes 	:= c.GetString("nodes","")
	roleId,_	:= c.GetInt("role_id", 1)

	//删除现有的节点
	roleNodeM := models.NewRoleNodeModel()
	if _,err := roleNodeM.DeleteByRoleId(roleId); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1,err.Error())
	}

	//创建列表
	roleNodeList := models.NewRoleNodeList()

	//切割nodes
	arr := strings.Split(nodes,",")
	for _,nodeId := range arr {

		roleNodeM := models.NewRoleNodeModel()
		roleNodeM.RoleId = roleId
		roleNodeM.NodeId,_ = strconv.Atoi(nodeId)

		roleNodeList = append(roleNodeList, roleNodeM)
	}

	//添加
	if _,err := roleNodeList.Add(); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1,err.Error())
	}


	//删除下缓存
	cacheRolePath := fmt.Sprintf(cache.C_RolePath, roleId)
	cache.Delete(cacheRolePath)

	c.JsonResult(0,"ok")
}

// @Title 获取角色下的权限列表
// @Description 获取角色下的权限列表
// @Param	role_id		formData	string	true  "角色id"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /getRoleNodes [post]
func (c *RoleController) GetRoleNodes()  {

	roleId,_ := c.GetInt("role_id", 1)

	//获取角色信息
	roleM := models.NewRoleModel()
	if err := roleM.GetById(roleId);err != nil {
		beego.Error("获取角色信息失败：",err.Error())
		c.JsonResult(1, "获取角色信息失败")
	}

	//查询下面的节点
	listM := models.NewRoleNodeList()
	if _,err := listM.GetByRoleId(roleId); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1,err.Error())
	}

	c.JsonResult(0, "ok", map[string]interface{}{
		"role": roleM,
		"node_list": listM,
	})
}


// @Title 删除角色
// @Description 删除角色
// @Param	role_id		formData	string	true  "角色id"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /delete [post]
func (c *RoleController) Delete() {

	roleId,_ := c.GetInt("role_id", 1)

	//获取角色信息
	roleM := models.NewRoleModel()
	if err := roleM.GetById(roleId);err != nil {
		beego.Error("获取角色信息失败：",err.Error())
		c.JsonResult(1, "获取角色信息失败")
	}

	//删除角色信息
	if _,err := models.NewRoleModel().DeleteById(roleId);err != nil {
		beego.Error("删除角色失败：", err.Error())
		c.JsonResult(1,"删除角色失败")
	}

	c.JsonResult(0,"ok")
}


// @Title 更新角色
// @Description 更新角色
// @Param	role_id		formData	string	true  "角色id"
// @Param	name		formData	string	true  "名字"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /update [post]
func (c *RoleController) Update() {
	roleId,_ := c.GetInt("role_id", 1)
	name := c.GetString("name","")

	//获取角色信息
	roleM := models.NewRoleModel()
	if err := roleM.GetById(roleId);err != nil {
		beego.Error("获取角色信息失败：",err.Error())
		c.JsonResult(1, "获取角色信息失败")
	}

	//赋值
	if name != "" { roleM.Name = name }

	//更新
	if _,err := roleM.UpdateById();err != nil {
		beego.Error("更新角色信息失败：", err.Error())
		c.JsonResult(1, "更新角色信息失败")
	}

	c.JsonResult(0, "ok")
}



// @Title 获取所有角色
// @Description 获取所有角色
// @Param	page		query	int	false  "页数"
// @Param	count		query	int	false  "每页数量"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /all [post]
func (c *RoleController) All() {

	page, _ := c.GetInt("page",1)
	count, _ := c.GetInt("count", 10)
	start := 0
	if page > 1 {
		start = (page - 1) * count
	}

	list := models.NewRoleList()
	_,err := list.All(start, count)
	if err != nil {
		beego.Error("获取全部角色失败：", err.Error())
		c.JsonResult(1,"获取全部角色失败")
	}

	c.JsonResult(0,"ok", list)
}