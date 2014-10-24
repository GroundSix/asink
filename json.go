// asink v0.1.1-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"./asink"
	"github.com/asink/typed"
)

type Json struct {
	taskMap typed.Typed
	tasks 	[]asink.Task
}

// Returns a slice of all asink tasks
func (j *Json) Tasks() []asink.Task {
	return j.tasks
}

// Parses the JSON into a typed.Typed object
// which acts as map[string]interface{}
func (j *Json) parse(body []byte) Parser {
	mapped, err := typed.Json(body)
	if (err != nil) {
		panic(err)
	}
	j.taskMap = mapped
	return j
}

// Creates and assigns tasks using the map
// from the Json struct
func (j *Json) assignTasks() Parser {
	t := j.taskMap.StringObject("tasks")
	tasks := []asink.Task{}
	for name, task := range t {
		c := asink.NewCommand(task["command"].(string))
		j.buildCommand(&c, task)
		tasks = append(tasks, j.buildTask(name, task, &c))
	}
	j.tasks = tasks
	return j
}

// Builds up the asink command using the parsed
// JSON data
func (j *Json) buildCommand(c *asink.Command, t typed.Typed) {
	j.setAsyncCount(c, t)
	j.setRelCount(c, t)
	j.setDir(c, t)
	j.setArgs(c, t)
}

// Build up the asink task using the parsed
// JSON data
func (j *Json) buildTask(name string, task typed.Typed, c *asink.Command) asink.Task {
	t := asink.NewTask(name, c)
	j.setRequire(&t, task)
	j.setGroup(&t, task)
	return t
}

// Sets the AsyncCount for the command object
func (j *Json) setAsyncCount(c *asink.Command, t typed.Typed) {
	c.AsyncCount = t.IntsOr("count", []int{1, 1})[0]
}

// Sets the RelCount for the command object
func (j *Json) setRelCount(c *asink.Command, t typed.Typed) {
	c.RelCount = t.IntsOr("count", []int{1, 1})[1]
}

// Sets the Dir for the command object
func (j *Json) setDir(c *asink.Command, t typed.Typed) {
	c.Dir = t.String("dir")
}

// Sets the Args for the command object
func (j *Json) setArgs(c *asink.Command, t typed.Typed) {
	c.Args = t.StringsOr("args", []string{})
}

// Sets the Require for the task object
func (j *Json) setRequire(t *asink.Task, task typed.Typed) {
	t.Require = task.StringOr("require", "")
}

// Sets the Group for the task object
func (j *Json) setGroup(t *asink.Task, task typed.Typed) {
	t.Group   = task.StringOr("group", "")
}
