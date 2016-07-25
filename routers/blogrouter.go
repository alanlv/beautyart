package routers

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/beautyart/controllers/admin/blog"
)

func blogrouter() {

	//文章管理
	beego.Router("/blog/category/index", &blog.CategoryController{}, "*:Index")
	beego.Router("/blog/category/addcategory", &blog.CategoryController{}, "*:AddCategory")
	beego.Router("/blog/category/updatecategory", &blog.CategoryController{}, "*:UpdateCategory")
	beego.Router("/blog/category/deletecategory", &blog.CategoryController{}, "*:DeleteCategory")

	beego.Router("/blog/paper/index", &blog.PaperController{}, "*:Index")
	beego.Router("/blog/paper/addpaper", &blog.PaperController{}, "*:AddPaper")
	beego.Router("/blog/paper/updatepaper", &blog.PaperController{}, "*:UpdatePaper")
	beego.Router("/blog/paper/deletepaper", &blog.PaperController{}, "*:DeletePaper")

	beego.Router("/blog/paper/realdelpaper", &blog.PaperController{}, "*:RealDelPaper")
	beego.Router("/blog/paper/rubbish", &blog.PaperController{}, "*:Rubbish")

	//上传文件
	beego.Router("/public/file/upload", &blog.UploadController{}, "*:UploadFile")
	beego.Router("/public/file/filemanage", &blog.UploadController{}, "*:FileManage")

	//下载文件
	beego.Router("/public/file/getfile", &blog.UploadController{}, "*:GetWebFile")


	//相册管理
	beego.Router("/picture/album/index", &blog.AlbumController{}, "*:Index")
	beego.Router("/picture/album/addalbum", &blog.AlbumController{}, "*:AddCategory")
	beego.Router("/picture/album/updatealbum", &blog.AlbumController{}, "*:UpdateCategory")
	beego.Router("/picture/album/deletealbum", &blog.AlbumController{}, "*:DeleteCategory")


	beego.Router("/picture/photo/index", &blog.PhotoController{}, "*:Index")
	beego.Router("/picture/photo/addpaper", &blog.PhotoController{}, "*:AddPaper")
	beego.Router("/picture/photo/updatepaper", &blog.PhotoController{}, "*:UpdatePaper")
	beego.Router("/picture/photo/deletepaper", &blog.PhotoController{}, "*:DeletePaper")

	beego.Router("/picture/photo/realdelpaper", &blog.PhotoController{}, "*:RealDelPaper")
	beego.Router("/picture/photo/rubbish", &blog.PhotoController{}, "*:Rubbish")

}
