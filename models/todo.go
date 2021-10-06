package models

type Todo struct {
	ID       string `json:"id" form:"id" example:"1"`
	Name     string `json:"name" form:"name" example:"Create POST user registration"`
	Complete bool   `json:"complete" form:"complete" example:"true"`
}
