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

func TestNewApt(t *testing.T) {
	ai := NewApt("install")

    tmi := reflect.TypeOf(ai).String()
    if tmi != "asink.Apt" {
        t.Error("Expected asink.Apt, got ", tmi)
    }

    au := NewApt("update")
    tmu := reflect.TypeOf(au).String()
    if tmu != "asink.Apt" {
        t.Error("Expected asink.Apt, got ", tmu)
    }
}
