package asink

import (
    "testing"
    "./asink"
)

func TestExecuteCommand(t *testing.T) {
    Asink := asink.New()

    command := "echo"
    args    := []string{"test"}

    if (Asink.ExecuteCommand(command, args, 2, 2, true) != true) {
        t.Error("Expected bool (true)")
    }
}

func TestExecute(t *testing.T) {
    Asink := asink.New()

    Asink.SetName("echo")
    Asink.SetAsyncCount(2)
    Asink.SetRelativeCount(2)
    Asink.SetArgs([]string{"test"})

    if (Asink.Execute() != true) {
        t.Error("Expected bool (true)")
    }
}
