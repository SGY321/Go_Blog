package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

// Article 对应一条文章数据
type Article struct {
	Title, Body string
	ID          int64
}

func main() {

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
