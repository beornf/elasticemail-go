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

// InboundPayload struct for InboundPayload
type InboundPayload struct {
	// Filter of the inbound data
	Filter string `json:"Filter"`
	// Name of this route
	Name string `json:"Name"`
	// Type of the filter
	FilterType InboundRouteFilterType `json:"FilterType"`
	// Type of action to take
	ActionType InboundRouteActionType `json:"ActionType"`
	// Email to forward the inbound to
	EmailAddress *string `json:"EmailAddress,omitempty"`
	// Address to notify about the inbound
	HttpAddress *string `json:"HttpAddress,omitempty"`
}

// NewInboundPayload instantiates a new InboundPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundPayload(filter string, name string, filterType InboundRouteFilterType, actionType InboundRouteActionType) *InboundPayload {
	this := InboundPayload{}
	this.Filter = filter
	this.Name = name
	this.FilterType = filterType
	this.ActionType = actionType
	return &this
}

// NewInboundPayloadWithDefaults instantiates a new InboundPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundPayloadWithDefaults() *InboundPayload {
	this := InboundPayload{}
	return &this
}

// GetFilter returns the Filter field value
func (o *InboundPayload) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *InboundPayload) SetFilter(v string) {
	o.Filter = v
}

// GetName returns the Name field value
func (o *InboundPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InboundPayload) SetName(v string) {
	o.Name = v
}

// GetFilterType returns the FilterType field value
func (o *InboundPayload) GetFilterType() InboundRouteFilterType {
	if o == nil {
		var ret InboundRouteFilterType
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetFilterTypeOk() (*InboundRouteFilterType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *InboundPayload) SetFilterType(v InboundRouteFilterType) {
	o.FilterType = v
}

// GetActionType returns the ActionType field value
func (o *InboundPayload) GetActionType() InboundRouteActionType {
	if o == nil {
		var ret InboundRouteActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetActionTypeOk() (*InboundRouteActionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *InboundPayload) SetActionType(v InboundRouteActionType) {
	o.ActionType = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *InboundPayload) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *InboundPayload) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *InboundPayload) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetHttpAddress returns the HttpAddress field value if set, zero value otherwise.
func (o *InboundPayload) GetHttpAddress() string {
	if o == nil || o.HttpAddress == nil {
		var ret string
		return ret
	}
	return *o.HttpAddress
}

// GetHttpAddressOk returns a tuple with the HttpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundPayload) GetHttpAddressOk() (*string, bool) {
	if o == nil || o.HttpAddress == nil {
		return nil, false
	}
	return o.HttpAddress, true
}

// HasHttpAddress returns a boolean if a field has been set.
func (o *InboundPayload) HasHttpAddress() bool {
	if o != nil && o.HttpAddress != nil {
		return true
	}

	return false
}

// SetHttpAddress gets a reference to the given string and assigns it to the HttpAddress field.
func (o *InboundPayload) SetHttpAddress(v string) {
	o.HttpAddress = &v
}

func (o InboundPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Filter"] = o.Filter
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["FilterType"] = o.FilterType
	}
	if true {
		toSerialize["ActionType"] = o.ActionType
	}
	if o.EmailAddress != nil {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if o.HttpAddress != nil {
		toSerialize["HttpAddress"] = o.HttpAddress
	}
	return json.Marshal(toSerialize)
}

type NullableInboundPayload struct {
	value *InboundPayload
	isSet bool
}

func (v NullableInboundPayload) Get() *InboundPayload {
	return v.value
}

func (v *NullableInboundPayload) Set(val *InboundPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundPayload(val *InboundPayload) *NullableInboundPayload {
	return &NullableInboundPayload{value: val, isSet: true}
}

func (v NullableInboundPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


