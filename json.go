package main

import (
	"github.com/asink/typed"
)

type Json struct {
	body   string
	mapped typed.Typed
}

func (j *Json) parse(body string) typed.Typed {
	mapped, err := typed.Json([]byte(body))
	if (err != nil) {
		panic(err)
	}
	j.body = body
	j.mapped = mapped
	return j.mapped
}
