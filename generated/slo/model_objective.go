/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the Objective type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Objective{}

// Objective Defines properties for objective
type Objective struct {
	// the target objective between 0 and 1 excluded
	Target float32 `json:"target"`
	// the target objective for each slice when using a timeslices budgeting method
	TimeslicesTarget *float32 `json:"timeslicesTarget,omitempty"`
	// the duration of each slice when using a timeslices budgeting method, as {duraton}{unit}
	TimeslicesWindow *string `json:"timeslicesWindow,omitempty"`
}

// NewObjective instantiates a new Objective object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjective(target float32) *Objective {
	this := Objective{}
	this.Target = target
	return &this
}

// NewObjectiveWithDefaults instantiates a new Objective object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectiveWithDefaults() *Objective {
	this := Objective{}
	return &this
}

// GetTarget returns the Target field value
func (o *Objective) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Objective) GetTargetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Objective) SetTarget(v float32) {
	o.Target = v
}

// GetTimeslicesTarget returns the TimeslicesTarget field value if set, zero value otherwise.
func (o *Objective) GetTimeslicesTarget() float32 {
	if o == nil || IsNil(o.TimeslicesTarget) {
		var ret float32
		return ret
	}
	return *o.TimeslicesTarget
}

// GetTimeslicesTargetOk returns a tuple with the TimeslicesTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Objective) GetTimeslicesTargetOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeslicesTarget) {
		return nil, false
	}
	return o.TimeslicesTarget, true
}

// HasTimeslicesTarget returns a boolean if a field has been set.
func (o *Objective) HasTimeslicesTarget() bool {
	if o != nil && !IsNil(o.TimeslicesTarget) {
		return true
	}

	return false
}

// SetTimeslicesTarget gets a reference to the given float32 and assigns it to the TimeslicesTarget field.
func (o *Objective) SetTimeslicesTarget(v float32) {
	o.TimeslicesTarget = &v
}

// GetTimeslicesWindow returns the TimeslicesWindow field value if set, zero value otherwise.
func (o *Objective) GetTimeslicesWindow() string {
	if o == nil || IsNil(o.TimeslicesWindow) {
		var ret string
		return ret
	}
	return *o.TimeslicesWindow
}

// GetTimeslicesWindowOk returns a tuple with the TimeslicesWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Objective) GetTimeslicesWindowOk() (*string, bool) {
	if o == nil || IsNil(o.TimeslicesWindow) {
		return nil, false
	}
	return o.TimeslicesWindow, true
}

// HasTimeslicesWindow returns a boolean if a field has been set.
func (o *Objective) HasTimeslicesWindow() bool {
	if o != nil && !IsNil(o.TimeslicesWindow) {
		return true
	}

	return false
}

// SetTimeslicesWindow gets a reference to the given string and assigns it to the TimeslicesWindow field.
func (o *Objective) SetTimeslicesWindow(v string) {
	o.TimeslicesWindow = &v
}

func (o Objective) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Objective) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target"] = o.Target
	if !IsNil(o.TimeslicesTarget) {
		toSerialize["timeslicesTarget"] = o.TimeslicesTarget
	}
	if !IsNil(o.TimeslicesWindow) {
		toSerialize["timeslicesWindow"] = o.TimeslicesWindow
	}
	return toSerialize, nil
}

type NullableObjective struct {
	value *Objective
	isSet bool
}

func (v NullableObjective) Get() *Objective {
	return v.value
}

func (v *NullableObjective) Set(val *Objective) {
	v.value = val
	v.isSet = true
}

func (v NullableObjective) IsSet() bool {
	return v.isSet
}

func (v *NullableObjective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjective(val *Objective) *NullableObjective {
	return &NullableObjective{value: val, isSet: true}
}

func (v NullableObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
