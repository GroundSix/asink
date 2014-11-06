package main

import (
	"log"
	"github.com/asink/go.crypto/ssh"
)

type Remote struct {
    Name     string
    Host     string
    Port     string
    User     string
    Password string
    Key      ssh.Signer
}

var remotes  map[string]Remote        = make(map[string]Remote)
var sessions map[string]*ssh.Session   = make(map[string]*ssh.Session)
var connections map[string]*ssh.Client = make(map[string]*ssh.Client)

// Inits remotes and sessions then
// returns a new instance of remote
func NewRemote() Remote {
    return Remote{}
}

func (r Remote) Add(remoteName string) {
	remotes[remoteName] = r
}

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

