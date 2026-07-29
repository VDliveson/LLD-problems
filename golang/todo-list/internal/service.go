package internal

import "fmt"

type TaskManager struct {
	userTasks map[int]map[int]Task
	cnt       int
}

func New() *TaskManager {
	return &TaskManager{
		userTasks: make(map[int]map[int]Task),
	}
}

func (t *TaskManager) AddTask(userId int, name string) int {
	t.cnt += 1
	if _, ok := t.userTasks[userId]; !ok {
		t.userTasks[userId] = make(map[int]Task)
	}
	t.userTasks[userId][t.cnt] = Task{
		id:     t.cnt,
		name:   name,
		status: false,
	}
	return t.userTasks[userId][t.cnt].id
}

func (t *TaskManager) MarkCompleted(userId int, taskId int) {
	if tasks, found := t.userTasks[userId]; found {
		if task, foundTask := tasks[taskId]; foundTask {
			task.status = true
			tasks[taskId] = task
		}
	}
}

func (t *TaskManager) DeleteTask(userId int, taskId int) {
	if _, found := t.userTasks[userId]; found {
		delete(t.userTasks[userId], taskId)
	}
}

func (t *TaskManager) ListTasks(userId int) {
	if _, found := t.userTasks[userId]; found {
		fmt.Printf("User id %d tasks\n", userId)
		for _, task := range t.userTasks[userId] {
			fmt.Printf("%s, status -> %v \n", task.name, task.status)
		}
		fmt.Println()
	}
}
