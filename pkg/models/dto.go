package models

import "time"

type Task struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Name  string    `json:"name"`
	Id    int64     `json:"id"`
}

type sortType string

const (
	ASC  sortType = "ASC"
	DESC sortType = "DESC"
)

type ReqTaskParams struct {
	Start time.Time `json:"start"`
	Name  string    `json:"name"`
	Limit int64     `json:"limit"`
}
