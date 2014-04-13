package asink

import (
    "testing"
    "./asink"
)

func TestExecuteCommand(t *testing.T) {
    Asink := asink.New()

    command := "echo"
    args    := []string{"test"}

    if (Asink.ExecuteCommand(command, args, 2, 2) != true) {
        t.Error("Expected bool (true)")
    }
}

func TestExecute(t *testing.T) {
    Asink := asink.New()
    
    command       := "echo"
    asyncCount    := float64(2)
    relativeCount := float64(2)
    args          := []string{"test"}

    argsInterface := make([]interface{}, len(args))
    for i, v := range args {
        argsInterface[i] = interface{}(v)
    }

    if (Asink.Execute(command, asyncCount, relativeCount, argsInterface) != true) {
        t.Error("Expected bool (true)")
    }
}
