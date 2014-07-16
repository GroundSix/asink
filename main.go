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
func initAsink() {
    configFile := asink.GetFirstCliParam()
    if configFile != "" {
        json_data  := jconfig.LoadConfig(configFile)
        startExecutionProcess(json_data)
    }
}

// Gets response from GET to parse
// the JSON
func initAsinkWithHttp(args []string) {
    if args[0] != "" {
        response, err := http.Get(args[0])
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
}

// Inits asink with only a JSON string
func initAsinkWithString(json string) {
    json_data := jconfig.LoadConfigString(json)
    startExecutionProcess(json_data)
}

// Used in both inits to start the execution
// process of a command or a task
//
// This will need reducing soon
func startExecutionProcess(json_data *jconfig.Config) {
    if detectExtend(json_data) == true {
        setupExtend(json_data)
    }
    if detectIncludes(json_data) == true {
        setupIncludes(json_data)
    }
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

// Initially sets up everything from
// config file in a new instance of Asink
func setupAsinkCommand(json_data *jconfig.Config) *asink.Command {
    name   := json_data.GetString("command")
    counts := convertCounts(json_data.GetArray("count"))
    args   := convertStringArray(json_data.GetArray("args"))

    command := createCommand(name, counts, args, ".")

    return command
}

// Detects if there are any tasks to be ran
func detectTasks(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("tasks")) > 0 {
        return true
    }
    return false
}

// If tasks are detected, they are configured here
func setupAsinkTasks(json_data *jconfig.Config) *asink.Task {
    task       := asink.NewTask()
    json_tasks := json_data.GetStringMap("tasks")

    for task_name, cmd := range json_tasks {
        block := validateBlock(cmd.(map[string]interface{}))
        
        name     := block["command"].(string)
        counts   := convertCounts(block["count"].([]interface{}))
        args     := convertStringArray(block["args"].([]interface{}))
        //includes := convertStringArray(block["include"].([]interface{}))
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

    return task
}

// Detects if there are any SSH remotes
// that need to be setup
func detectSshRemotes(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("ssh")) > 0 {
        return true
    }
    return false
}

// If SSH remotes are detected, they are setup
// here
func setupSshRemotes(json_data *jconfig.Config) {
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

// Remotely runs a command within the SSH session
func runInSshSession(remote string, command string) {
    if (remote != "") {
        StartSession(remote)
        RunRemoteCommand(remote, command)
    }
}

// Detects if there are any config files that
// should be extended
func detectExtend(json_data *jconfig.Config) bool {
    if len(json_data.GetString("extend")) > 0 {
        return true
    }
    return false
}

// Runs extended file if is is found
func setupExtend(json_data *jconfig.Config) {
    extend    := json_data.GetString("extend")
    file_name := extend + ".json"
    if _, err := os.Stat(file_name); err == nil {
        contents_bytes, _ := ioutil.ReadFile(file_name)
        contents := string(contents_bytes)
        initAsinkWithString(contents)
    }
}

// Detects if there are any config files that
// should be included
func detectIncludes(json_data *jconfig.Config) bool {
    if len(json_data.GetArray("include")) > 0 {
        return true
    }
    return false
}

// Sets up asink to run the included tasks
// if they are needed
func setupIncludes(json_data *jconfig.Config) {
    
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
