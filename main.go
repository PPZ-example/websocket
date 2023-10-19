package main

import (
	"fmt"
	"net/http"
	"time"

	"log/slog"

	"github.com/gorilla/websocket"
)

func main() {
	// 静态文件
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // strip: 除去、撕掉

	// 测试普通的 http 接口
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello, yo!"))
	})

	upgrade_ws := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// websocket
	http.HandleFunc("/chat", func(res http.ResponseWriter, req *http.Request) {
		conn, err := upgrade_ws.Upgrade(res, req, nil)
		if err != nil {
			slog.Error("建立 ws 连接失败", "err", err)
			return
		}
		// read
		go func() {
			for {
				m_type, msg, err := conn.ReadMessage()
				if err != nil {
					slog.Error("读取 msg 失败", "err", err)
					continue
				}
				slog.Info("读取 msg",
					"m_type", m_type,
					"msg", msg,
				)
			}
		}()
		// write
		go func() {
			i := 0
			for {
				i += 1
				err := conn.WriteMessage(websocket.TextMessage, []byte(
					fmt.Sprintf("writing msg #%d", i),
				))
				if err != nil {
					slog.Error("写入 msg 失败", "err", err)
				}
				time.Sleep(time.Second * 5)
			}
		}()
	})

	http.ListenAndServe(":8080", nil)
}
