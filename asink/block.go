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

package asink

import (
    "sync"
)

type Block struct {
    block      func()
    AsyncCount int
    RelCount   int
}

// Creates a new instance of Block with some
// default values. The block func is the
// only initial value that is required
func NewBlock(block func()) Block {
    return Block{block, 1, 1}
}

// Implemented to satisfy the task's Execer
// interface. Loops through the AsyncCount
// to concurrently execute the block
func (b Block) Exec() bool {
    var wg sync.WaitGroup

    block := make(chan Block)

    for i := 0; i != b.AsyncCount; i++ {
        wg.Add(1)
        go runBlock(block, &wg)
        block <- b
    }

    close(block)
    wg.Wait()
    return true
}

// Is called within Exec, the actual block
// execution happens in here
func runBlock(block chan Block, wg *sync.WaitGroup) {
    defer wg.Done()

    b := <- block

    for j := 0; j != b.RelCount; j++ {
        b.block()
    }
}
