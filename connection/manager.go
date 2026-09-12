package connection

import (
	"sync"

	"github.com/sandertv/gophertunnel/minecraft"
)

var mutex sync.RWMutex

type ConnectionManager struct {
	connections map[string]*minecraft.Conn
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{}
}

func (cm ConnectionManager) GetConn(id string) *minecraft.Conn {
	mutex.Lock()
	defer mutex.Unlock()
	return cm.connections[id]
}

func (cm ConnectionManager) addConn(id string, conn *minecraft.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	cm.connections[id] = conn
}

func (cm ConnectionManager) removeConn(id string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(cm.connections, id)
}
