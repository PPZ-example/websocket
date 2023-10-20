package ws

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrade_ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handle(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrade_ws.Upgrade(res, req, nil)
	if err != nil {
		slog.Error("建立 ws 连接失败", "err", err)
		return
	}

	sc := safeclose{value: false}
	worker := worker{
		sc:   &sc,
		conn: conn,
	}
	// read
	go func() {
		for {
			if worker.read() {
				return
			}
		}
	}()

	// write
	go func() {
		for {
			if worker.write() {
				return
			}
			time.Sleep(time.Second * 3)
		}
	}()

	// close
	go func() {
		time.Sleep(time.Second * 12)
		sc.mu.Lock()
		sc.value = true
		slog.Info("wx 关闭")
		conn.Close()
		sc.mu.Unlock()
	}()
}
