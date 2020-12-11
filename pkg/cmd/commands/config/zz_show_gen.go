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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package config

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *showParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *showParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Name, "name", "", p.Name, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *showParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func (p *showParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("config", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		sets = append(sets, &core.FlagSet{
			Title: "Config-specific options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *showParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *showParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
