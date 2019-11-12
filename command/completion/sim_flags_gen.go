// Copyright 2017-2019 The Usacloud Authors
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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func SIMListCompleteFlags(ctx command.Context, params *params.ListSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["SIM"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMCreateCompleteFlags(ctx command.Context, params *params.CreateSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "iccid":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("iccid")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "passcode":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("passcode")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disabled":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("disabled")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "imei":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("imei")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "carrier":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("carrier")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["SIM"].Commands["create"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMReadCompleteFlags(ctx command.Context, params *params.ReadSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMUpdateCompleteFlags(ctx command.Context, params *params.UpdateSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMDeleteCompleteFlags(ctx command.Context, params *params.DeleteSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "force", "f":
		param := define.Resources["SIM"].Commands["delete"].BuildedParams().Get("force")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["SIM"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMCarrierInfoCompleteFlags(ctx command.Context, params *params.CarrierInfoSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["carrier-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["carrier-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMCarrierUpdateCompleteFlags(ctx command.Context, params *params.CarrierUpdateSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["carrier-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["carrier-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "carrier":
		param := define.Resources["SIM"].Commands["carrier-update"].BuildedParams().Get("carrier")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMActivateCompleteFlags(ctx command.Context, params *params.ActivateSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["activate"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["activate"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMDeactivateCompleteFlags(ctx command.Context, params *params.DeactivateSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["deactivate"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["deactivate"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMImeiLockCompleteFlags(ctx command.Context, params *params.ImeiLockSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["imei-lock"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["imei-lock"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "imei":
		param := define.Resources["SIM"].Commands["imei-lock"].BuildedParams().Get("imei")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMIpAddCompleteFlags(ctx command.Context, params *params.IpAddSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["ip-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["ip-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ip":
		param := define.Resources["SIM"].Commands["ip-add"].BuildedParams().Get("ip")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMImeiUnlockCompleteFlags(ctx command.Context, params *params.ImeiUnlockSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["imei-unlock"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["imei-unlock"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMIpDeleteCompleteFlags(ctx command.Context, params *params.IpDeleteSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["SIM"].Commands["ip-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["ip-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMLogsCompleteFlags(ctx command.Context, params *params.LogsSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "follow", "f":
		param := define.Resources["SIM"].Commands["logs"].BuildedParams().Get("follow")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "refresh-interval":
		param := define.Resources["SIM"].Commands["logs"].BuildedParams().Get("refresh-interval")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["SIM"].Commands["logs"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["logs"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func SIMMonitorCompleteFlags(ctx command.Context, params *params.MonitorSIMParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "start":
		param := define.Resources["SIM"].Commands["monitor"].BuildedParams().Get("start")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "end":
		param := define.Resources["SIM"].Commands["monitor"].BuildedParams().Get("end")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "key-format":
		param := define.Resources["SIM"].Commands["monitor"].BuildedParams().Get("key-format")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["SIM"].Commands["monitor"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["SIM"].Commands["monitor"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
