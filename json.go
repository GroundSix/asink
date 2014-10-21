package main

import (
	//"fmt"
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
		t := asink.NewTask(taskName, c)
		tasksSlice = append(tasksSlice, t)
	}
	j.tasks = tasksSlice
	return j
}
