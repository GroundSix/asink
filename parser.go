package main

import (
	"strings"
)

type Parser interface {
	parse(body []byte) map[string]interface{}
}

func createParserFromFileType(fileName string) Parser {
	if strings.Contains(fileName, "yml") || strings.Contains(fileName, "yaml") {
		return Yaml{}
	}
	
	// Fall back to JSON if all else fails
	return Json{}
}