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
	Command *Command
	Require string
	Group   string
}

var tasks []*Task = nil

/**
 * Creates a new instance of the Task
 * struct
 *
 * @return *Task a new task
 */
func NewTask() *Task {
	return new(Task)
}

func (t *Task) AddTask(command *Command, require string, group string) *Task {
	task := new(Task)

	task.Command = command
	task.Require = require
	task.Group   = group

	tasks = append(tasks, task)

	return new(Task)
}

func (t *Task) Execute() {
	for _,v := range tasks {
		command := v.Command
		command.Execute()
	}
}
