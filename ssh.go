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
    "fmt"
	"log"
    "io/ioutil"
    "./vendor/color"
	"./vendor/go.crypto/ssh"
    "os"
    "os/user"
    "strings"
)

type Remote struct {
    Name     string
    Host     string
    Port     string
    User     string
    Password string
    Key      ssh.Signer
}

var remotes  map[string]*Remote        = nil
var sessions map[string]*ssh.Session   = nil
var connections map[string]*ssh.Client = nil

// Inits remotes and sessions then
// returns a new instance of remote
func NewRemote() *Remote {
    remotes  = make(map[string]*Remote)
    sessions = make(map[string]*ssh.Session)
    connections = make(map[string]*ssh.Client)
    return new(Remote)
}

// Adds a new remote to the map of
// remotes with a string key
func (r *Remote) AddRemote(name string, host string, port string, user string, password string) {
    remote := new(Remote)

    remote.Name     = name
    remote.Host     = host
    remote.Port     = port
    remote.User     = user
    remote.Password = password

    remotes[name] = remote
}

// Corrects a ~ with the users home directory
// This is similar to above to needs to be
// abstracted
func validateKeyPath(key_path string) string {
    return strings.Replace(key_path, "~", getHomeDirectory(), -1)
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

// Parses then adds the key to our remote struct
func (r *Remote) AddSshKey(name string, file string) {
    remote    := remotes[name]
    remote.Key = parseKey(validateKeyPath(file))
}

// Starts a new SSH session which is then
// stored in the sessions map with a given
// key
func StartSession(name string) {
    remote := remotes[name]

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

    connections[name] = conn
    sessions[name]    = session
}

// Runs the remote command given the session
// key
func RunRemoteCommand(name string, command string) {
    session  := sessions[name]

    format := color.New(color.FgCyan).SprintFunc()
    fmt.Printf("%s ", format("> " + name + ":"))

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
