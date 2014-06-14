/**
 * asink v0.0.2-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.0.2-dev
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
    "fmt"
    //"strings"
    //"encoding/json"
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
    var startCommand = &cobra.Command{
        Use:   "start [JSON configuration file]",
        Short: "Start your asink processes",
        Long:  `start running a command the specified amount of times from your configuration file`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsink()
        },
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
        json_data  := jconfig.LoadConfig(configFile)
        if detectTasks(json_data) == true {
            task := setupAsinkTasks(json_data)
            task.Execute()
        } else {
            command := setupAsinkCommand(json_data)
            command.Execute()   
        }
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
func setupAsinkCommand(json_data *jconfig.Config) *asink.Command {
    command := asink.New()

    counts := convertCounts(json_data.GetArray("count"))
    args   := convertArgs(json_data.GetArray("args"))

    command.Name = json_data.GetString("command")
    command.AsyncCount = counts[0]
    command.RelativeCount = counts[1]
    command.Args = args
    command.Output = json_data.GetBool("output")

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

func detectTasks(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("tasks")) > 0 {
        return true
    }
    return false
}

func setupAsinkTasks(json_data *jconfig.Config) *asink.Task {
    task       := asink.NewTask()
    json_tasks := json_data.GetStringMap("tasks")

    for key, cmd := range json_tasks {
        command := asink.New()
        
        command_string := cmd.(map[string]interface{})

        fmt.Println("key:", key, "   value:", command_string["command"])
/*
        counts := convertCounts(cmd.GetArray("count"))
        args   := convertArgs(cmd.GetArray("args"))

        command.Name = cmd.GetString("command")
        command.AsyncCount = counts[0]
        command.RelativeCount = counts[1]
        command.Args = args
        command.Output = cmd.GetBool("output")
        */
        task.AddTask(command, "test", "test")
    }

    return task
}