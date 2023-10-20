package ws

import (
	"fmt"
	"log/slog"

	"github.com/gorilla/websocket"
)

type worker struct {
	i    int
	sc   *safeclose
	conn *websocket.Conn
}

func (worker_ *worker) write() bool {
	closed := worker_.sc.lock()
	defer worker_.sc.mu.Unlock()

	worker_.i++

	if closed {
		slog.Info("write 时，遇到 ws 已关闭，停止 write 循环")
		return true
	}

	err := worker_.conn.WriteMessage(websocket.TextMessage, []byte(
		fmt.Sprintf("writing msg #%d", worker_.i),
	))
	if err != nil {
		slog.Error("写入 msg 失败", "err", err)
	}
	return false
}

func (worker_ *worker) read() bool {
	closed := worker_.sc.lock()
	defer worker_.sc.unlock()
	if closed {
		slog.Info("read 时，遇到 ws 已关闭，停止 read 循环")
		return true
	}

	m_type, msg, err := worker_.conn.ReadMessage()
	if err != nil {
		slog.Error("读取 msg 失败", "err", err)
		return true
	}
	slog.Info("读取 msg",
		"m_type", m_type,
		"msg", msg,
	)
	return false
}
