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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-params'; DO NOT EDIT

package params

import (
	"fmt"
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validation"
)

// ListZoneParam is input parameters for the sacloud API
type ListZoneParam struct {
	Name              []string
	Id                []sacloud.ID
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string

	config *config.Config
	input  Input
}

// NewListZoneParam return new ListZoneParam
func NewListZoneParam() *ListZoneParam {
	return &ListZoneParam{}
}

// Initialize init ListZoneParam
func (p *ListZoneParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListZoneParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListZoneParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if util.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if util.IsEmpty(p.From) {
		p.From = 0
	}
	if util.IsEmpty(p.Max) {
		p.Max = 0
	}
	if util.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if util.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *ListZoneParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Zone"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ListZoneParam) ResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ListZoneParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListZoneParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListZoneParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListZoneParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListZoneParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListZoneParam) GetResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ListZoneParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListZoneParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListZoneParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListZoneParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListZoneParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListZoneParam) SetName(v []string) {
	p.Name = v
}

func (p *ListZoneParam) GetName() []string {
	return p.Name
}
func (p *ListZoneParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListZoneParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListZoneParam) SetFrom(v int) {
	p.From = v
}

func (p *ListZoneParam) GetFrom() int {
	return p.From
}
func (p *ListZoneParam) SetMax(v int) {
	p.Max = v
}

func (p *ListZoneParam) GetMax() int {
	return p.Max
}
func (p *ListZoneParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListZoneParam) GetSort() []string {
	return p.Sort
}
func (p *ListZoneParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListZoneParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListZoneParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListZoneParam) GetParameters() string {
	return p.Parameters
}
func (p *ListZoneParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListZoneParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListZoneParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListZoneParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListZoneParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListZoneParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListZoneParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListZoneParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListZoneParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListZoneParam) GetColumn() []string {
	return p.Column
}
func (p *ListZoneParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListZoneParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListZoneParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListZoneParam) GetFormat() string {
	return p.Format
}
func (p *ListZoneParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListZoneParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListZoneParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListZoneParam) GetQuery() string {
	return p.Query
}
func (p *ListZoneParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListZoneParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListZoneParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ReadZoneParam is input parameters for the sacloud API
type ReadZoneParam struct {
	Assumeyes         bool
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string
	Id                sacloud.ID

	config *config.Config
	input  Input
}

// NewReadZoneParam return new ReadZoneParam
func NewReadZoneParam() *ReadZoneParam {
	return &ReadZoneParam{}
}

// Initialize init ReadZoneParam
func (p *ReadZoneParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config

	if len(args) == 0 {
		return fmt.Errorf("argument <ID> is required")
	}
	p.Id = sacloud.StringID(args[0])
	if p.Id.IsEmpty() {
		return fmt.Errorf("argument <ID> is required")
	}
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadZoneParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadZoneParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if util.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if util.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadZoneParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Zone"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := cli.ValidateOutputOption(p, p.config.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ReadZoneParam) ResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ReadZoneParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadZoneParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadZoneParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadZoneParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadZoneParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ReadZoneParam) GetResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ReadZoneParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadZoneParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadZoneParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadZoneParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadZoneParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadZoneParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadZoneParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadZoneParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadZoneParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadZoneParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadZoneParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadZoneParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadZoneParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadZoneParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadZoneParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadZoneParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadZoneParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadZoneParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadZoneParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadZoneParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadZoneParam) GetColumn() []string {
	return p.Column
}
func (p *ReadZoneParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadZoneParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadZoneParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadZoneParam) GetFormat() string {
	return p.Format
}
func (p *ReadZoneParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadZoneParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadZoneParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadZoneParam) GetQuery() string {
	return p.Query
}
func (p *ReadZoneParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadZoneParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadZoneParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadZoneParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadZoneParam) Changed(name string) bool {
	return p.input.Changed(name)
}
