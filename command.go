// asink v0.0.3-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.0.3-dev
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
)

// Creates a default command
func createCommand(name string, counts []float64, args[]string, dir string) *asink.Command {
	command := asink.New()

	command.Name 		  = name
	command.AsyncCount    = counts[0]
    command.RelativeCount = counts[1]
    command.Args 		  = args
    command.Dir           = dir

	return command
}

// Converts jconfigs []interface into
// []string for asink
func convertStringArray(args []interface{}) []string {
    argsSlice := make([]string, len(args))
    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    return argsSlice
}

// Converts jconfigs []interface into
// []float64 for asink
func convertCounts(counts []interface{}) []float64 {
    argsSlice := make([]float64, len(counts))
    for i, s := range counts {
        argsSlice[i] = s.(float64)
    }

    return argsSlice
}
