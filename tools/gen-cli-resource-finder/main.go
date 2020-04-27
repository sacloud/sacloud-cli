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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/sacloud/usacloud/tools"
)

var (
	destination = "src/github.com/sacloud/usacloud/pkg/cmd"
	ctx         = tools.NewGenerateContext()
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprint(os.Stderr, "\tgen-cli-resource-finder\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-cli-resource-finder: ")

	for _, resource := range ctx.Resources {
		filePath := filepath.Join(destination, resource.CLIResourceFinderSourceFileName())
		fileFullPath := filepath.Join(ctx.Gopath(), filePath)

		src, err := generateSource(resource)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}

		err = ioutil.WriteFile(fileFullPath, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
		fmt.Printf("generated: %s\n", filePath)
	}
}

func generateSource(resource *tools.Resource) (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, resource)
	return buf.String(), err
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-resource-finder'; DO NOT EDIT

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

{{ range .Commands }}{{ if .MultipleArgToIdParams }}
func {{ .ArgToIdFunc }}(ctx cli.Context, param *params.{{.InputParameterTypeName}}) ([]types.ID, error) {
	var ids []types.ID
	args := ctx.Args()
	client := sacloud.New{{.TargetAPIName}}Op(ctx.Client())

	if len(args) == 0 {
		{{ if .NoSelector -}}
		return ids, fmt.Errorf("ID or Name argument is required")
		{{ else -}}
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		res, err := client.Find(ctx{{if not .IsGlobal}}, ctx.Zone(){{ end }}, &sacloud.FindCondition{})
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.{{.FindResultFieldName}} {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.ID)
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
		{{ end -}}
	} else {
		{{if .SkipAfterSecondArgs -}}
		arg := args[0]
		{{ else -}}
		for _, arg := range args {
		{{ end -}}
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := types.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					res, err := client.Find(ctx{{if not .IsGlobal}}, ctx.Zone(){{ end }}, &sacloud.FindCondition{
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
					for _, v := range res.{{.FindResultFieldName}} {
						{{ if not .NoSelector }}if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) { {{ end }}
							ids = append(ids, v.ID)
						{{ if not .NoSelector }}} {{ end }}
					}
				}
			}
		{{ if not .SkipAfterSecondArgs }}
		}
		{{ end }}
	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	{{ if .RequireSingleID -}}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}
	{{ end }}
	return ids, nil
}
{{ end }}{{ end }}
`
