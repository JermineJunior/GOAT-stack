package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GOAT-stack-todo-app/components"
	"github.com/GOAT-stack-todo-app/storage"
	"github.com/GOAT-stack-todo-app/utils"
)

var tasks []utils.Task
var db = storage.NewStorage[[]utils.Task]("data.json")

// landing page
func HomePage(w http.ResponseWriter, r *http.Request) {
	_ = db.Load(&tasks)
	fmt.Println("homepage visitor")

	// rendering all tasks
	home := components.HomePage(tasks)
	_ = home.Render(r.Context(), w)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getting all tasks")

	if len(tasks) == 0 {
		fmt.Println("no tasks to show")
		_, _ = w.Write([]byte("no tasks to show"))
		return
	}

	// rendering all tasks
	component := components.AllTasks(tasks)
	_ = component.Render(r.Context(), w)

}
func AddTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adding a task to the list")

	input := r.FormValue("task")

	if input == "" {
		fmt.Println("form is empty")
		_, _ = w.Write([]byte("form is empty"))
		return
	}

	// creating a new task and added to tasks
	task := utils.Task{strconv.Itoa(len(tasks) + 1), input, false} //? using the len of arr + 1 as id
	tasks = append(tasks, task)
	_ = db.Save(tasks)
	fmt.Println("tasks are: ", tasks)

	singleTask := components.SingleTask(task)
	_ = singleTask.Render(r.Context(), w)

}
func FinishByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("completing  a task from the list.")

	id := r.PathValue("id")
	index := utils.LinearSearch(id, tasks)

	if index == -1 {
		fmt.Println("id not found")
		_, _ = w.Write([]byte("id not found"))
		return
	}

	tasks[index].Completed = true
	fmt.Println("task was added.")
	_ = db.Save(tasks)
	w.Header().Set("HX-Trigger", "tasksChanged")
	// Re-render the complete tasks list to show updated status
	tasksComponent := components.AllTasks(tasks)
	tasksComponent.Render(r.Context(), w)

}
func DeleteByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("deleting  a task from the list.")

	id := r.PathValue("id")
	index := utils.LinearSearch(id, tasks)

	if index == -1 {
		fmt.Println("id not found")
		w.Write([]byte("id not found"))
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("task was deleted.")
	_ = db.Save(tasks)
	w.Header().Set("HX-Trigger", "tasksChanged")
	home := components.AllTasks(tasks)
	_ = home.Render(r.Context(), w)

}
