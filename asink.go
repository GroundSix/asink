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
	blockTask.Require = "do-ls"

	ls 	   := asink.NewCommand("ls")
	lstask := asink.NewTask("do-ls", ls)

	tasks := []asink.Task{blockTask, lstask}

	asink.ExecMulti(tasks)
}