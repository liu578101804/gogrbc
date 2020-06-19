package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/all`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/all`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetRoleNodes",
			Router: `/getRoleNodes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:RoleController"],
		beego.ControllerComments{
			Method: "UpdateNode",
			Router: `/updateNode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Regist",
			Router: `/regist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/liu578101804/gorbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateRoles",
			Router: `/updateRoles`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
