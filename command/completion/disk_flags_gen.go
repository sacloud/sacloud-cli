// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DiskListCompleteFlags(ctx command.Context, params *params.ListDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "scope":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("scope")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-archive-id":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("source-archive-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-disk-id":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("source-disk-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "storage":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("storage")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["Disk"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskCreateCompleteFlags(ctx command.Context, params *params.CreateDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "plan":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("plan")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "connection":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("connection")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-archive-id":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("source-archive-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-disk-id":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("source-disk-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "size":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("size")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "distant-from":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("distant-from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["Disk"].Commands["create"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskReadCompleteFlags(ctx command.Context, params *params.ReadDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Disk"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskUpdateCompleteFlags(ctx command.Context, params *params.UpdateDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "connection":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("connection")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskDeleteCompleteFlags(ctx command.Context, params *params.DeleteDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Disk"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskEditCompleteFlags(ctx command.Context, params *params.EditDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "hostname":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("hostname")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "password":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("password")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ssh-key-ids":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("ssh-key-ids")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disable-password-auth", "disable-pw-auth":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("disable-password-auth")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ipaddress", "ip":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "default-route", "gateway":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("default-route")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "nw-masklen", "network-masklen":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("nw-masklen")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "startup-script-ids", "note-ids":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("startup-script-ids")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["edit"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskReinstallFromArchiveCompleteFlags(ctx command.Context, params *params.ReinstallFromArchiveDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "source-archive-id":
		param := define.Resources["Disk"].Commands["reinstall-from-archive"].BuildedParams().Get("source-archive-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "distant-from":
		param := define.Resources["Disk"].Commands["reinstall-from-archive"].BuildedParams().Get("distant-from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["reinstall-from-archive"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["reinstall-from-archive"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskReinstallFromDiskCompleteFlags(ctx command.Context, params *params.ReinstallFromDiskDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "source-disk-id":
		param := define.Resources["Disk"].Commands["reinstall-from-disk"].BuildedParams().Get("source-disk-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "distant-from":
		param := define.Resources["Disk"].Commands["reinstall-from-disk"].BuildedParams().Get("distant-from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["reinstall-from-disk"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["reinstall-from-disk"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskReinstallToBlankCompleteFlags(ctx command.Context, params *params.ReinstallToBlankDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "distant-from":
		param := define.Resources["Disk"].Commands["reinstall-to-blank"].BuildedParams().Get("distant-from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["reinstall-to-blank"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["reinstall-to-blank"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskServerConnectCompleteFlags(ctx command.Context, params *params.ServerConnectDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-id":
		param := define.Resources["Disk"].Commands["server-connect"].BuildedParams().Get("server-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["Disk"].Commands["server-connect"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["server-connect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskServerDisconnectCompleteFlags(ctx command.Context, params *params.ServerDisconnectDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Disk"].Commands["server-disconnect"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["server-disconnect"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskMonitorCompleteFlags(ctx command.Context, params *params.MonitorDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Disk"].Commands["monitor"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "end":
		param := define.Resources["Disk"].Commands["monitor"].BuildedParams().Get("end")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["monitor"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "key-format":
		param := define.Resources["Disk"].Commands["monitor"].BuildedParams().Get("key-format")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "start":
		param := define.Resources["Disk"].Commands["monitor"].BuildedParams().Get("start")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DiskWaitForCopyCompleteFlags(ctx command.Context, params *params.WaitForCopyDiskParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Disk"].Commands["wait-for-copy"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Disk"].Commands["wait-for-copy"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
