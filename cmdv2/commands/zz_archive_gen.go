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

// archiveCmd represents the command to manage SAKURA Cloud Archive
func archiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "archive",
		Short: "A manage commands of Archive",
		Long:  `A manage commands of Archive`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func archiveListCmd() *cobra.Command {
	archiveListParam := params.NewListArchiveParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "select"},
		Short:        "List Archive",
		Long:         `List Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveListParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveListParam)
			}

			return funcs.ArchiveList(ctx, archiveListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &archiveListParam.Id), "id", "", "set filter by id(s)")
	fs.StringVarP(&archiveListParam.Scope, "scope", "", "", "set filter by scope('user' or 'shared')")
	fs.StringSliceVarP(&archiveListParam.Tags, "tags", "", []string{}, "set filter by tags(AND) (aliases: selector)")
	fs.VarP(newIDValue(0, &archiveListParam.SourceArchiveId), "source-archive-id", "", "set filter by source-archive-id")
	fs.VarP(newIDValue(0, &archiveListParam.SourceDiskId), "source-disk-id", "", "set filter by source-disk-id")
	fs.IntVarP(&archiveListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&archiveListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&archiveListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&archiveListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(archiveListNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveListFlagOrder(cmd))
	return cmd
}

func archiveCreateCmd() *cobra.Command {
	archiveCreateParam := params.NewCreateArchiveParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create Archive",
		Long:         `Create Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveCreateParam)
			}

			// confirm
			if !archiveCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ArchiveCreate(ctx, archiveCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &archiveCreateParam.SourceDiskId), "source-disk-id", "", "set source disk ID")
	fs.VarP(newIDValue(0, &archiveCreateParam.SourceArchiveId), "source-archive-id", "", "set source archive ID")
	fs.IntVarP(&archiveCreateParam.Size, "size", "", 0, "set archive size(GB)")
	fs.StringVarP(&archiveCreateParam.ArchiveFile, "archive-file", "", "", "set archive image file")
	fs.StringVarP(&archiveCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&archiveCreateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&archiveCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &archiveCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&archiveCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(archiveCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveCreateFlagOrder(cmd))
	return cmd
}

