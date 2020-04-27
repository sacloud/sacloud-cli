// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-resource-finder'; DO NOT EDIT

package cmd

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
)

func findArchiveReadTargets(ctx cli.Context, param *params.ReadArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findArchiveUpdateTargets(ctx cli.Context, param *params.UpdateArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findArchiveDeleteTargets(ctx cli.Context, param *params.DeleteArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findArchiveUploadTargets(ctx cli.Context, param *params.UploadArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findArchiveDownloadTargets(ctx cli.Context, param *params.DownloadArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findArchiveFTPOpenTargets(ctx cli.Context, param *params.FTPOpenArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findArchiveFTPCloseTargets(ctx cli.Context, param *params.FTPCloseArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findArchiveWaitForCopyTargets(ctx cli.Context, param *params.WaitForCopyArchiveParam) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.NewArchiveOp(ctx.Client())

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Archives {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx, ctx.Zone(), &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key("Name"): search.ExactMatch(idOrName),
						},
					})
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Archives {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.ID)
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}
