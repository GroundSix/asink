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
    "io/ioutil"
    "./asink"
)

func main() {
    // Creates the root and sub commands defined
    // in options.go using cobra
    createRootCommand()
}

func initAsinkWithFile(args []string) {
    if validateArguments(args) == true {
        fileName := args[0]
        p := createParserFromFileType(fileName)

        contents, err := ioutil.ReadFile(fileName)
        if (err != nil) {
            panic(err)
        }
        p = p.parse(contents)
        p.assignTasks()
        p.assignRemotes()
        asink.ExecMulti(p.Tasks())
    }
}

func initAsinkWithRequest(request []byte) {
    p := createJsonParser()
    p = p.parse(request)
    p.assignTasks()
    p.assignRemotes()
    asink.ExecMulti(p.Tasks())
}

func validateArguments(args []string) bool {
    if len(args) == 0 {
        fmt.Println("Arguments needed, 0 passed")
        fmt.Println("Use 'asink help' to see list of available commands")
        return false
    }
    return true
}
