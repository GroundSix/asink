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
    "./vendor/pb"
)

var progressBar   *pb.ProgressBar = nil

// Creates the progress bar on the
// listen init event
func createProgressBar(count int) {
    progressBar = pb.StartNew(count)
}

// Increments the progress bar
// by one on the listen progress
// event
func incrementProgressBar() {
    progressBar.Increment()
}

// Stops the progress bar on the
// listen finish event
func endProgressBar() {
    progressBar.FinishPrint("Finished.")
}
