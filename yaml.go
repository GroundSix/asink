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

func (y *Yaml) buildTasks() Parser {
	return y
}

func (y *Yaml) Tasks() []asink.Task {
	return []asink.Task{}
}
