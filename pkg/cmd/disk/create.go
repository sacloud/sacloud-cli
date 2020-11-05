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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/base"
)

type CreateParameter struct {
	// TODO これらはコマンドのコンテキストでパラメーターに含めないべき？要検討
	*base.ExecContext `cli:"-" mapconv:"-"`

	Zone string `cli:"-" validate:"required"`

	Name            string     `cli:",category=disk"`
	Description     string     `cli:",category=disk"`
	Tags            []string   `cli:",category=disk"`
	IconID          types.ID   `cli:",category=disk"`
	DiskPlanID      string     `cli:",category=disk" validate:"oneof=ssd hdd"` // TODO パラメータ変換ロジック(string -> types.ID)をコード生成と組み合わせる
	Connection      string     `cli:",category=disk" validate:"oneof=virtio ide"`
	SourceDiskID    types.ID   `cli:",category=disk"`
	SourceArchiveID types.ID   `cli:",category=disk"`
	ServerID        types.ID   `cli:",category=disk"`
	SizeGB          int        `cli:"size,category=disk"`
	DistantFrom     []types.ID `cli:",category=disk"`
	OSType          string     `cli:",category=disk"` // TODO 設定できる値の例示をどうするか

	*base.OutputParameter `cli:",squash" mapconv:"-"`
}

func NewCreateParameter() *CreateParameter {
	return &CreateParameter{
		// TODO デフォルト値はここで設定する
	}
}