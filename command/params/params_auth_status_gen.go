// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ShowAuthStatusParam is input parameters for the sacloud API
type ShowAuthStatusParam struct {
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

// NewShowAuthStatusParam return new ShowAuthStatusParam
func NewShowAuthStatusParam() *ShowAuthStatusParam {
	return &ShowAuthStatusParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ShowAuthStatusParam) FillValueToSkeleton() {
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
func (p *ShowAuthStatusParam) Validate() []error {
	errors := []error{}

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

func (p *ShowAuthStatusParam) GetResourceDef() *schema.Resource {
	return define.Resources["AuthStatus"]
}

func (p *ShowAuthStatusParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["show"]
}

func (p *ShowAuthStatusParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ShowAuthStatusParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ShowAuthStatusParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ShowAuthStatusParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ShowAuthStatusParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ShowAuthStatusParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ShowAuthStatusParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ShowAuthStatusParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ShowAuthStatusParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ShowAuthStatusParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ShowAuthStatusParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ShowAuthStatusParam) GetOutputType() string {
	return p.OutputType
}
func (p *ShowAuthStatusParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ShowAuthStatusParam) GetColumn() []string {
	return p.Column
}
func (p *ShowAuthStatusParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ShowAuthStatusParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ShowAuthStatusParam) SetFormat(v string) {
	p.Format = v
}

func (p *ShowAuthStatusParam) GetFormat() string {
	return p.Format
}
func (p *ShowAuthStatusParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ShowAuthStatusParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ShowAuthStatusParam) SetQuery(v string) {
	p.Query = v
}

func (p *ShowAuthStatusParam) GetQuery() string {
	return p.Query
}
func (p *ShowAuthStatusParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ShowAuthStatusParam) GetQueryFile() string {
	return p.QueryFile
}
