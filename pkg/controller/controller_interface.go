package controller

import (
	"simpleTimeTracker/pkg/models"
)

type App interface {
	StartTask(name string) error
	StopTask() (int64, error)

	SelectTask(id int64) (models.Task, error)
	SearchTasks(params models.ReqTaskParams) ([]models.Task, error)
	ShutDown() error
}
