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
    "testing"
    "reflect"
    "fmt"
)

func TestNewTask(t *testing.T) {
    c := NewCommand("echo")
    c.Args = []string{"'Hello, World!'"}

    ta := NewTask("do-echo", c)
    tp := reflect.TypeOf(ta).String()
    if tp != "asink.Task" {
        t.Error("Expected asink.Task, got ", tp)
    }
}

func TestExecTask(t *testing.T) {
    c := NewCommand("echo")
    c.Args = []string{"'Hello, World!'"}

    ta := NewTask("do-echo", c)
    result := ta.Exec()
    if result != true {
        t.Error("Expected true, got ", result)
    }
}

func TextExecTaskMulti(t *testing.T) {
    c := NewCommand("echo")
    c.Args = []string{"'Hello, World!'"}

    tac := NewTask("do-echo", c)
    tac.Require = "do-print"

    b := NewBlock(func() {
        fmt.Println("Hello, World!")
    });

    tab := NewTask("do-print", b)

    tasks := []Task{tac, tab}

    result := ExecMulti(tasks)
    if result != true {
        t.Error("Expected true, got ", result)
    }
}
