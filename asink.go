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

    app.Commands = []cli.Command{
        {
            Name:  "run",
            Usage: "[tasks.yml] pass through your tasks file",
            Action: func(c *cli.Context) {
                conn := Connection{}
                if c.IsSet("r") == false {
                    if err := initAsinkWithFile(os.Args[2]); err != nil {
                        fmt.Println(err); os.Exit(1)
                    }
                } else {
                    conn.remote = c.String("r")
                    if c.IsSet("i") {
                        conn.privateKeyPath = c.String("i")
                    }
                    contents, err := ioutil.ReadFile(os.Args[2])
                    if err != nil {
                        panic(err)
                    }
                    conn.loadPrivateKey()
                    conn.signRequest(contents)
                    conn.makeRequest()

                }
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
            Name:  "server",
            Usage: "starts up a small server listening on port 3000",
            Action: func(c *cli.Context) {
                s := Server{}
                s.Port = "3000"
                if c.IsSet("a") {
                    s.AuthorizedKeysPath = c.String("a")
                }
                if c.IsSet("p") {
                    s.Port = c.String("p")
                }
                s.Start()
            },
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "port, p",
                    Usage: "port for asink's server to listen on",
                },
                cli.StringFlag{
                    Name:  "authorized-keys, a",
                    Usage: "path to asink's authorized_keys file",
                },
            },
        },
        {
            Name:  "keygen",
            Usage: "generates public/private key pair",
            Action: func(c *cli.Context) {
                hd, err := homedir.Dir()
                if err != nil {
                    panic(err)
                }
                p := hd + "/.asink"
                if c.IsSet("d") {
                    p = c.String("d")
                }
                k := newKeys(p)
                k.generate()
                if k.exists() {

                }

            },
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "directory, d",
                    Usage: "path to create public/private key pair to use with asink",
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
