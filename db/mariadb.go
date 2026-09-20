package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/syseditor/libeloula/memcached"
)

type MariaDBProvider struct {
	DB      *sql.DB
	queries *map[string]interface{}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *MariaDBProvider) InitializeDB(username string, password string) {
	var err error
	p.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost)/Libeloula", username, password))
	check(err)

	p.DB.SetConnMaxLifetime(time.Minute)
	p.DB.SetMaxOpenConns(30)
	p.DB.SetMaxIdleConns(30)

	err = p.DB.Ping()
	check(err)

	//Load all available db queries
	p.queries = new(map[string]interface{}{
		"createTable": map[string]string{
			"player": "CREATE TABLE IF NOT EXISTS Players (UUID VARCHARACTER, Username VARCHARACTER, JoinedAt VARCHARACTER, BlocksBroken SMALLINT);",
		},
	})

	fmt.Printf("%sSuccessfully connected to database!\n", text.Blue)
}

func (p MariaDBProvider) AddPlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) UpdatePlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) RemovePlayer(uuid string) {

}
