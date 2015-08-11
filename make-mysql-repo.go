package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB     string = "claire"
	DBUSER string = "root"
	DBPASS string = "" // eg ":mypasswd"
	DBHOST string = "mysql"
)

// Make mysql repository
// Implements IMakeRepository interface
type MakeMysqlRepo struct {
	Db *sql.DB
}

func (r *MakeMysqlRepo) Init() {
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

func (r *MakeMysqlRepo) Deinit() {
	r.Db.Close()
}

func (r MakeMysqlRepo) Get(id int) Make {
	return Make{1, "foo", "bar", "photo"}
}

func (r MakeMysqlRepo) GetAll(limit int) (result []Make) {
	query := fmt.Sprintf("select * from make limit %d", limit)

	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	make := Make{}

	for rows.Next() {
		rows.Scan(
			&make.Id,
			&make.Name,
			&make.Description,
			&make.Photo)

		result = append(result, make)
	}

	return
}

func (r MakeMysqlRepo) Query(query string) (makes []Make, err error) {
	return
}
