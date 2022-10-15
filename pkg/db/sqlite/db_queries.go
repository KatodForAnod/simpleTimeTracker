package sqlite

import (
	"log"
	"simpleTimeTracker/pkg/models"
)

const saveTaskQuery = `
	INSERT into notes (name, start, end) VALUES ($1, $2, $3)
`

func (l *SqlLite) SaveTask(task models.Task) (int64, error) {
	result, err := l.conn.Exec(saveTaskQuery, task.Name, task.Start, task.End)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id, err
}

const searchTaskQuery = ``

func (l *SqlLite) SearchTasks(params models.ReqTaskParams) ([]models.Task, error) {
	panic("")
}

const selectTaskQuery = ``

func (l *SqlLite) SelectTask(id int64) (models.Task, error) {
	panic("")
}
