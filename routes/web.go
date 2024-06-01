package routes

import (
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	//01静态页面
	pc := new(controllers.PagesController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound) // 404 页面

	//02文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

	//03显示所有的文章
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	//04创建和保存文章
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store") // 保存表单数据的路由
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")

	//05编辑和更新文章
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")

	//06 删除文章
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	//静态资源01：
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	//中间件01 ：强制内容类型为 HTML
	//r.Use(middlewares.ForceHTML)
}
