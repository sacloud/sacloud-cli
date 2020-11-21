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

import "github.com/sacloud/usacloud/pkg/output"

var defaultColumnDefs = []output.ColumnDef{
	{Name: "Zone"},
	{Name: "ID"},
	{Name: "Name"},
	{
		Name:     "Server",
		Template: `{{ if ne .ServerID "0" }}{{ .ServerID }}({{ .ServerName }}){{ end }}`,
	},
	{
		Name:     "Plan",
		Template: "{{ disk_plan_to_key .DiskPlanID }}",
	},
	{
		Name:     "Size",
		Template: "{{ mib_to_gib .SizeMB }}GB",
	},
	{Name: "Connection"},
	{
		Name:     "Storage",
		Template: "{{ .Storage.Name }}",
	},
}