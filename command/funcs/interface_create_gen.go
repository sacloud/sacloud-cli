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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfaceCreate(ctx command.Context, params *params.CreateInterfaceParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInterfaceAPI()
	p := api.New()

	// set params

	p.SetServerID(params.ServerId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("InterfaceCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
