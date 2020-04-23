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
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/flags"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validation"
)

// ListLicenseParam is input parameters for the sacloud API
type ListLicenseParam struct {
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

	options *flags.Flags
	input   Input
}

// NewListLicenseParam return new ListLicenseParam
func NewListLicenseParam() *ListLicenseParam {
	return &ListLicenseParam{}
}

// Initialize init ListLicenseParam
func (p *ListLicenseParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListLicenseParam) FillValueToSkeleton() {
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

func (p *ListLicenseParam) validate() error {
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
		validator := define.Resources["License"].Commands["list"].Params["id"].ValidateFunc
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
		errs := cli.ValidateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ListLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ListLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ListLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListLicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListLicenseParam) GetName() []string {
	return p.Name
}
func (p *ListLicenseParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListLicenseParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListLicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListLicenseParam) GetFrom() int {
	return p.From
}
func (p *ListLicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListLicenseParam) GetMax() int {
	return p.Max
}
func (p *ListLicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListLicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *ListLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ListLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ListLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ListLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListLicenseParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// CreateLicenseParam is input parameters for the sacloud API
type CreateLicenseParam struct {
	LicenseInfoId     sacloud.ID
	Name              string
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

	options *flags.Flags
	input   Input
}

// NewCreateLicenseParam return new CreateLicenseParam
func NewCreateLicenseParam() *CreateLicenseParam {
	return &CreateLicenseParam{}
}

// Initialize init CreateLicenseParam
func (p *CreateLicenseParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *CreateLicenseParam) FillValueToSkeleton() {
	if util.IsEmpty(p.LicenseInfoId) {
		p.LicenseInfoId = sacloud.ID(0)
	}
	if util.IsEmpty(p.Name) {
		p.Name = ""
	}
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

}

func (p *CreateLicenseParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
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
		errs := cli.ValidateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *CreateLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *CreateLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *CreateLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *CreateLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateLicenseParam) SetLicenseInfoId(v sacloud.ID) {
	p.LicenseInfoId = v
}

func (p *CreateLicenseParam) GetLicenseInfoId() sacloud.ID {
	return p.LicenseInfoId
}
func (p *CreateLicenseParam) SetName(v string) {
	p.Name = v
}

func (p *CreateLicenseParam) GetName() string {
	return p.Name
}
func (p *CreateLicenseParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateLicenseParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CreateLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *CreateLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CreateLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CreateLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *CreateLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateLicenseParam) GetFormat() string {
	return p.Format
}
func (p *CreateLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateLicenseParam) GetQuery() string {
	return p.Query
}
func (p *CreateLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CreateLicenseParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *CreateLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ReadLicenseParam is input parameters for the sacloud API
type ReadLicenseParam struct {
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

	options *flags.Flags
	input   Input
}

// NewReadLicenseParam return new ReadLicenseParam
func NewReadLicenseParam() *ReadLicenseParam {
	return &ReadLicenseParam{}
}

// Initialize init ReadLicenseParam
func (p *ReadLicenseParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadLicenseParam) FillValueToSkeleton() {
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

func (p *ReadLicenseParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateSakuraID
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
		errs := cli.ValidateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ReadLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ReadLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ReadLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ReadLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ReadLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ReadLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ReadLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadLicenseParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadLicenseParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadLicenseParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// UpdateLicenseParam is input parameters for the sacloud API
type UpdateLicenseParam struct {
	Name              string
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

	options *flags.Flags
	input   Input
}

// NewUpdateLicenseParam return new UpdateLicenseParam
func NewUpdateLicenseParam() *UpdateLicenseParam {
	return &UpdateLicenseParam{}
}

// Initialize init UpdateLicenseParam
func (p *UpdateLicenseParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *UpdateLicenseParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Name) {
		p.Name = ""
	}
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

func (p *UpdateLicenseParam) validate() error {
	var errors []error

	{
		validator := define.Resources["License"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateSakuraID
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
		errs := cli.ValidateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *UpdateLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *UpdateLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *UpdateLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *UpdateLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateLicenseParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateLicenseParam) GetName() string {
	return p.Name
}
func (p *UpdateLicenseParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateLicenseParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *UpdateLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *UpdateLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *UpdateLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *UpdateLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateLicenseParam) GetFormat() string {
	return p.Format
}
func (p *UpdateLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateLicenseParam) GetQuery() string {
	return p.Query
}
func (p *UpdateLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *UpdateLicenseParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *UpdateLicenseParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *UpdateLicenseParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *UpdateLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// DeleteLicenseParam is input parameters for the sacloud API
type DeleteLicenseParam struct {
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

	options *flags.Flags
	input   Input
}

// NewDeleteLicenseParam return new DeleteLicenseParam
func NewDeleteLicenseParam() *DeleteLicenseParam {
	return &DeleteLicenseParam{}
}

// Initialize init DeleteLicenseParam
func (p *DeleteLicenseParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *DeleteLicenseParam) FillValueToSkeleton() {
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

func (p *DeleteLicenseParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateSakuraID
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
		errs := cli.ValidateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *DeleteLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *DeleteLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *DeleteLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *DeleteLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteLicenseParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteLicenseParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteLicenseParam) GetFormat() string {
	return p.Format
}
func (p *DeleteLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteLicenseParam) GetQuery() string {
	return p.Query
}
func (p *DeleteLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteLicenseParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *DeleteLicenseParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *DeleteLicenseParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *DeleteLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}
