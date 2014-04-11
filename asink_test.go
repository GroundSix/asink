package asink

import (
    "testing"
    "./asink"
)

func TestExecuteCommand(t *testing.T) {
    command := "echo"
    args    := []string{"test"}

    if (asink.ExecuteCommand(command, args, 2, 2) != true) {
        t.Error("Expected bool (true)")
    }
}

func TestExecute(t *testing.T) {
    command       := "echo"
    asyncCount    := float64(2)
    relativeCount := float64(2)
    args          := []string{"test"}

    argsInterface := make([]interface{}, len(args))
    for i, v := range args {
        argsInterface[i] = interface{}(v)
    }

    if (asink.Execute(command, asyncCount, relativeCount, argsInterface) != true) {
        t.Error("Expected bool (true)")
    }
}
