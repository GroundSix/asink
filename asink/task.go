/**
 * asink v0.0.2-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.1-dev
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package asink

import ()

type Task struct {
	Command *Command
	Then    string
}

/**
 * Creates a new instance of the Task
 * struct
 *
 * @return *Task a new task
 */
func NewTask() *Task {
	return new(Task)
}

func (t *Task) Execute() {
	
}
