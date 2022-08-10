package main

import (
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"
)

type Result interface {
	LastInsertId() (int64, error) // 使用 INSERT 向数据插入记录，数据表有自增 id 时，该函数有返回值
	RowsAffected() (int64, error) // 表示影响的数据表行数
}

func main() {

	bootstrap.SetupDB()

	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
