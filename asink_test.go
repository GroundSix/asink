package asink

import (
    "testing"
    "./asink"
)

func TestSetupCommand(t *testing.T) {
    command       := "echo"
    asyncCount    := float64(2)
    relativeCount := float64(2)
    args          := []string{"test"}

    argsInterface := make([]interface{}, len(args))
    for i, v := range args {
        argsInterface[i] = interface{}(v)
    }

    if (asink.SetupCommand(command, asyncCount, relativeCount, argsInterface) != true) {
        t.Error("Expected bool (true)")
    }
}
