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
    "os"
    "net/http"
    "io/ioutil"
    "./asink"
    "./vendor/jconfig"
)

// Entry point for asink. Runs the command
// and follows general instructions as
// specefied in the JSON configuration
// file
func main() {
    // Located in help.go
    executeRootCommand()
    
    closeSshSessions()
}

// Sets up the configuration for asink
// and executes the command
func initAsinkWithFile(config_file_path string) {
    json_data  := jconfig.LoadConfig(config_file_path)
    startExecutionProcess(json_data)
}

// Inits asink with only a JSON string
func initAsinkWithString(json string) {
    json_data := jconfig.LoadConfigString(json)
    startExecutionProcess(json_data)
}

// Gets response from GET to parse
// the JSON
func initAsinkWithHttp(url string) {
    response, err := http.Get(url)
    if err != nil {
        panic(err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            panic(err)
        }
        json_data := jconfig.LoadConfigString(string(contents))
        startExecutionProcess(json_data)
    }
}

// Used in both inits to start the execution
// process of a command or a task
func startExecutionProcess(json_data *jconfig.Config) {
    setupExtend(json_data)
    setupIncludes(json_data)
    setupSshRemotes(json_data)

    // Keep the original legacy command API
    if len(json_data.GetStringMap("tasks")) > 0 {
        task := setupAsinkTasks(json_data)
        task.Execute()
    } else {
        command := setupAsinkCommand(json_data)
        command.Execute()
    }
}

// Initially sets up everything from
// config file in a new instance of Asink
func setupAsinkCommand(json_data *jconfig.Config) *asink.Command {
    name   := json_data.GetString("command")
    counts := convertCounts(json_data.GetArray("count"))
    args   := convertStringArray(json_data.GetArray("args"))

    command := createCommand(name, counts, args, ".")

    return command
}

// If tasks are detected, they are configured here
func setupAsinkTasks(json_data *jconfig.Config) *asink.Task {
    task       := asink.NewTask()
    if len(json_data.GetStringMap("tasks")) > 0 {
        json_tasks := json_data.GetStringMap("tasks")

        for task_name, cmd := range json_tasks {
            block := validateBlock(cmd.(map[string]interface{}))
            
            name     := block["command"].(string)
            counts   := convertCounts(block["count"].([]interface{}))
            args     := convertStringArray(block["args"].([]interface{}))
            require  := block["require"].(string)
            group    := block["group"].(string)
            remote   := block["remote"].(string)
            dir      := block["dir"].(string)

            command := createCommand(name, counts, args, dir)

            command.SetManualCallback(func(name string) {
                runInSshSession(remote, name)
            });
            
            task.AddTask(task_name, command, require, group)
            task.SetRemote(task_name, remote)
        }
    }
    return task
}

// If SSH remotes are detected, they are setup
// here
func setupSshRemotes(json_data *jconfig.Config) {
    if len(json_data.GetStringMap("ssh")) > 0 {
        remote       := NewRemote()
        json_remotes := json_data.GetStringMap("ssh")

        for remote_name, config := range json_remotes {
            block := config.(map[string]interface{})
            block  = validateBlock(block)

            host     := block["host"].(string)
            port     := block["port"].(string)
            user     := block["user"].(string)
            password := block["password"].(string)
            key      := block["key"].(string)

            remote.AddRemote(remote_name, host, port, user, password)
            if (password == "") {
                remote.AddSshKey(remote_name, key)
            }
        }
    }
}

// Remotely runs a command within the SSH session
func runInSshSession(remote string, command string) {
    if (remote != "") {
        StartSession(remote)
        RunRemoteCommand(remote, command)
    }
}

// Runs extended file if is is found
func setupExtend(json_data *jconfig.Config) {
    if len(json_data.GetString("extend")) > 0 {
        extend    := json_data.GetString("extend")
        file_name := extend + ".json"
        if _, err := os.Stat(file_name); err == nil {
            contents_bytes, _ := ioutil.ReadFile(file_name)
            contents := string(contents_bytes)
            initAsinkWithString(contents)
        }
    }
}

// Sets up asink to run the included tasks
// if they are needed
func setupIncludes(json_data *jconfig.Config) {
    if len(json_data.GetArray("include")) > 0 {
        include := json_data.GetArray("include")
        for _, value := range include {
            file_name := value.(string) + ".json"
            if _, err := os.Stat(file_name); err == nil {
                contents_bytes, _ := ioutil.ReadFile(file_name)
                contents := string(contents_bytes)
                contents_json := jconfig.LoadConfigString(contents)
                json_tasks := contents_json.GetStringMap("tasks")
                for key, task := range json_tasks {
                    task_interface := task.(map[string]interface{})
                    task_interface["tag"] = value
                    json_tasks[key] = task_interface
                }
                
            }
        }
    }
}

// Returns the current working directory
// as a string
func getWorkingDirectory() string {
    dir, err := os.Getwd()
    if (err != nil) {
        panic(err)
    }
    return dir
}
