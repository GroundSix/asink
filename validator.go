package main

import (

)

func validateBlock(block map[string]interface{}) map[string]interface{} {
	finalBlock := validateName(block)
	finalBlock  = validateCount(finalBlock)
	finalBlock  = validateArgs(finalBlock)
	finalBlock  = validateOutput(finalBlock)
	finalBlock  = validateRequire(finalBlock)
	finalBlock  = validateGroup(finalBlock)

	return finalBlock
}

func validateName(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["command"]; !ok {
	    block["command"] = ""
	}
	return block
}

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

func validateArgs(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["args"]; !ok {
		var defaults []interface{}
		block["args"] = defaults
	}
	return block
}

func validateOutput(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["output"]; !ok {
	    block["output"] = false
	}
	return block
}

func validateRequire(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["require"]; !ok {
	    block["require"] = ""
	}
	return block
}

func validateGroup(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["group"]; !ok {
	    block["group"] = ""
	}
	return block
}

func validateDir(block map[string]interface{}) map[string]interface{} {
	if _,ok := block["dir"]; !ok {
	    block["dir"] = "."
	}
	return block
}