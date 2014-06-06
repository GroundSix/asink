/**
 * asink v0.1-dev
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
    //"strings"
    //"encoding/json"
    "./asink"
    "./vendor/pb"
    "./vendor/jconfig"
    //"./vendor/jsonq"
)

/**
 * @var *pb.ProgressBar asink's progress indicator
 */
var progressBar *pb.ProgressBar = nil

/**
 * Entry point for asink. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
    configFile := asink.GetConfigFile()
    if configFile != "" {
        json := jconfig.LoadConfig(configFile)
        if (detectTasks(json) == true) {
            tasks := setupAsinkTasks(json)
            tasks.Execute()
        } else {
            command := setupAsinkCommand(json)
            command.ListenForInit(createProgressBar)
            command.ListenForProgress(incrementProgressBar)
            command.ListenForFinish(endProgressBar)
            command.Execute()
        }
    }
}

/**
 * Creates the progress bar on the
 * listen init event
 *
 * @param int number of commands
 *
 * @return nil
 */
func createProgressBar(count int) {
    progressBar = pb.StartNew(count)
}

/**
 * Increments the progress bar
 * by one on the listen progress
 * event
 *
 * @return nil
 */
func incrementProgressBar() {
    progressBar.Increment()
}

/**
 * Stops the progress bar on the
 * listen finish event
 *
 * @return nil
 */
func endProgressBar() {
    progressBar.FinishPrint("Finished.")
}

/**
 * Initially sets up everything from
 * config file in a new instance of Asink
 *
 * @param string path to config file
 *
 * @return *asink.Command configured instance of
 * asink
 */
func setupAsinkCommand(json_data *jconfig.Config) *asink.Command {
    command := asink.New()

    counts := convertCounts(json_data.GetArray("count"))
    args   := convertArgs(json_data.GetArray("args"))

    command.Name = json_data.GetString("command")
    command.AsyncCount = counts[0]
    command.RelativeCount = counts[1]
    command.Args = args
    command.Output = json_data.GetBool("output")

    return command
}

/**
 * Converts jconfigs []interface into
 * []string for asink
 *
 * @param []interface{} jconfig's array
 *
 * @return []string asink's array
 */
func convertArgs(args []interface{}) []string {
    argsSlice := make([]string, len(args))
    for i, s := range args {
        argsSlice[i] = s.(string)
    }

    return argsSlice
}

/**
 * Converts jconfigs []interface into
 * []float64 for asink
 *
 * @param []interface{} jconfig's array
 *
 * @return []float64 asink's array
 */
func convertCounts(counts []interface{}) []float64 {
    argsSlice := make([]float64, len(counts))
    for i, s := range counts {
        argsSlice[i] = s.(float64)
    }

    return argsSlice
}

func detectTasks(json_data *jconfig.Config) bool {
    if len(json_data.GetStringMap("tasks")) > 0 {
        return true
    }
    return false
}

func setupAsinkTasks(json_data *jconfig.Config) *asink.Task {
    //tasks   := map[string]interface{}{}
    //decoder := json.NewDecoder
    tasks := asink.NewTask()

    return tasks
}
