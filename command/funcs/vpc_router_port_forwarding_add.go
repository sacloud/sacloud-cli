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

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterPortForwardingAdd(ctx command.Context, params *params.PortForwardingAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddPortForwarding(
		params.Protocol,
		fmt.Sprintf("%d", params.GlobalPort),
		params.PrivateIpaddress,
		fmt.Sprintf("%d", params.PrivatePort),
		params.Description,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", err)
	}
	return nil

}
