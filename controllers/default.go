package controllers

import (
	"redfind/util"
	"fmt"
)

type MainController struct {
	BaseController
}
var queryp= util.P{}

func (c *MainController) Get() {
	//下面显示多少个页码框，一定要用奇数，3，5，7
	number:=5
	explorerSearch:=c.GetString("explorerSearch")
	page:=c.GetString("page")
	curlpage:=util.ToInt(page)
	if curlpage<1{
		curlpage=1
	}
	queryp["__v"]="0"

	a:=util.D("test").Find(queryp).Page(0,10).AllData()
	totalcount:=util.D("test").Count()
	fmt.Print(util.ToString(a))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["page"]=PagerHtml(totalcount,10,curlpage,explorerSearch,number)
	c.Data["datalist"]=a
	c.TplName = "home.html"

}
