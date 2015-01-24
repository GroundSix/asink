// asink v0.1.1-dev
//
// (c) Ground Six 2015
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: GNU GPL v2.0
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
    "fmt"
    "github.com/asink/cli"
    "github.com/asink/libasink"
    "github.com/asink/go-homedir"
    "io/ioutil"
    "os"
)

// Application entry point
func main() {
    app := cli.NewApp()

    // constants located in meta.go
    app.Name = appName
    app.Version = version
    app.Usage = usage
    app.Author = author
    app.Email = email

    setEnvVars()

    app.Commands = []cli.Command{
        {
            Name:  "run",
            Usage: "[tasks.yml] pass through your tasks file",
            Action: func(c *cli.Context) {                
                if err := initAsinkWithFile(os.Args[2]); err != nil {
                    fmt.Println(err); os.Exit(1)
                }
            },
        },
    }
    app.Run(os.Args)
}

// Creates the assigner object to deal with
// the parsed instructions for execution
func initAsink(p Parser) {
    a := new(Assigner)
    a.TaskMap = p.TaskMap()
    a.assignTasks()
    a.assignRemotes()
    asink.ExecMulti(a.Tasks())
    defer closeSshSessions()
}

// Inits Asink using a file, either JSON or YAML
// A parser is generated based on the file
// extension
func initAsinkWithFile(filename string) error {
    // Loads file contents of task file
    c, err := loadTasksFile(filename)
    if err != nil {
        return err
    }
    p, err := parserFromFileType(filename)
    if err != nil {
        return err
    }
    p = p.parse(c)

    initAsink(p)
    return nil
}

// Inits Asink over the HTTP, will accept JSON
// only
func initAsinkWithRequest(request []byte) {
    p := createJsonParser()
    p = p.parse(request)

    initAsink(p)
}

// Loads task file from fs to be used to run Asink
func loadTasksFile(filename string) ([]byte, error) {
    c, err := ioutil.ReadFile(filename)
    if err != nil {
        return []byte{}, fmt.Errorf("File '%s' could not be found", filename)
    }
    return c, nil
}

func setEnvVars() {
    hd, _ := homedir.Dir()
    os.Setenv("HOME", hd)
    os.Setenv("PWD", getWorkingDirectory())
}
