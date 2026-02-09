package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Anh467/task-tracker/core/constants"
	models "github.com/Anh467/task-tracker/models"
	services "github.com/Anh467/task-tracker/services"
)

func main() {
	taskStorage := services.TaskStorage{Tasks: []models.Task{}}
	var command string
	var data1 string
	var data2 string
	var task models.Task
	var tasks []models.Task
	var err error = nil
	var id int

	if(len(os.Args) == 1){
		fmt.Println("Data have to use command task-cli <command> {add, update, delete, mark-in-progress, mark-done, list}")
		return
	}
	command = os.Args[1]

	if(len(os.Args) >= 3){
		data1 = os.Args[2]
	}
	if(len(os.Args) >= 4){
		data2 = os.Args[3]
	}

	taskStorage.ReadTasks()
	switch(command){
	case "add":
		if(data1 == ""){
			fmt.Println("Add command require data")
			return
		}

		task, err =taskStorage.Create(models.TaskCreate{Description: data1, Status: constants.STATUS_TODO})
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}else{
			fmt.Printf("Success: Create successful %s\n", task.ToString())
		}
	case "update":
		if(data1 == "" || data2 == ""){
			fmt.Println("Update command require 2 data")
			return
		}
		id, err = strconv.Atoi(data1)
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}

		task, err := taskStorage.Update(models.TaskUpdate{Id: id, Description: data2})
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}else{
			fmt.Printf("Success: Update successful %s\n", task.ToString())
		}
	case "mark-in-progress":
		if(data1 == ""){
			fmt.Println("mark-in-progress command require data")
			return
		}

		id, err = strconv.Atoi(data1)
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}

		task, err = taskStorage.Update(models.TaskUpdate{Id: id, Description:  "", Status: constants.STATUS_INPROGRESS})
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}else{
			fmt.Printf("Success: Update successful %s\n", task.ToString())
		}
	case "mark-done":
		if(data1 == ""){
			fmt.Println("mark-done command require data")
			return
		}

		id, err = strconv.Atoi(data1)
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}

		task, err = taskStorage.Update(models.TaskUpdate{Id: id, Description:  "", Status: constants.STATUS_DONE})
		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}else{
			fmt.Printf("Success: Update successful %s\n", task.ToString())
		}
	case "list":
		tasks, err = taskStorage.GetAllElement(constants.Status(data1))

		if(err != nil){
			fmt.Printf("Fail: %s\n", err)
		}else{
			fmt.Printf("Success: Get all tasks successful %s\n", models.TasksToString(tasks))
		}

	}
	taskStorage.SaveTasks()
}