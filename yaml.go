package main

type Yaml struct {
	body []byte
}

func (y Yaml) parse(body []byte) map[string]interface{} {
	v := map[string]interface{}{"hi" : "d"}
	return v
}