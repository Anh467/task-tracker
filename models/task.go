package models

import (
	"fmt"
	"strings"

	constants "github.com/Anh467/task-tracker/core/constants"
)
type TaskList []Task


type Task struct {
	Id          int
	Description string
	Status      constants.Status
	CreatedAt   string
	UpdatedAt    string
}

func (task *Task) ToString() string{
	return fmt.Sprintf("Id: %d, Description: %s, Status: %s, Updated At: %s, Created At: %s", task.Id, task.Description, task.Status, task.UpdatedAt, task.CreatedAt)
}

func TasksToString(tasks []Task) string{
	var stringBuilder strings.Builder
	for _, ele := range tasks{
		stringBuilder.WriteString(ele.ToString())
		stringBuilder.WriteString("\n")
	}

	return stringBuilder.String()
}