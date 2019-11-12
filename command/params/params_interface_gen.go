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

// ListInterfaceParam is input parameters for the sacloud API
type ListInterfaceParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
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

// NewListInterfaceParam return new ListInterfaceParam
func NewListInterfaceParam() *ListInterfaceParam {
	return &ListInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListInterfaceParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
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
func (p *ListInterfaceParam) Validate() []error {
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
		validator := define.Resources["Interface"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *ListInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListInterfaceParam) SetName(v []string) {
	p.Name = v
}

func (p *ListInterfaceParam) GetName() []string {
	return p.Name
}
func (p *ListInterfaceParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListInterfaceParam) GetId() []int64 {
	return p.Id
}
func (p *ListInterfaceParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInterfaceParam) GetFrom() int {
	return p.From
}
func (p *ListInterfaceParam) SetMax(v int) {
	p.Max = v
}

func (p *ListInterfaceParam) GetMax() int {
	return p.Max
}
func (p *ListInterfaceParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListInterfaceParam) GetSort() []string {
	return p.Sort
}
func (p *ListInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListInterfaceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListInterfaceParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListInterfaceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListInterfaceParam) GetColumn() []string {
	return p.Column
}
func (p *ListInterfaceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListInterfaceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListInterfaceParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListInterfaceParam) GetFormat() string {
	return p.Format
}
func (p *ListInterfaceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListInterfaceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListInterfaceParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListInterfaceParam) GetQuery() string {
	return p.Query
}
func (p *ListInterfaceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListInterfaceParam) GetQueryFile() string {
	return p.QueryFile
}

// PacketFilterConnectInterfaceParam is input parameters for the sacloud API
type PacketFilterConnectInterfaceParam struct {
	PacketFilterId    int64  `json:"packet-filter-id"`
	Assumeyes         bool   `json:"assumeyes"`
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
	Id                int64  `json:"id"`
}

// NewPacketFilterConnectInterfaceParam return new PacketFilterConnectInterfaceParam
func NewPacketFilterConnectInterfaceParam() *PacketFilterConnectInterfaceParam {
	return &PacketFilterConnectInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PacketFilterConnectInterfaceParam) FillValueToSkeleton() {
	if isEmpty(p.PacketFilterId) {
		p.PacketFilterId = 0
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
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *PacketFilterConnectInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-connect"].Params["packet-filter-id"].ValidateFunc
		errs := validator("--packet-filter-id", p.PacketFilterId)
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

	return errors
}

func (p *PacketFilterConnectInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *PacketFilterConnectInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["packet-filter-connect"]
}

func (p *PacketFilterConnectInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PacketFilterConnectInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PacketFilterConnectInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PacketFilterConnectInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PacketFilterConnectInterfaceParam) SetPacketFilterId(v int64) {
	p.PacketFilterId = v
}

func (p *PacketFilterConnectInterfaceParam) GetPacketFilterId() int64 {
	return p.PacketFilterId
}
func (p *PacketFilterConnectInterfaceParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PacketFilterConnectInterfaceParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PacketFilterConnectInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PacketFilterConnectInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PacketFilterConnectInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PacketFilterConnectInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PacketFilterConnectInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PacketFilterConnectInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PacketFilterConnectInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *PacketFilterConnectInterfaceParam) GetId() int64 {
	return p.Id
}

// CreateInterfaceParam is input parameters for the sacloud API
type CreateInterfaceParam struct {
	ServerId          int64    `json:"server-id"`
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

// NewCreateInterfaceParam return new CreateInterfaceParam
func NewCreateInterfaceParam() *CreateInterfaceParam {
	return &CreateInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CreateInterfaceParam) FillValueToSkeleton() {
	if isEmpty(p.ServerId) {
		p.ServerId = 0
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
func (p *CreateInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--server-id", p.ServerId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["create"].Params["server-id"].ValidateFunc
		errs := validator("--server-id", p.ServerId)
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

func (p *CreateInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *CreateInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateInterfaceParam) SetServerId(v int64) {
	p.ServerId = v
}

func (p *CreateInterfaceParam) GetServerId() int64 {
	return p.ServerId
}
func (p *CreateInterfaceParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateInterfaceParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateInterfaceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateInterfaceParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateInterfaceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateInterfaceParam) GetColumn() []string {
	return p.Column
}
func (p *CreateInterfaceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateInterfaceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateInterfaceParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateInterfaceParam) GetFormat() string {
	return p.Format
}
func (p *CreateInterfaceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateInterfaceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateInterfaceParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateInterfaceParam) GetQuery() string {
	return p.Query
}
func (p *CreateInterfaceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CreateInterfaceParam) GetQueryFile() string {
	return p.QueryFile
}

// PacketFilterDisconnectInterfaceParam is input parameters for the sacloud API
type PacketFilterDisconnectInterfaceParam struct {
	PacketFilterId    int64  `json:"packet-filter-id"`
	Assumeyes         bool   `json:"assumeyes"`
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
	Id                int64  `json:"id"`
}

// NewPacketFilterDisconnectInterfaceParam return new PacketFilterDisconnectInterfaceParam
func NewPacketFilterDisconnectInterfaceParam() *PacketFilterDisconnectInterfaceParam {
	return &PacketFilterDisconnectInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PacketFilterDisconnectInterfaceParam) FillValueToSkeleton() {
	if isEmpty(p.PacketFilterId) {
		p.PacketFilterId = 0
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
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *PacketFilterDisconnectInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-disconnect"].Params["packet-filter-id"].ValidateFunc
		errs := validator("--packet-filter-id", p.PacketFilterId)
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

	return errors
}

func (p *PacketFilterDisconnectInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *PacketFilterDisconnectInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["packet-filter-disconnect"]
}

func (p *PacketFilterDisconnectInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PacketFilterDisconnectInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PacketFilterDisconnectInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PacketFilterDisconnectInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PacketFilterDisconnectInterfaceParam) SetPacketFilterId(v int64) {
	p.PacketFilterId = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetPacketFilterId() int64 {
	return p.PacketFilterId
}
func (p *PacketFilterDisconnectInterfaceParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PacketFilterDisconnectInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PacketFilterDisconnectInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PacketFilterDisconnectInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PacketFilterDisconnectInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetId() int64 {
	return p.Id
}

// ReadInterfaceParam is input parameters for the sacloud API
type ReadInterfaceParam struct {
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

// NewReadInterfaceParam return new ReadInterfaceParam
func NewReadInterfaceParam() *ReadInterfaceParam {
	return &ReadInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadInterfaceParam) FillValueToSkeleton() {
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
func (p *ReadInterfaceParam) Validate() []error {
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

func (p *ReadInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *ReadInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadInterfaceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadInterfaceParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadInterfaceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadInterfaceParam) GetColumn() []string {
	return p.Column
}
func (p *ReadInterfaceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadInterfaceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadInterfaceParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadInterfaceParam) GetFormat() string {
	return p.Format
}
func (p *ReadInterfaceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadInterfaceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadInterfaceParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadInterfaceParam) GetQuery() string {
	return p.Query
}
func (p *ReadInterfaceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadInterfaceParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadInterfaceParam) GetId() int64 {
	return p.Id
}

// UpdateInterfaceParam is input parameters for the sacloud API
type UpdateInterfaceParam struct {
	UserIpaddress     string   `json:"user-ipaddress"`
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

// NewUpdateInterfaceParam return new UpdateInterfaceParam
func NewUpdateInterfaceParam() *UpdateInterfaceParam {
	return &UpdateInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *UpdateInterfaceParam) FillValueToSkeleton() {
	if isEmpty(p.UserIpaddress) {
		p.UserIpaddress = ""
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
func (p *UpdateInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Interface"].Commands["update"].Params["user-ipaddress"].ValidateFunc
		errs := validator("--user-ipaddress", p.UserIpaddress)
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

func (p *UpdateInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *UpdateInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateInterfaceParam) SetUserIpaddress(v string) {
	p.UserIpaddress = v
}

func (p *UpdateInterfaceParam) GetUserIpaddress() string {
	return p.UserIpaddress
}
func (p *UpdateInterfaceParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateInterfaceParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateInterfaceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateInterfaceParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateInterfaceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateInterfaceParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateInterfaceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateInterfaceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateInterfaceParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateInterfaceParam) GetFormat() string {
	return p.Format
}
func (p *UpdateInterfaceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateInterfaceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateInterfaceParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateInterfaceParam) GetQuery() string {
	return p.Query
}
func (p *UpdateInterfaceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *UpdateInterfaceParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *UpdateInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateInterfaceParam) GetId() int64 {
	return p.Id
}

// DeleteInterfaceParam is input parameters for the sacloud API
type DeleteInterfaceParam struct {
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

// NewDeleteInterfaceParam return new DeleteInterfaceParam
func NewDeleteInterfaceParam() *DeleteInterfaceParam {
	return &DeleteInterfaceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteInterfaceParam) FillValueToSkeleton() {
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
func (p *DeleteInterfaceParam) Validate() []error {
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

func (p *DeleteInterfaceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *DeleteInterfaceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteInterfaceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteInterfaceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteInterfaceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteInterfaceParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteInterfaceParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteInterfaceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteInterfaceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteInterfaceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteInterfaceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteInterfaceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteInterfaceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteInterfaceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteInterfaceParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteInterfaceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteInterfaceParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteInterfaceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteInterfaceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteInterfaceParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteInterfaceParam) GetFormat() string {
	return p.Format
}
func (p *DeleteInterfaceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteInterfaceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteInterfaceParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteInterfaceParam) GetQuery() string {
	return p.Query
}
func (p *DeleteInterfaceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteInterfaceParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *DeleteInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteInterfaceParam) GetId() int64 {
	return p.Id
}
