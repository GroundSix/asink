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
    name   := json_data.GetString("command")
    counts := convertCounts(json_data.GetArray("count"))
    args   := convertArgs(json_data.GetArray("args"))
    output := json_data.GetBool("output")

    command := createCommand(name, counts, args, output)
    command  = attachCallbacks(command)

    return command
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

    for _, cmd := range json_tasks {
        block := validateBlock(cmd.(map[string]interface{}))
        
        name    := block["command"].(string)
        counts  := convertCounts(block["count"].([]interface{}))
        args    := convertArgs(block["args"].([]interface{}))
        output  := block["output"].(bool)
        require := block["require"].(string)
        group   := block["group"].(string)

        command := createCommand(name, counts, args, output)
        command  = attachCallbacks(command)
        
        task.AddTask(command, require, group)
    }

    return task
}
