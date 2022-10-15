package sqlite

import (
	"log"
	"simpleTimeTracker/pkg/models"
)

const saveTaskQuery = `
	INSERT into notes 
	(name, start, end) 
	VALUES ($1, $2, $3)
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

const searchTaskQuery = `
	SELECT (id, name, start, end)
	FROM notes WHERE start >= $1
	ORDER BY start
`

func (l *SqlLite) SearchTasks(params models.ReqTaskParams) ([]models.Task, error) {
	rows, err := l.conn.Query(searchTaskQuery, params.Start)
	if err != nil {
		log.Println(err)
		return []models.Task{}, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(task.Id, task.Name, task.Start, task.End)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return tasks, nil
}

const selectTaskQuery = `
	SELECT (id, name, start, end)
	FROM notes 
	WHERE id = $1
`

func (l *SqlLite) SelectTask(id int64) (models.Task, error) {
	rows, err := l.conn.Query(selectTaskQuery, id)
	if err != nil {
		log.Println(err)
		return models.Task{}, err
	}

	var task models.Task
	for rows.Next() {
		err := rows.Scan(task.Id, task.Name, task.Start, task.End)
		if err != nil {
			log.Println(err)
			return models.Task{}, err
		}
	}

	return task, nil
}
