/**
 * asink v0.0.2
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.0.2
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package asink

import (
    "os"
)

/**
 * Gets params and flags from cli
 * and calls appropriate func
 *
 * @return String 
 */
func GetFirstCliParam() string {
	if len(os.Args) > 2 {
		if (os.Args[1] == "start") {
			return CliStart()
		}
	}
	return ""
}


/**
 * Returns string from cli to start
 * asink
 *
 * @return String config file name
 */
func CliStart() string {
	filePath := os.Args[2]
	if _, err := os.Stat(filePath); err == nil {
		return filePath
	}
	return ""
}
