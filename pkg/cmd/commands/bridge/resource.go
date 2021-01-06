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

package bridge

import (
	"reflect"

	"github.com/sacloud/libsacloud/v2/helper/service/bridge"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var Resource = &core.Resource{
	Name:        "bridge",
	ServiceType: reflect.TypeOf(&bridge.Service{}),
	Category:    core.ResourceCategoryNetworking,
}
