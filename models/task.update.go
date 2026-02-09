package models

import constants "github.com/Anh467/task-tracker/core/constants"

type TaskUpdate struct{
	Id int
	Description string
	Status constants.Status
}