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
    //"fmt"
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
            c := asink.NewCommand(task["command"].(string))
            c.AsyncCount = task.IntsOr("count", []int{1, 1})[0]
            c.RelCount   = task.IntsOr("count", []int{1, 1})[1]
            c.Dir        = task.StringOr("dir", ".")
            c.Args       = task.StringsOr("args", []string{})

            c.Callback = func(command string) {
                r := "$local: " + command
                color.Cyan(r)
            }

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
    remotes := a.TaskMap.StringObject("remotes")

    for name, remote := range remotes {
        r := NewRemote(name)
        r.Add(name)
        a.buildRemote(r, remote).Connect()
    }
    return a
}

// Builds up the asink command using the parsed
// JSON data
func (a *Assigner) buildCommand(c *asink.Command, t typed.Typed) {
    a.setRemote(c, t)
    a.setRemotes(c, t)
}

// Build up the asink task using the parsed
// JSON data
func (a *Assigner) buildTask(name string, task typed.Typed, c *asink.Command) asink.Task {
    t := asink.NewTask(name, c)
    return t
}

// Build up the asink remotes using the parsed
// JSON data
func (a *Assigner) buildRemote(r *Remote, remote typed.Typed) *Remote {
    a.setHost(r, remote)
    a.setPort(r, remote)
    a.setUser(r, remote)
    return a.setKey(r, remote)
}

// Settings for commands

// Sets the Args for the command object
func (a *Assigner) setRemote(c *asink.Command, t typed.Typed) {
    r := t.StringOr("remote", "")
    if (r != "") {
        c.Dummy = true
        c.Callback = func(command string) {
            runRemoteCommand(r, "cd " + c.Dir + " && " + command)
        }
    }
}

func (a *Assigner) setRemotes(c *asink.Command, t typed.Typed) {
    // TODO
}

// Settings for tasks

// Settings for remote machines

// Sets the Host for the remote object
func (a *Assigner) setHost(r *Remote, remote typed.Typed) {
    r.Host = remote.StringOr("host", "localhost")
}

// Sets the Port for the remote object
func (a *Assigner) setPort(r *Remote, remote typed.Typed) {
    r.Port = remote.StringOr("port", "2222")
}

// Sets the User for the remote object
func (a *Assigner) setUser(r *Remote, remote typed.Typed) {
    r.User = remote.StringOr("user", "root")
}


// Sets the Key for the remote object
func (a *Assigner) setKey(r *Remote, remote typed.Typed) *Remote {
    k := remote.StringOr("key", "")
    if (k != "") {
        return r.AddSshKey(r.Name, remote.String("key"))
    }
    return r
}
