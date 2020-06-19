package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strings"
	"io"
	"compress/gzip"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) JsonResult(errCode int, errMsg string, data ...interface{}) {

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
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

	if strings.Contains(strings.ToLower(c.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		c.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		w := gzip.NewWriter(c.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJson)
		w.Flush()
	} else {
		io.WriteString(c.Ctx.ResponseWriter, string(returnJson))
	}
	c.StopRun()
}

