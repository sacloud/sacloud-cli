// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func SimpleMonitorReadCompleteArgs(ctx Context, params *ReadSimpleMonitorParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSimpleMonitorAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.SimpleMonitors {
		fmt.Println(res.SimpleMonitors[i].ID)
	}

}
