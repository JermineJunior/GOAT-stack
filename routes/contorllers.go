package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GOAT-stack-todo-app/components"
	"github.com/GOAT-stack-todo-app/utils"
)

var tasks []utils.Task

// landing page
func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("homepage visitor")

	// rendering all tasks
	home := components.HomePage(tasks)
	home.Render(r.Context(), w)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getting all tasks")

	if len(tasks) == 0 {
		fmt.Println("no tasks to show")
		w.Write([]byte("no tasks to show"))
		return
	}

	// rendering all tasks
	component := components.AllTasks(tasks)
	component.Render(r.Context(), w)

}
func AddTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adding a task to the list")

	input := r.FormValue("task")

	if input == "" {
		fmt.Println("form is empty")
		w.Write([]byte("form is empty"))
		return
	}

	// creating a new task and added to tasks
	task := utils.Task{strconv.Itoa(len(tasks) + 1), input, false} //? using the len of arr + 1 as id
	tasks = append(tasks, task)

	fmt.Println("tasks are: ", tasks)

	singleTask := components.SingleTask(task)
	singleTask.Render(r.Context(), w)

}
func FinishByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adding a task to the list")

	input := r.FormValue("task")

	if input == "" {
		fmt.Println("form is empty")
		w.Write([]byte("form is empty"))
		return
	}

	// creating a new task and added to tasks
	task := utils.Task{input, strconv.Itoa(len(tasks) + 1), false} //? using the len of arr + 1 as id
	tasks = append(tasks, task)

	fmt.Println("task was added.")

}

func DeleteByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adding a task to the list")

	input := r.FormValue("task")

	if input == "" {
		fmt.Println("form is empty")
		w.Write([]byte("form is empty"))
		return
	}

	// creating a new task and added to tasks
	task := utils.Task{input, strconv.Itoa(len(tasks) + 1), false} //? using the len of arr + 1 as id
	tasks = append(tasks, task)

	fmt.Println("task was added.")

}
