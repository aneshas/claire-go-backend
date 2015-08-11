package main

import (
	"database/sql"
	"fmt"
	"log"
)

type MysqlDb struct {
	Db *sql.DB
}

func (r *MysqlDb) Init() {
	connectionString := fmt.Sprintf("%s%s@tcp(%s:3306)/%s", DBUSER, DBPASS, DBHOST, DB)

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

func (r *MysqlDb) Deinit() {
	r.Db.Close()
}
