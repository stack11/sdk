/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.15
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceRegistrationFlowWithPasswordMethod SubmitSelfServiceRegistrationFlowWithPasswordMethod is used to decode the registration form payload when using the password method.
type SubmitSelfServiceRegistrationFlowWithPasswordMethod struct {
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `password` when using the password method.
	Method string `json:"method"`
	// Password to sign the user up with
	Password *string `json:"password,omitempty"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits,omitempty"`
}

// NewSubmitSelfServiceRegistrationFlowWithPasswordMethod instantiates a new SubmitSelfServiceRegistrationFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceRegistrationFlowWithPasswordMethod(method string) *SubmitSelfServiceRegistrationFlowWithPasswordMethod {
	this := SubmitSelfServiceRegistrationFlowWithPasswordMethod{}
	this.Method = method
	return &this
}

// NewSubmitSelfServiceRegistrationFlowWithPasswordMethodWithDefaults instantiates a new SubmitSelfServiceRegistrationFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceRegistrationFlowWithPasswordMethodWithDefaults() *SubmitSelfServiceRegistrationFlowWithPasswordMethod {
	this := SubmitSelfServiceRegistrationFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = &v
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetTraits() map[string]interface{} {
	if o == nil || o.Traits == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || o.Traits == nil {
		return nil, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) HasTraits() bool {
	if o != nil && o.Traits != nil {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *SubmitSelfServiceRegistrationFlowWithPasswordMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o SubmitSelfServiceRegistrationFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Traits != nil {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod struct {
	value *SubmitSelfServiceRegistrationFlowWithPasswordMethod
	isSet bool
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) Get() *SubmitSelfServiceRegistrationFlowWithPasswordMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) Set(val *SubmitSelfServiceRegistrationFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceRegistrationFlowWithPasswordMethod(val *SubmitSelfServiceRegistrationFlowWithPasswordMethod) *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod {
	return &NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


