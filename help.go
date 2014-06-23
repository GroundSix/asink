/**
 * asink v0.0.2-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.0.2-dev
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
    "./vendor/cobra"
)

/**
 * Prompts the cobra help screen with
 * a list of all available commands
 *
 * @return nil
 */
func executeRootCommand() {
    var startCommand = &cobra.Command{
        Use:   "start [JSON configuration file]",
        Short: "Start your asink processes",
        Long:  `start running a command the specified amount of times from your configuration file`,
        Run: func(cmd *cobra.Command, args []string) {
            initAsink()
        },
    }

    var rootCmd = &cobra.Command{Use: "asink"}
    rootCmd.AddCommand(startCommand)
    rootCmd.Execute()
}
