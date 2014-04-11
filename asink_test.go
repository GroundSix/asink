package asink

import (
    "testing"
    "./asink"
)

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
