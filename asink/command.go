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
    "os"
    "os/exec"
    "sync"
    "strings"
)

type Command struct {
    Name       string
    AsyncCount int
    RelCount   int
    Dir        string
    Args       []string
    Callback   func(command string)
    Dummy      bool
}

// Creates a new instance of Command with some
// default values. The command string is the
// only initial value that is required
func NewCommand(name string) Command {
    return Command{
        name,                   // Root command (required)
        1,                      // AsyncCount
        1,                      // RelCount
        getWorkingDirectory(),  // Directory
        []string{},             // Command arguments
        func(command string){}, // Callback func
        false}                  // Dummy command or not?
}

// Implemented to satisfy the task's Execer
// interface. Loops through the AsyncCount
// to concurrently execute the command
func (c Command) Exec() bool {
    var wg sync.WaitGroup

    command := make(chan Command)

    validateDirectoryName(&c)
    os.Chdir(getWorkingDirectory())
    os.Chdir(c.Dir)

    for i := 0; i != c.AsyncCount; i++ {
        wg.Add(1)
        go runCommand(command, &wg)
        command <- c
    }

    close(command)
    wg.Wait()
    return true
}

// Is called within Exec, the actual command
// execution happens in here
func runCommand(command chan Command, wg *sync.WaitGroup) {
    defer wg.Done()

    c := <- command

    for j := 0; j != c.RelCount; j++ {
        c.Callback(c.Name + " " + strings.Join(c.Args, " "))
        if c.Dummy == false {
            cmd := exec.Command(c.Name, c.Args...)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
        }
    }
}
