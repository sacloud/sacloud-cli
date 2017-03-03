// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func SwitchUpdateCompleteArgs(ctx Context, params *UpdateSwitchParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
	}

}
