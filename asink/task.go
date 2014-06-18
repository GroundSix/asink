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
	for _, task := range tasks {
		command := task.Command
		
		if detectRequiredTask(task) == true {
			executeRequiredTask(task)
		}
		command.Execute()
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
		command.Execute()
		delete(tasks, task.Require)
	}
}
