package main

import (
	"strings"
)

type Parser interface {
	parse(body []byte) Parser
	assignTasks() Parser
}

func createParserFromFileType(fileName string) Parser {
	if strings.Contains(fileName, "yml") || strings.Contains(fileName, "yaml") {
		return new(Yaml)
	}
	
	// Fall back to JSON if all else fails
	return new(Json)
}