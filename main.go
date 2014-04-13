/**
 * asink v0.1-dev
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
    "./asink"
    "./vendor/jconfig"
)

/**
 * Entry point for asink. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
    configFile := asink.GetConfigFile()
    if configFile != "" {
        Asink := asink.New()

        config  := jconfig.LoadConfig(configFile)
        command := config.GetString("command")
        counts  := config.GetArray("count")
        args    := config.GetArray("args")

        Asink.SetOutput(config.GetBool("output"))
        Asink.Execute(command, counts[0].(float64), counts[1].(float64), args)
    }
}
