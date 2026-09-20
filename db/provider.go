package db

import "github.com/syseditor/libeloula/memcached"

type DataProvider interface {
	InitializeDB(username string, password string)
	AddPlayer(session memcached.PlayerSession)
	UpdatePlayer(session memcached.PlayerSession)
	RemovePlayer(uuid string)
	CheckTables() error
}

func NewDataProvider(source string) DataProvider {
	switch source {
	case "mariadb":
		return &MariaDBProvider{}
	default:
		return nil
	}
}
