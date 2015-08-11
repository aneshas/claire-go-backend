package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MakeMysqlRepo mysql repository
// Implements IMakeRepository interface
type MakeMysqlRepo struct {
	*MysqlDb
}

// Helper funcs

func (r MakeMysqlRepo) Map(row *sql.Row) (make Make, err error) {
	err = row.Scan(
		&make.Id,
		&make.Name,
		&make.Description,
		&make.Photo)

	return
}

func (r MakeMysqlRepo) MapRows(rows *sql.Rows) (makes []Make, err error) {
	make := Make{}

	for rows.Next() {
		err = rows.Scan(
			&make.Id,
			&make.Name,
			&make.Description,
			&make.Photo)

		if err != nil {
			return
		}

		makes = append(makes, make)
	}

	return
}

// Database access funcs

func (r MakeMysqlRepo) Get(id int) (make Make, err error) {
	query := fmt.Sprintf("select * from make where id=%d", id)
	row := r.Db.QueryRow(query)
	make, err = r.Map(row)

	return
}

func (r MakeMysqlRepo) GetAll(limit int) (result []Make, err error) {
	query := fmt.Sprintf("select * from make limit %d", limit)
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	result, err = r.MapRows(rows)

	return
}

func (r MakeMysqlRepo) Query(query string) (makes []Make, err error) {
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	makes, err = r.MapRows(rows)

	return
}
