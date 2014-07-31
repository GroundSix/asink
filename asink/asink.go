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

package asink

import (
    "log"
    "os"
    "os/exec"
    "os/user"
    "strings"
    "sync"
)

var initial_directory string = ""

type Command struct {
    Name          string
    AsyncCount    float64
    RelativeCount float64
    Args          []string
    Dir           string
    Manual        bool

    progressInit   func(count int)
    progressAdd    func()
    progressEnd    func()
    manualCallback func(command string)
}

// Acting as a constructor to
// return a new instance of
// Command
func New() *Command {
    command := new(Command)
    command.Name = ""
    command.AsyncCount = 0
    command.Args = []string{}
    command.RelativeCount = 0
    command.Dir = getWorkingDirectory()
    command.Manual = false

    command.progressInit = func(count int) {}
    command.progressAdd = func() {}
    command.progressEnd = func() {}
    command.manualCallback = func(command string) {}

    initial_directory = getWorkingDirectory()

    return command
}

// An optional callback public function
// that gets called on job creation
func (c *Command) ListenForInit(callback func(count int)) {
    c.progressInit = callback
}

// An optional callback public function
// that gets called everytime a job
// has finished
func (c *Command) ListenForProgress(callback func()) {
    c.progressAdd = callback
}

// An optional callback public function
// that gets called when all jobs have
// been finished
func (c *Command) ListenForFinish(callback func()) {
    c.progressEnd = callback
}

func (c *Command) SetManualCallback(callback func(command string)) {
    c.manualCallback = callback
}

// An alias function that nicely
// interfaces the Execute function
// to make it easier to use in external Go
// programs
func (c *Command) ExecuteCommand(name string, args []string, asyncCount int, relativeCount int) bool {
    Asink := new(Command)
    Asink.Name = name
    Asink.AsyncCount = float64(asyncCount)
    Asink.RelativeCount = float64(relativeCount)
    Asink.Args = args

    // Default all callbacks
    Asink.ListenForInit(func(count int) {})
    Asink.ListenForProgress(func() {})
    Asink.ListenForFinish(func() {})

    return Asink.Execute()
}

// Creates the command channel and sets
// up everything ready for execution
func (c *Command) Execute() bool {
    commandChan := make(chan *Command)

    var wg sync.WaitGroup

    c.progressInit(int(c.AsyncCount * c.RelativeCount))
    c = validateDirectoryName(c)

    // Reset to initial directory and then move
    // to the new one
    os.Chdir(initial_directory)
    os.Chdir(c.Dir)

    for i := 0; i != int(c.AsyncCount); i++ {
        wg.Add(1)
        go runConcurrently(commandChan, &wg)
        commandChan <- c
    }

    close(commandChan)
    wg.Wait()

    c.progressEnd()

    return true
}

// Returns the current working directory
// as a string
func getWorkingDirectory() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return dir
}

// Returns the current user's home directory
// as a string
func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    return usr.HomeDir
}

// Corrects a ~ with the users home directory
func validateDirectoryName(command *Command) *Command {
    if command.Manual == false {
        command.Dir = strings.Replace(command.Dir, "~", getHomeDirectory(), -1)
    }
    return command
}

// Generates a full command as a string for manual
// use
func generateCommandWithDirectory(command string, args []string, directory string) string {
    if directory == initial_directory {
        directory = "."
    }
    full_command := "cd " + directory + " && "
    full_command = full_command + command + " "
    full_command = full_command + strings.Join(args, " ")

    return full_command
}

// Executes command a given amount
// of times as specefied in the
// JSON configuration file
func runConcurrently(command chan *Command, wg *sync.WaitGroup) {
    defer wg.Done()

    commandData := <-command

    for c := 0; c != int(commandData.RelativeCount); c++ {
        if commandData.Manual == true {
            full_command := generateCommandWithDirectory(commandData.Name, commandData.Args, commandData.Dir)
            commandData.manualCallback(full_command)
        } else {
            cmd := exec.Command(commandData.Name, commandData.Args...)

            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
        }
        commandData.progressAdd()
    }
}
