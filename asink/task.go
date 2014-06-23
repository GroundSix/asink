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

/**
 * @var String task name
 * @var *Command initial command
 * @var String required task name
 * @var String task group name
 */
type Task struct {
	Name     string
	Command  *Command
	Require  string
	Group    string
	Remote   string
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

/**
 * Adds a new tasks to the map
 *
 * @param String name of task
 * @param *Command the initial command
 * @param String the required task to run
 * @param String the group of tasks to be ran
 *
 * @return nil
 */
func (t *Task) AddTask(name string, command *Command, require string, group string) {
	task := new(Task)

	task.Name     = name
	task.Command  = command
	task.Require  = require
	task.Group    = group

	tasks[name] = task
}

/**
 * Runs all tasks, required and grouped
 *
 * @return nil
 */
func (t *Task) Execute() {
	for name, task := range tasks {
		runTasks(task, name)
	}
}

/**
 * Checks for groups and required
 * tasks
 *
 * @param *Task
 * @param String name of task
 *
 * @return nil
 */
func runTasks(task *Task, task_name string) {
	command := task.Command
	if detectRequiredTask(task) == true {
		executeRequiredTask(task)
	}

	if detectGroupedTasks(task) == true {
		executeGroupedTasks(task)
	} else {
		command.Execute()
		delete(tasks, task_name)
	}
}

/**
 * Checks to see if there is a required
 * task before running it's parent
 *
 * @param *Task
 *
 * @return Bool
 */
func detectRequiredTask(task *Task) bool {
	if (task.Require != "") {
		return true
	}
	return false
}

/**
 * If a required task is found it
 * is ran
 *
 * @param *Task
 *
 * @return nil
 */
func executeRequiredTask(task *Task) {
	required_task := tasks[task.Require]
	if (required_task != nil) {
		runTasks(required_task, task.Require)
	}
}

/**
 * Checks to see if there is a grouped
 * task so they can be ran concurrently
 *
 * @param *Task
 *
 * @return Bool
 */
func detectGroupedTasks(task *Task) bool {
	if (task.Group != "") {
		return true
	}
	return false
}

/**
 * If a grouped task is found, it is
 * ran
 *
 * @param *Task
 *
 * @return nil
 */
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

/**
 * Allows tasks to be ran without
 * any blocking
 *
 * @param *Task
 * @param *sync.WaitGroup
 */
func executeGroupConcurrently(task *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	command := task.Command
	command.Execute()
	delete(tasks, task.Name)
}
