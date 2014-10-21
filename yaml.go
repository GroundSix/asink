package main

import (
	"./asink"
)

type Yaml struct {
	taskMap map[string]interface{}
	tasks   []asink.Task
}

// Parses the YAML into a typed.Typed object
// which acts as map[string]interface{}
func (y *Yaml) parse(body []byte) Parser {
	return y
}

// Creates and assigns tasks using the map
// from the Yaml struct
func (y *Yaml) assignTasks() Parser {
	return y
}