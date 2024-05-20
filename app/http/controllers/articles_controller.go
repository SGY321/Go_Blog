package controllers

import (
	"database/sql"
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
	"text/template"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	//1. 获取 URL 参数
	// vars := mux.Vars(r) // 从HTTP请求中获取路由参数的值
	// id := vars["id"]
	id := getRouteVariable("id", r)

	//2. 读取对应的文章数据
	// article := Article{}
	// query := "SELECT * FROM articles WHERE id = ?"
	// err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	//QueryRow() 来读取单条数据；Scan() 将查询结果赋值到我们的 article struct 中，传参应与数据表字段的顺序保持一致。
	article, err := getArticleByID(id)

	//3 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows { //Scan() 发现没有返回数据的话，会返回 sql.ErrNoRows 类型的错误
			//3.1 数据未找到
			w.WriteHeader(404) //http.StatusNotFound
			fmt.Fprint(w, "404 文章未找到")
		} else {
			//3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		//4. 读取成功
		// tmpl, err := template.ParseFiles("resources/views/articles/show.gohtml") // 加载模板文件show.gohtml,后缀名也可以是,tmpl
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
			"Int64ToString": types.Int64ToString,
		}).ParseFiles("resources/views/articles/show.gohtml")

		logger.LogError(err)
		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}
