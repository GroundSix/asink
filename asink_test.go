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

    Asink.Name = "echo"
    Asink.AsyncCount = 2
    Asink.RelativeCount = 2
    Asink.Args = []string{"test"}

    if (Asink.Execute() != true) {
        t.Error("Expected bool (true)")
    }
}

func TestExecuteWithCallbacks(t *testing.T) {
    Asink := asink.New()

    Asink.Name = "echo"
    Asink.AsyncCount = 2
    Asink.RelativeCount = 2
    Asink.Args = []string{"test"}

    // Set callback functions
    Asink.ListenForInit(func(count int){})
    Asink.ListenForProgress(func(){})
    Asink.ListenForFinish(func(){})

    if (Asink.Execute() != true) {
        t.Error("Expected bool (true)")
    }
}
