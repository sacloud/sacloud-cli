// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func PacketFilterUpdateCompleteArgs(ctx Context, params *UpdatePacketFilterParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.PacketFilters {
		fmt.Println(res.PacketFilters[i].ID)
	}

}
