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

/**
 * Validates all keys for tasks and returns
 * default values
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
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

	return finalBlock
}

/**
 * Validates and defaults the command name
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateName(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["command"]; !ok {
	    block["command"] = ""
	}
	return block
}

/**
 * Validates and defaults the command count
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
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

/**
 * Validates and defaults the command arguments
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateArgs(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["args"]; !ok {
		var defaults []interface{}
		block["args"] = defaults
	}
	return block
}

/**
 * Validates and defaults the command output
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateOutput(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["output"]; !ok {
	    block["output"] = false
	}
	return block
}

/**
 * Validates and defaults the command require field
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateRequire(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["require"]; !ok {
	    block["require"] = ""
	}
	return block
}

/**
 * Validates and defaults the command group field
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateGroup(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["group"]; !ok {
	    block["group"] = ""
	}
	return block
}

/**
 * Validates and defaults the command directory path
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateDir(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["dir"]; !ok {
	    block["dir"] = getWorkingDirectory()
	}
	return block
}

/**
 * Validates and defaults the command ssh remote field
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateRemote(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["remote"]; !ok {
	    block["remote"] = ""
	}
	return block
}

/**
 * Validates and defaults the command ssh password
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateSshPassword(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["password"]; !ok {
	    block["password"] = ""
	}
	return block
}

/**
 * Validates and defaults the command ssh key
 *
 * @param map[string]interface{} block of keys and values
 *
 * @return map[string]interface{} defaulted block of keys and values
 */
func validateSshKey(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["key"]; !ok {
	    block["key"] = ""
	}
	return block
}
