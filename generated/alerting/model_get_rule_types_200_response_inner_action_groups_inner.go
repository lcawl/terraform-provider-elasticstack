/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
)

// checks if the GetRuleTypes200ResponseInnerActionGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRuleTypes200ResponseInnerActionGroupsInner{}

// GetRuleTypes200ResponseInnerActionGroupsInner struct for GetRuleTypes200ResponseInnerActionGroupsInner
type GetRuleTypes200ResponseInnerActionGroupsInner struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGetRuleTypes200ResponseInnerActionGroupsInner instantiates a new GetRuleTypes200ResponseInnerActionGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRuleTypes200ResponseInnerActionGroupsInner() *GetRuleTypes200ResponseInnerActionGroupsInner {
	this := GetRuleTypes200ResponseInnerActionGroupsInner{}
	return &this
}

// NewGetRuleTypes200ResponseInnerActionGroupsInnerWithDefaults instantiates a new GetRuleTypes200ResponseInnerActionGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRuleTypes200ResponseInnerActionGroupsInnerWithDefaults() *GetRuleTypes200ResponseInnerActionGroupsInner {
	this := GetRuleTypes200ResponseInnerActionGroupsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetRuleTypes200ResponseInnerActionGroupsInner) SetName(v string) {
	o.Name = &v
}

func (o GetRuleTypes200ResponseInnerActionGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRuleTypes200ResponseInnerActionGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetRuleTypes200ResponseInnerActionGroupsInner struct {
	value *GetRuleTypes200ResponseInnerActionGroupsInner
	isSet bool
}

func (v NullableGetRuleTypes200ResponseInnerActionGroupsInner) Get() *GetRuleTypes200ResponseInnerActionGroupsInner {
	return v.value
}

func (v *NullableGetRuleTypes200ResponseInnerActionGroupsInner) Set(val *GetRuleTypes200ResponseInnerActionGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRuleTypes200ResponseInnerActionGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRuleTypes200ResponseInnerActionGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRuleTypes200ResponseInnerActionGroupsInner(val *GetRuleTypes200ResponseInnerActionGroupsInner) *NullableGetRuleTypes200ResponseInnerActionGroupsInner {
	return &NullableGetRuleTypes200ResponseInnerActionGroupsInner{value: val, isSet: true}
}

func (v NullableGetRuleTypes200ResponseInnerActionGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRuleTypes200ResponseInnerActionGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}