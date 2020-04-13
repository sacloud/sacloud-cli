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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"errors"
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// gslbCmd represents the command to manage SAKURA Cloud GSLB
func gslbCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gslb",
		Short: "A manage commands of GSLB",
		Long:  `A manage commands of GSLB`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func gslbListCmd() *cobra.Command {
	gslbListParam := params.NewListGSLBParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "selector"},
		Short:        "List GSLB",
		Long:         `List GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbListParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbListParam)
			}

			return funcs.GSLBList(ctx, gslbListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&gslbListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &gslbListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&gslbListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&gslbListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&gslbListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&gslbListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&gslbListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(gslbListNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbListFlagOrder(cmd))
	return cmd
}

func gslbServerInfoCmd() *cobra.Command {
	gslbServerInfoParam := params.NewServerInfoGSLBParam()
	cmd := &cobra.Command{
		Use:          "server-info",
		Aliases:      []string{"server-list"},
		Short:        "ServerInfo GSLB",
		Long:         `ServerInfo GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbServerInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbServerInfoParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbServerInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbServerInfoParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBServerInfoTargets(ctx, gslbServerInfoParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbServerInfoParam.SetId(id)
				go func(p *params.ServerInfoGSLBParam) {
					err := funcs.GSLBServerInfo(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbServerInfoParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&gslbServerInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&gslbServerInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbServerInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbServerInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbServerInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbServerInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbServerInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbServerInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbServerInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbServerInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbServerInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbServerInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbServerInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbServerInfoParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbServerInfoNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbServerInfoFlagOrder(cmd))
	return cmd
}

func gslbCreateCmd() *cobra.Command {
	gslbCreateParam := params.NewCreateGSLBParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create GSLB",
		Long:         `Create GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbCreateParam)
			}

			// confirm
			if !gslbCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.GSLBCreate(ctx, gslbCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&gslbCreateParam.Protocol, "protocol", "", "ping", "set healthcheck protocol[http/https/ping/tcp]")
	fs.StringVarP(&gslbCreateParam.HostHeader, "host-header", "", "", "set host header of http/https healthcheck request")
	fs.StringVarP(&gslbCreateParam.Path, "path", "", "/", "set path of http/https healthcheck request")
	fs.IntVarP(&gslbCreateParam.ResponseCode, "response-code", "", 200, "set response-code of http/https healthcheck request")
	fs.IntVarP(&gslbCreateParam.Port, "port", "", 0, "set port of tcp healthcheck")
	fs.IntVarP(&gslbCreateParam.DelayLoop, "delay-loop", "", 10, "set delay-loop of healthcheck")
	fs.BoolVarP(&gslbCreateParam.Weighted, "weighted", "", true, "enable weighted")
	fs.StringVarP(&gslbCreateParam.SorryServer, "sorry-server", "", "", "set sorry-server hostname/ipaddress")
	fs.StringVarP(&gslbCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&gslbCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&gslbCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &gslbCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&gslbCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(gslbCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbCreateFlagOrder(cmd))
	return cmd
}

func gslbServerAddCmd() *cobra.Command {
	gslbServerAddParam := params.NewServerAddGSLBParam()
	cmd := &cobra.Command{
		Use: "server-add",

		Short:        "ServerAdd GSLB",
		Long:         `ServerAdd GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbServerAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbServerAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbServerAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbServerAddParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBServerAddTargets(ctx, gslbServerAddParam)
			if err != nil {
				return err
			}

			// confirm
			if !gslbServerAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("server-add", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbServerAddParam.SetId(id)
				go func(p *params.ServerAddGSLBParam) {
					err := funcs.GSLBServerAdd(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbServerAddParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&gslbServerAddParam.Ipaddress, "ipaddress", "", "", "set target ipaddress")
	fs.BoolVarP(&gslbServerAddParam.Disabled, "disabled", "", false, "set disabled")
	fs.IntVarP(&gslbServerAddParam.Weight, "weight", "", 0, "set weight")
	fs.StringSliceVarP(&gslbServerAddParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&gslbServerAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbServerAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbServerAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbServerAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbServerAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbServerAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbServerAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbServerAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbServerAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbServerAddParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbServerAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbServerAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbServerAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbServerAddParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbServerAddNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbServerAddFlagOrder(cmd))
	return cmd
}

func gslbReadCmd() *cobra.Command {
	gslbReadParam := params.NewReadGSLBParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read GSLB",
		Long:         `Read GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBReadTargets(ctx, gslbReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbReadParam.SetId(id)
				go func(p *params.ReadGSLBParam) {
					err := funcs.GSLBRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&gslbReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&gslbReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbReadNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbReadFlagOrder(cmd))
	return cmd
}

func gslbServerUpdateCmd() *cobra.Command {
	gslbServerUpdateParam := params.NewServerUpdateGSLBParam()
	cmd := &cobra.Command{
		Use: "server-update",

		Short:        "ServerUpdate GSLB",
		Long:         `ServerUpdate GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbServerUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbServerUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbServerUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbServerUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBServerUpdateTargets(ctx, gslbServerUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !gslbServerUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("server-update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbServerUpdateParam.SetId(id)
				go func(p *params.ServerUpdateGSLBParam) {
					err := funcs.GSLBServerUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbServerUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&gslbServerUpdateParam.Index, "index", "", 0, "index of target server")
	fs.StringVarP(&gslbServerUpdateParam.Ipaddress, "ipaddress", "", "", "set target ipaddress")
	fs.BoolVarP(&gslbServerUpdateParam.Disabled, "disabled", "", false, "set disabled")
	fs.IntVarP(&gslbServerUpdateParam.Weight, "weight", "", 0, "set weight")
	fs.StringSliceVarP(&gslbServerUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&gslbServerUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbServerUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbServerUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbServerUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbServerUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbServerUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbServerUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbServerUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbServerUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbServerUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbServerUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbServerUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbServerUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbServerUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbServerUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbServerUpdateFlagOrder(cmd))
	return cmd
}

func gslbServerDeleteCmd() *cobra.Command {
	gslbServerDeleteParam := params.NewServerDeleteGSLBParam()
	cmd := &cobra.Command{
		Use: "server-delete",

		Short:        "ServerDelete GSLB",
		Long:         `ServerDelete GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbServerDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbServerDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbServerDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbServerDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBServerDeleteTargets(ctx, gslbServerDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !gslbServerDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete server", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbServerDeleteParam.SetId(id)
				go func(p *params.ServerDeleteGSLBParam) {
					err := funcs.GSLBServerDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbServerDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&gslbServerDeleteParam.Index, "index", "", 0, "index of target server")
	fs.StringSliceVarP(&gslbServerDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&gslbServerDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbServerDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbServerDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbServerDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbServerDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbServerDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbServerDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbServerDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbServerDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbServerDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbServerDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbServerDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbServerDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbServerDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbServerDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbServerDeleteFlagOrder(cmd))
	return cmd
}

func gslbUpdateCmd() *cobra.Command {
	gslbUpdateParam := params.NewUpdateGSLBParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update GSLB",
		Long:         `Update GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBUpdateTargets(ctx, gslbUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !gslbUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbUpdateParam.SetId(id)
				go func(p *params.UpdateGSLBParam) {
					err := funcs.GSLBUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&gslbUpdateParam.Protocol, "protocol", "", "", "set healthcheck protocol[http/https/ping/tcp]")
	fs.StringVarP(&gslbUpdateParam.HostHeader, "host-header", "", "", "set host header of http/https healthcheck request")
	fs.StringVarP(&gslbUpdateParam.Path, "path", "", "", "set path of http/https healthcheck request")
	fs.IntVarP(&gslbUpdateParam.ResponseCode, "response-code", "", 0, "set response-code of http/https healthcheck request")
	fs.IntVarP(&gslbUpdateParam.Port, "port", "", 0, "set port of tcp healthcheck")
	fs.IntVarP(&gslbUpdateParam.DelayLoop, "delay-loop", "", 0, "set delay-loop of healthcheck")
	fs.BoolVarP(&gslbUpdateParam.Weighted, "weighted", "", false, "enable weighted")
	fs.StringVarP(&gslbUpdateParam.SorryServer, "sorry-server", "", "", "set sorry-server hostname/ipaddress")
	fs.StringSliceVarP(&gslbUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&gslbUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&gslbUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&gslbUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &gslbUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&gslbUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbUpdateFlagOrder(cmd))
	return cmd
}

func gslbDeleteCmd() *cobra.Command {
	gslbDeleteParam := params.NewDeleteGSLBParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete GSLB",
		Long:         `Delete GSLB`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return gslbDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, gslbDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if gslbDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, gslbDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findGSLBDeleteTargets(ctx, gslbDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !gslbDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				gslbDeleteParam.SetId(id)
				go func(p *params.DeleteGSLBParam) {
					err := funcs.GSLBDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(gslbDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&gslbDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&gslbDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&gslbDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&gslbDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&gslbDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&gslbDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&gslbDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&gslbDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&gslbDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&gslbDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&gslbDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&gslbDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&gslbDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&gslbDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &gslbDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(gslbDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, gslbDeleteFlagOrder(cmd))
	return cmd
}
