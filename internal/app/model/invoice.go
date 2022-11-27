package model

import "strings"

type InvoiceResp struct {
	Data string `json:"data" db:"data"`
}

type Invoice struct {
	HeadFak     string `json:"head_fak" db:"head_fak"`
	LastCounter int    `json:"last_counter" db:"last_counter"`
}

var TranslateMonthHeadFak = strings.NewReplacer(
	"January", "A",
	"February", "B",
	"March", "C",
	"April", "D",
	"May", "E",
	"June", "F",
	"July", "G",
	"August", "H",
	"September", "I",
	"October", "J",
	"November", "K",
	"December", "L",
)
