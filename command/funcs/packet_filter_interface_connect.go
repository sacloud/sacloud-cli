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

func PacketFilterInterfaceConnect(ctx command.Context, params *params.InterfaceConnectPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	interfaceAPI := client.GetInterfaceAPI()

	// read packet filter
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", e)
	}

	// read interface
	nic, e := interfaceAPI.Read(params.InterfaceId)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", e)
	}

	if nic.PacketFilter != nil && nic.PacketFilter.ID > sacloud.EmptyID {
		return fmt.Errorf("interface(%d) is already connected packet filter", nic.ID)
	}

	// call manipurate functions
	_, err := interfaceAPI.ConnectToPacketFilter(params.InterfaceId, params.Id)
	if err != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", err)
	}
	return nil

}
