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

func DiskServerDisconnect(ctx command.Context, params *params.ServerDisconnectDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", e)
	}

	// disk is connected to server?
	if p.Server == nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", "Disk needs to be connected to server")
	}

	// server is stopped?
	if !p.Server.IsDown() {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", "Server needs to be stopped")
	}

	// call manipurate functions
	_, err := api.DisconnectFromServer(params.Id)
	if err != nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", err)
	}

	return nil
}
