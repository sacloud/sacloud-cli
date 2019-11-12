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
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DiskReinstallFromDisk(ctx command.Context, params *params.ReinstallFromDiskDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallFromDisk is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still installing[ID:%d]...", params.Id),
		fmt.Sprintf("Reinstall disk[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			_, err := api.ReinstallFromDisk(params.Id, params.SourceDiskId, params.DistantFrom...)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DiskReinstallFromDisk is failed: %s", err)
	}

	return nil
}
