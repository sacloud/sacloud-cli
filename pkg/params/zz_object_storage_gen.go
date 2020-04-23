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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/flags"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
)

// ListObjectStorageParam is input parameters for the sacloud API
type ListObjectStorageParam struct {
	AccessKey         string
	SecretKey         string
	Bucket            string
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

// NewListObjectStorageParam return new ListObjectStorageParam
func NewListObjectStorageParam() *ListObjectStorageParam {
	return &ListObjectStorageParam{}
}

// Initialize init ListObjectStorageParam
func (p *ListObjectStorageParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListObjectStorageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListObjectStorageParam) FillValueToSkeleton() {
	if util.IsEmpty(p.AccessKey) {
		p.AccessKey = ""
	}
	if util.IsEmpty(p.SecretKey) {
		p.SecretKey = ""
	}
	if util.IsEmpty(p.Bucket) {
		p.Bucket = ""
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

func (p *ListObjectStorageParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--secret-key", p.SecretKey)
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

func (p *ListObjectStorageParam) ResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *ListObjectStorageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListObjectStorageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListObjectStorageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListObjectStorageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListObjectStorageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *ListObjectStorageParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListObjectStorageParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListObjectStorageParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListObjectStorageParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *ListObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *ListObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *ListObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *ListObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *ListObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *ListObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListObjectStorageParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListObjectStorageParam) GetParameters() string {
	return p.Parameters
}
func (p *ListObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListObjectStorageParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListObjectStorageParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListObjectStorageParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListObjectStorageParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListObjectStorageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListObjectStorageParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListObjectStorageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListObjectStorageParam) GetColumn() []string {
	return p.Column
}
func (p *ListObjectStorageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListObjectStorageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListObjectStorageParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListObjectStorageParam) GetFormat() string {
	return p.Format
}
func (p *ListObjectStorageParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListObjectStorageParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListObjectStorageParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListObjectStorageParam) GetQuery() string {
	return p.Query
}
func (p *ListObjectStorageParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListObjectStorageParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListObjectStorageParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// PutObjectStorageParam is input parameters for the sacloud API
type PutObjectStorageParam struct {
	AccessKey         string
	ContentType       string
	Recursive         bool
	SecretKey         string
	Bucket            string
	Assumeyes         bool
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool

	options *flags.Flags
	input   Input
}

// NewPutObjectStorageParam return new PutObjectStorageParam
func NewPutObjectStorageParam() *PutObjectStorageParam {
	return &PutObjectStorageParam{
		ContentType: "application/octet-stream"}
}

// Initialize init PutObjectStorageParam
func (p *PutObjectStorageParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PutObjectStorageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *PutObjectStorageParam) FillValueToSkeleton() {
	if util.IsEmpty(p.AccessKey) {
		p.AccessKey = ""
	}
	if util.IsEmpty(p.ContentType) {
		p.ContentType = ""
	}
	if util.IsEmpty(p.Recursive) {
		p.Recursive = false
	}
	if util.IsEmpty(p.SecretKey) {
		p.SecretKey = ""
	}
	if util.IsEmpty(p.Bucket) {
		p.Bucket = ""
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

}

func (p *PutObjectStorageParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--secret-key", p.SecretKey)
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

func (p *PutObjectStorageParam) ResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *PutObjectStorageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["put"]
}

func (p *PutObjectStorageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PutObjectStorageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PutObjectStorageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PutObjectStorageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *PutObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *PutObjectStorageParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["put"]
}

func (p *PutObjectStorageParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PutObjectStorageParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PutObjectStorageParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PutObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PutObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *PutObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *PutObjectStorageParam) SetContentType(v string) {
	p.ContentType = v
}

func (p *PutObjectStorageParam) GetContentType() string {
	return p.ContentType
}
func (p *PutObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *PutObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *PutObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *PutObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *PutObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *PutObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *PutObjectStorageParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PutObjectStorageParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PutObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PutObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PutObjectStorageParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *PutObjectStorageParam) GetParameters() string {
	return p.Parameters
}
func (p *PutObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PutObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PutObjectStorageParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PutObjectStorageParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PutObjectStorageParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PutObjectStorageParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// Changed usacloud v0系との互換性維持のための実装
func (p *PutObjectStorageParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// GetObjectStorageParam is input parameters for the sacloud API
type GetObjectStorageParam struct {
	AccessKey         string
	Recursive         bool
	SecretKey         string
	Bucket            string
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool

	options *flags.Flags
	input   Input
}

// NewGetObjectStorageParam return new GetObjectStorageParam
func NewGetObjectStorageParam() *GetObjectStorageParam {
	return &GetObjectStorageParam{}
}

// Initialize init GetObjectStorageParam
func (p *GetObjectStorageParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *GetObjectStorageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *GetObjectStorageParam) FillValueToSkeleton() {
	if util.IsEmpty(p.AccessKey) {
		p.AccessKey = ""
	}
	if util.IsEmpty(p.Recursive) {
		p.Recursive = false
	}
	if util.IsEmpty(p.SecretKey) {
		p.SecretKey = ""
	}
	if util.IsEmpty(p.Bucket) {
		p.Bucket = ""
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

}

func (p *GetObjectStorageParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--secret-key", p.SecretKey)
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

func (p *GetObjectStorageParam) ResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *GetObjectStorageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["get"]
}

func (p *GetObjectStorageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *GetObjectStorageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *GetObjectStorageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *GetObjectStorageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *GetObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *GetObjectStorageParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["get"]
}

func (p *GetObjectStorageParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *GetObjectStorageParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *GetObjectStorageParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *GetObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *GetObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *GetObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *GetObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *GetObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *GetObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *GetObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *GetObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *GetObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *GetObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *GetObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *GetObjectStorageParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *GetObjectStorageParam) GetParameters() string {
	return p.Parameters
}
func (p *GetObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *GetObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *GetObjectStorageParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *GetObjectStorageParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *GetObjectStorageParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *GetObjectStorageParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// Changed usacloud v0系との互換性維持のための実装
func (p *GetObjectStorageParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// DeleteObjectStorageParam is input parameters for the sacloud API
type DeleteObjectStorageParam struct {
	AccessKey         string
	Recursive         bool
	SecretKey         string
	Bucket            string
	Assumeyes         bool
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool

	options *flags.Flags
	input   Input
}

// NewDeleteObjectStorageParam return new DeleteObjectStorageParam
func NewDeleteObjectStorageParam() *DeleteObjectStorageParam {
	return &DeleteObjectStorageParam{}
}

// Initialize init DeleteObjectStorageParam
func (p *DeleteObjectStorageParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteObjectStorageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *DeleteObjectStorageParam) FillValueToSkeleton() {
	if util.IsEmpty(p.AccessKey) {
		p.AccessKey = ""
	}
	if util.IsEmpty(p.Recursive) {
		p.Recursive = false
	}
	if util.IsEmpty(p.SecretKey) {
		p.SecretKey = ""
	}
	if util.IsEmpty(p.Bucket) {
		p.Bucket = ""
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

}

func (p *DeleteObjectStorageParam) validate() error {
	var errors []error

	{
		validator := cli.ValidateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := cli.ValidateRequired
		errs := validator("--secret-key", p.SecretKey)
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

func (p *DeleteObjectStorageParam) ResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *DeleteObjectStorageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteObjectStorageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteObjectStorageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteObjectStorageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteObjectStorageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *DeleteObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *DeleteObjectStorageParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteObjectStorageParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteObjectStorageParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteObjectStorageParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *DeleteObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *DeleteObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *DeleteObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *DeleteObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *DeleteObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *DeleteObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *DeleteObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *DeleteObjectStorageParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteObjectStorageParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteObjectStorageParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteObjectStorageParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteObjectStorageParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteObjectStorageParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteObjectStorageParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteObjectStorageParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// Changed usacloud v0系との互換性維持のための実装
func (p *DeleteObjectStorageParam) Changed(name string) bool {
	return p.input.Changed(name)
}
