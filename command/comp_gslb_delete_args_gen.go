// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func GSLBDeleteCompleteArgs(ctx Context, params *DeleteGSLBParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
	}

}
