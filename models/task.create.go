package models

import constants "github.com/Anh467/task-tracker/core/constants"

type TaskCreate struct{
	Description string
	Status constants.Status
}