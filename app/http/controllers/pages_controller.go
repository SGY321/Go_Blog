// 应用控制层
package controllers

import (
	"fmt"
	"net/http"
)

// 处理静态页面
type PagesController struct {
}

// Home 首页
func (pc *PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello,欢迎来到happyBlog</h1>")
}

// About 关于我们页面
func (pc *PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，其余问题请自行"+"<a href=\"www.baidu.com\">百度一下</a>")
}

// NotFound 404 页面
func (pc *PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到:(</h><p>如有疑惑，请联系我们。</p>)")
}
