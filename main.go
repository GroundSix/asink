/**
 * asink v0.0.1
 *
 * (c) Ground Six
 *
 * @package asink
<<<<<<< HEAD
 * @version 0.0.2-dev
=======
 * @version 0.0.1
>>>>>>> master
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
    "./vendor/cobra"
    "./vendor/jconfig"
)

/**
 * Entry point for asink. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
<<<<<<< HEAD
    configFile := asink.GetConfigFile()
    if configFile != "" {
        command := setupAsinkCommand(configFile)

        command.ListenForInit(createProgressBar)
        command.ListenForProgress(incrementProgressBar)
        command.ListenForFinish(endProgressBar)
        command.Execute()
=======
    var startCommand = &cobra.Command{
        Use:   "start [JSON configuration file]",
        Short: "Start your asink processes",
        Long:  `start running a command the specified amount of times from your configuration file`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsink()
        },
>>>>>>> master
    }
    var rootCmd = &cobra.Command{Use: "asink"}
    rootCmd.AddCommand(startCommand)
    rootCmd.Execute()
}

/**
 * Sets up the configuration for asink
 * and executes the command
 *
 * @return nil
 */
func initAsink() {
    configFile := asink.GetConfigFile()
    if configFile != "" {
        command := setupAsinkCommand(configFile)
        command.Execute()
    }
}

/**
 * Initially sets up everything from
 * config file in a new instance of Asink
 *
 * @param string path to config file
 *
 * @return *asink.Command configured instance of
 * asink
 */
func setupAsinkCommand(configFile string) *asink.Command {
    command := asink.New()
    config := jconfig.LoadConfig(configFile)

    counts := convertCounts(config.GetArray("count"))
    args := convertArgs(config.GetArray("args"))

    command.Name = config.GetString("command")
    command.AsyncCount = counts[0]
    command.RelativeCount = counts[1]
    command.Args = args
    command.Output = config.GetBool("output")

    if (command.Output == false) {
        command.ListenForInit(createProgressBar)
        command.ListenForProgress(incrementProgressBar)
        command.ListenForFinish(endProgressBar)
    }

    return command
}

/**
 * Converts jconfigs []interface into
 * []string for asink
 *
 * @param []interface{} jconfig's array
 *
 * @return []string asink's array
 */
func convertArgs(args []interface{}) []string {
    argsSlice := make([]string, len(args))
    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    return argsSlice
}

/**
 * Converts jconfigs []interface into
 * []float64 for asink
 *
 * @param []interface{} jconfig's array
 *
 * @return []float64 asink's array
 */
func convertCounts(counts []interface{}) []float64 {
    argsSlice := make([]float64, len(counts))
    for i, s := range counts {
        argsSlice[i] = s.(float64)
    }

    return argsSlice
}
