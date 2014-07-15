// asink v0.0.3-dev
//
// (c) Ground Six
//
// @package asink
// @version 0.0.3-dev
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
    "./vendor/cobra"
)

// Prompts the cobra help screen with
// a list of all available commands
func executeRootCommand() {
    var rootCmd = &cobra.Command{Use: "asink"}
    rootCmd.AddCommand(cobraVersionCommand())
    rootCmd.AddCommand(cobraStartCommand())
    rootCmd.AddCommand(cobraGetCommand())
    rootCmd.AddCommand(cobraServerCommand())
    rootCmd.Execute()
}

/**
 * Returns the 'version' subcommand for asink
 *
 * @return *cobra.Command
 */
 func cobraVersionCommand() *cobra.Command {
    var versionCommand = &cobra.Command{
        Use:   "version",
        Short: "Shows asink version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Asink version 0.0.3-dev - Created by Ground Six")
        },
    }

    return versionCommand
}

/**
 * Returns the 'start' subcommand for asink
 *
 * @return *cobra.Command
 */
 func cobraStartCommand() *cobra.Command {
    var startCommand = &cobra.Command{
        Use:   "start  [JSON config file]",
        Short: "Start your asink processes",
        Long:  `start running a command the specified amount of times from your configuration file`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsink()
        },
    }

    return startCommand
}

/**
 * Returns the 'get' subcommand for asink
 *
 * @return *cobra.Command
 */
func cobraGetCommand() *cobra.Command {
    var getCommand = &cobra.Command{
        Use:   "get    [config URL]",
        Short: "Start asink using remote configuration",
        Long:  `use an external / remote configuration file to start asink rather than one on your file system`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsinkWithHttp(args)
        },
    }

    return getCommand
}

/**
 * Returns the 'server' subcommand for asink
 *
 * @return *cobra.Command
 */
func cobraServerCommand() *cobra.Command {
    var serverCommand = &cobra.Command{
        Use:   "server",
        Short: "Starts a small http server",
        Long:  `a small server can be used to interface asink by sending JSON in the request body`,
        Run: func(cmd *cobra.Command, args []string) {
            startServer()
        },
    }

    return serverCommand
}
