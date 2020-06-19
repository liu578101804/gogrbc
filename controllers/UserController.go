package controllers

import (
	"github.com/liu578101804/gorbac/models"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"github.com/liu578101804/gorbac/cache"
	"fmt"
)

type UserController struct {
	BaseController
}


// @Title 登录
// @Description 登录
// @Param	account		formData	string	true  "账号"
// @Param	password	formData	string	true  "密码"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {

	password 	:= c.GetString("password","")
	account 	:= c.GetString("account","")

	//查找用户
	mu := models.NewUserModel()
	err := mu.FindByAccount(account)
	if err != nil {
		beego.Error("找不到账号：",err.Error())
		c.JsonResult(1,"找不到账号")
	}

	if mu.Password != password {
		beego.Error("密码错误：")
		c.JsonResult(1,"密码错误")
	}

	c.JsonResult(0,"登录成功")
}

// @Title 注册
// @Description 注册
// @Param	account		formData	string	true  "账号"
// @Param	password	formData	string	true  "密码"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /regist [post]
func (c *UserController) Regist() {

	password 	:= c.GetString("password","")
	account 	:= c.GetString("account","")

	//查找用户
	mu := models.NewUserModel()
	err := mu.FindByAccount(account)
	if err == nil {
		beego.Error(account+" 账号已经存在")
		c.JsonResult(1,"账号已经存在")
	}

	//添加用户
	mu.Password = password
	mu.Account = account
	if _,err := mu.Add();err != nil{
		beego.Error("注册失败：",err.Error())
		c.JsonResult(1, "注册失败")
	}

	//添加用户角色
	mru := models.NewRoleUserModel()
	mru.UserId = mu.Id
	mru.RoleId = 1
	if _,err := mru.Add();err != nil{
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}

	c.JsonResult(0, "ok")
}

// @Title 更新用户角色
// @Description 更新用户角色
// @Param	user_id		formData	string	true  "用户id"
// @Param	roles	formData	string	true  "角色id字符串，用户,（英文逗号）分割"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /updateRoles [post]
func (c *UserController) UpdateRoles(){

	roles 	:= c.GetString("roles","")
	userId,_	:= c.GetInt("user_id", 1)

	//删除现有的权限
	roleUserM := models.NewRoleUserModel()
	if _,err := roleUserM.DeleteByUserId(userId); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1,err.Error())
	}

	//创建列表
	roleUserList := models.NewRoleUserList()

	//切割nodes
	arr := strings.Split(roles,",")
	for _,rolesId := range arr {

		roleUserM := models.NewRoleUserModel()
		roleUserM.UserId = userId
		roleUserM.RoleId,_ = strconv.Atoi(rolesId)

		roleUserList = append(roleUserList, roleUserM)
	}

	//添加
	if _,err := roleUserList.Add(); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1,err.Error())
	}

	//删除下缓存
	cacheUserRole := fmt.Sprintf(cache.C_UserRole, userId)
	cache.Delete(cacheUserRole)

	c.JsonResult(0,"ok")
}
