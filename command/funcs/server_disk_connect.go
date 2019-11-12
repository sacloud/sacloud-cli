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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerDiskConnect(ctx command.Context, params *params.DiskConnectServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskConnect is failed: %s", e)
	}

	if len(p.Disks) >= sacloud.ServerMaxDiskLen {
		return fmt.Errorf("Server already connected maximum count of disks")
	}

	if p.IsUp() {
		return fmt.Errorf("ServerDiskConnect is failed: %s", "server is running")
	}

	// call manipurate functions
	_, err := client.GetDiskAPI().ConnectToServer(params.DiskId, params.Id)
	if err != nil {
		return fmt.Errorf("ServerDiskConnect is failed: %s", err)
	}

	return nil

}
