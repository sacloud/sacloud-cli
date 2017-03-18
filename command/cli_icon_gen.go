// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	createParam := NewCreateIconParam()
	deleteParam := NewDeleteIconParam()
	listParam := NewListIconParam()
	readParam := NewReadIconParam()
	updateParam := NewUpdateIconParam()

	cliCommand := &cli.Command{
		Name:  "icon",
		Usage: "A manage commands of Icon",
		Subcommands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create Icon",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "image",
						Usage:       "[Required] set icon image",
						Destination: &createParam.Image,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &createParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &createParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &createParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, createParam)

					// Set option values for slice
					createParam.Tags = c.StringSlice("tags")
					createParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								IconCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										IconCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									IconCreateCompleteFlags(ctx, createParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							IconCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					createParam.Tags = c.StringSlice("tags")
					createParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := createParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), createParam)

					// Run command with params
					return IconCreate(ctx, createParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Icon",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "force",
						Aliases:     []string{"f"},
						Destination: &deleteParam.Force,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &deleteParam.Id,
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &deleteParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &deleteParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &deleteParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, deleteParam)

					// Set option values for slice
					deleteParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								IconDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										IconDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									IconDeleteCompleteFlags(ctx, deleteParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							IconDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					deleteParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := deleteParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), deleteParam)

					// confirm
					if !deleteParam.Force && !confirmContinue("delete this") {
						return nil
					}

					// Run command with params
					return IconDelete(ctx, deleteParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Icon",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.StringFlag{
						Name:        "scope",
						Usage:       "set filter by scope('user' or 'shared')",
						Destination: &listParam.Scope,
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &listParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &listParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &listParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, listParam)

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								IconListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										IconListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									IconListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							IconListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return IconList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Icon",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &readParam.Id,
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &readParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &readParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &readParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, readParam)

					// Set option values for slice
					readParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								IconReadCompleteArgs(ctx, readParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										IconReadCompleteArgs(ctx, readParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									IconReadCompleteFlags(ctx, readParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							IconReadCompleteArgs(ctx, readParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					readParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), readParam)

					// Run command with params
					return IconRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Icon",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &updateParam.Id,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "set resource display name",
						Destination: &updateParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &updateParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &updateParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &updateParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, updateParam)

					// Set option values for slice
					updateParam.Tags = c.StringSlice("tags")
					updateParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								IconUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										IconUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									IconUpdateCompleteFlags(ctx, updateParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							IconUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					updateParam.Tags = c.StringSlice("tags")
					updateParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := updateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), updateParam)

					// Run command with params
					return IconUpdate(ctx, updateParam)
				},
			},
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("icon", &schema.Category{
		Key:         "commonitem",
		DisplayName: "Common items",
		Order:       60,
	})

	// build Category-Command mapping

	appendCommandCategoryMap("icon", "create", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("icon", "delete", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("icon", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("icon", "read", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("icon", "update", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	appendFlagCategoryMap("icon", "create", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "image", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "create", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "force", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "delete", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "from", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "max", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "scope", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "list", "sort", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "read", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "read", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "read", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "read", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("icon", "update", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