func archiveReadCmd() *cobra.Command {
	archiveReadParam := params.NewReadArchiveParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read Archive",
		Long:         `Read Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveReadTargets(ctx, archiveReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveReadParam.SetId(id)
				go func(p *params.ReadArchiveParam) {
					err := funcs.ArchiveRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&archiveReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &archiveReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveReadNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveReadFlagOrder(cmd))
	return cmd
}

func archiveUpdateCmd() *cobra.Command {
	archiveUpdateParam := params.NewUpdateArchiveParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update Archive",
		Long:         `Update Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveUpdateTargets(ctx, archiveUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveUpdateParam.Assumeyes {
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
				archiveUpdateParam.SetId(id)
				go func(p *params.UpdateArchiveParam) {
					err := funcs.ArchiveUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&archiveUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&archiveUpdateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&archiveUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &archiveUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&archiveUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &archiveUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveUpdateFlagOrder(cmd))
	return cmd
}

func archiveDeleteCmd() *cobra.Command {
	archiveDeleteParam := params.NewDeleteArchiveParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete Archive",
		Long:         `Delete Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveDeleteTargets(ctx, archiveDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveDeleteParam.Assumeyes {
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
				archiveDeleteParam.SetId(id)
				go func(p *params.DeleteArchiveParam) {
					err := funcs.ArchiveDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&archiveDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &archiveDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveDeleteFlagOrder(cmd))
	return cmd
}

func archiveUploadCmd() *cobra.Command {
	archiveUploadParam := params.NewUploadArchiveParam()
	cmd := &cobra.Command{
		Use: "upload",

		Short:        "Upload Archive",
		Long:         `Upload Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveUploadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveUploadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveUploadParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveUploadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveUploadTargets(ctx, archiveUploadParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveUploadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("upload", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveUploadParam.SetId(id)
				go func(p *params.UploadArchiveParam) {
					err := funcs.ArchiveUpload(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveUploadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&archiveUploadParam.ArchiveFile, "archive-file", "", "", "set archive image file")
	fs.StringSliceVarP(&archiveUploadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&archiveUploadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveUploadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveUploadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveUploadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveUploadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveUploadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveUploadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveUploadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveUploadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveUploadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveUploadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveUploadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveUploadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &archiveUploadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveUploadNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveUploadFlagOrder(cmd))
	return cmd
}

func archiveDownloadCmd() *cobra.Command {
	archiveDownloadParam := params.NewDownloadArchiveParam()
	cmd := &cobra.Command{
		Use: "download",

		Short:        "Download Archive",
		Long:         `Download Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveDownloadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveDownloadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveDownloadParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveDownloadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveDownloadTargets(ctx, archiveDownloadParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveDownloadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("download", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveDownloadParam.SetId(id)
				go func(p *params.DownloadArchiveParam) {
					err := funcs.ArchiveDownload(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveDownloadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&archiveDownloadParam.FileDestination, "file-destination", "", "", "set file destination path")
	fs.StringSliceVarP(&archiveDownloadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&archiveDownloadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveDownloadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveDownloadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveDownloadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveDownloadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveDownloadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &archiveDownloadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveDownloadNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveDownloadFlagOrder(cmd))
	return cmd
}

func archiveFTPOpenCmd() *cobra.Command {
	archiveFTPOpenParam := params.NewFTPOpenArchiveParam()
	cmd := &cobra.Command{
		Use: "ftp-open",

		Short:        "FTPOpen Archive",
		Long:         `FTPOpen Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveFTPOpenParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveFTPOpenParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveFTPOpenParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveFTPOpenParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveFTPOpenTargets(ctx, archiveFTPOpenParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveFTPOpenParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ftp-open", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveFTPOpenParam.SetId(id)
				go func(p *params.FTPOpenArchiveParam) {
					err := funcs.ArchiveFTPOpen(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveFTPOpenParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveFTPOpenParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&archiveFTPOpenParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveFTPOpenParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveFTPOpenParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveFTPOpenParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveFTPOpenParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveFTPOpenParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&archiveFTPOpenParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&archiveFTPOpenParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&archiveFTPOpenParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&archiveFTPOpenParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&archiveFTPOpenParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&archiveFTPOpenParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&archiveFTPOpenParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &archiveFTPOpenParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveFTPOpenNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveFTPOpenFlagOrder(cmd))
	return cmd
}

func archiveFTPCloseCmd() *cobra.Command {
	archiveFTPCloseParam := params.NewFTPCloseArchiveParam()
	cmd := &cobra.Command{
		Use: "ftp-close",

		Short:        "FTPClose Archive",
		Long:         `FTPClose Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveFTPCloseParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveFTPCloseParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveFTPCloseParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveFTPCloseParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveFTPCloseTargets(ctx, archiveFTPCloseParam)
			if err != nil {
				return err
			}

			// confirm
			if !archiveFTPCloseParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ftp-close", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveFTPCloseParam.SetId(id)
				go func(p *params.FTPCloseArchiveParam) {
					err := funcs.ArchiveFTPClose(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveFTPCloseParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveFTPCloseParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&archiveFTPCloseParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&archiveFTPCloseParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveFTPCloseParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveFTPCloseParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveFTPCloseParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveFTPCloseParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &archiveFTPCloseParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveFTPCloseNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveFTPCloseFlagOrder(cmd))
	return cmd
}

func archiveWaitForCopyCmd() *cobra.Command {
	archiveWaitForCopyParam := params.NewWaitForCopyArchiveParam()
	cmd := &cobra.Command{
		Use: "wait-for-copy",

		Short:        "WaitForCopy Archive",
		Long:         `WaitForCopy Archive`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return archiveWaitForCopyParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, archiveWaitForCopyParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if archiveWaitForCopyParam.GenerateSkeleton {
				return generateSkeleton(ctx, archiveWaitForCopyParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findArchiveWaitForCopyTargets(ctx, archiveWaitForCopyParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				archiveWaitForCopyParam.SetId(id)
				go func(p *params.WaitForCopyArchiveParam) {
					err := funcs.ArchiveWaitForCopy(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(archiveWaitForCopyParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&archiveWaitForCopyParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&archiveWaitForCopyParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&archiveWaitForCopyParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&archiveWaitForCopyParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&archiveWaitForCopyParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&archiveWaitForCopyParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &archiveWaitForCopyParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(archiveWaitForCopyNormalizeFlagNames)
	buildFlagsUsage(cmd, archiveWaitForCopyFlagOrder(cmd))
	return cmd
}
