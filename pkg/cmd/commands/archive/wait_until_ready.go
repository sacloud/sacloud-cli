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

package archive

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var waitUntilReadyCommand = &core.Command{
	Name:         "wait-until-ready",
	Aliases:      []string{"wait", "wait-for-copy"}, // v0との互換用
	Category:     "other",
	Order:        10,
	SelectorType: core.SelectorTypeRequireMulti,

	ServiceFuncAltName: "WaitReady",

	ParameterInitializer: func() interface{} {
		return newWaitUntilReadyParameter()
	},
}

type waitUntilReadyParameter struct {
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
}

func newWaitUntilReadyParameter() *waitUntilReadyParameter {
	return &waitUntilReadyParameter{}
}

func init() {
	Resource.AddCommand(waitUntilReadyCommand)
}
