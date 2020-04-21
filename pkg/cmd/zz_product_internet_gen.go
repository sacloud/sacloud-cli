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

// productInternetCmd represents the command to manage SAKURA Cloud ProductInternet
func productInternetCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "product-internet",
		Aliases: []string{"internet-plan"},
		Short:   "A manage commands of ProductInternet",
		Long:    `A manage commands of ProductInternet`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "list")
		},
	}
}

func productInternetListCmd() *cobra.Command {
	productInternetListParam := params.NewListProductInternetParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List ProductInternet (default)",
		Long:         `List ProductInternet (default)`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, productInternetListParam)
			if err != nil {
				return err
			}
			if err := productInternetListParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if productInternetListParam.GenerateSkeleton {
				return generateSkeleton(ctx, productInternetListParam)
			}

			return funcs.ProductInternetList(ctx, productInternetListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&productInternetListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &productInternetListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&productInternetListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&productInternetListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&productInternetListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&productInternetListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productInternetListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productInternetListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productInternetListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productInternetListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productInternetListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&productInternetListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&productInternetListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productInternetListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&productInternetListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productInternetListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productInternetListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(productInternetListNormalizeFlagNames)
	buildFlagsUsage(cmd, productInternetListFlagOrder(cmd))
	return cmd
}

func productInternetReadCmd() *cobra.Command {
	productInternetReadParam := params.NewReadProductInternetParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read ProductInternet",
		Long:         `Read ProductInternet`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, productInternetReadParam)
			if err != nil {
				return err
			}
			if err := productInternetReadParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if productInternetReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, productInternetReadParam)
			}

			// confirm
			if !productInternetReadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("read", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ProductInternetRead(ctx, productInternetReadParam)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&productInternetReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&productInternetReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productInternetReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productInternetReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productInternetReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productInternetReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productInternetReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&productInternetReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&productInternetReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productInternetReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&productInternetReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productInternetReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productInternetReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &productInternetReadParam.Id), "id", "", "set resource ID")
	fs.SetNormalizeFunc(productInternetReadNormalizeFlagNames)
	buildFlagsUsage(cmd, productInternetReadFlagOrder(cmd))
	return cmd
}
