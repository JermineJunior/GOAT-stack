package routes

import (
	"fmt"
	"net/http"

	"github.com/GOAT-stack-todo-app/components"
)

func Tasks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("inside tasks")

	tasks := components.Task()
	tasks.Render(r.Context(), w)

}
