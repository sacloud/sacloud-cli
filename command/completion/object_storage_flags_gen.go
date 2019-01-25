// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ObjectStorageListCompleteFlags(ctx command.Context, params *params.ListObjectStorageParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "access-key":
		param := define.Resources["ObjectStorage"].Commands["list"].BuildedParams().Get("access-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "secret-key":
		param := define.Resources["ObjectStorage"].Commands["list"].BuildedParams().Get("secret-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "bucket":
		param := define.Resources["ObjectStorage"].Commands["list"].BuildedParams().Get("bucket")
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

func ObjectStoragePutCompleteFlags(ctx command.Context, params *params.PutObjectStorageParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "access-key":
		param := define.Resources["ObjectStorage"].Commands["put"].BuildedParams().Get("access-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "content-type":
		param := define.Resources["ObjectStorage"].Commands["put"].BuildedParams().Get("content-type")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "recursive", "r":
		param := define.Resources["ObjectStorage"].Commands["put"].BuildedParams().Get("recursive")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "secret-key":
		param := define.Resources["ObjectStorage"].Commands["put"].BuildedParams().Get("secret-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "bucket":
		param := define.Resources["ObjectStorage"].Commands["put"].BuildedParams().Get("bucket")
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

func ObjectStorageGetCompleteFlags(ctx command.Context, params *params.GetObjectStorageParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "access-key":
		param := define.Resources["ObjectStorage"].Commands["get"].BuildedParams().Get("access-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "recursive", "r":
		param := define.Resources["ObjectStorage"].Commands["get"].BuildedParams().Get("recursive")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "secret-key":
		param := define.Resources["ObjectStorage"].Commands["get"].BuildedParams().Get("secret-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "bucket":
		param := define.Resources["ObjectStorage"].Commands["get"].BuildedParams().Get("bucket")
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

func ObjectStorageDeleteCompleteFlags(ctx command.Context, params *params.DeleteObjectStorageParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "access-key":
		param := define.Resources["ObjectStorage"].Commands["delete"].BuildedParams().Get("access-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "recursive", "r":
		param := define.Resources["ObjectStorage"].Commands["delete"].BuildedParams().Get("recursive")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "secret-key":
		param := define.Resources["ObjectStorage"].Commands["delete"].BuildedParams().Get("secret-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "bucket":
		param := define.Resources["ObjectStorage"].Commands["delete"].BuildedParams().Get("bucket")
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
