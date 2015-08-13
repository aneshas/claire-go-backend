package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "io/ioutil"
	"log"
	_ "os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	Db *sql.DB
}

const (
	MAX_MYSQL_RESULTS int = 20
)

var ErrNoResults = errors.New("No results matching your query.")

// Init initializes mysql databse and creates connectins
func (r *MysqlDb) Init() {
	connectionString := fmt.Sprintf("%s%s@tcp(%s:3306)%s", MYSQL_DBUSER, MYSQL_DBPASS, MYSQL_DBHOST, MYSQL_DB)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Database connection established")

	r.Db = db
}

// Deinit cleans up all open connections
func (r *MysqlDb) Deinit() {
	r.Db.Close()
}
