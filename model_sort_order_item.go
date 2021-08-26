/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
)

// SortOrderItem Change the ordering of this inbound route for when matching the inbound
type SortOrderItem struct {
	// ID of the route to change the order of
	PublicInboundId string `json:"PublicInboundId"`
	// 1 - route will be used first
	SortOrder int32 `json:"SortOrder"`
}

// NewSortOrderItem instantiates a new SortOrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortOrderItem(publicInboundId string, sortOrder int32) *SortOrderItem {
	this := SortOrderItem{}
	this.PublicInboundId = publicInboundId
	this.SortOrder = sortOrder
	return &this
}

// NewSortOrderItemWithDefaults instantiates a new SortOrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortOrderItemWithDefaults() *SortOrderItem {
	this := SortOrderItem{}
	return &this
}

// GetPublicInboundId returns the PublicInboundId field value
func (o *SortOrderItem) GetPublicInboundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicInboundId
}

// GetPublicInboundIdOk returns a tuple with the PublicInboundId field value
// and a boolean to check if the value has been set.
func (o *SortOrderItem) GetPublicInboundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicInboundId, true
}

// SetPublicInboundId sets field value
func (o *SortOrderItem) SetPublicInboundId(v string) {
	o.PublicInboundId = v
}

// GetSortOrder returns the SortOrder field value
func (o *SortOrderItem) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *SortOrderItem) GetSortOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *SortOrderItem) SetSortOrder(v int32) {
	o.SortOrder = v
}

func (o SortOrderItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PublicInboundId"] = o.PublicInboundId
	}
	if true {
		toSerialize["SortOrder"] = o.SortOrder
	}
	return json.Marshal(toSerialize)
}

type NullableSortOrderItem struct {
	value *SortOrderItem
	isSet bool
}

func (v NullableSortOrderItem) Get() *SortOrderItem {
	return v.value
}

func (v *NullableSortOrderItem) Set(val *SortOrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSortOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSortOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortOrderItem(val *SortOrderItem) *NullableSortOrderItem {
	return &NullableSortOrderItem{value: val, isSet: true}
}

func (v NullableSortOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


