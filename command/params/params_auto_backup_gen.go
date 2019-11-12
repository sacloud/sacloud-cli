// Copyright 2017-2019 The Usacloud Authors
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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListAutoBackupParam is input parameters for the sacloud API
type ListAutoBackupParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	Tags              []string `json:"tags"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewListAutoBackupParam return new ListAutoBackupParam
func NewListAutoBackupParam() *ListAutoBackupParam {
	return &ListAutoBackupParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListAutoBackupParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *ListAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
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
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListAutoBackupParam) GetResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ListAutoBackupParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListAutoBackupParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListAutoBackupParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListAutoBackupParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListAutoBackupParam) SetName(v []string) {
	p.Name = v
}

func (p *ListAutoBackupParam) GetName() []string {
	return p.Name
}
func (p *ListAutoBackupParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListAutoBackupParam) GetId() []int64 {
	return p.Id
}
func (p *ListAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *ListAutoBackupParam) SetFrom(v int) {
	p.From = v
}

func (p *ListAutoBackupParam) GetFrom() int {
	return p.From
}
func (p *ListAutoBackupParam) SetMax(v int) {
	p.Max = v
}

func (p *ListAutoBackupParam) GetMax() int {
	return p.Max
}
func (p *ListAutoBackupParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListAutoBackupParam) GetSort() []string {
	return p.Sort
}
func (p *ListAutoBackupParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListAutoBackupParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListAutoBackupParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListAutoBackupParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *ListAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *ListAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *ListAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}

// CreateAutoBackupParam is input parameters for the sacloud API
type CreateAutoBackupParam struct {
	DiskId            int64    `json:"disk-id"`
	Weekdays          []string `json:"weekdays"`
	Generation        int      `json:"generation"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Tags              []string `json:"tags"`
	IconId            int64    `json:"icon-id"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewCreateAutoBackupParam return new CreateAutoBackupParam
func NewCreateAutoBackupParam() *CreateAutoBackupParam {
	return &CreateAutoBackupParam{

		Weekdays:   []string{"all"},
		Generation: 1,
	}
}

// FillValueToSkeleton fill values to empty fields
func (p *CreateAutoBackupParam) FillValueToSkeleton() {
	if isEmpty(p.DiskId) {
		p.DiskId = 0
	}
	if isEmpty(p.Weekdays) {
		p.Weekdays = []string{""}
	}
	if isEmpty(p.Generation) {
		p.Generation = 0
	}
	if isEmpty(p.Name) {
		p.Name = ""
	}
	if isEmpty(p.Description) {
		p.Description = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.IconId) {
		p.IconId = 0
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *CreateAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--disk-id", p.DiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["disk-id"].ValidateFunc
		errs := validator("--disk-id", p.DiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["weekdays"].ValidateFunc
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["generation"].ValidateFunc
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateAutoBackupParam) GetResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *CreateAutoBackupParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateAutoBackupParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateAutoBackupParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateAutoBackupParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateAutoBackupParam) SetDiskId(v int64) {
	p.DiskId = v
}

func (p *CreateAutoBackupParam) GetDiskId() int64 {
	return p.DiskId
}
func (p *CreateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *CreateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *CreateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *CreateAutoBackupParam) GetGeneration() int {
	return p.Generation
}
func (p *CreateAutoBackupParam) SetName(v string) {
	p.Name = v
}

func (p *CreateAutoBackupParam) GetName() string {
	return p.Name
}
func (p *CreateAutoBackupParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateAutoBackupParam) GetDescription() string {
	return p.Description
}
func (p *CreateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *CreateAutoBackupParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateAutoBackupParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateAutoBackupParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateAutoBackupParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateAutoBackupParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateAutoBackupParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *CreateAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *CreateAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *CreateAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CreateAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadAutoBackupParam is input parameters for the sacloud API
type ReadAutoBackupParam struct {
	Selector          []string `json:"selector"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
	Id                int64    `json:"id"`
}

