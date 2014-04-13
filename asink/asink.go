/**
 * asink v0.1-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.1-dev
 * 
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package asink

import (
    "fmt"
    "log"
    "os/exec"
    "sync"
)

/**
 * @var string name of the command
 * @var float64 number of async iterations
 * @var float64 number of sync iterations
 * @var []string command arguments
 */
type Command struct {
    name          string
    asyncCount    float64
    relativeCount float64
    args          []string
    output        bool
}

/**
 * Acting as a constructor to
 * return a new instance of
 * Command
 *
 * @return *Command
 */
func New() *Command {
    command := new(Command)
    command.SetOutput(false)

    return command
}

/**
 * Allows you to choose whether or
 * not to show the output for
 * each command that is ran
 *
 * @param bool output switch
 *
 * @return nil
 */
func (c *Command) SetOutput(output bool) {
    c.output = output
}

/**
 * An alias function that nicely
 * interfaces the Execute function
 * to make it easier to use in external Go
 * programs
 *
 * @param string the command name
 * @param []string command arguments
 * @param int number of async iterations
 * @param int number of sync iterations
 *
 * @return bool
 */
func (c *Command) ExecuteCommand(command string, args []string, asyncCount int, relativeCount int) bool {
    argsInterface := make([]interface{}, len(args))
    for i, v := range args {
        argsInterface[i] = interface{}(v)
    }
    Asink := new(Command)
    return Asink.Execute(command, float64(asyncCount), float64(relativeCount), argsInterface)
}

/**
 * Creates the command channel and sets
 * up everything ready for execution
 *
 * @param string the command name
 * @param float64 number of async iterations
 * @param float64 number of sync iterations
 * @param []interface{} command arguments
 *
 * @return bool
 */
func (c *Command) Execute(command string, asyncCount float64, relativeCount float64, args []interface{}) bool {
    commandChan := make(chan *Command)

    var wg sync.WaitGroup

    c.name = command
    c.asyncCount = asyncCount
    c.relativeCount = relativeCount

    argsSlice := make([]string, len(args))

    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    c.args = argsSlice

    for i := 0; i != int(asyncCount); i++ {
        wg.Add(1)
        go runConcurrently(commandChan, &wg)
        commandChan <- c
    }

    close(commandChan)
    wg.Wait()

    return true
}

/**
 * Executes command a given amount
 * of times as specefied in the
 * JSON configuration file
 *
 * @param Command in instance of Command struct
 * @param WaitGroup our async wait group for the channel
 *
 * @return nil
 */
func runConcurrently(command chan *Command, wg *sync.WaitGroup) {
    defer wg.Done()

    commandData := <-command

    for c := 0; c != int(commandData.relativeCount); c++ {
        out, err := exec.Command(commandData.name, commandData.args...).Output()
        if err != nil {
            log.Fatal(err)
        }
        if (commandData.output == true) {
            fmt.Printf("%s\n", out)
        }
    }
}
