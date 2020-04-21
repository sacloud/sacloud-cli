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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cmd

import (
	"errors"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/funcs"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// zoneCmd represents the command to manage SAKURA Cloud Zone
func zoneCmd() *cobra.Command {
	return &cobra.Command{
		Use: "zone",

		Short: "A manage commands of Zone",
		Long:  `A manage commands of Zone`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "list")
		},
	}
}

func zoneListCmd() *cobra.Command {
	zoneListParam := params.NewListZoneParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List Zone (default)",
		Long:         `List Zone (default)`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, zoneListParam)
			if err != nil {
				return err
			}
			if err := zoneListParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if zoneListParam.GenerateSkeleton {
				return generateSkeleton(ctx, zoneListParam)
			}

			return funcs.ZoneList(ctx, zoneListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&zoneListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &zoneListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&zoneListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&zoneListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&zoneListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&zoneListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&zoneListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&zoneListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&zoneListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&zoneListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&zoneListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&zoneListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&zoneListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&zoneListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&zoneListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&zoneListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&zoneListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(zoneListNormalizeFlagNames)
	buildFlagsUsage(cmd, zoneListFlagOrder(cmd))
	return cmd
}

func zoneReadCmd() *cobra.Command {
	zoneReadParam := params.NewReadZoneParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read Zone",
		Long:         `Read Zone`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, zoneReadParam)
			if err != nil {
				return err
			}
			if err := zoneReadParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if zoneReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, zoneReadParam)
			}

			// confirm
			if !zoneReadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("read", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ZoneRead(ctx, zoneReadParam)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&zoneReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&zoneReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&zoneReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&zoneReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&zoneReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&zoneReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&zoneReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&zoneReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&zoneReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&zoneReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&zoneReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&zoneReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&zoneReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &zoneReadParam.Id), "id", "", "set resource ID")
	fs.SetNormalizeFunc(zoneReadNormalizeFlagNames)
	buildFlagsUsage(cmd, zoneReadFlagOrder(cmd))
	return cmd
}
