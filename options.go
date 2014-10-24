// asink v0.1.1-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.1.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"fmt"
	"github.com/asink/cobra"
)

// Creates the root command and sub commands
// for asink to start
func createRootCommand() {
	var rootCmd = &cobra.Command{Use: "asink"}
	rootCmd.AddCommand(createVersionCommand())
	rootCmd.AddCommand(createStartCommand())
	rootCmd.AddCommand(createGetCommand())
	rootCmd.AddCommand(createServerCommand())
	rootCmd.Execute()
}

// Creates the version sub command to display
// which version of asink is being used
 func createVersionCommand() *cobra.Command {
    var versionCommand = &cobra.Command{
        Use:   "version",
        Short: "Shows asink version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Asink version 0.1.1-dev")
            fmt.Println("Created by Ground Six")
        },
    }
    return versionCommand
}

// Creates the start sub command which uses a
// JSON or YAML file to start asink
func createStartCommand() *cobra.Command {
    var startCommand = &cobra.Command{
        Use:   "start  [JSON config file]",
        Short: "Start your asink processes",
        Long:  `start running a command the specified amount of times from your configuration file`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsinkWithFile(args)
        },
    }
    return startCommand
}

// Creates the get sub command used for remote
// configuration files
func createGetCommand() *cobra.Command {
    var getCommand = &cobra.Command{
        Use:   "get    [config URL]",
        Short: "Start asink using remote configuration",
        Long:  `use an external / remote configuration file to start asink rather than one on your file system`,
        Run: func(cmd *cobra.Command, args []string) {
            //initAsinkWithFile(args)
        },
    }
    return getCommand
}

// Creates the server command for starting up
// asink as a small internal server
func createServerCommand() *cobra.Command {
    var serverCommand = &cobra.Command{
        Use:   "server",
        Short: "Starts a small http server",
        Long:  `a small server can be used to interface asink by sending JSON in the request body`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsinkWithServer(args)
        },
    }
    return serverCommand
}
