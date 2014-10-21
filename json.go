package main

import (
	"github.com/asink/typed"
)

type Json struct {
	body   []byte
	mapped typed.Typed
}

func (j Json) parse(body []byte) map[string]interface{} {
	mapped, err := typed.Json(body)
	if (err != nil) {
		panic(err)
	}
	j.body   = body
	j.mapped = mapped
	return j.mapped
}
