package main

import (
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

var remotes  map[string]*Remote        = make(map[string]*Remote)
var sessions map[string]*ssh.Session   = make(map[string]*ssh.Session)
var connections map[string]*ssh.Client = make(map[string]*ssh.Client)

// Inits remotes and sessions then
// returns a new instance of remote
func NewRemote() Remote {
    return Remote{}
}

