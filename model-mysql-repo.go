package main

import (
	"fmt"
	"github.com/kisielk/sqlstruct"
	"log"
)

// ModelMysqlRepo mysql repository
// Implements IModelRepository interface
type ModelMysqlRepo struct {
	*MysqlDb
}

func (r ModelMysqlRepo) Get(id int, joins []string) (model Model, err error) {
	query := fmt.Sprintf("select * from model where id=%d", id)
	rows, err := r.Db.Query(query)
	if err != nil {
		return
	}

	if !rows.Next() {
		err = ErrNoResults
		return
	}

	sqlstruct.Scan(&model, rows)
	err = rows.Err()

	rows.Close()

	return
}

func (r ModelMysqlRepo) GetAll(limit int, joins []string) (result []Model, err error) {
	if limit == 0 {
		limit = MAX_MYSQL_RESULTS
	}

	query := fmt.Sprintf("select * from model limit %d", limit)
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	model := Model{}
	for rows.Next() {
		err = sqlstruct.Scan(&model, rows)
		result = append(result, model)
	}
	rows.Close()

	return
}

func (r ModelMysqlRepo) Query(query string) (models []Model, err error) {
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	model := Model{}
	for rows.Next() {
		err = sqlstruct.Scan(&model, rows)
		models = append(models, model)
	}
	rows.Close()

	return
}
