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
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
)

// CsvBillParam is input parameters for the sacloud API
type CsvBillParam struct {
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	NoHeader          bool
	BillOutput        string
	BillId            sacloud.ID

	config *config.Config
	input  Input
}

// NewCsvBillParam return new CsvBillParam
func NewCsvBillParam() *CsvBillParam {
	return &CsvBillParam{}
}

// Initialize init CsvBillParam
func (p *CsvBillParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CsvBillParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *CsvBillParam) FillValueToSkeleton() {
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
	if util.IsEmpty(p.NoHeader) {
		p.NoHeader = false
	}
	if util.IsEmpty(p.BillOutput) {
		p.BillOutput = ""
	}
	if util.IsEmpty(p.BillId) {
		p.BillId = sacloud.ID(0)
	}

}

func (p *CsvBillParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Bill"].Commands["csv"].Params["bill-id"].ValidateFunc
		errs := validator("--bill-id", p.BillId)
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
	return util.FlattenErrors(errors)
}

func (p *CsvBillParam) ResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *CsvBillParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["csv"]
}

func (p *CsvBillParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CsvBillParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CsvBillParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CsvBillParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *CsvBillParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *CsvBillParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["csv"]
}

func (p *CsvBillParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CsvBillParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CsvBillParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CsvBillParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CsvBillParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CsvBillParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CsvBillParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CsvBillParam) GetParameters() string {
	return p.Parameters
}
func (p *CsvBillParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CsvBillParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CsvBillParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CsvBillParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CsvBillParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CsvBillParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CsvBillParam) SetNoHeader(v bool) {
	p.NoHeader = v
}

func (p *CsvBillParam) GetNoHeader() bool {
	return p.NoHeader
}
func (p *CsvBillParam) SetBillOutput(v string) {
	p.BillOutput = v
}

func (p *CsvBillParam) GetBillOutput() string {
	return p.BillOutput
}
func (p *CsvBillParam) SetBillId(v sacloud.ID) {
	p.BillId = v
}

func (p *CsvBillParam) GetBillId() sacloud.ID {
	return p.BillId
}

// Changed usacloud v0系との互換性維持のための実装
func (p *CsvBillParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ListBillParam is input parameters for the sacloud API
type ListBillParam struct {
	Year              int
	Month             int
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

// NewListBillParam return new ListBillParam
func NewListBillParam() *ListBillParam {
	return &ListBillParam{}
}

// Initialize init ListBillParam
func (p *ListBillParam) Initialize(in Input, args []string, config *config.Config) error {
	p.input = in
	p.config = config
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListBillParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListBillParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Year) {
		p.Year = 0
	}
	if util.IsEmpty(p.Month) {
		p.Month = 0
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

func (p *ListBillParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Bill"].Commands["list"].Params["year"].ValidateFunc
		errs := validator("--year", p.Year)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Bill"].Commands["list"].Params["month"].ValidateFunc
		errs := validator("--month", p.Month)
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

func (p *ListBillParam) ResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *ListBillParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListBillParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListBillParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListBillParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListBillParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListBillParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *ListBillParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListBillParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListBillParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListBillParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListBillParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListBillParam) SetYear(v int) {
	p.Year = v
}

func (p *ListBillParam) GetYear() int {
	return p.Year
}
func (p *ListBillParam) SetMonth(v int) {
	p.Month = v
}

func (p *ListBillParam) GetMonth() int {
	return p.Month
}
func (p *ListBillParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListBillParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListBillParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListBillParam) GetParameters() string {
	return p.Parameters
}
func (p *ListBillParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListBillParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListBillParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListBillParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListBillParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListBillParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListBillParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListBillParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListBillParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListBillParam) GetColumn() []string {
	return p.Column
}
func (p *ListBillParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListBillParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListBillParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListBillParam) GetFormat() string {
	return p.Format
}
func (p *ListBillParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListBillParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListBillParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListBillParam) GetQuery() string {
	return p.Query
}
func (p *ListBillParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListBillParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListBillParam) Changed(name string) bool {
	return p.input.Changed(name)
}
