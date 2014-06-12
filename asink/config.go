/**
* asink v0.0.2-dev
*
* (c) Ground Six
*
* @package asink
* @version 0.0.1
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
* Gets the name of your config file
* from the param passed through when
* the program is ran
*
* e.g. asink config.json
*
* @return string file path or empty string
 */
func GetConfigFile() string {
	if len(os.Args) > 2 {
		if (os.Args[1] == "start") {
			filePath := os.Args[2]
			if _, err := os.Stat(filePath); err == nil {
				return filePath
			}
		}
	}
	return ""
}
