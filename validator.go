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

// Validates all keys for tasks and returns
// default values
func validateBlock(block map[string]interface{}) map[string]interface{} {
	finalBlock := validateName(block)
	finalBlock  = validateCount(finalBlock)
	finalBlock  = validateArgs(finalBlock)
	finalBlock  = validateOutput(finalBlock)
	finalBlock  = validateRequire(finalBlock)
	finalBlock  = validateGroup(finalBlock)
	finalBlock  = validateDir(finalBlock)
	finalBlock  = validateRemote(finalBlock)
	finalBlock  = validateSshPassword(finalBlock)
	finalBlock  = validateSshKey(finalBlock)
	finalBlock  = validateInclude(finalBlock)

	return finalBlock
}

// Validates and defaults the command name
func validateName(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["command"]; !ok {
	    block["command"] = ""
	}
	return block
}

// Validates and defaults the command count
func validateCount(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["count"]; !ok {
		var defaults []interface{}
		var single float64 = 1

		defaults = append(defaults, single)
		defaults = append(defaults, single)

	    block["count"] = defaults
	}
	return block
}

// Validates and defaults the command arguments
func validateArgs(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["args"]; !ok {
		var defaults []interface{}
		block["args"] = defaults
	}
	return block
}

// Validates and defaults the command output
func validateOutput(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["output"]; !ok {
	    block["output"] = false
	}
	return block
}

// Validates and defaults the command require field
func validateRequire(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["require"]; !ok {
	    block["require"] = ""
	}
	return block
}

// Validates and defaults the command group field
func validateGroup(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["group"]; !ok {
	    block["group"] = ""
	}
	return block
}

// Validates and defaults the command directory path
func validateDir(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["dir"]; !ok {
	    block["dir"] = getWorkingDirectory()
	}
	return block
}

// Validates and defaults the command ssh remote field
func validateRemote(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["remote"]; !ok {
	    block["remote"] = ""
	}
	return block
}

// Validates and defaults the command ssh password
func validateSshPassword(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["password"]; !ok {
	    block["password"] = ""
	}
	return block
}

// Validates and defaults the command ssh key
func validateSshKey(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["key"]; !ok {
	    block["key"] = ""
	}
	return block
}

// Validates and defaults the includes arguments
func validateInclude(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["include"]; !ok {
		var defaults []interface{}
		block["include"] = defaults
	}
	return block
}
