package connection

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft"
)

type connectionListener struct {
	server.Listener
	connectionManager *ConnectionManager
}

func (l connectionListener) Accept() (session.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	if mcConn, ok := conn.(*minecraft.Conn); ok {
		id := mcConn.IdentityData().Identity
		l.connectionManager.addConn(id, mcConn)
	}

	return conn, err
}

func (l connectionListener) Disconnect(conn session.Conn, reason string) error {
	if mcConn, ok := conn.(*minecraft.Conn); ok {
		id := mcConn.IdentityData().Identity
		l.connectionManager.removeConn(id)
	}

	return l.Listener.Disconnect(conn.(*minecraft.Conn), reason)
}
