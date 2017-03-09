// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func AutoBackupCreateCompleteFlags(ctx Context, params *CreateAutoBackupParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["description"].CompleteFunc
	case "disk-id":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["disk-id"].CompleteFunc
	case "generation":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["generation"].CompleteFunc
	case "icon-id":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["icon-id"].CompleteFunc
	case "name":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["name"].CompleteFunc
	case "start-hour":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["start-hour"].CompleteFunc
	case "tags":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["tags"].CompleteFunc
	case "weekdays":
		comp = define.Resources["AutoBackup"].Commands["create"].Params["weekdays"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}