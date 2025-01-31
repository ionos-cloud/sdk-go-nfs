/*
 * IONOS Cloud - Network File Storage API
 *
 * The RESTful API for managing Network File Storage.
 *
 * API version: 0.1.3
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ShareClientGroups struct for ShareClientGroups
type ShareClientGroups struct {
	// Optional description for the clients groups.
	Description *string               `json:"description,omitempty"`
	IpNetworks  *[]string             `json:"ipNetworks,omitempty"`
	Hosts       *[]string             `json:"hosts,omitempty"`
	Nfs         *ShareClientGroupsNfs `json:"nfs,omitempty"`
}

// NewShareClientGroups instantiates a new ShareClientGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareClientGroups() *ShareClientGroups {
	this := ShareClientGroups{}

	return &this
}

// NewShareClientGroupsWithDefaults instantiates a new ShareClientGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareClientGroupsWithDefaults() *ShareClientGroups {
	this := ShareClientGroups{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ShareClientGroups) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareClientGroups) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *ShareClientGroups) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *ShareClientGroups) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetIpNetworks returns the IpNetworks field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ShareClientGroups) GetIpNetworks() *[]string {
	if o == nil {
		return nil
	}

	return o.IpNetworks

}

// GetIpNetworksOk returns a tuple with the IpNetworks field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareClientGroups) GetIpNetworksOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.IpNetworks, true
}

// SetIpNetworks sets field value
func (o *ShareClientGroups) SetIpNetworks(v []string) {

	o.IpNetworks = &v

}

// HasIpNetworks returns a boolean if a field has been set.
func (o *ShareClientGroups) HasIpNetworks() bool {
	if o != nil && o.IpNetworks != nil {
		return true
	}

	return false
}

// GetHosts returns the Hosts field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ShareClientGroups) GetHosts() *[]string {
	if o == nil {
		return nil
	}

	return o.Hosts

}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareClientGroups) GetHostsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Hosts, true
}

// SetHosts sets field value
func (o *ShareClientGroups) SetHosts(v []string) {

	o.Hosts = &v

}

// HasHosts returns a boolean if a field has been set.
func (o *ShareClientGroups) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// GetNfs returns the Nfs field value
// If the value is explicit nil, the zero value for ShareClientGroupsNfs will be returned
func (o *ShareClientGroups) GetNfs() *ShareClientGroupsNfs {
	if o == nil {
		return nil
	}

	return o.Nfs

}

// GetNfsOk returns a tuple with the Nfs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShareClientGroups) GetNfsOk() (*ShareClientGroupsNfs, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nfs, true
}

// SetNfs sets field value
func (o *ShareClientGroups) SetNfs(v ShareClientGroupsNfs) {

	o.Nfs = &v

}

// HasNfs returns a boolean if a field has been set.
func (o *ShareClientGroups) HasNfs() bool {
	if o != nil && o.Nfs != nil {
		return true
	}

	return false
}

func (o ShareClientGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.IpNetworks != nil {
		toSerialize["ipNetworks"] = o.IpNetworks
	}

	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}

	if o.Nfs != nil {
		toSerialize["nfs"] = o.Nfs
	}

	return json.Marshal(toSerialize)
}

type NullableShareClientGroups struct {
	value *ShareClientGroups
	isSet bool
}

func (v NullableShareClientGroups) Get() *ShareClientGroups {
	return v.value
}

func (v *NullableShareClientGroups) Set(val *ShareClientGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableShareClientGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableShareClientGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareClientGroups(val *ShareClientGroups) *NullableShareClientGroups {
	return &NullableShareClientGroups{value: val, isSet: true}
}

func (v NullableShareClientGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareClientGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
