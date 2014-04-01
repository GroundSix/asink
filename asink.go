package main

import (
    "./vendor/jconfig"
    "fmt"
    "log"
    "os"
    "os/exec"
    "sync"
)

type Command struct {
    name          string
    asyncCount    float64
    relativeCount float64
    args          []string
}

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

func getConfigFile() string {
    if len(os.Args) > 1 {
        filePath := os.Args[1]
        if _, err := os.Stat(filePath); err == nil {
            return filePath
        }
    }
    return ""
}

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