/**
 * asink v0.0.1-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.1-dev
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
    "./vendor/pb"
)

/**
 * @var *pb.ProgressBar asink's progress indicator
 */
var progressBar   *pb.ProgressBar = nil

/**
 * Creates the progress bar on the
 * listen init event
 *
 * @param int number of commands
 *
 * @return nil
 */
func createProgressBar(count int) {
    progressBar = pb.StartNew(count)
}

/**
 * Increments the progress bar
 * by one on the listen progress
 * event
 *
 * UI progress is well off here
 * just testing with it at the
 * moment
 *
 * @return nil
 */
func incrementProgressBar() {
    progressBar.Increment()
}

/**
 * Stops the progress bar on the
 * listen finish event
 *
 * @return nil
 */
func endProgressBar() {
    progressBar.FinishPrint("Finished.")
}
