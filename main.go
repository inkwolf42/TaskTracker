package main

import (
	"os"
	"time"
)

type TaskStatus int

const (
	ToDo TaskStatus = iota
	InProgress
	Done
)

type task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type saveFile struct {
	Current_ID int    `json:"current_id"`
	Task_list  []task `json:"task_list"`
}

func (t *task) getTaskStatusString() string {
	switch t.Status {
	case ToDo:
		return "To Do"
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	default:
		panic("Error")
	}

}
func getTaskStatusFromString(status string) TaskStatus {
	switch status {
	case "todo":
		return ToDo
	case "in-progress":
		return InProgress
	case "done":
		return Done
	default:
		panic("Error")
	}

}

func main() {

	if len(os.Args) < 2 {
		panic("Command Is Requierd!!!")
	}

	data := Load()

	switch os.Args[1] {

	case "add":
		Add(&data)
	case "update":
		Update(&data)
	case "delete":
		Delete(&data)
	case "mark-done":
		Mark(&data, Done)
	case "mark-in-progress":
		Mark(&data, InProgress)
	case "mark-todo":
		Mark(&data, ToDo)
	case "list":
		List(&data)
	default:
		panic("Invalid command.")
	}

	Save(data)
}
