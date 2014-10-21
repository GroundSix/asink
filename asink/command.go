package asink

import (
	"os"
	"os/exec"
	"sync"
)

type Command struct {
	Name 	   string
	AsyncCount int
	RelCount   int
	Dir 	   string
	Args 	   []string
}

func NewCommand(name string) Command {
	return Command{name, 1, 1, getWorkingDirectory(), []string{}}
}

func (c Command) Exec() {
	var wg sync.WaitGroup

	command := make(chan Command)

	validateDirectoryName(&c)
	os.Chdir(getWorkingDirectory())
	os.Chdir(c.Dir)

    for i := 0; i != c.AsyncCount; i++ {
        wg.Add(1)
        go runCommand(command, &wg)
        command <- c
    }

    close(command)
    wg.Wait()
}

// Executes command a given amount
// of times as specefied in the
// JSON configuration file
func runCommand(command chan Command, wg *sync.WaitGroup) {
    defer wg.Done()

    c := <- command

    for j := 0; j != c.RelCount; j++ {
        cmd := exec.Command(c.Name, c.Args...)

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()
	}
}
