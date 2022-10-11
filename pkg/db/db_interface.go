package db

import (
	"simpleTimeTracker/pkg/models"
)

type Db interface {
	SaveTask(task models.Task) (int64, error)

	SearchTasks(params models.ReqTaskParams) ([]models.Task, error)
	SelectTask(id int64) (models.Task, error)
}
