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
    "fmt"
	"log"
	"./vendor/go.crypto/ssh"
)

type Remote struct {
    Name     string
    Host     string
    Port     string
    User     string
    Password string
}

var remotes  map[string]*Remote = nil
var sessions map[string]*ssh.Session = nil

func NewRemote() *Remote {
    remotes = make(map[string]*Remote)
    sessions = make(map[string]*ssh.Session)
    return new(Remote)
}

func (r *Remote) AddRemote(name string, host string, port string, user string, password string) {
    remote := new(Remote)

    remote.Name     = name
    remote.Host     = host
    remote.Port     = port
    remote.User     = user
    remote.Password = password

    remotes[name] = remote
}

func StartSession(name string) {
    remote := remotes[name]

    config := &ssh.ClientConfig{
        User: remote.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(remote.Password),
        },
    }

    hostname := remote.Host + ":" + remote.Port

    conn, err := ssh.Dial("tcp", hostname, config)
    if err != nil {
        log.Fatalf("unable to connect: %s", err)
    }

   //defer conn.Close()
    // Create a session
    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }
    //defer session.Close()
    // Set up terminal modes
    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    // Request pseudo terminal
    if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
        log.Fatalf("request for pseudo terminal failed: %s", err)
    }

    sessions[name] = session
}

func RunRemoteCommand(name string, command string) {
    session  := sessions[name]
    res, err := session.Output(command);

    if (err == nil) {
        s := string(res[:])
        fmt.Println(s)
    }
}
