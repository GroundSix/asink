// asink v0.1.1-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.1.1-dev
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
    "fmt"
    "./asink"
    "io/ioutil"
    "github.com/asink/cli"
)

// Application entry point
func main() {
    app := cli.NewApp()

    // constants located in meta.go
    app.Name    = appName
    app.Version = version
    app.Usage   = usage
    app.Author  = author
    app.Email   = email

    app.Commands = []cli.Command{
        {
            Name: "start",
            Usage: "<tasks.yml> pass through your tasks file",
            Action: func (c *cli.Context) {
                initAsinkWithFile(os.Args)
            },
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "remote, r",
                    Usage: "remote to connect to",
                },
                cli.StringFlag{
                    Name:  "identity-file, i",
                    Usage: "path to private key",
                },
            },
        },
        {
            Name: "server",
            Usage: "starts up a small server listening on port 3000",
            Action: func (c *cli.Context) {
                startServer()
            },
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "authorized-keys, a",
                    Usage: "path to asink's authorized_keys file",
                },
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
func initAsinkWithFile(args []string) {
    if validateArguments(args) == true {
        fileName := args[2]
        p := createParserFromFileType(fileName)

        contents, err := ioutil.ReadFile(fileName)
        if (err != nil) {
            panic(err)
        }
        p = p.parse(contents)

        initAsink(p)
    }
}

// Inits Asink over the HTTP, will accept JSON
// only
func initAsinkWithRequest(request []byte) {
    p := createJsonParser()
    p = p.parse(request)

    initAsink(p)
}

// A rough way of validating args passed through
// when using Asink via CLI
func validateArguments(args []string) bool {
    if len(args) == 0 {
        fmt.Println("Arguments needed, 0 passed")
        fmt.Println("Use 'asink help' to see list of available commands")
        return false
    }
    return true
}
