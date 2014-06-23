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
    Name          string
    AsyncCount    float64
    RelativeCount float64
    Args          []string
    Output        bool
    Manual        bool

    progressInit  func(count int)
    progressAdd   func()
    progressEnd   func()
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
    command.Name          = ""
    command.AsyncCount    = 0
    command.Args          = []string{}
    command.RelativeCount = 0
    command.Output        = false
    command.Manual        = false

    command.progressInit = func(count int){}
    command.progressAdd  = func(){}
    command.progressEnd  = func(){}

    return command
}

/**
 * An optional callback public function
 * that gets called on job creation
 *
 * @param func(count int) callback function
 *
 * @return nil
 */
func (c *Command) ListenForInit(callback func(count int)) {
    c.progressInit = callback
}

/**
 * An optional callback public function
 * that gets called everytime a job
 * has finished
 *
 * @param func() callback function
 *
 * @return nil
 */
func (c *Command) ListenForProgress(callback func()) {
    c.progressAdd = callback
}

/**
 * An optional callback public function
 * that gets called when all jobs have
 * been finished
 *
 * @param func() callback function
 *
 * @return nil
 */
func (c *Command) ListenForFinish(callback func()) {
    c.progressEnd = callback
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
    Asink.Name = name
    Asink.AsyncCount = float64(asyncCount)
    Asink.RelativeCount = float64(relativeCount)
    Asink.Args = args
    Asink.Output = output

    // Default all callbacks
    Asink.ListenForInit(func(count int){})
    Asink.ListenForProgress(func(){})
    Asink.ListenForFinish(func(){})

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

    c.progressInit(int(c.AsyncCount * c.RelativeCount))

    for i := 0; i != int(c.AsyncCount); i++ {
        wg.Add(1)
        go runConcurrently(commandChan, &wg)
        commandChan <- c
    }

    close(commandChan)
    wg.Wait()

    c.progressEnd()

    return true
}

/**
 * Executes command a given amount
 * of times as specefied in the
 * JSON configuration file
 *
 * @param *Command an instance of Command struct
 * @param WaitGroup our async wait group for the channel
 *
 * @return nil
 */
func runConcurrently(command chan *Command, wg *sync.WaitGroup) {
    defer wg.Done()

    commandData := <-command

    for c := 0; c != int(commandData.RelativeCount); c++ {
        if commandData.Manual == true {
            fmt.Println("Hello!")
        } else {
            out, err := exec.Command(commandData.Name, commandData.Args...).Output()
            if err != nil {
                log.Fatal(err)
            }
            if commandData.Output == true {
                fmt.Printf("%s\n", out)
            }
        }
        commandData.progressAdd()
    }
}
