package main

import (
	"fmt"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/functions"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newFunctionsCommand() cli.Command {
	return cli.Command{
		Name:  "functions",
		Usage: "Functions management commands",
		Subcommands: []cli.Command{
			newFunctionsListCommand(),
			newFunctionsCreateCommand(),
			newFunctionsGetCommand(),
			newFunctionsDeleteCommand(),
			newFunctionsUpdateCommand(),
			newFunctionsInvokeCommand(),
		},
	}
}

func newFunctionsListCommand() cli.Command {
	var params functions.ListParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:    "list",
		Usage:   "List functions",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "limit",
				Usage:       "Limit the number of results returned",
				Destination: &params.Limit,
			},
			&cli.StringFlag{
				Name:        "offset",
				Usage:       "Offset for pagination",
				Destination: &params.Offset,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := functions.Client{
				CLI:          getCLI(ctx),
				FunctionsApi: api.FunctionsApi.OnlyCloud(),
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newFunctionsCreateCommand() cli.Command {
	var params functions.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:   "create",
		Usage:  "Create a new managed function",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name of the new managed function",
				Destination: &params.Name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "Description of the new managed function",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "script",
				Usage:       "Script to be executed by the function",
				Destination: &params.Script,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "language",
				Usage:       "Language of the function's script",
				Destination: &params.Language,
				Required:    true,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := functions.Client{
				CLI:              getCLI(ctx),
				FunctionsApi:     api.FunctionsApi.OnlyCloud(),
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newFunctionsGetCommand() cli.Command {
	var params functions.GetParams

	return cli.Command{
		Name:   "get",
		Usage:  "Get the information about a single function",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the function to get",
				Destination: &params.ID,
				Required:    true,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := functions.Client{
				CLI:          getCLI(ctx),
				FunctionsApi: api.FunctionsApi.OnlyCloud(),
			}
			return client.Get(getContext(ctx), &params)
		},
	}
}

func newFunctionsDeleteCommand() cli.Command {
	var params functions.DeleteParams

	return cli.Command{
		Name:   "delete",
		Usage:  "Delete a managed function by ID",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the function to delete",
				Destination: &params.ID,
				Required:    true,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := functions.Client{
				CLI:          getCLI(ctx),
				FunctionsApi: api.FunctionsApi.OnlyCloud(),
			}
			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newFunctionsInvokeCommand() cli.Command {
	var params functions.InvokeParams

	return cli.Command{
		Name:   "invoke",
		Usage:  "Invoke a managed function by ID",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the function to invoke",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.StringSliceFlag{
				Name:  "params",
				Usage: "Params to use for invoking the function, in the form of key1=val1,key2=val2 etc.",
			},
		),
		Action: func(ctx *cli.Context) error {
			// This is the list of values to be passed as the "params" key for
			// invoking the function; not to be confused with the CLI parameters for
			// the invoke command which also include the function ID
			rawFuncParams := ctx.StringSlice("params")

			// If any params were passed in like --params key1=val1,key2=val2 split
			// those using the comma
			splitParams := []string{}
			for _, p := range rawFuncParams {
				splitParams = append(splitParams, strings.Split(p, ",")...)
			}

			funcParamsMap := map[string]interface{}{}

			for _, p := range splitParams {
				param := strings.Split(p, "=")
				if len(param) != 2 {
					// Error out if the parameter is not in the form of key=val.
					return fmt.Errorf("unable to parse parameter %q", p)
				}
				funcParamsMap[param[0]] = param[1]
			}

			params.FuncParams = funcParamsMap

			api := getAPI(ctx)
			client := functions.Client{
				CLI:          getCLI(ctx),
				FunctionsApi: api.FunctionsApi.OnlyCloud(),
			}
			return client.Invoke(getContext(ctx), &params)
		},
	}
}

func newFunctionsUpdateCommand() cli.Command {
	var params functions.UpdateParams

	return cli.Command{
		Name:   "update",
		Usage:  "Update a managed function by ID",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the function to update",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Updated name of the function",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "Updated description of the function",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "script",
				Usage:       "Updated script to be executed by the function",
				Destination: &params.Script,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := functions.Client{
				CLI:          getCLI(ctx),
				FunctionsApi: api.FunctionsApi.OnlyCloud(),
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}
