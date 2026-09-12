package connection

import (
	"sync"

	"github.com/sandertv/gophertunnel/minecraft"
)

var mutex sync.RWMutex
var connectionManager map[string]*minecraft.Conn

func GetConn(id string) *minecraft.Conn {
	mutex.Lock()
	defer mutex.Unlock()
	return connectionManager[id]
}

func addConn(id string, conn *minecraft.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	connectionManager[id] = conn
}

func removeConn(id string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(connectionManager, id)
}
