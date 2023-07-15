package model

import "database/sql"

type Beer struct {
	ID     int
	Name   string
	Type   sql.NullString
	Detail sql.NullString
	URL    sql.NullString
}

type BeerResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	URL    string `json:"url"`
}
