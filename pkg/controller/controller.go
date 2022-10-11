package controller

import (
	"fmt"
	"pkg/db"
	"pkg/models"
	"time"
)

type Controller struct {
	currTask *models.Task
	db       db.Db
}

func InitController(db db.Db) Controller {
	return Controller{db: db}
}

func (c *Controller) StartTask(name string) error {
	if c.currTask != nil {
		return fmt.Errorf("stop the previus task: %s", c.currTask.Name)
	}

	newTask := models.Task{
		Start: time.Now(),
		End:   time.Time{},
		Name:  name,
	}
	c.currTask = &newTask

	return nil
}

func (c *Controller) StopTask() (int64, error) {
	if c.currTask == nil {
		return -1, fmt.Errorf("no runnig task")
	}

	c.currTask.End = time.Now()
	id, err := c.db.SaveTask(*c.currTask)
	if err != nil {
		return -1, fmt.Errorf("controller, StopTask err: %v", err)
	}

	c.currTask = nil
	return id, nil
}

func (c *Controller) SelectTask(id int64) (models.Task, error) {
	task, err := c.db.SelectTask(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("controller, SelectTask err: %v", err)
	}
	return task, nil
}

func (c *Controller) SearchTasks(params models.ReqTaskParams) ([]models.Task, error) {
	tasks, err := c.db.SearchTasks(params)
	if err != nil {
		return []models.Task{}, fmt.Errorf("controller, SearchTasks err: %v", err)
	}
	return tasks, nil
}
