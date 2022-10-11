package models

import "time"

type Task struct {
	Start time.Time
	End   time.Time
	Name  string
	Id    int64
}

type ReqTaskParams struct {
}
