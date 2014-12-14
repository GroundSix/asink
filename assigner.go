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
    "strconv"
    "./asink"
    "github.com/asink/typed"
    "github.com/asink/color"
)

type Assigner struct {
    TaskMap typed.Typed
    tasks   []asink.Task
}

// Returns a slice of all asink tasks
func (a *Assigner) Tasks() []asink.Task {
    return a.tasks
}

// Creates and assigns tasks using the map
// from the Json struct
func (a *Assigner) assignTasks() *Assigner {
    t := a.TaskMap.StringObject("tasks")
    tasks := []asink.Task{}

    for n, task := range t {
        if task["command"] != nil {

            // Creates a new Asink command with values that have
            // been converted into Typed objects
            c := asink.NewCommand(task["command"].(string))
            c.AsyncCount = task.IntsOr("count", []int{1, 1})[0]
            c.RelCount   = task.IntsOr("count", []int{1, 1})[1]
            c.Dir        = task.StringOr("dir", ".")
            c.Args       = task.StringsOr("args", []string{})

            // Set a default callback as we don't know if this will
            // be ran on a remote machine yet or not
            c.Callback = func(command string) {
                r := "$local: " + command
                color.Cyan(r)
            }

            r := task.StringOr("remote", "")
            if (r != "") {
                c.Dummy = true
                c.Callback = func(command string) {
                    remotes[r].Connect()
                    runRemoteCommand(r, "cd " + c.Dir + " && " + command)
                }
            }

            // Start to build the Asink task up
            at := asink.NewTask(n, c)
            at.Require = task.StringOr("require", "")
            at.Group   = task.StringOr("group", "")

            tasks = append(tasks, at)
        }
    }
    a.tasks = tasks

    return a
}

// Creates and assigns remotes using the map
// from the Json struct
func (a *Assigner) assignRemotes() *Assigner {
    r := a.TaskMap.StringObject("remotes")
    for n, remote := range r {
        r := NewRemote(n)

        r.Host = remote.StringOr("host", "localhost")
        r.Port = strconv.Itoa(remote.IntOr("port", 22))
        r.User = remote.StringOr("user", "root")

        r.Add(n)
        r.AddSshKey(n, remote.String("key"))
    }

    return a
}
