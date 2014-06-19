/**
 * asink v0.0.2-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.0.2-dev
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package asink

import (
	"sync"
)

type Task struct {
	Name     string
	Command  *Command
	Require  string
	Group    string
}

var tasks map[string]*Task = nil

/**
 * Creates a new instance of the Task
 * struct
 *
 * @return *Task a new task
 */
func NewTask() *Task {
	tasks = make(map[string]*Task)
	return new(Task)
}

func (t *Task) AddTask(name string, command *Command, require string, group string) {
	task := new(Task)

	task.Name     = name
	task.Command  = command
	task.Require  = require
	task.Group    = group

	tasks[name] = task
}

func (t *Task) Execute() {
	for name, task := range tasks {
		command := task.Command

		if detectRequiredTask(task) == true {
			executeRequiredTask(task)
		}

		if detectGroupedTasks(task) == true {
			executeGroupedTasks(task)
		} else {
			command.Execute()
			delete(tasks, name)
		}
	}
}

func detectRequiredTask(task *Task) bool {
	if (task.Require != "") {
		return true
	}
	return false
}

func executeRequiredTask(task *Task) {
	required_task := tasks[task.Require]
	if (required_task != nil) {
		command := required_task.Command
		
		if detectRequiredTask(required_task) == true {
			executeRequiredTask(required_task)
		}

		if detectGroupedTasks(required_task) == true {
			executeGroupedTasks(required_task)
		} else {
			command.Execute()
			delete(tasks, task.Require)
		}
	}
}

func detectGroupedTasks(task *Task) bool {
	if (task.Group != "") {
		return true
	}
	return false
}

func executeGroupedTasks(task *Task) {
	group := task.Group
	var wg sync.WaitGroup
	for _, block := range tasks {
		if block.Group == group {
			wg.Add(1)
			go executeGroupConcurrently(block, &wg)
		}
	}
	wg.Wait()
}

func executeGroupConcurrently(task *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	command := task.Command
	command.Execute()
	delete(tasks, task.Name)
}
