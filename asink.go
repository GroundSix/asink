package main

import (
	"./asink"
)

func main() {
	ls := asink.NewCommand("ls")
	ls.Exec()
}