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
)

func TestNewCommand(t *testing.T) {
    c := NewCommand("echo")
    c.Args = []string{"'Hello, World!"}
    c.AsyncCount = 1
    c.RelCount   = 1
    c.Dir        = "~"

    tp := reflect.TypeOf(c).String()
    if tp != "asink.Command" {
        t.Error("Expected asink.Command, got ", tp)
    }
}

func TestExecCommand(t *testing.T) {
    c := NewCommand("echo")
    c.Args = []string{"Hello, World!"}
    result := c.Exec()
    if result != true {
        t.Error("Expected true, got", result)
    }
}
