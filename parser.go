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
	"path/filepath"
	"fmt"
)

type Parser interface {
	parse(body []byte) Parser
	TaskMap() typed.Typed
}

// Creates a parser using the file extension
// as a way of determining what parser is
// needed
func parserFromFileType(filename string) (Parser, error) {
	ext := filepath.Ext(filename)
	if (ext == ".json") {
		return new(Json), nil
	}
	if (ext == ".yml" || ext == ".yaml") {
		return new(Yaml), nil
	}
	return nil, fmt.Errorf("Could not recognize file type with extension '%s'", ext)
}

// Returns a new instance of the JSON parser
func createJsonParser() Parser {
	return new(Json)
}

// Returns a new instance of the YAML parser
func createYamlParser() Parser {
	return new(Yaml)
}
