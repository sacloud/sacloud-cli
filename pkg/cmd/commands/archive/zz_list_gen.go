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

package archive

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *listParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *listParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.IntVarP(&p.Count, "count", "", p.Count, "(aliases: --max, --limit)")
	fs.IntVarP(&p.From, "from", "", p.From, "(aliases: --offset)")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.StringSliceVarP(&p.Names, "names", "", p.Names, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.StringVarP(&p.OSType, "os-type", "", p.OSType, "options: [centos/centos8/centos7/ubuntu/ubuntu2004/ubuntu1804/ubuntu1604/debian/debian10/debian9/coreos/rancheros/k3os/kusanagi/freebsd/windows2016/windows2016-rds/windows2016-rds-office/windows2016-sql-web/windows2016-sql-standard/windows2016-sql-standard-all/windows2016-sql2017-standard/windows2016-sql2017-enterprise/windows2016-sql2017-standard-all/windows2019/windows2019-rds/windows2019-rds-office2019/windows2019-sql2017-web/windows2019-sql2019-web/windows2019-sql2017-standard/windows2019-sql2019-standard/windows2019-sql2017-enterprise/windows2019-sql2019-enterprise/windows2019-sql2017-standard-all/windows2019-sql2019-standard-all]")
	fs.StringVarP(&p.Scope, "scope", "", p.Scope, "options: [user/shared]")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *listParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "max":
		name = "count"
	case "limit":
		name = "count"
	case "offset":
		name = "from"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *listParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("archive", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		sets = append(sets, &core.FlagSet{
			Title: "Archive options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("count"))
		fs.AddFlag(cmd.LocalFlags().Lookup("from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("names"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("os-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("scope"))
		sets = append(sets, &core.FlagSet{
			Title: "Filter options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *listParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("os-type", util.FlagCompletionFunc("centos", "centos8", "centos7", "ubuntu", "ubuntu2004", "ubuntu1804", "ubuntu1604", "debian", "debian10", "debian9", "coreos", "rancheros", "k3os", "kusanagi", "freebsd", "windows2016", "windows2016-rds", "windows2016-rds-office", "windows2016-sql-web", "windows2016-sql-standard", "windows2016-sql-standard-all", "windows2016-sql2017-standard", "windows2016-sql2017-enterprise", "windows2016-sql2017-standard-all", "windows2019", "windows2019-rds", "windows2019-rds-office2019", "windows2019-sql2017-web", "windows2019-sql2019-web", "windows2019-sql2017-standard", "windows2019-sql2019-standard", "windows2019-sql2017-enterprise", "windows2019-sql2019-enterprise", "windows2019-sql2017-standard-all", "windows2019-sql2019-standard-all"))
	cmd.RegisterFlagCompletionFunc("scope", util.FlagCompletionFunc("user", "shared"))

}

func (p *listParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
