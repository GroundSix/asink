package asink

import (
	"sync"
)

type Block struct {
	block func()
	AsyncCount int
	RelCount   int
}

func NewBlock(block func()) Block {
	return Block{block, 1, 1}
}

func (b Block) Exec() {
	var wg sync.WaitGroup

	block := make(chan Block)

    for i := 0; i != b.AsyncCount; i++ {
        wg.Add(1)
        go runBlock(block, &wg)
        block <- b
    }

    close(block)
    wg.Wait()
}

func runBlock(block chan Block, wg *sync.WaitGroup) {
    defer wg.Done()

    b := <- block

    for j := 0; j != b.RelCount; j++ {
    	b.block()
	}
}