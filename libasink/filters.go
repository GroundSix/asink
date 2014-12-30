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

type Filter struct {
    Dummy    bool
    commands []string
}

// A list of software packages defined for commands
// or configuration to be ran before the install
var packages map[string]func(f *Filter) = map[string]func(f *Filter){

}

// Creates a new instance of Filter with a
// default value. The task package string.
func NewFilter() Filter {
    return Filter{false, []string{}}
}

// Applies the filter before the package is
// installed
func (f Filter) Apply(installs []string) {
    for _, p := range installs {
        packages[p](&f)
    }
}

func (f Filter) Commands() []string {
    return f.commands
}

// Package Filters

