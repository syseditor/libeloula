package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syseditor/libeloula/memcached"
)

type MariaDBProvider struct {
	DB      *sql.DB
	queries *mariaDBQueries
}

type mariaDBQueries map[string]interface{}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *MariaDBProvider) InitializeDB(username string, password string) {
	var err error
	p.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost)/Libeloula", username, password))
	check(err)

	err = p.DB.Ping()
	check(err)

	//Load all available db queries from queries.json
	file, err := os.Open("queries.json")
	check(err)

	dec := json.NewDecoder(file)
	dec.Decode(p.queries)
}

func (p MariaDBProvider) AddPlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) UpdatePlayer(session memcached.PlayerSession) {

}

func (p MariaDBProvider) RemovePlayer(uuid string) {

}
