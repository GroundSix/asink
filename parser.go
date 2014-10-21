package main

type Parser interface {
	parse(body string) map[string]interface{}
}
