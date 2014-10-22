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

		c.AsyncCount = task.Ints("count")[0]
		c.RelCount   = task.Ints("count")[1]
		c.Dir  = task.String("dir")
		c.Args = task.StringsOr("args", []string{})

		t := asink.NewTask(taskName, c)
		tasksSlice = append(tasksSlice, t)
	}
	j.tasks = tasksSlice
	return j
}

func (j *Json) Tasks() []asink.Task {
	return j.tasks
}
