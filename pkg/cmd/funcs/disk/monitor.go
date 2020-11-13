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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var monitorCommand = &core.Command{
	Name:       "monitor-disk",
	Aliases:    []string{"monitor"},
	Category:   "monitor",
	Order:      10,
	NoConfirm:  true,
	NoProgress: true,

	ColumnDefs: []output.ColumnDef{
		{Name: "Time"},
		{Name: "Read"},
		{Name: "Write"},
	},

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newMonitorParameter()
	},
}

type monitorParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	core.IDParameter   `cli:",squash" mapconv:",squash"`

	core.MonitorParameter `cli:",squash" mapconv:",squash"`

	core.OutputParameter `cli:",squash" mapconv:"-"`
}

func newMonitorParameter() *monitorParameter {
	return &monitorParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(monitorCommand)
}
