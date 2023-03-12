package model

type SystemConf struct {
	ID    string `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
}
