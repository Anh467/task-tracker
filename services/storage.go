package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	constants "github.com/Anh467/task-tracker/core/constants"
	models "github.com/Anh467/task-tracker/models"
)


type TaskStorage struct{
	index int
	Tasks []models.Task
}

func (v *TaskStorage) Create(taskCreate models.TaskCreate) (task models.Task, err error){
	task = models.Task{
		Id: v.index + 1,
		Description: taskCreate.Description,
		Status: taskCreate.Status,
		UpdatedAt: time.Now().String(),
		CreatedAt: time.Now().String(),
	}

	v.Tasks = append(v.Tasks, task)
	v.index = v.index + 1

	return
} 

func (v *TaskStorage) Delete(id int) (task models.Task, err error){
	index := constants.EMPTY_INDEX
	err = nil

	for i, ele := range v.Tasks{
		if(ele.Id == id){
			index = i
			task = ele
			break
		}
	}

	if(index == constants.EMPTY_INDEX){
		err = fmt.Errorf(constants.MESSAGE_NOT_FOUND, id)
		return
	}

	v.Tasks = append(v.Tasks[:index], v.Tasks[index+1:]...)

	return
}

func (v *TaskStorage) Update(taskUpdate models.TaskUpdate) (task models.Task, err error){
	index := constants.EMPTY_INDEX
	err = nil

	for i, ele := range v.Tasks {
		if(ele.Id == taskUpdate.Id){
			index = i
			task = ele
			break
		}
	}

	if(index == constants.EMPTY_INDEX){
		err = fmt.Errorf(constants.MESSAGE_NOT_FOUND, taskUpdate.Id)
		return
	}

	if( taskUpdate.Description != ""){
		v.Tasks[index].Description = taskUpdate.Description
	}
	v.Tasks[index].Status = taskUpdate.Status
	v.Tasks[index].UpdatedAt = time.Now().String()

	task = v.Tasks[index]

	return
}

func (v *TaskStorage) GetById(id int) (task models.Task, err error){
	err = nil
	
	for _,ele := range v.Tasks{
		if (ele.Id == id){
			task = ele
			return 
		}
	}
	
	err = fmt.Errorf(constants.MESSAGE_NOT_FOUND, id)

	return
}

func (v *TaskStorage) GetAllElement(status constants.Status) ([]models.Task, error){
	var tasks []models.Task
	var err error = nil

	if len(v.Tasks) == 0{
		return v.Tasks, fmt.Errorf(constants.MESSAGE_LIST_EMPTY)
	}

	if(status != ""){
	 	for _, ele := range v.Tasks {
			if(ele.Status == status){
				tasks = append(tasks, ele)
			}
		}
	}else{
		tasks  = v.Tasks
	}

	return tasks, err
}

func (v *TaskStorage) SaveTasks() error{
	createPath()
	if(len(v.Tasks) == 0){
		return fmt.Errorf(constants.MESSAGE_NO_DATA_IS_SAVE)
	}

	data, err := json.Marshal(v.Tasks)

	if err != nil{
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err 
	}
	dataPath := dir + "\\" + constants.FILE_TASKS
	err = os.WriteFile(dataPath, data, 0644)

	return err
}

func (v *TaskStorage) ReadTasks() (err error) {
	createPath()
	var data []byte
	var index int
	dir, err := os.Getwd()
	if err != nil {
		return err 
	}
	dataPath := dir + "//" + constants.FILE_TASKS
	data, err = os.ReadFile(dataPath)

	if err != nil {
		return err 
	}

	var tasks []models.Task 

	if len(data) == 0{
		v.Tasks = []models.Task {};
		return nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil{
		return
	}
	if(len(tasks)== 0){
		index = 0
	}else{
		index = tasks[len(tasks) - 1].Id
	}

	v.Tasks = tasks
	v.index = index
	return nil
}

func createPath() (string, error) {
	wd, err := os.Getwd()
    if err != nil {
        return "", err
    }

    dir := filepath.Join(wd, ".task-cli")
    if err := os.MkdirAll(dir, 0755); err != nil {
        return "", err
    }

    file := filepath.Join(dir, "tasks.json")
    if _, err := os.Stat(file); os.IsNotExist(err) {
        if err := os.WriteFile(file, []byte("[]"), 0644); err != nil {
            return "", err
        }
    }

    return file, nil
}