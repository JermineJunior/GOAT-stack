package utils

type Task struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	Completed bool   `json:"completed"`
}

func LinearSearch(id string, tasks []Task) int {

	for index, task := range tasks {
		if id == task.ID {
			return index
		}
	}
	return -1
}
