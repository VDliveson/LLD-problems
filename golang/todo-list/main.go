package main

import "todo-list/internal"

func main() {
	taskManager := internal.New()
	taskManager.AddTask(1, "task1")
	taskManager.AddTask(1, "task2")
	taskManager.ListTasks(1)
	taskManager.AddTask(2, "task3")
	taskManager.MarkCompleted(1, 2)
	taskManager.ListTasks(1)
	taskManager.DeleteTask(1, 2)
	taskManager.ListTasks(1)
}
