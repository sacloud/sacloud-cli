// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductLicenseParam is input parameters for the sacloud API
type ListProductLicenseParam struct {
	From       int
	Id         []int64
	Max        int
	Name       []string
	Sort       []string
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewListProductLicenseParam return new ListProductLicenseParam
func NewListProductLicenseParam() *ListProductLicenseParam {
	return &ListProductLicenseParam{}
}

// Validate checks current values in model
func (p *ListProductLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["ProductLicense"].Commands["list"].Params["id"].ValidateFunc
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
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
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

func (p *ListProductLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ListProductLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListProductLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListProductLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListProductLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListProductLicenseParam) GetOutputFormat() string {
	return "table"
}

func (p *ListProductLicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductLicenseParam) GetFrom() int {
	return p.From
}
func (p *ListProductLicenseParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListProductLicenseParam) GetId() []int64 {
	return p.Id
}
func (p *ListProductLicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductLicenseParam) GetMax() int {
	return p.Max
}
func (p *ListProductLicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductLicenseParam) GetName() []string {
	return p.Name
}
func (p *ListProductLicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductLicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductLicenseParam) GetFormat() string {
	return p.Format
}

// ReadProductLicenseParam is input parameters for the sacloud API
type ReadProductLicenseParam struct {
	Id         int64
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewReadProductLicenseParam return new ReadProductLicenseParam
func NewReadProductLicenseParam() *ReadProductLicenseParam {
	return &ReadProductLicenseParam{}
}

// Validate checks current values in model
func (p *ReadProductLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductLicense"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
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

func (p *ReadProductLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ReadProductLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadProductLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadProductLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadProductLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadProductLicenseParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadProductLicenseParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadProductLicenseParam) GetId() int64 {
	return p.Id
}
func (p *ReadProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductLicenseParam) GetFormat() string {
	return p.Format
}
