// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func ProductLicenseReadCompleteArgs(ctx Context, params *ReadProductLicenseParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductLicenseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.LicenseInfo {
		fmt.Println(res.LicenseInfo[i].ID)
	}

}
