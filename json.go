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
    "./asink"
    "github.com/asink/typed"
    "github.com/asink/color"
)

var callbackRemoteName string = ""

type Json struct {
    taskMap typed.Typed
    tasks   []asink.Task
}

// Returns a slice of all asink tasks
func (j *Json) Tasks() []asink.Task {
    return j.tasks
}

// Parses the JSON into a typed.Typed object
// which acts as map[string]interface{}
func (j *Json) parse(body []byte) Parser {
    mapped, err := typed.Json(body)
    if (err != nil) {
        panic(err)
    }
    j.taskMap = mapped
    return j
}

// Creates and assigns tasks using the map
// from the Json struct
func (j *Json) assignTasks() Parser {
    t := j.taskMap.StringObject("tasks")
    tasks := []asink.Task{}
    for name, task := range t {
        c := asink.NewCommand(task["command"].(string))
        j.buildCommand(&c, task)
        tasks = append(tasks, j.buildTask(name, task, &c))
    }
    j.tasks = tasks
    return j
}

// Creates and assigns remotes using the map
// from the Json struct
func (j *Json) assignRemotes() Parser {
    remotes := j.taskMap.StringObject("ssh")

    for name, remote := range remotes {
        r := NewRemote(name)
        r.Add(name)
        j.buildRemote(r, remote).Connect()
    }
    return j
}

// Builds up the asink command using the parsed
// JSON data
func (j *Json) buildCommand(c *asink.Command, t typed.Typed) {
    j.setAsyncCount(c, t)
    j.setRelCount(c, t)
    j.setDir(c, t)
    j.setArgs(c, t)
    j.setCallback(c, t)
    j.setRemote(c, t)
    j.setRemotes(c, t)
}

// Build up the asink task using the parsed
// JSON data
func (j *Json) buildTask(name string, task typed.Typed, c *asink.Command) asink.Task {
    t := asink.NewTask(name, c)
    j.setRequire(&t, task)
    j.setGroup(&t, task)
    return t
}

// Build up the asink remotes using the parsed
// JSON data
func (j *Json) buildRemote(r *Remote, remote typed.Typed) *Remote {
    j.setHost(r, remote)
    j.setPort(r, remote)
    j.setUser(r, remote)
    j.setPassword(r, remote)
    return j.setKey(r, remote)
}

// Settings for commands

// Sets the AsyncCount for the command object
func (j *Json) setAsyncCount(c *asink.Command, t typed.Typed) {
    c.AsyncCount = t.IntsOr("count", []int{1, 1})[0]
}

// Sets the RelCount for the command object
func (j *Json) setRelCount(c *asink.Command, t typed.Typed) {
    c.RelCount = t.IntsOr("count", []int{1, 1})[1]
}

// Sets the Dir for the command object
func (j *Json) setDir(c *asink.Command, t typed.Typed) {
    c.Dir = t.String("dir")
}

// Sets the Args for the command object
func (j *Json) setArgs(c *asink.Command, t typed.Typed) {
    c.Args = t.StringsOr("args", []string{})
}

// Sets the Callback for the command object
func (j *Json) setCallback(c *asink.Command, t typed.Typed) {
    c.Callback = func(command string) {
        r := "$local: " + command
        color.Cyan(r)
    }
}

// Sets the Args for the command object
func (j *Json) setRemote(c *asink.Command, t typed.Typed) {
    r := t.StringOr("remote", "")
    if (r != "") {
        c.Dummy = true
        c.Callback = func(command string) {
            runRemoteCommand(r, command)
        }
    }
}

func (j *Json) setRemotes(c *asink.Command, t typed.Typed) {
    /*
    r := t.StringsOr("remotes", []string{})
    c.Dummy = true
    c.Callback = func(command string) {
        tasks := []asink.Task{}
        for _, remoteName := range r {
            callbackRemoteName = remoteName
            task := asink.NewTask("task-"+remoteName, asink.NewBlock(func() {
                runRemoteCommand(callbackRemoteName, command)
            }))
            task.Group = "exec-remotes"
            tasks = append(tasks, task)
        }
        fmt.Println(tasks)
        asink.ExecMulti(tasks)
    }*/
}

// Settings for tasks

// Sets the Require for the task object
func (j *Json) setRequire(t *asink.Task, task typed.Typed) {
    t.Require = task.StringOr("require", "")
}

// Sets the Group for the task object
func (j *Json) setGroup(t *asink.Task, task typed.Typed) {
    t.Group   = task.StringOr("group", "")
}

// Settings for remote machines

// Sets the Host for the remote object
func (j *Json) setHost(r *Remote, remote typed.Typed) {
    r.Host = remote.StringOr("host", "localhost")
}

// Sets the Port for the remote object
func (j *Json) setPort(r *Remote, remote typed.Typed) {
    r.Port = remote.StringOr("port", "2222")
}

// Sets the User for the remote object
func (j *Json) setUser(r *Remote, remote typed.Typed) {
    r.User = remote.StringOr("user", "root")
}

// Sets the Password for the remote object
func (j *Json) setPassword(r *Remote, remote typed.Typed) {
    r.Password = remote.StringOr("password", "")
}

// Sets the Key for the remote object
func (j *Json) setKey(r *Remote, remote typed.Typed) *Remote {
    k := remote.StringOr("key", "")
    if (k != "") {
        return r.AddSshKey(r.Name, remote.String("key"))
    }
    return r
}
