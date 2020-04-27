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

package ipv6

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/search"
)

func List(ctx cli.Context, params *params.ListIPv6Param) error {
	client := sacloud.NewIPv6AddrOp(ctx.Client())
	res, err := client.Find(ctx, ctx.Zone(), search.FindCondition(params))
	if err != nil {
		return fmt.Errorf("IPv6List is failed: %s", err)
	}

	var list []interface{}
	for i := range res.IPv6Addrs {
		// TODO ルータIDでのフィルタを実装

		if !params.IPv6netId.IsEmpty() && res.IPv6Addrs[i].IPv6NetID != params.IPv6netId {
			continue
		}

		list = append(list, &res.IPv6Addrs[i])
	}
	return ctx.Output().Print(list...)
}
