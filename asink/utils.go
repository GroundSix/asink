package asink

import (
	"os"
	"os/user"
	"strings"
)

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
        panic(err)
    }
    return usr.HomeDir
}

// Corrects a ~ with the users home directory
func validateDirectoryName(c *Command) {
	c.Dir = strings.Replace(c.Dir, "~", getHomeDirectory(), -1)
}
