package repository

import (
	"beerstore/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

//go:generate mockery --with-expecter --name "IRepository" --output $PWD/mocks
type IRepository interface {
	Find(req model.GetRequest) ([]model.Beer, error)
	Add(fields []string, values []interface{}) (int, error)
	Update(id int, fields []string) error
	Delete(id int) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Find(req model.GetRequest) ([]model.Beer, error) {

	// frist pagination
	offset := (req.Page - 1) * req.PageSize

	// create SQL query
	query := "SELECT id, name, type ,detail, url FROM beer"

	// name checker
	if req.Name != "" {
		query += " WHERE name = ?"
	}

	// pagination condition
	query += " LIMIT ? OFFSET ?"
	fmt.Println(query)

	// query
	var rows *sql.Rows
	var err error
	if req.Name != "" {
		rows, err = r.db.Query(query, req.Name, req.PageSize, offset)
	} else {
		rows, err = r.db.Query(query, req.PageSize, offset)
	}
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	beers := []model.Beer{}

	// loop from query
	for rows.Next() {
		var beer model.Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Detail, &beer.URL)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		beers = append(beers, beer)
	}
	return beers, nil
}

func (r *Repository) Add(fields []string, values []interface{}) (int, error) {

	// create SQL query
	query := "INSERT INTO beer ("

	// create statement
	query += strings.Join(fields, ", ") + ") VALUES (?" + strings.Repeat(", ?", len(fields)-1) + ")"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	defer stmt.Close()

	// create data in MariaDB
	res, err := stmt.Exec(values...)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	// get the ID of the newly inserted beer
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return int(id), nil
}

func (r *Repository) Update(id int, setFields []string) error {

	// create SQL query
	query := "UPDATE beer SET "

	// create statement
	query += fmt.Sprintf("%s WHERE id = ?", strings.Join(setFields, ", "))
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	// update data
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result, err := res.RowsAffected()
	if result == 0 {
		return errors.New("have no data to change")
	}

	return nil
}

func (r *Repository) Delete(id int) error {

	// create SQL query
	query := "DELETE FROM beer WHERE id = ?"

	// create statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	// delete
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result, err := res.RowsAffected()
	if result == 0 {
		return errors.New("have no data to delete")
	}
	return nil
}
