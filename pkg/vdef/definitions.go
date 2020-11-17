// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vdef

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// definitions usacloudで使う名称(key)/値(value)のペア
var definitions = map[string]map[interface{}]interface{}{
	"disk_plan": {
		"ssd": types.DiskPlans.SSD,
		"hdd": types.DiskPlans.HDD,
	},
	"disk_connection": {
		types.DiskConnections.VirtIO.String(): types.DiskConnections.VirtIO.String(),
		types.DiskConnections.IDE.String():    types.DiskConnections.IDE.String(),
	},
	"os_type": ostypeDefinition(),
}

func ostypeDefinition() map[interface{}]interface{} {
	def := map[interface{}]interface{}{}
	for _, name := range ostype.OSTypeShortNames {
		def[name] = ostype.StrToOSType(name)
	}
	return def
}

func init() {
	registerFunctions()
}

func registerFunctions() {
	// definitionsから各種定義を登録(Note: 同名のものがあった場合は上書き)
	registerConverterFilters()
	registerTemplateFuncMap()
	registerValidators()
	registerCLITagOptions()
}

func registerConverterFilters() {
	for name, defs := range definitions {
		ConverterFilters[name+"_to_value"] = convertFuncToValue(name, defs)
		ConverterFilters[name+"_to_key"] = convertFuncToKey(name, defs)
	}
}

func registerTemplateFuncMap() {
	for name, defs := range definitions {
		TemplateFuncMap[name+"_to_value"] = templateFuncToValue(defs)
		TemplateFuncMap[name+"_to_key"] = templateFuncToKey(defs)
	}
}

func registerValidators() {
	// definitionsの各値からキーを取り出し、"oneof=keyのスペース区切り"というルールを登録する
	for name, defs := range definitions {
		var allows []string
		for key := range defs {
			switch s := key.(type) {
			case string:
				allows = append(allows, s)
			case fmt.Stringer:
				allows = append(allows, s.String())
			}
		}
		validatorAliases[name] = fmt.Sprintf("oneof=%s", joinWithSpace(allows))
	}
}

func registerCLITagOptions() {
	// definitionsの各値からキーを取り出し、FlagOptionsMapに登録する
	for name, defs := range definitions {
		var allows []string
		for key := range defs {
			switch s := key.(type) {
			case string:
				allows = append(allows, s)
			case fmt.Stringer:
				allows = append(allows, s.String())
			}
		}
		FlagOptionsMap[name] = allows
	}
}

func convertFuncToValue(defName string, def map[interface{}]interface{}) mapconv.FilterFunc {
	return func(v interface{}) (interface{}, error) {
		var result interface{}
		for key, value := range def {
			if reflect.DeepEqual(v, key) {
				result = value
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("key %v not found in %s", v, defName)
		}
		return result, nil
	}
}

func convertFuncToKey(defName string, def map[interface{}]interface{}) mapconv.FilterFunc {
	return func(v interface{}) (interface{}, error) {
		var result interface{}
		for key, value := range def {
			if reflect.DeepEqual(v, value) {
				result = key
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("value %v not found in %s", v, defName)
		}
		return result, nil
	}
}

func templateFuncToValue(def map[interface{}]interface{}) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := ""
		if v, ok := raw.(json.Number); ok {
			in = v.String()
		}
		var result interface{}
		for key, value := range def {
			switch ky := key.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, ky.String()) {
					result = value
					break
				}
			default:
				if reflect.DeepEqual(in, key) {
					result = value
					break
				}
			}
		}
		return result
	}
}

func templateFuncToKey(def map[interface{}]interface{}) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := ""
		if v, ok := raw.(json.Number); ok {
			in = v.String()
		}
		var result interface{}
		for key, value := range def {
			switch val := value.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, val.String()) {
					result = key
					break
				}
			default:
				if reflect.DeepEqual(in, value) {
					result = key
					break
				}
			}
		}
		return result
	}
}