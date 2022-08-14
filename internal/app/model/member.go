package model

type Member struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name" validate:"max=255"`
	Phone string `json:"phone" db:"phone" validate:"max=20"`
}
