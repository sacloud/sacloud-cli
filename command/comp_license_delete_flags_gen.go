// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func LicenseDeleteCompleteFlags(ctx Context, params *DeleteLicenseParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "force", "f":
		comp = define.Resources["License"].Commands["delete"].Params["force"].CompleteFunc
	case "id":
		comp = define.Resources["License"].Commands["delete"].Params["id"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
