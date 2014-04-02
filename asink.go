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

package main

import (
    "./vendor/jconfig"
    "fmt"
    "log"
    "os"
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
}

/**
 * Entry point for asink. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
    configFile := getConfigFile()
    if configFile != "" {
        config := jconfig.LoadConfig(getConfigFile())
        command := config.GetString("command")
        counts := config.GetArray("count")
        args := config.GetArray("args")

        setupCommand(command, counts[0].(float64), counts[1].(float64), args)
    }
}

/**
 * Gets the name of your config file
 * from the param passed through when
 * the program is ran
 *
 * e.g. asink config.json
 *
 * @return string file path or empty string
 */
func getConfigFile() string {
    if len(os.Args) > 1 {
        filePath := os.Args[1]
        if _, err := os.Stat(filePath); err == nil {
            return filePath
        }
    }
    return ""
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
 * @return nil
 */
func setupCommand(command string, asyncCount float64, relativeCount float64, args []interface{}) {
    commandChan := make(chan *Command)
    commandStruct := new(Command)

    var wg sync.WaitGroup

    commandStruct.name = command
    commandStruct.asyncCount = asyncCount
    commandStruct.relativeCount = relativeCount

    argsSlice := make([]string, len(args))

    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    commandStruct.args = argsSlice

    for i := 0; i != int(asyncCount); i++ {
        wg.Add(1)
        go executeCommand(commandChan, &wg)
        commandChan <- commandStruct
    }

    close(commandChan)
    wg.Wait()
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
func executeCommand(command chan *Command, wg *sync.WaitGroup) {
    defer wg.Done()

    commandData := <-command

    for c := 0; c != int(commandData.relativeCount); c++ {
        out, err := exec.Command(commandData.name, commandData.args...).Output()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s\n", out)
    }
}
