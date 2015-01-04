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
)

type Json struct {
	taskMap typed.Typed
}

// Parses the JSON into a typed.Typed object
// which acts as map[string]interface{}
func (j *Json) parse(body []byte) Parser {
	mapped, err := typed.Json(body)
	if err != nil {
		panic(err)
	}
	j.taskMap = mapped
	return j
}

// Returns a map of parsed tasks
func (j Json) TaskMap() typed.Typed {
	return j.taskMap
}
