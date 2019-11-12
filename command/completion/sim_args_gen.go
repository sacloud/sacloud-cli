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
)

func SIMListCompleteArgs(ctx command.Context, params *params.ListSIMParam, cur, prev, commandName string) {

}

func SIMCreateCompleteArgs(ctx command.Context, params *params.CreateSIMParam, cur, prev, commandName string) {

}

func SIMReadCompleteArgs(ctx command.Context, params *params.ReadSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMUpdateCompleteArgs(ctx command.Context, params *params.UpdateSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMDeleteCompleteArgs(ctx command.Context, params *params.DeleteSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMCarrierInfoCompleteArgs(ctx command.Context, params *params.CarrierInfoSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMCarrierUpdateCompleteArgs(ctx command.Context, params *params.CarrierUpdateSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMActivateCompleteArgs(ctx command.Context, params *params.ActivateSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMDeactivateCompleteArgs(ctx command.Context, params *params.DeactivateSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMImeiLockCompleteArgs(ctx command.Context, params *params.ImeiLockSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMIpAddCompleteArgs(ctx command.Context, params *params.IpAddSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMImeiUnlockCompleteArgs(ctx command.Context, params *params.ImeiUnlockSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMIpDeleteCompleteArgs(ctx command.Context, params *params.IpDeleteSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMLogsCompleteArgs(ctx command.Context, params *params.LogsSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SIMMonitorCompleteArgs(ctx command.Context, params *params.MonitorSIMParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSIMAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceSIMItems {
		fmt.Println(res.CommonServiceSIMItems[i].ID)
		var target interface{} = &res.CommonServiceSIMItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
