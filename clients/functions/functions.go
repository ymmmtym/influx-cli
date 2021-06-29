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
	api.OrganizationsApi
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
		return fmt.Errorf("failed to list functions: %w", err)
	}

	c.printFunctions(functionPrintOpts{functions: functions.GetFunctions()})

	return nil
}

type CreateParams struct {
	clients.OrgParams
	Name        string
	Description string
	Script      string
	Language    string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	reqBody := api.FunctionCreateRequest{
		Name:   params.Name,
		Script: params.Script,
	}

	var lang api.FunctionLanguage
	if err := lang.UnmarshalJSON([]byte(fmt.Sprintf("%q", params.Language))); err != nil {
		return err
	}
	reqBody.Language = lang

	if params.Description != "" {
		reqBody.Description = &params.Description
	}

	// The org ID will be obtained based on the org name if an org name is
	// provided but not an org ID.
	if params.OrgID.Valid() {
		reqBody.OrgID = params.OrgID.String()
	} else {
		orgName := params.OrgName
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		res, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return fmt.Errorf("failed to look up ID for org %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no org found with name %q", orgName)
		}
		reqBody.OrgID = res.GetOrgs()[0].GetId()
	}

	function, err := c.PostFunctions(ctx).FunctionCreateRequest(reqBody).Execute()

	if err != nil {
		return fmt.Errorf("failed to create function: %w", err)
	}

	c.printFunctions(functionPrintOpts{function: &function})

	return nil
}

type GetParams struct {
	ID string
}

func (c Client) Get(ctx context.Context, params *GetParams) error {
	function, err := c.GetFunctionsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to get function %q: %w", params.ID, err)
	}

	c.printFunctions(functionPrintOpts{function: &function})

	return nil
}

type DeleteParams struct {
	ID string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	function, err := c.GetFunctionsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete function %q: %w", params.ID, err)
	}

	err = c.DeleteFunctionsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete function %q: %w", params.ID, err)
	}

	c.printFunctions(functionPrintOpts{function: &function})

	return nil
}

type UpdateParams struct {
	ID          string
	Name        string
	Description string
	Script      string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	reqBody := api.FunctionUpdateRequest{}

	if params.Name != "" {
		reqBody.Name = &params.Name
	}
	if params.Description != "" {
		reqBody.Description = &params.Description // TODO: updating the description doesn't seem to work
	}
	if params.Script != "" {
		reqBody.Script = &params.Script
	}

	req := c.PatchFunctionsID(ctx, params.ID)
	function, err := req.FunctionUpdateRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to update managed function %q: %w", params.ID, err)
	}

	c.printFunctions(functionPrintOpts{function: &function})

	return nil
}

type InvokeParams struct {
	ID         string
	FuncParams map[string]interface{}
}

func (c Client) Invoke(ctx context.Context, params *InvokeParams) error {
	reqBody := api.FunctionInvocationParams{}

	if len(params.FuncParams) > 0 {
		reqBody.SetParams(params.FuncParams)
	}

	req := c.PostFunctionsIDInvoke(ctx, params.ID)
	res, err := req.FunctionInvocationParams(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to execute managed function %q: %w", params.ID, err)
	}

	// Placeholder for more sophisticated result handling. Could be adapted from
	// the query result printer perhaps.
	fmt.Println(res)

	return nil
}

func (c Client) printFunctions(opts functionPrintOpts) error {
	if opts.function != nil {
		opts.functions = append(opts.functions, *opts.function)
	}

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
		"Language",
	}

	var rows []map[string]interface{}
	for _, t := range opts.functions {
		row := map[string]interface{}{
			"ID":              *t.Id,
			"Name":            t.Name,
			"Description":     *t.Description,
			"Organization ID": t.OrgID,
			"Language":        *t.Language,
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}
