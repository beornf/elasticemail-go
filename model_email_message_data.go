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

// EmailMessageData Email data
type EmailMessageData struct {
	// List of recipients (visible to others)
	Recipients []EmailRecipient `json:"Recipients"`
	// Proper e-mail content
	Content *EmailContent `json:"Content,omitempty"`
	// E-mail configuration
	Options *Options `json:"Options,omitempty"`
}

// NewEmailMessageData instantiates a new EmailMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailMessageData(recipients []EmailRecipient) *EmailMessageData {
	this := EmailMessageData{}
	this.Recipients = recipients
	return &this
}

// NewEmailMessageDataWithDefaults instantiates a new EmailMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailMessageDataWithDefaults() *EmailMessageData {
	this := EmailMessageData{}
	return &this
}

// GetRecipients returns the Recipients field value
func (o *EmailMessageData) GetRecipients() []EmailRecipient {
	if o == nil {
		var ret []EmailRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *EmailMessageData) GetRecipientsOk() (*[]EmailRecipient, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *EmailMessageData) SetRecipients(v []EmailRecipient) {
	o.Recipients = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *EmailMessageData) GetContent() EmailContent {
	if o == nil || o.Content == nil {
		var ret EmailContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageData) GetContentOk() (*EmailContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *EmailMessageData) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given EmailContent and assigns it to the Content field.
func (o *EmailMessageData) SetContent(v EmailContent) {
	o.Content = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EmailMessageData) GetOptions() Options {
	if o == nil || o.Options == nil {
		var ret Options
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageData) GetOptionsOk() (*Options, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EmailMessageData) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Options and assigns it to the Options field.
func (o *EmailMessageData) SetOptions(v Options) {
	o.Options = &v
}

func (o EmailMessageData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Recipients"] = o.Recipients
	}
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if o.Options != nil {
		toSerialize["Options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableEmailMessageData struct {
	value *EmailMessageData
	isSet bool
}

func (v NullableEmailMessageData) Get() *EmailMessageData {
	return v.value
}

func (v *NullableEmailMessageData) Set(val *EmailMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMessageData(val *EmailMessageData) *NullableEmailMessageData {
	return &NullableEmailMessageData{value: val, isSet: true}
}

func (v NullableEmailMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


