package connection

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft"
)

type ConnectionListener struct {
	server.Listener
}

func (l ConnectionListener) Accept() (session.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	if mcConn, ok := conn.(*minecraft.Conn); ok {
		id := mcConn.IdentityData().Identity
		addConn(id, mcConn)
	}

	return conn, err
}

func (l ConnectionListener) Disconnect(conn session.Conn, reason string) error {
	if mcConn, ok := conn.(*minecraft.Conn); ok {
		id := mcConn.IdentityData().Identity
		removeConn(id)
	}

	return l.Listener.Disconnect(conn.(*minecraft.Conn), reason)
}
