package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liu578101804/gorbac/models"
)

type NodeController struct {
	BaseController
}

// @Title 添加权限
// @Description 添加权限
// @Param	name		formData	string	true  "权限名字"
// @Param	controller	formData	string	true  "控制器名字"
// @Param	function	formData	string	true  "方法"
// @Param	path		formData	string	true  "路径"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /add [post]
func (c *NodeController) Add() {

	name 	:= c.GetString("name","")
	controller 	:= c.GetString("controller","")
	function 	:= c.GetString("function","")
	path 	:= c.GetString("path","")

	m := models.NewNodeModel()
	m.Name = name
	m.Controller = controller
	m.Function = function
	m.Path = path

	if _,err := m.Add();err != nil{
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}
	c.JsonResult(0, "ok")

}


// @Title 更新权限
// @Description 更新权限
// @Param	node_id		formData	int	true  "权限id"
// @Param	name		formData	string	true  "权限名字"
// @Param	controller	formData	string	true  "控制器名字"
// @Param	function	formData	string	true  "方法"
// @Param	path		formData	string	true  "路径"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /update [post]
func (c *NodeController) Update() {

	nodeId,_ 	:= c.GetInt("node_id",0)
	name 	:= c.GetString("name","")
	controller 	:= c.GetString("controller","")
	function 	:= c.GetString("function","")
	path 	:= c.GetString("path","")

	//查询
	m := models.NewNodeModel()
	if err := m.GetById(nodeId);err != nil {
		beego.Debug("获取node失败：",err.Error())
		c.JsonResult(1, "获取node失败")
	}

	if name != "" { m.Name = name }
	if controller != "" { m.Controller = controller }
	if function != "" { m.Function = function }
	if path != "" { m.Path = path }

	if _,err := m.UpdateById();err != nil{
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}
	c.JsonResult(0, "ok")

}

// @Title 删除权限
// @Description 删除权限
// @Param	node_id		formData	int	true  "权限id"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /delete [post]
func (c *NodeController) Delete() {

	nodeId,_ 	:= c.GetInt("node_id",0)

	//查询
	m := models.NewNodeModel()
	if err := m.GetById(nodeId);err != nil {
		beego.Debug("获取node失败：",err.Error())
		c.JsonResult(1, "获取node失败")
	}

	if _,err := m.DeleteById(nodeId);err != nil{
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}
	c.JsonResult(0, "ok")

}

// @Title 查询权限
// @Description 查询权限
// @Param	node_id		formData	int	true  "权限id"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /get [post]
func (c *NodeController) Get() {

	nodeId,_ 	:= c.GetInt("node_id",0)

	//查询
	m := models.NewNodeModel()
	if err := m.GetById(nodeId);err != nil {
		beego.Debug("获取node失败：",err.Error())
		c.JsonResult(1, "获取node失败")
	}

	c.JsonResult(0, "ok", m)
}

// @Title 全部权限
// @Description 全部权限
// @Param	page		query	int	false  "页数"
// @Param	count		query	int	false  "每页数量"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /all [post]
func (c *NodeController) All() {

	page, _ := c.GetInt("page",1)
	count, _ := c.GetInt("count", 10)
	start := 0
	if page > 1 {
		start = (page - 1) * count
	}

	list := models.NewNodeList()
 	_,err := list.All(start, count)
	if err != nil {
		beego.Error("获取全部节点失败：", err.Error())
		c.JsonResult(1,"获取全部节点失败")
	}

	c.JsonResult(0,"ok", list)
}
