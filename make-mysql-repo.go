package main

import (
	"fmt"
	"github.com/kisielk/sqlstruct"
	"log"
)

// MakeMysqlRepo mysql repository
// Implements IMakeRepository interface
type MakeMysqlRepo struct {
	*MysqlDb
}

func (r MakeMysqlRepo) Get(id int, joins []string) (make Make, err error) {
	entities := "make"

	query := fmt.Sprintf("select * from %s where id=%d", entities, id)
	rows, err := r.Db.Query(query)
	if err != nil {
		return
	}

	if !rows.Next() {
		err = ErrNoResults
		return
	}
	sqlstruct.Scan(&make, rows)

	rows.Close()

	return
}

func (r MakeMysqlRepo) GetAll(limit int, joins []string) (result []Make, err error) {
	if limit == 0 {
		limit = MAX_MYSQL_RESULTS
	}

	query := fmt.Sprintf("select * from make limit %d", limit)
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	make := Make{}
	for rows.Next() {
		err = sqlstruct.Scan(&make, rows)
		result = append(result, make)
	}
	rows.Close()

	return
}

func (r MakeMysqlRepo) Query(query string) (makes []Make, err error) {
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	make := Make{}
	for rows.Next() {
		err = sqlstruct.Scan(&make, rows)
		makes = append(makes, make)
	}
	rows.Close()

	return
}
