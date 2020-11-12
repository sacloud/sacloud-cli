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

//go:generate go run github.com/sacloud/usacloud/tools/gen-commands/
package cmd

import (
	"github.com/sacloud/usacloud/pkg/cmd/base"
	"github.com/sacloud/usacloud/pkg/cmd/disk"
	"github.com/sacloud/usacloud/pkg/cmd/root"
)

var Resources = []*base.Resource{
	disk.Resource,
}

func initCommands() {
	for _, r := range Resources {
		root.Command.AddCommand(r.CLICommand())
	}
}