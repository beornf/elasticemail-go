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
	"time"
)

// ContactHistory struct for ContactHistory
type ContactHistory struct {
	// Type of event occured on this Contact.
	EventType *ContactHistEventType `json:"EventType,omitempty"`
	// Formatted date of event.
	EventDate *time.Time `json:"EventDate,omitempty"`
	// Name of channel this event occured on
	ChannelName *string `json:"ChannelName,omitempty"`
	// Name of template this event occured on
	TemplateName *string `json:"TemplateName,omitempty"`
	// IP Address of the event.
	IPAddress *string `json:"IPAddress,omitempty"`
	// Country of the event.
	Country *string `json:"Country,omitempty"`
	// Additional information about the event
	Data *string `json:"Data,omitempty"`
}

// NewContactHistory instantiates a new ContactHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactHistory() *ContactHistory {
	this := ContactHistory{}
	return &this
}

// NewContactHistoryWithDefaults instantiates a new ContactHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactHistoryWithDefaults() *ContactHistory {
	this := ContactHistory{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ContactHistory) GetEventType() ContactHistEventType {
	if o == nil || o.EventType == nil {
		var ret ContactHistEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetEventTypeOk() (*ContactHistEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ContactHistory) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given ContactHistEventType and assigns it to the EventType field.
func (o *ContactHistory) SetEventType(v ContactHistEventType) {
	o.EventType = &v
}

// GetEventDate returns the EventDate field value if set, zero value otherwise.
func (o *ContactHistory) GetEventDate() time.Time {
	if o == nil || o.EventDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetEventDateOk() (*time.Time, bool) {
	if o == nil || o.EventDate == nil {
		return nil, false
	}
	return o.EventDate, true
}

// HasEventDate returns a boolean if a field has been set.
func (o *ContactHistory) HasEventDate() bool {
	if o != nil && o.EventDate != nil {
		return true
	}

	return false
}

// SetEventDate gets a reference to the given time.Time and assigns it to the EventDate field.
func (o *ContactHistory) SetEventDate(v time.Time) {
	o.EventDate = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *ContactHistory) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *ContactHistory) HasChannelName() bool {
	if o != nil && o.ChannelName != nil {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *ContactHistory) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *ContactHistory) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *ContactHistory) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *ContactHistory) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise.
func (o *ContactHistory) GetIPAddress() string {
	if o == nil || o.IPAddress == nil {
		var ret string
		return ret
	}
	return *o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetIPAddressOk() (*string, bool) {
	if o == nil || o.IPAddress == nil {
		return nil, false
	}
	return o.IPAddress, true
}

// HasIPAddress returns a boolean if a field has been set.
func (o *ContactHistory) HasIPAddress() bool {
	if o != nil && o.IPAddress != nil {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given string and assigns it to the IPAddress field.
func (o *ContactHistory) SetIPAddress(v string) {
	o.IPAddress = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ContactHistory) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ContactHistory) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ContactHistory) SetCountry(v string) {
	o.Country = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContactHistory) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactHistory) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContactHistory) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ContactHistory) SetData(v string) {
	o.Data = &v
}

func (o ContactHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["EventType"] = o.EventType
	}
	if o.EventDate != nil {
		toSerialize["EventDate"] = o.EventDate
	}
	if o.ChannelName != nil {
		toSerialize["ChannelName"] = o.ChannelName
	}
	if o.TemplateName != nil {
		toSerialize["TemplateName"] = o.TemplateName
	}
	if o.IPAddress != nil {
		toSerialize["IPAddress"] = o.IPAddress
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableContactHistory struct {
	value *ContactHistory
	isSet bool
}

func (v NullableContactHistory) Get() *ContactHistory {
	return v.value
}

func (v *NullableContactHistory) Set(val *ContactHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableContactHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableContactHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactHistory(val *ContactHistory) *NullableContactHistory {
	return &NullableContactHistory{value: val, isSet: true}
}

func (v NullableContactHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


