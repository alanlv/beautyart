package blog

import (
	// "fmt"
	//"github.com/astaxie/beego"
	. "github.com/hunterhug/beautyart/lib"
	//"github.com/hunterhug/beautyart/models/admin"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/beautyart/models/blog"
)

type RollController struct {
	baseController
}

func (this *RollController) Index() {
	roll := new(blog.Roll)
	rolls := []orm.Params{}
	roll.Query().OrderBy("-Status", "-Sort", "Createtime").Values(&rolls)
	this.Data["roll"] = rolls
	this.TplName = this.GetTemplate() + "/blog/rollindex.html"
}

func (this *RollController) AddRoll() {
	roll := new(blog.Roll)
	num, _ := roll.Query().Count()
	if num >= 6 {
		this.Rsp(false,"图片不能超过六张")
	}

	title := this.GetString("title", "")
	content := this.GetString("content", "")
	url := this.GetString("url", "")
	sort, _ := this.GetInt64("sort", 0)
	photo := this.GetString("photo")
	roll.Content = content
	roll.Createtime = GetTime()
	roll.Photo = photo
	roll.Status = 0
	roll.Title = title
	roll.Url = url
	roll.Sort = sort
	roll.View = 0
	err := roll.Insert()
	if err != nil {
		this.Rsp(false, err.Error())
	}else {
		this.Rsp(true, "增加成功")
	}
}

func (this *RollController) UpdateRoll() {
}
func (this *RollController) DeleteRoll() {
	id,_:=this.GetInt64("id",-1)
	if id==-1{
		this.Rsp(false,"id参数问题")
	}
	roll:=new(blog.Roll)
	roll.Id=id
	err:=roll.Delete()
	if err!=nil{
		this.Rsp(false,err.Error())
	}else{
		this.Rsp(true,"删除成功")
	}
}
