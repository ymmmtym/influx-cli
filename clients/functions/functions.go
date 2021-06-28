package functions

import (
	"context"
	"fmt"
	"strconv"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.FunctionsApi
}

type functionPrintOpts struct {
	function  *api.Function
	functions []api.Function
}

type ListParams struct {
	clients.OrgParams
	Limit  string
	Offset string
}

// This operation doesn't work on the server right now. How are we going to
// handle pagination?
func (c Client) List(ctx context.Context, params *ListParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	req := c.GetFunctions(ctx)
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}

	if params.Limit != "" {
		l, err := strconv.Atoi(params.Limit)
		if err != nil {
			return fmt.Errorf("invalid value provided for limit: %w", err)
		}

		req = req.Limit(int32(l))
	}
	if params.Offset != "" {
		o, err := strconv.Atoi(params.Offset)
		if err != nil {
			return fmt.Errorf("invalid value provided for offset: %w", err)
		}

		req = req.Offset(int32(o))
	}

	functions, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list dbrps: %w", err)
	}

	c.printFunctions(functionPrintOpts{functions: functions.GetFunctions()})

	return nil
}

func (c Client) printFunctions(opts functionPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{} = opts.functions
		if opts.functions != nil {
			v = opts.functions
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Name",
		"Description",
		"Organization ID",
		"Script", // We probably won't actually want this to be output with the list
		"Language",
		"URL",
		"Created At",
		"Updated At",
	}

	if opts.function != nil {
		opts.functions = append(opts.functions, *opts.function)
	}

	var rows []map[string]interface{}
	for _, t := range opts.functions {
		row := map[string]interface{}{
			"ID":              t.Id,
			"Name":            t.Name,
			"Description":     t.Description,
			"Organization ID": t.OrgID,
			"Script":          t.Script, // We probably won't actually want this to be output with the list
			"Language":        t.Language,
			"URL":             t.Url,
			"Created At":      t.CreatedAt,
			"Updated At":      t.UpdatedAt,
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}
