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

type Apt struct {
    Action   string
    Package  string
    Callback func(command string)
}

// Creates a new instance of Apt with some
// default values. The action string is the
// only initial value that is required
func NewApt(action string) Apt {
    a := Apt{}
    a.Action  = action
    a.Package = ""
    return a
}

// Generates the command string to be ran based
// on apt-get. Currently supports 'update' or
// 'install'
func (a Apt) Exec() bool {
    if a.Action == "update" {
        a.Callback("apt-get update")
    } else if (a.Action == "install") {
        a.Callback("apt-get install -y " + a.Package)
    } else {
        return false
    }
    return true
}

// Adds a package to install, needs to be changed
// into a slice soon rather than concating a
// string like this!
func (a Apt) AddPackage(p string) {
    a.Package = a.Package + " " + p
}
