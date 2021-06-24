/*
 * Influx API Service (V1 compatible endpoints)
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// InfluxQLContentType the model 'InfluxQLContentType'
type InfluxQLContentType string

// List of InfluxQLContentType
const (
	INFLUXQLCONTENTTYPE_TEXT_CSV         InfluxQLContentType = "text/csv"
	INFLUXQLCONTENTTYPE_APPLICATION_JSON InfluxQLContentType = "application/json"
)

func (v *InfluxQLContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InfluxQLContentType(value)
	for _, existing := range []InfluxQLContentType{"text/csv", "application/json"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InfluxQLContentType", value)
}

// Ptr returns reference to InfluxQLContentType value
func (v InfluxQLContentType) Ptr() *InfluxQLContentType {
	return &v
}

type NullableInfluxQLContentType struct {
	value *InfluxQLContentType
	isSet bool
}

func (v NullableInfluxQLContentType) Get() *InfluxQLContentType {
	return v.value
}

func (v *NullableInfluxQLContentType) Set(val *InfluxQLContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableInfluxQLContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableInfluxQLContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfluxQLContentType(val *InfluxQLContentType) *NullableInfluxQLContentType {
	return &NullableInfluxQLContentType{value: val, isSet: true}
}

func (v NullableInfluxQLContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfluxQLContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}