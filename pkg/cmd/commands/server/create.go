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

package server

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	CPU        int    `cli:"cpu,aliases=core" validate:"required"`
	Memory     int    `cli:"memory" mapconv:"MemoryMB,filters=gib_to_mib" validate:"required"`
	Commitment string `cli:",options=server_plan_commitment" mapconv:"ServerPlanCommitment,filters=server_plan_commitment_to_value" validate:"required,server_plan_commitment"`
	Generation string `cli:",options=server_plan_generation" mapconv:"ServerPlanGeneration,filters=server_plan_generation_to_value" validate:"required,server_plan_generation"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		CPU:        1,
		Memory:     1,
		Commitment: types.Commitments.Standard.String(),
		Generation: "default",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}
