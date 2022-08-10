package main

import (
	"database/sql"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Result interface {
	LastInsertId() (int64, error) // 使用 INSERT 向数据插入记录，数据表有自增 id 时，该函数有返回值
	RowsAffected() (int64, error) // 表示影响的数据表行数
}

var router *mux.Router
var db *sql.DB

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置标头
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	database.Initialize()
	db = database.DB

	router = bootstrap.SetupRoute()

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
