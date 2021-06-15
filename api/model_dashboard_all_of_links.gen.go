/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DashboardAllOfLinks struct for DashboardAllOfLinks
type DashboardAllOfLinks struct {
	// URI of resource.
	Self *string `json:"self,omitempty"`
	// URI of resource.
	Cells *string `json:"cells,omitempty"`
	// URI of resource.
	Members *string `json:"members,omitempty"`
	// URI of resource.
	Owners *string `json:"owners,omitempty"`
	// URI of resource.
	Labels *string `json:"labels,omitempty"`
	// URI of resource.
	Org *string `json:"org,omitempty"`
}

// NewDashboardAllOfLinks instantiates a new DashboardAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardAllOfLinks() *DashboardAllOfLinks {
	this := DashboardAllOfLinks{}
	return &this
}

// NewDashboardAllOfLinksWithDefaults instantiates a new DashboardAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardAllOfLinksWithDefaults() *DashboardAllOfLinks {
	this := DashboardAllOfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *DashboardAllOfLinks) SetSelf(v string) {
	o.Self = &v
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetCells() string {
	if o == nil || o.Cells == nil {
		var ret string
		return ret
	}
	return *o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetCellsOk() (*string, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasCells() bool {
	if o != nil && o.Cells != nil {
		return true
	}

	return false
}

// SetCells gets a reference to the given string and assigns it to the Cells field.
func (o *DashboardAllOfLinks) SetCells(v string) {
	o.Cells = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetMembers() string {
	if o == nil || o.Members == nil {
		var ret string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetMembersOk() (*string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given string and assigns it to the Members field.
func (o *DashboardAllOfLinks) SetMembers(v string) {
	o.Members = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetOwners() string {
	if o == nil || o.Owners == nil {
		var ret string
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetOwnersOk() (*string, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given string and assigns it to the Owners field.
func (o *DashboardAllOfLinks) SetOwners(v string) {
	o.Owners = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetLabels() string {
	if o == nil || o.Labels == nil {
		var ret string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetLabelsOk() (*string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given string and assigns it to the Labels field.
func (o *DashboardAllOfLinks) SetLabels(v string) {
	o.Labels = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *DashboardAllOfLinks) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOfLinks) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *DashboardAllOfLinks) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *DashboardAllOfLinks) SetOrg(v string) {
	o.Org = &v
}

func (o DashboardAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardAllOfLinks struct {
	value *DashboardAllOfLinks
	isSet bool
}

func (v NullableDashboardAllOfLinks) Get() *DashboardAllOfLinks {
	return v.value
}

func (v *NullableDashboardAllOfLinks) Set(val *DashboardAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardAllOfLinks(val *DashboardAllOfLinks) *NullableDashboardAllOfLinks {
	return &NullableDashboardAllOfLinks{value: val, isSet: true}
}

func (v NullableDashboardAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}