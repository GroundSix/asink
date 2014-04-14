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
 * Sets the command name
 *
 * @param string command name
 *
 * @return nil
 */
func (c *Command) SetName(name string) {
    c.name = name
}

/**
 * Sets the command args
 *
 * @param []string command args
 *
 * @return nil
 */
func (c *Command) SetArgs(args []string) {
    c.args = args
}

/**
 * Sets the command async count
 *
 * @param int command async count
 *
 * @return nil
 */
func (c *Command) SetAsyncCount(asyncCount int) {
    c.asyncCount = float64(asyncCount)
}

/**
 * Sets the command relative count
 *
 * @param int command relative count
 *
 * @return nil
 */
func (c *Command) SetRelativeCount(relativeCount int) {
    c.relativeCount = float64(relativeCount)
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
 * @param bool flag to show command output
 *
 * @return bool
 */
func (c *Command) ExecuteCommand(name string, args []string, asyncCount int, relativeCount int, output bool) bool {
    Asink := new(Command)
    Asink.SetName(name)
    Asink.SetAsyncCount(asyncCount)
    Asink.SetRelativeCount(relativeCount)
    Asink.SetArgs(args)
    Asink.SetOutput(output)

    return Asink.Execute()
}

/**
 * Creates the command channel and sets
 * up everything ready for execution
 *
 * @return bool
 */
func (c *Command) Execute() bool {
    commandChan := make(chan *Command)

    var wg sync.WaitGroup

    for i := 0; i != int(c.asyncCount); i++ {
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
