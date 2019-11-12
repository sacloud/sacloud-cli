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
	"time"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/helper/printer"
)

func VPCRouterLogs(ctx command.Context, params *params.LogsVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()

	// validate params
	if params.Follow && params.LogName == "all" {
		fmt.Fprintf(command.GlobalOption.Err, "[WARN] -f/--follow option can not use with --log-name=%q, ignored", "all")
		params.Follow = false
	}
	if ctx.IsSet("refresh-interval") && params.LogName == "all" {
		fmt.Fprintf(command.GlobalOption.Err, "[WARN] --refresh-interval option can not use with --log-name=%q, ignored", "all")
	}

	logBuf := internal.NewHashQueue(500)
	out := command.GlobalOption.Out

	if params.ListLogNames {
		for _, name := range define.AllowVPCRouterLogNames {
			fmt.Fprintln(out, name)
		}
		return nil
	}

	for {
		// call Read(id)
		res, err := api.Status(params.Id)
		if err != nil {
			return fmt.Errorf("VPCRouterLog is failed: %s", err)
		}

		logs := map[string][]string{}
		switch params.LogName {
		case "all":
			logs["vpn"] = res.VPNLogs
			logs["firewall-send"] = res.FirewallSendLogs
			logs["firewall-receive"] = res.FirewallReceiveLogs
		case "vpn":
			logs["vpn"] = res.VPNLogs
		case "firewall-send":
			logs["firewall-send"] = res.FirewallSendLogs
		case "firewall-receive":
			logs["firewall-receive"] = res.FirewallReceiveLogs
		}

		for key, lines := range logs {
			if params.LogName == "all" {
				printer.Fprintf(out, color.New(color.FgHiGreen), "\n==> [%s]:start\n", key)
			}
			for _, line := range lines {
				if logBuf.PutIfAbsent(line) {
					fmt.Fprintf(out, "%s\n", line)
				}
			}
			if params.LogName == "all" {
				printer.Fprintf(out, color.New(color.FgHiGreen), "\n<== [%s]:end\n", key)
			}
		}

		if !params.Follow {
			break
		}
		time.Sleep(time.Duration(params.RefreshInterval) * time.Second)
	}
	return nil
}
