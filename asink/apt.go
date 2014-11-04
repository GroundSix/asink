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

import "fmt"

type Apt struct {
    Action        string
    Packages      []string
    Callback      func(command string)
    Dummy         bool
    CommandString string
}

// Creates a new instance of Apt with some
// default values. The action string is the
// only initial value that is required
func NewApt(action string) Apt {
    a := Apt{}
    a.Action  = action
    a.Packages = []string{}
    a.Callback = func(command string){}
    a.Dummy    = false
    a.CommandString = ""
    return a
}

// Generates the command string to be ran based
// on apt-get. Currently supports 'update' or
// 'install'
func (a Apt) Exec() bool {
    a.CommandString = "apt-get"
    a.CommandString = a.appendAptAction(a.CommandString)

    if a.Dummy == false {
        c := NewCommand("apt-get")
        args := []string{a.Action}
        for _, pa := range a.Packages {
            args = append(args, pa)
        }
        c.Args = args
        c.Callback = func(command string) {
            fmt.Println(command)
        }
        c.Exec()
    }
    return true
}

// Adds a package to install
func (a *Apt) AddPackage(p string) {
    a.Packages = append(a.Packages, p)
}

// Adds multiple packages to install
func (a *Apt) AddPackages(p []string) {
    for _, pa := range p {
        a.Packages = append(a.Packages, pa)
    }
}

func (a Apt) appendAptAction(command string) string {
    if a.Action == "update" {
        command = command + " update"
    } else if (a.Action == "install") {
        command = command + " install -y"
        for _, p := range a.Packages {
            command = command + " " + p
        }
    }
    return command
}
