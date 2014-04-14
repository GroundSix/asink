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
        command := setupAsinkCommand(configFile)
        command.Execute()
    }
}

func setupAsinkCommand(configFile string) *asink.Command {
    command := asink.New()
    config  := jconfig.LoadConfig(configFile)

    counts := convertCounts(config.GetArray("count"))
    args   := convertArgs(config.GetArray("args"))

    command.SetName(config.GetString("command"))
    command.SetAsyncCount(int(counts[0]))
    command.SetRelativeCount(int(counts[1]))
    command.SetArgs(args)
    command.SetOutput(config.GetBool("output"))

    return command
}

func convertArgs(args []interface{}) []string {
    argsSlice := make([]string, len(args))
    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    return argsSlice
}

func convertCounts(counts []interface{}) []float64 {
    argsSlice := make([]float64, len(counts))
    for i, s := range counts {
        argsSlice[i] = s.(float64)
    }
    
    return argsSlice
}
