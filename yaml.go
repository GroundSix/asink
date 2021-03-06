// asink v0.1.1-dev
//
// (c) Ground Six 2015
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: GNU GPL v2.0
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"github.com/asink/typed"
	"github.com/asink/yaml"
)

type Yaml struct {
	taskMap typed.Typed
}

// Parses the YAML into a typed.Typed object
// which acts as map[string]interface{}
func (y *Yaml) parse(body []byte) Parser {
	var parsed interface{}
	err := yaml.Unmarshal(body, &parsed)
	if err != nil {
		panic(err)
	}
	parsedMap := parsed.(map[string]interface{})
	y.taskMap = typed.New(parsedMap)

	return y
}

// Returns a map of parsed tasks
func (y Yaml) TaskMap() typed.Typed {
	return y.taskMap
}
