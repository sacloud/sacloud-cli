// Copyright 2017-2021 The Usacloud Authors
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

package sim

import (
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var deleteCommand = &core.Command{
	Name:         "delete",
	Aliases:      []string{"rm"},
	Category:     "basic",
	Order:        50,
	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newDeleteParameter()
	},
}

type deleteParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.FailIfNotFoundParameter `cli:",squash" mapconv:",squash"`
	cflag.WaitForReleaseParameter `cli:",squash" mapconv:",squash"`
	Zones                         []string `cli:"-"` // Note: Customizeの中でctxから受け取る
}

func newDeleteParameter() *deleteParameter {
	return &deleteParameter{}
}

func init() {
	Resource.AddCommand(deleteCommand)
}

// Customize パラメータ変換処理
func (p *deleteParameter) Customize(ctx cli.Context) error {
	var zones []string
	for _, zone := range ctx.Option().Zones {
		if zone != "all" {
			zones = append(zones, zone)
		}
	}
	p.Zones = zones
	return nil
}
