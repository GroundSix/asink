package asink

import (
	"sync"
)

type Execer interface {
	Exec()
}

var TasksMap map[string]Task = nil

type Task struct {
	Name 	string
	Process Execer
	Require string
	Group   string
}

func NewTask(name string, process Execer) Task {
	t := Task{}
	t.Name 	  = name
	t.Process = process
	return t
}

func (t Task) Exec() {
	p := t.Process

	executeRequiredTask(t)

	if executeGroupedTasks(t) == true {
	} else {
		if (p != nil) {
			p.Exec()
			delete(TasksMap, t.Name)
		}
	}
}

func ExecMulti(taskSlice []Task) {
	TasksMap = createTasksMap(taskSlice)
	for _, t := range TasksMap {
		t.Exec()
	}
}

func createTasksMap(tasks []Task) map[string]Task {
	tasksMap := make(map[string]Task)
	for _, task := range tasks {
		tasksMap[task.Name] = task
	}
	return tasksMap
}

// If a required task is found it
// is ran
func executeRequiredTask(t Task) {
	if (t.Require != "") {
		task := TasksMap[t.Require]
		task.Exec()
	}
}

// If a grouped task is found, it is
// ran
func executeGroupedTasks(task Task) bool {
	if (task.Group != "") {
		group := task.Group
		var wg sync.WaitGroup
		for _, block := range TasksMap {
			if block.Group == group {
				wg.Add(1)
				go executeGroupConcurrently(block, &wg)
			}
		}
		wg.Wait()
		return true
	}
	return false
}

// Allows tasks to be ran without
// any blocking
func executeGroupConcurrently(t Task, wg *sync.WaitGroup) {
	defer wg.Done()
	process := t.Process
	process.Exec()
	delete(TasksMap, t.Name)
}
