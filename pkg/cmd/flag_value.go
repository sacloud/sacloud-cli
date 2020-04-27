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

package cmd

import (
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type idValue types.ID

func newIDValue(val types.ID, p *types.ID) *idValue {
	*p = val
	return (*idValue)(p)
}

func (v *idValue) String() string {
	return types.ID(*v).String()
}

func (v *idValue) Set(id string) error {
	*v = idValue(types.StringID(id))
	return nil
}

func (v *idValue) Type() string {
	return "types.ID"
}

type idSliceValue []types.ID

func newIDSliceValue(val []types.ID, p *[]types.ID) *idSliceValue {
	*p = val
	return (*idSliceValue)(p)
}

func (v *idSliceValue) String() string {
	var ids []string
	for _, id := range []types.ID(*v) {
		ids = append(ids, id.String())
	}
	return strings.Join(ids, ",")
}

func (v *idSliceValue) Set(ids string) error {
	values := strings.Split(ids, ",")
	*v = []types.ID{}
	for _, val := range values {
		*v = append(*v, types.StringID(val))
	}
	return nil
}

func (v *idSliceValue) Type() string {
	return "types.IDSlice"
}
