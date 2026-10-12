package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func PrintTaskHeader() {
	fmt.Printf("Id  |%-50s | %19s | %19s | %s\n", "Description", "Created At", "Updated At", "Status")
	fmt.Println(strings.Repeat("-", 5+52+22+22+7))
}

func PrintTask(t *task) {
	fmt.Printf("%-3d |%-50s | %s | %s | %s\n", t.Id, t.Description, t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"), t.getTaskStatusString())
}

func FindIndexOfTaskById(id int, tasks *[]task) (index int, taskRef *task, err error) {

	for i := 0; i < len(*tasks); i++ {
		t := &(*tasks)[i]
		if t.Id == int(id) {
			return i, t, nil
		}
	}
	return -1, nil, errors.New("Task Not Found.")
}

func TaskFilterBuilder() func(t *task) bool {
	if len(os.Args) == 2 {
		return func(t *task) bool { return true }
	}
	alt := func(t *task) bool { return false }
	for i := 2; i < len(os.Args); i++ {
		statusT := getTaskStatusFromString(os.Args[i])
		prev := alt
		alt = func(t *task) bool { return t.Status == statusT || prev(t) }
	}
	return alt
}

func Add(data *saveFile) {
	PanicIfArgsNumNotEqual(3)
	taskName := os.Args[2]
	t := task{
		Id:          data.Current_ID,
		Description: taskName,
		Status:      ToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	data.Task_list = append(data.Task_list, t)
	data.Current_ID++

	PrintTaskHeader()
	PrintTask(&t)
	fmt.Println("")
	fmt.Println("Task Has Been Added")
}

func Update(data *saveFile) {
	PanicIfArgsNumNotEqual(4)
	taskIdStr := os.Args[2]
	taskName := os.Args[3]

	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	PanicIfError(err)

	_, t, err := FindIndexOfTaskById(int(taskId), &data.Task_list)
	PanicIfError(err)

	t.Description = taskName

	PrintTaskHeader()
	PrintTask(t)
	fmt.Println("")
	fmt.Println("Task Has Been Updated!")
}

func Delete(data *saveFile) {
	PanicIfArgsNumNotEqual(3)
	taskIdStr := os.Args[2]

	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	PanicIfError(err)

	index, t, err := FindIndexOfTaskById(int(taskId), &data.Task_list)
	PanicIfError(err)

	taskTmp := *t
	data.Task_list = slices.Delete(data.Task_list, index, index+1)

	fmt.Println("DELETED : ")
	PrintTaskHeader()
	PrintTask(&taskTmp)
	fmt.Println("")
	fmt.Println("Task Has Been Updated!")
}

func Mark(data *saveFile, status TaskStatus) {

	PanicIfArgsNumNotEqual(3)
	taskIdStr := os.Args[2]

	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	PanicIfError(err)

	_, t, err := FindIndexOfTaskById(int(taskId), &data.Task_list)
	PanicIfError(err)

	t.Status = status

	PrintTaskHeader()
	PrintTask(t)
	fmt.Println("")
	fmt.Println("Task Has Been Updated!")
}

func List(data *saveFile) {
	if len(data.Task_list) == 0 {
		fmt.Println("No Tasks Found.")
		return
	}

	filter := TaskFilterBuilder()

	PrintTaskHeader()

	for _, t := range data.Task_list {
		if filter(&t) {
			PrintTask(&t)
		}
	}
}