// NewReadAutoBackupParam return new ReadAutoBackupParam
func NewReadAutoBackupParam() *ReadAutoBackupParam {
	return &ReadAutoBackupParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadAutoBackupParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *ReadAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
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
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadAutoBackupParam) GetResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *ReadAutoBackupParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadAutoBackupParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadAutoBackupParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadAutoBackupParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadAutoBackupParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadAutoBackupParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadAutoBackupParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadAutoBackupParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *ReadAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *ReadAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *ReadAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadAutoBackupParam) GetId() int64 {
	return p.Id
}

// UpdateAutoBackupParam is input parameters for the sacloud API
type UpdateAutoBackupParam struct {
	Weekdays          []string `json:"weekdays"`
	Generation        int      `json:"generation"`
	Selector          []string `json:"selector"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Tags              []string `json:"tags"`
	IconId            int64    `json:"icon-id"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
	Id                int64    `json:"id"`
}

// NewUpdateAutoBackupParam return new UpdateAutoBackupParam
func NewUpdateAutoBackupParam() *UpdateAutoBackupParam {
	return &UpdateAutoBackupParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *UpdateAutoBackupParam) FillValueToSkeleton() {
	if isEmpty(p.Weekdays) {
		p.Weekdays = []string{""}
	}
	if isEmpty(p.Generation) {
		p.Generation = 0
	}
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Name) {
		p.Name = ""
	}
	if isEmpty(p.Description) {
		p.Description = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.IconId) {
		p.IconId = 0
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *UpdateAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["weekdays"].ValidateFunc
		errs := validator("--weekdays", p.Weekdays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["generation"].ValidateFunc
		errs := validator("--generation", p.Generation)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["AutoBackup"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
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
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateAutoBackupParam) GetResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *UpdateAutoBackupParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateAutoBackupParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateAutoBackupParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateAutoBackupParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateAutoBackupParam) SetWeekdays(v []string) {
	p.Weekdays = v
}

func (p *UpdateAutoBackupParam) GetWeekdays() []string {
	return p.Weekdays
}
func (p *UpdateAutoBackupParam) SetGeneration(v int) {
	p.Generation = v
}

func (p *UpdateAutoBackupParam) GetGeneration() int {
	return p.Generation
}
func (p *UpdateAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *UpdateAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *UpdateAutoBackupParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateAutoBackupParam) GetName() string {
	return p.Name
}
func (p *UpdateAutoBackupParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateAutoBackupParam) GetDescription() string {
	return p.Description
}
func (p *UpdateAutoBackupParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateAutoBackupParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateAutoBackupParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateAutoBackupParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateAutoBackupParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateAutoBackupParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateAutoBackupParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateAutoBackupParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *UpdateAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *UpdateAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *UpdateAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *UpdateAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateAutoBackupParam) GetId() int64 {
	return p.Id
}

// DeleteAutoBackupParam is input parameters for the sacloud API
type DeleteAutoBackupParam struct {
	Selector          []string `json:"selector"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
	Id                int64    `json:"id"`
}

// NewDeleteAutoBackupParam return new DeleteAutoBackupParam
func NewDeleteAutoBackupParam() *DeleteAutoBackupParam {
	return &DeleteAutoBackupParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteAutoBackupParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *DeleteAutoBackupParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
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
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteAutoBackupParam) GetResourceDef() *schema.Resource {
	return define.Resources["AutoBackup"]
}

func (p *DeleteAutoBackupParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteAutoBackupParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteAutoBackupParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteAutoBackupParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteAutoBackupParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteAutoBackupParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *DeleteAutoBackupParam) GetSelector() []string {
	return p.Selector
}
func (p *DeleteAutoBackupParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteAutoBackupParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteAutoBackupParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteAutoBackupParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteAutoBackupParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteAutoBackupParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteAutoBackupParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteAutoBackupParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteAutoBackupParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteAutoBackupParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteAutoBackupParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteAutoBackupParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteAutoBackupParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteAutoBackupParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteAutoBackupParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteAutoBackupParam) GetFormat() string {
	return p.Format
}
func (p *DeleteAutoBackupParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteAutoBackupParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteAutoBackupParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteAutoBackupParam) GetQuery() string {
	return p.Query
}
func (p *DeleteAutoBackupParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteAutoBackupParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *DeleteAutoBackupParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteAutoBackupParam) GetId() int64 {
	return p.Id
}
