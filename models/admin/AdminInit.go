// 后台模型数据填充包
package admin

import (
	"fmt"
	. "github.com/hunterhug/beautyart/lib"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = "admin"
	u.Nickname = "画室"
	u.Password = Pwdhash("admin")
	u.Email = "569929309@qq.com"
	u.Remark = "最高权限的王"
	u.Status = 2
	u.Createtime = GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "user"
	u1.Nickname = "测试用户"
	u1.Password = Pwdhash("user")
	u1.Email = "569929309@qq.com"
	u1.Remark = "测试用户"
	u1.Status = 2
	u1.Createtime = GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "画室官网"
	g.Title = "后台管理"
	g.Sort = 1
	g.Status = 1
	g.Insert()
	g1 := new(Group)
	g1.Name = "画室官网"
	g1.Title = "文章管理"
	g1.Sort = 2
	g1.Status = 1
	g1.Insert()
	g2 := new(Group)
	g2.Name = "画室官网"
	g2.Title = "图片管理"
	g2.Sort = 3
	g2.Status = 1
	g2.Insert()
	fmt.Println("insert group end")
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	r.Insert()
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	g1 := new(Group)
	g1.Id = 2
	g2 := new(Group)
	g2.Id = 3
	//nodes := make([20]Node)
	nodes := [53]Node{
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 25, Name: "account", Title: "配置中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 26, Name: "option/index", Title: "网站配置", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 27, Name: "Index", Title: "网站配置首页", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 28, Name: "UpdateOption", Title: "更新网站配置", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 29, Name: "user/index", Title: "个人信息", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 30, Name: "Index", Title: "个人信息首页", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 31, Name: "UpdateUser", Title: "更新个人信息", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 32, Name: "blog", Title: "文章中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g1},
		{Id: 33, Name: "post/index", Title: "文章列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 34, Name: "Index", Title: "文章列表首页", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 35, Name: "AddPost", Title: "增加文章", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 36, Name: "UpdatePost", Title: "修改文章", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 37, Name: "DelPost", Title: "回收文章", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 38, Name: "RealDelPost", Title: "删除文章", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 39, Name: "category/index", Title: "目录列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 40, Name: "Index", Title: "目录列表首页", Remark: "", Level: 3, Pid: 39, Status: 1, Group: g1},
		{Id: 41, Name: "AddCategory", Title: "增加目录", Remark: "", Level: 3, Pid: 39, Status: 1, Group: g1},
		{Id: 42, Name: "UpdateCategory", Title: "修改目录", Remark: "", Level: 3, Pid: 39, Status: 1, Group: g1},
		{Id: 43, Name: "picture", Title: "图片中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g2},
		{Id: 44, Name: "album/index", Title: "相册管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 45, Name: "Index", Title: "相册首页", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 46, Name: "AddAlbum", Title: "增加相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 47, Name: "DelAlbum", Title: "删除相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 48, Name: "UpdateAlbum", Title: "修改相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 49, Name: "photo/index", Title: "图片管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 50, Name: "Index", Title: "图片首页", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 51, Name: "AddPhoto", Title: "增加图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 52, Name: "DelPhoto", Title: "删除图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 53, Name: "UpdatePhoto", Title: "修改图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		n.Insert()
	}
	fmt.Println("insert node end")
}
