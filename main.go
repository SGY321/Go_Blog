package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8") //设置标头响应
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello,欢迎来到happyBlog</h1>")
	} else {
		w.WriteHeader(404) //w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到:(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，其余问题请自行"+"<a href=\"www.baidu.com\">百度一下</a>")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)
	//文章详情
	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "创建新的文章")
		case "POST":
			fmt.Fprint(w, "访问文章列表")
		}
	})
	http.ListenAndServe(":3000", router)
}
