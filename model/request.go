package model

type UpdateRequest struct {
	ID     int
	User   string `json:"user" binding:"required"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	URL    string `json:"url"`
}

type AddRequest struct {
	User   string `json:"user" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	URL    string `json:"url"`
}

type GetRequest struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
}

type DeleteRequest struct {
	User string `json:"user" binding:"required"`
}
