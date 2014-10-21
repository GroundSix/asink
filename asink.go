package main

import (
	"fmt"
	"./asink"
)

func main() {
	b := asink.NewBlock(func() {
		fmt.Println("Yo!")
	});

	b.AsyncCount = 3
	b.RelCount   = 3

	blockTask := asink.NewTask("block", b)

	blockTask.Exec()
}