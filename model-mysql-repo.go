package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ModelMysqlRepo mysql repository
// Implements IModelRepository interface
type ModelMysqlRepo struct {
	*MysqlDb
}

// Helper funcs

func (r ModelMysqlRepo) Map(row *sql.Row) (model Model, err error) {
	err = row.Scan(
		&model.Id,
		&model.Name,
		&model.Description,
		&model.Photo,
		&model.MakeId)

	return
}

func (r ModelMysqlRepo) MapRows(rows *sql.Rows) (models []Model, err error) {
	model := Model{}

	for rows.Next() {
		err = rows.Scan(
			&model.Id,
			&model.Name,
			&model.Description,
			&model.Photo,
			&model.MakeId)

		if err != nil {
			return
		}

		models = append(models, model)
	}

	return
}

// Database access funcs

func (r ModelMysqlRepo) Get(id int) (model Model, err error) {
	query := fmt.Sprintf("select * from model where id=%d", id)
	row := r.Db.QueryRow(query)
	model, err = r.Map(row)

	return
}

func (r ModelMysqlRepo) GetAll(limit int) (result []Model, err error) {
	query := fmt.Sprintf("select * from model limit %d", limit)
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	result, err = r.MapRows(rows)

	return
}

func (r ModelMysqlRepo) Query(query string) (models []Model, err error) {
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	models, err = r.MapRows(rows)

	return
}
