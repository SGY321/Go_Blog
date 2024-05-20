// Package route 路由相关
package route

import (
	"github.com/gorilla/mux"
)

// Name2URL 通过路由来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	var route *mux.Router
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}
	return url.String()
}
