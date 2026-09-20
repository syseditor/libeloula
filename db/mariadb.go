package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/syseditor/libeloula/memcached"
	"github.com/syseditor/libeloula/utils"
)

type MariaDBProvider struct {
	DB      *sql.DB
	queries map[string]interface{}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func criticalError(err error) {
	fmt.Printf("%s[Critical] %s-> %s%s\n%s", text.ANSI(text.DarkRed), text.ANSI(text.DarkGrey), text.ANSI(text.Red), err, text.ANSI(text.Reset))
	utils.Server.Close()
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
	p.queries = *new(map[string]interface{}{
		"createTable": map[string]string{
			"Players": "CREATE TABLE IF NOT EXISTS Libeloula.Players (uuid VARCHAR2 NOT NULL, username VARCHAR2 NOT NULL, joined_at VARCHAR2 NOT NULL, blocks_broken INT DEFAULT 0, PRIMARY KEY (uuid));",
		},
	})

	fmt.Printf("%sSuccessfully connected to database!\n", text.ANSI(text.Blue))

	err = p.CheckTables()
	if err != nil {
		criticalError(err)
	}
}

func (p MariaDBProvider) CheckTables() error {
	if tables, ok := p.queries["createTable"].(map[string]string); ok {
		for key, value := range tables {
			_, err := p.DB.Exec(value)
			if err != nil {
				return err
			}
			fmt.Printf("%sTable %s is properly registered in the database!\n", text.ANSI(text.Green), key)
		}
	} else {
		return errors.New("'createTable' queries are not type of map[string]string.")
	}
	return nil
}

func (p MariaDBProvider) AddPlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) UpdatePlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) RemovePlayer(uuid string) {

}
