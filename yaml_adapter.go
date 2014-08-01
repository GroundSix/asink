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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"./vendor/yaml"
	"strconv"
)

func processYamlContentFromFile(filename string) string {
	var data interface{}
	input, err := ioutil.ReadFile(filename)
	if (err != nil) {
		panic(err)
	}
	err = yaml.Unmarshal(input, &data)
	if (err != nil) {
		panic(err)
	}
	data, err = transformData(data)
	if (err != nil) {
		panic(err)
	}
	output, err := json.Marshal(data)
	if (err != nil) {
		panic(err)
	}

	return string(output)
}

func transformData(in interface{}) (out interface{}, err error) {
	switch in.(type) {
	case map[interface{}]interface{}:
		o := make(map[string]interface{})
		for k, v := range in.(map[interface{}]interface{}) {
			sk := ""
			switch k.(type) {
			case string:
				sk = k.(string)
			case int:
				sk = strconv.Itoa(k.(int))
			default:
				return nil, errors.New(
					fmt.Sprintf("type not match: expect map key string or int get: %T", k))
			}
			v, err = transformData(v)
			if err != nil {
				return nil, err
			}
			o[sk] = v
		}
		return o, nil
	case []interface{}:
		in1 := in.([]interface{})
		len1 := len(in1)
		o := make([]interface{}, len1)
		for i := 0; i < len1; i++ {
			o[i], err = transformData(in1[i])
			if err != nil {
				return nil, err
			}
		}
		return o, nil
	default:
		return in, nil
	}
	return in, nil
}
