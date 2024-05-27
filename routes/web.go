package routes

import (
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	//静态页面
	pc := new(controllers.PagesController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound) // 404 页面

	//文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

	//显示所有的文章
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	// RegisterWebRoutes 注册网页相关路由
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store") // 保存表单数据的路由
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
}