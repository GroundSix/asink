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
	"io/ioutil"
	"log"
	"os"
	"os/user"
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
			Name:  "start",
			Usage: "[tasks.yml] pass through your tasks file",
			Action: func(c *cli.Context) {
				conn := Connection{}
				if c.IsSet("r") == false {
					initAsinkWithFile(os.Args)
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
				usr, err := user.Current()
				if err != nil {
					log.Fatal(err)
				}
				p := usr.HomeDir + "/.asink"
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
func initAsinkWithFile(args []string) {
	if validateArguments(args) == true {
		fileName := args[2]
		p := createParserFromFileType(fileName)

		contents, err := ioutil.ReadFile(fileName)
		if err != nil {
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
