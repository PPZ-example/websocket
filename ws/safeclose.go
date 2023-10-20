package ws

import (
	"sync"
)

type safeclose struct {
	value bool
	mu    sync.Mutex
}

func (sc_ *safeclose) lock() bool {
	sc_.mu.Lock()
	return sc_.value
}
func (sc_ *safeclose) unlock() {
	sc_.mu.Unlock()
}
