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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-normalize-flag-name'; DO NOT EDIT

package commands

import (
	"github.com/spf13/pflag"
)

func ipv6ListNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "offset":
		name = "from"
	case "limit":
		name = "max"
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func ipv6PtrAddNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func ipv6PtrReadNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func ipv6PtrUpdateNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func ipv6PtrDeleteNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}
