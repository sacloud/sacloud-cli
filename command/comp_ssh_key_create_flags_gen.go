// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func SSHKeyCreateCompleteFlags(ctx Context, params *CreateSSHKeyParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["SSHKey"].Commands["create"].Params["description"].CompleteFunc
	case "name":
		comp = define.Resources["SSHKey"].Commands["create"].Params["name"].CompleteFunc
	case "public-key":
		comp = define.Resources["SSHKey"].Commands["create"].Params["public-key"].CompleteFunc
	case "public-key-content":
		comp = define.Resources["SSHKey"].Commands["create"].Params["public-key-content"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}