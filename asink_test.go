package asink

import (
	"testing"
	"./asink"
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

func TestSetupCommand(t *testing.T) {
	command       := "echo"
	asyncCount    := float64(2)
	relativeCount := float64(2)
	args 		  := []string{"test"}

	argsInterface := make([]interface{}, len(args))
	for i, v := range args {
	    argsInterface[i] = interface{}(v)
	}

	if (asink.SetupCommand(command, asyncCount, relativeCount, argsInterface) != true) {
		t.Error("Expected bool (true)")
	}
}
