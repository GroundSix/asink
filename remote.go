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
	"log"
	"io/ioutil"
	"github.com/asink/go.crypto/ssh"
	"github.com/asink/color"
)

type Remote struct {
    Name     string
    Host     string
    Port     string
    User     string
    Password string
    Key      ssh.Signer
}

var remotes  map[string]Remote         = make(map[string]Remote)
var sessions map[string]*ssh.Session   = make(map[string]*ssh.Session)
var connections map[string]*ssh.Client = make(map[string]*ssh.Client)

// Inits remotes and sessions then
// returns a new instance of remote
func NewRemote(name string) Remote {
    r := Remote{}
    r.Name = name
    return r
}

// Adds a new remote to the remote map
func (r Remote) Add(remoteName string) {
	remotes[remoteName] = r
}

// Parses then adds the key to our remote struct
func (r Remote) AddSshKey(remoteName string, filePath string) {
    remote := remotes[remoteName]
    remote.Key = parseKey(filePath)
}

// Makes a connection to the remote machine
func (r Remote) Connect(remoteName string) {
    remote := remotes[remoteName]

    config := &ssh.ClientConfig{
        User: remote.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(remote.Password),
            ssh.PublicKeys(remote.Key),
        },
    }

    hostname := remote.Host + ":" + remote.Port

    conn, err := ssh.Dial("tcp", hostname, config)
    if err != nil {
        log.Fatalf("unable to connect: %s", err)
    }

    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,
        ssh.TTY_OP_ISPEED: 14400,
        ssh.TTY_OP_OSPEED: 14400,
    }

    if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
        log.Fatalf("request for pseudo terminal failed: %s", err)
    }

    connections[remoteName] = conn
    sessions[remoteName]    = session
}

// Runs the remote command given the session
// key
func runRemoteCommand(remoteName string, command string) {
    session  := sessions[remoteName]

    format := color.New(color.FgCyan).SprintFunc()
    fmt.Printf("%s ", format("$" + remoteName + ":"))

    session.Stdout = os.Stdout
    session.Stderr = os.Stderr

    err  := session.Run(command)

    if (err != nil) {
        fmt.Println("Failed to run:", command)
    }
}

// Parses the key for the client so
// we can SSH into the remote
func parseKey(file string) ssh.Signer {
    privateBytes, err := ioutil.ReadFile(file)
    if err != nil {
        panic("Failed to load private key")
    }

    private, err := ssh.ParsePrivateKey(privateBytes)
    if err != nil {
        panic("Failed to parse private key")
    }
    return private
}

// Closes all SSH sessions and connections
func closeSshSessions() {
    for _, session := range sessions {
        session.Close()
    }
    for _, connection := range connections {
        connection.Close()
    }
}
