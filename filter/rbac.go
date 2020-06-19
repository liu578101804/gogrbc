package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"encoding/json"
	"net/http"
	"github.com/liu578101804/gorbac/models"
	"strings"
	"fmt"
	"strconv"
	"github.com/liu578101804/gorbac/cache"
	"errors"
)

//缓存时间
const cacheTime = 60*60*24

func InitRBAC()  {

	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(c *context.Context) {

		//获取请求的url
		reqPath := c.Request.URL.Path
		beego.Info("请求的url：", reqPath)

		//获取UserId，这里可以从jwt解析获得，或者是session里面获得，这里简化了直接传，生产可不能这么做
		userIds := c.Request.Header["User-Id"]
		if len(userIds) == 0 {
			beego.Debug("请求头里面没找到User-Id")
			JsonResult(c, 1, "你没权限")
			return
		}
		//这里把传过来的转字符串然后转int
		userId,err := strconv.Atoi(fmt.Sprintf("%v",userIds[0]))
		if err != nil {
			beego.Debug("传过来的User-Id异常")
			JsonResult(c, 1, "你没权限")
			return
		}
		beego.Debug("请求的userId：",userId)

		//判断是否是超级管理员，这里判断id是为1，可以有其他判断逻辑
		if userId == 1 {
			beego.Debug("来自管理员的请求")
			return
		}

		roleId,err := getUserRoleIds(userId)
		allPathString,err := getRoleAllPath(roleId)

		//添加免鉴权的url
		allowPathString := beego.AppConfig.DefaultString("allowPathString","")
		allPathString = allowPathString + "," + allPathString

		beego.Debug("allPathString:", allPathString)

		//全部转换成小写才
		beego.Debug("比较的两个字符串：",strings.ToLower(allPathString), strings.ToLower(reqPath))
		index := strings.Index(strings.ToLower(allPathString), strings.ToLower(reqPath))
		beego.Debug("index:",index)
		if index < 0 {
			JsonResult(c, 1,"你没权限")
			return
		}

	})

}

func getUserRoleIds(userId int) (string,error) {

	cacheUserRole := fmt.Sprintf(cache.C_UserRole, userId)

	//查找用户的权限缓存
	if cache.IsExist(cacheUserRole) {
		allIdString := cache.Get(cacheUserRole).(string)
		beego.Debug("【取缓存】获取用户角色 userId：", userId, "roleIds：", allIdString)
		return allIdString,nil
	}

	//获取用户的所有角色组
	roleUserListM := models.NewRoleUserList()
	_,err := roleUserListM.GetByUserId(userId)
	if err != nil {
		beego.Error("获取用户的权限组失败:", err.Error())
		return "", err
	}

	//遍历用户的所有权限组合
	allRoleList := make([]string, 0)
	for _,role := range roleUserListM {
		allRoleList = append(allRoleList, fmt.Sprintf("%d",role.RoleId))
	}
	allIdString := strings.Join(allRoleList, ",")

	//写入缓存
	cache.Set(cacheUserRole, allIdString, cacheTime)

	beego.Debug("【走数据库】获取用户角色 userId：", userId, "roleIds：", allIdString)
	return allIdString,nil
}

func getRoleAllPath(roleIds string) (string, error) {

	beego.Debug("角色组字符：", roleIds)

	allPath := make([]string,0)
	//切割角色组字符
	roleArr := strings.Split(roleIds,",")
	if len(roleArr) < 1 {
		beego.Error("没找到角色组")
		return "", errors.New("没找到角色组")
	}

	//循环角色组
	for _,roleId := range roleArr {

		//缓存id
		cacheRolePath := fmt.Sprintf(cache.C_RolePath, roleId)
		//查看缓存是否存在
		if cache.IsExist(cacheRolePath) {
			allPath = append(allPath, cache.Get(cacheRolePath).(string))
			beego.Debug("【取缓存】获取角色权限 roleId:",roleId)
		}else{
			//转int
			id,_ := strconv.Atoi(roleId)

			//获取角色的权限
			roleNodeListM := models.NewRoleNodeList()
			_,err := roleNodeListM.GetByRoleId(id)
			if err != nil {
				beego.Error("获取角色权限失败",err.Error())
				return "", err
			}

			//遍历所有的节点
			allNode := make([]int,0)
			for _,node := range roleNodeListM  {
				allNode = append(allNode, node.NodeId)
			}

			//获取所有node的path
			nodeList,err := models.NewNodeModel().AllInIds(allNode)
			if err != nil {
				beego.Error("获取节点的path报错:",err.Error())
				return "",err
			}
			tmpPath := make([]string,0)
			for _,node := range nodeList  {
				tmpPath = append(tmpPath, node.Path)
				allPath = append(tmpPath, node.Path)
			}

			//设置缓存
			cache.Set(cacheRolePath, strings.Join(tmpPath,","), cacheTime)

			beego.Debug("【走数据库】roleId:",roleId)
		}
	}

	return strings.Join(allPath,","), nil
}


func JsonResult(c *context.Context, errCode int, errMsg string, data ...interface{})  {

	jsonData := make(map[string]interface{}, 3)
	jsonData["err_code"] = errCode
	jsonData["message"] = errMsg

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJson, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}

	http.Error(c.ResponseWriter, string(returnJson), http.StatusUnauthorized)

}


