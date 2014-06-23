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
    "./vendor/jconfig"
)

/**
 * Entry point for asink. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
    // Located in help.go
    executeRootCommand()
}

/**
 * Sets up the configuration for asink
 * and executes the command
 *
 * @return nil
 */
func initAsink() {
    configFile := asink.GetFirstCliParam()
    if configFile != "" {
        json_data  := jconfig.LoadConfig(configFile)
        if detectSshRemotes(json_data) == true {
            setupSshRemotes(json_data)
        }
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

/**
 * Detects if there are any tasks to be ran
 *
 * @param *jconfig.Config json data
 *
 * @return Bool
 */
func detectTasks(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("tasks")) > 0 {
        return true
    }
    return false
}

/**
 * If tasks are detected, they are configured here
 *
 * @param *jconfig.Config json data
 *
 * @return *asink.Task configured task
 */
func setupAsinkTasks(json_data *jconfig.Config) *asink.Task {
    task       := asink.NewTask()
    json_tasks := json_data.GetStringMap("tasks")

    for task_name, cmd := range json_tasks {
        block := validateBlock(cmd.(map[string]interface{}))
        
        name    := block["command"].(string)
        counts  := convertCounts(block["count"].([]interface{}))
        args    := convertArgs(block["args"].([]interface{}))
        output  := block["output"].(bool)
        require := block["require"].(string)
        group   := block["group"].(string)
        remote  := block["remote"].(string)

        command := createCommand(name, counts, args, output)
        command  = attachCallbacks(command)

        command.SetManualCallback(func(name string) {
            runInSshSession(remote, name)
        });
        
        task.AddTask(task_name, command, require, group)
        task.SetRemote(task_name, remote)
    }

    return task
}

/**
 * Detects if there are any SSH remotes
 * that need to be setup
 *
 * @param *jconfig.Config json data
 *
 * @return Bool
 */
func detectSshRemotes(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("ssh")) > 0 {
        return true
    }
    return false
}

/**
 * If SSH remotes are detected, they are setup
 * here
 *
 * @param *jconfig.Config json data
 *
 * @return nil
 */
func setupSshRemotes(json_data *jconfig.Config) {
    remote       := NewRemote()
    json_remotes := json_data.GetStringMap("ssh")

    for remote_name, config := range json_remotes {
        block := config.(map[string]interface{})

        host     := block["host"].(string)
        port     := block["port"].(string)
        user     := block["user"].(string)
        password := block["password"].(string)

        remote.AddRemote(remote_name, host, port, user, password)
    }
}

/**
 * Remotely runs a command within the SSH session
 *
 * @param String remote name
 * @param String command name and args
 *
 * @return nil
 */
func runInSshSession(remote string, command string) {
    if (remote != "") {
        StartSession(remote)
        RunRemoteCommand(remote, command)
    }
}
