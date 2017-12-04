package routers

import (
	"bloggo/controllers"
	"bloggo/controllers/index"
	"bloggo/controllers/admin"
	"github.com/astaxie/beego"
	//"github.com/derekparker/delve/service/api"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/index", &index.IndexController{})
	beego.Router("/ueidtor", &controllers.UeditorController{},"*:Ueditor")
	beego.Router("/attachment/*", &controllers.AttachController{})
	admin :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/", &admin.IndexController{}),
			beego.NSRouter("/upload", &admin.UploadController{},"post:PostUpload"),
			beego.NSRouter("/success", &admin.SuccessController{}),
			beego.NSNamespace("/article",
				beego.NSRouter("/list", &admin.ArticleController{}, "get:ListArticle"),
				beego.NSRouter("/add", &admin.ArticleController{}, "get:AddArticle;post:AddArticle"),
				beego.NSRouter("/detail/:id", &admin.ArticleController{}, "get:GetDetailArticle"),
				beego.NSRouter("/update/?:id", &admin.ArticleController{}, "get:UpdateArticle;post:UpdateArticle"),
				//beego.NSRouter("/update/", &admin.ArticleController{}, "post:UpdateArticle"),
			),
		)
	//注册 namespace
	beego.AddNamespace(admin)
}
