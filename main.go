package main

import (
	"net/http"

	"zzz.ppz/ws"
)

func main() {
	// 静态文件
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // strip: 除去、撕掉

	// 测试普通的 http 接口
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello, yo!"))
	})

	// websocket
	http.HandleFunc("/chat", ws.Handle)

	http.ListenAndServe(":8080", nil)
}
