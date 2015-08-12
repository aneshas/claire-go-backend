package main

import (
	"database/sql"
	"fmt"
	"log"
)

type MysqlDb struct {
	Db *sql.DB
}

// Init initializes mysql databse and creates connectins
func (r *MysqlDb) Init() {
	connectionString := fmt.Sprintf("%s%s@tcp(%s:3306)/%s", MYSQL_DBUSER, MYSQL_DBPASS, MYSQL_DBHOST, MYSQL_DB)

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
