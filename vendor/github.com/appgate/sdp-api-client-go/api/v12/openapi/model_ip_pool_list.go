/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// IpPoolList struct for IpPoolList
type IpPoolList struct {
	ResultList
	// List of IP Pools.
	Data *[]IpPool `json:"data,omitempty"`
}

// NewIpPoolList instantiates a new IpPoolList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpPoolList() *IpPoolList {
	this := IpPoolList{}
	return &this
}

// NewIpPoolListWithDefaults instantiates a new IpPoolList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPoolListWithDefaults() *IpPoolList {
	this := IpPoolList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IpPoolList) GetData() []IpPool {
	if o == nil || o.Data == nil {
		var ret []IpPool
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolList) GetDataOk() (*[]IpPool, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IpPoolList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IpPool and assigns it to the Data field.
func (o *IpPoolList) SetData(v []IpPool) {
	o.Data = &v
}

func (o IpPoolList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedResultList, errResultList := json.Marshal(o.ResultList)
	if errResultList != nil {
		return []byte{}, errResultList
	}
	errResultList = json.Unmarshal([]byte(serializedResultList), &toSerialize)
	if errResultList != nil {
		return []byte{}, errResultList
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableIpPoolList struct {
	value *IpPoolList
	isSet bool
}

func (v NullableIpPoolList) Get() *IpPoolList {
	return v.value
}

func (v *NullableIpPoolList) Set(val *IpPoolList) {
	v.value = val
	v.isSet = true
}

func (v NullableIpPoolList) IsSet() bool {
	return v.isSet
}

func (v *NullableIpPoolList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpPoolList(val *IpPoolList) *NullableIpPoolList {
	return &NullableIpPoolList{value: val, isSet: true}
}

func (v NullableIpPoolList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpPoolList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
