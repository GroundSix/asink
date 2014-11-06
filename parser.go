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
    "strings"
)

type Parser interface {
    Tasks()            []asink.Task
    parse(body []byte) Parser
    assignTasks()      Parser
    assignRemotes()    Parser
}

// Creates a parser using the file extension
// as a way of determining what parser is
// needed
func createParserFromFileType(fileName string) Parser {
    if strings.Contains(fileName, "yml") || strings.Contains(fileName, "yaml") {
        return new(Yaml)
    }
    
    // Fall back to JSON if all else fails
    return new(Json)
}

func createJsonParser() Parser {
    return new(Json)
}

func createYamlParser() Parser {
    return new(Yaml)
}
