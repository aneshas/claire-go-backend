package main

import (
	"fmt"
	"github.com/kisielk/sqlstruct"
	"log"
)

// TagMysqlRepo mysql repository
// Implements ITagRepository interface
type TagMysqlRepo struct {
	*MysqlDb
}

func (r TagMysqlRepo) Get(id int, joins []string) (tag Tag, err error) {
	query := fmt.Sprintf("select * from tag where id=%d", id)
	rows, err := r.Db.Query(query)
	if err != nil {
		return
	}

	if !rows.Next() {
		err = ErrNoResults
		return
	}

	sqlstruct.Scan(&tag, rows)
	err = rows.Err()

	rows.Close()

	return
}

func (r TagMysqlRepo) GetAll(limit int, joins []string) (result []Tag, err error) {
	if limit == 0 {
		limit = MAX_MYSQL_RESULTS
	}
	query := fmt.Sprintf("select * from tag limit %d", limit)
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	tag := Tag{}
	for rows.Next() {
		err = sqlstruct.Scan(&tag, rows)
		result = append(result, tag)
	}
	rows.Close()

	return
}

func (r TagMysqlRepo) Query(query string) (tags []Tag, err error) {
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	tag := Tag{}
	for rows.Next() {
		err = sqlstruct.Scan(&tag, rows)
		tags = append(tags, tag)
	}
	rows.Close()

	return
}
