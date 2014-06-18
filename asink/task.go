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
	Name    string
	Command *Command
	Require string
	Group   string
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

func (t *Task) AddTask(name string, command *Command, require string, group string) *Task {
	task := new(Task)

	task.Name    = name
	task.Command = command
	task.Require = require
	task.Group   = group

	tasks[name] = task

	return new(Task)
}

func (t *Task) Execute() {
	for name, task := range tasks {
		command := task.Command

		// check for require and groups ect...
		command.Execute()
	}
}
