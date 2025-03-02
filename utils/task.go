package utils

type Task struct {
	ID        string
	Value     string
	Completed bool
}

func LinearSearch(id string, tasks []Task) int {

	for index, task := range tasks {
		if id == task.ID {
			return index
		}
	}
	return -1
}
