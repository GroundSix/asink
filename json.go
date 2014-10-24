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
	tasks 	   := j.taskMap.StringObject("tasks")
	tasksSlice := []asink.Task{}
	for taskName, task := range tasks {
		c := asink.NewCommand(task["command"].(string))

		j.buildCommand(&c, task)

		t := j.buildTask(taskName, task, &c)
		tasksSlice = append(tasksSlice, t)
	}
	j.tasks = tasksSlice
	return j
}

// Builds up the asink command using the parsed
// JSON data
func (j *Json) buildCommand(c *asink.Command, t typed.Typed) {
	c.AsyncCount = t.IntsOr("count", []int{1, 1})[0]
	c.RelCount = t.IntsOr("count", []int{1, 1})[1]
	c.Dir = t.String("dir")
	c.Args = t.StringsOr("args", []string{})
}

// Build up the asink task using the parsed
// JSON data
func (j *Json) buildTask(taskName string, task typed.Typed, c *asink.Command) asink.Task {
	t := asink.NewTask(taskName, c)
	t.Require = task.StringOr("require", "")
	t.Group   = task.StringOr("group", "")
	return t
}

