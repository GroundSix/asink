// asink v0.0.3-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.0.3-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package asink

import (
    "sync"
)

type Task struct {
	Name     string
	Command  *Command
	Require  string
	Group    string
	Remote   string
	Tag      string
}

var tasks map[string]*Task = nil

// Creates a new instance of the Task
// struct
func NewTask() *Task {
    tasks = make(map[string]*Task)
    return new(Task)
}

// Adds a new tasks to the map
func (t *Task) AddTask(name string, command *Command, require string, group string) {
    task := new(Task)
    
	task.Name     = name
	task.Command  = command
	task.Require  = require
	task.Group    = group
	task.Tag      = ""

    tasks[name] = task
}

// Sets a new remote from the ssh block
// in the JSON config
func (t *Task) SetRemote(name string, remote string) {
    task := tasks[name]
    task.Remote = remote
    if remote != "" {
        command := task.Command
        command.Manual = true
    }
}

// Sets a custom tag to a task which can be
// used to identify a particular group of tasks
// before execution
func (t *Task) SetTag(name string, tag string) {
	task := tasks[name]
	task.Tag = tag
}

// Runs all tasks, required and grouped
func (t *Task) Execute() {
    for name, task := range tasks {
        runTasks(task, name)
    }
}

// Checks for groups and required
// tasks
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

// Checks to see if there is a required
// task before running it's parent
func detectRequiredTask(task *Task) bool {
    if (task.Require != "") {
        return true
    }
    return false
}

// If a required task is found it
// is ran
func executeRequiredTask(task *Task) {
    required_task := tasks[task.Require]
    if (required_task != nil) {
        runTasks(required_task, task.Require)
    }
}

// Checks to see if there is a grouped
// task so they can be ran concurrently
func detectGroupedTasks(task *Task) bool {
    if (task.Group != "") {
        return true
    }
    return false
}

// If a grouped task is found, it is
// ran
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

// Allows tasks to be ran without
// any blocking
func executeGroupConcurrently(task *Task, wg *sync.WaitGroup) {
    defer wg.Done()
    command := task.Command
    command.Execute()
    delete(tasks, task.Name)
}
