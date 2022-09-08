/*
Send Payment Transfer API

Mastercard Send Payment Transfer API

API version: 1.0.1
Contact: apisupport@mastercard.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ResponseParticipant Information about a participant that may be passed along to the network. 
type ResponseParticipant struct {
	// Classification of the merchant's type of business or service. The value in the response is the code passed to the network and it may be different from the value provided in the request. Details- numeric, 4.
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
	// The IRD used for clearing the transaction on the Mastercard network. Details - Alphanumeric, length 2 characters.
	InterchangeRateDesignator *string `json:"interchange_rate_designator,omitempty"`
	// Code identifying the card acceptor. Please contact your Mastercard representative to enable the usage of fields in this section. Details- alphanumeric, 10-15.
	CardAcceptorId *string `json:"card_acceptor_id,omitempty"`
	// Customer service phone number that may be displayed on the recipient's bank or card statement. For U.S. or Canada regions, this must be a 10-digit phone number (e.g. 8005551234). For other regions, any length up to 13 digits is permitted. The phone number must be numeric without special characters such as spaces, - or ( ). To enable usage of this field, please contact your Mastercard representative. Details - numeric, 1-13.
	CustomerServiceContactInfo *string                       `json:"customer_service_contact_info,omitempty"`
	TransferAcceptorAddress *ResponseTransferAcceptorAddress `json:"transfer_acceptor_address,omitempty"`
}

// NewResponseParticipant instantiates a new ResponseParticipant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseParticipant() *ResponseParticipant {
	this := ResponseParticipant{}
	return &this
}

// NewResponseParticipantWithDefaults instantiates a new ResponseParticipant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseParticipantWithDefaults() *ResponseParticipant {
	this := ResponseParticipant{}
	return &this
}

// GetMerchantCategoryCode returns the MerchantCategoryCode field value if set, zero value otherwise.
func (o *ResponseParticipant) GetMerchantCategoryCode() string {
	if o == nil || o.MerchantCategoryCode == nil {
		var ret string
		return ret
	}
	return *o.MerchantCategoryCode
}

// GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParticipant) GetMerchantCategoryCodeOk() (*string, bool) {
	if o == nil || o.MerchantCategoryCode == nil {
		return nil, false
	}
	return o.MerchantCategoryCode, true
}

// HasMerchantCategoryCode returns a boolean if a field has been set.
func (o *ResponseParticipant) HasMerchantCategoryCode() bool {
	if o != nil && o.MerchantCategoryCode != nil {
		return true
	}

	return false
}

// SetMerchantCategoryCode gets a reference to the given string and assigns it to the MerchantCategoryCode field.
func (o *ResponseParticipant) SetMerchantCategoryCode(v string) {
	o.MerchantCategoryCode = &v
}

// GetInterchangeRateDesignator returns the InterchangeRateDesignator field value if set, zero value otherwise.
func (o *ResponseParticipant) GetInterchangeRateDesignator() string {
	if o == nil || o.InterchangeRateDesignator == nil {
		var ret string
		return ret
	}
	return *o.InterchangeRateDesignator
}

// GetInterchangeRateDesignatorOk returns a tuple with the InterchangeRateDesignator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParticipant) GetInterchangeRateDesignatorOk() (*string, bool) {
	if o == nil || o.InterchangeRateDesignator == nil {
		return nil, false
	}
	return o.InterchangeRateDesignator, true
}

// HasInterchangeRateDesignator returns a boolean if a field has been set.
func (o *ResponseParticipant) HasInterchangeRateDesignator() bool {
	if o != nil && o.InterchangeRateDesignator != nil {
		return true
	}

	return false
}

// SetInterchangeRateDesignator gets a reference to the given string and assigns it to the InterchangeRateDesignator field.
func (o *ResponseParticipant) SetInterchangeRateDesignator(v string) {
	o.InterchangeRateDesignator = &v
}

// GetCardAcceptorId returns the CardAcceptorId field value if set, zero value otherwise.
func (o *ResponseParticipant) GetCardAcceptorId() string {
	if o == nil || o.CardAcceptorId == nil {
		var ret string
		return ret
	}
	return *o.CardAcceptorId
}

// GetCardAcceptorIdOk returns a tuple with the CardAcceptorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParticipant) GetCardAcceptorIdOk() (*string, bool) {
	if o == nil || o.CardAcceptorId == nil {
		return nil, false
	}
	return o.CardAcceptorId, true
}

// HasCardAcceptorId returns a boolean if a field has been set.
func (o *ResponseParticipant) HasCardAcceptorId() bool {
	if o != nil && o.CardAcceptorId != nil {
		return true
	}

	return false
}

// SetCardAcceptorId gets a reference to the given string and assigns it to the CardAcceptorId field.
func (o *ResponseParticipant) SetCardAcceptorId(v string) {
	o.CardAcceptorId = &v
}

// GetCustomerServiceContactInfo returns the CustomerServiceContactInfo field value if set, zero value otherwise.
func (o *ResponseParticipant) GetCustomerServiceContactInfo() string {
	if o == nil || o.CustomerServiceContactInfo == nil {
		var ret string
		return ret
	}
	return *o.CustomerServiceContactInfo
}

// GetCustomerServiceContactInfoOk returns a tuple with the CustomerServiceContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParticipant) GetCustomerServiceContactInfoOk() (*string, bool) {
	if o == nil || o.CustomerServiceContactInfo == nil {
		return nil, false
	}
	return o.CustomerServiceContactInfo, true
}

// HasCustomerServiceContactInfo returns a boolean if a field has been set.
func (o *ResponseParticipant) HasCustomerServiceContactInfo() bool {
	if o != nil && o.CustomerServiceContactInfo != nil {
		return true
	}

	return false
}

// SetCustomerServiceContactInfo gets a reference to the given string and assigns it to the CustomerServiceContactInfo field.
func (o *ResponseParticipant) SetCustomerServiceContactInfo(v string) {
	o.CustomerServiceContactInfo = &v
}

// GetTransferAcceptorAddress returns the TransferAcceptorAddress field value if set, zero value otherwise.
func (o *ResponseParticipant) GetTransferAcceptorAddress() ResponseTransferAcceptorAddress {
	if o == nil || o.TransferAcceptorAddress == nil {
		var ret ResponseTransferAcceptorAddress
		return ret
	}
	return *o.TransferAcceptorAddress
}

// GetTransferAcceptorAddressOk returns a tuple with the TransferAcceptorAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParticipant) GetTransferAcceptorAddressOk() (*ResponseTransferAcceptorAddress, bool) {
	if o == nil || o.TransferAcceptorAddress == nil {
		return nil, false
	}
	return o.TransferAcceptorAddress, true
}

// HasTransferAcceptorAddress returns a boolean if a field has been set.
func (o *ResponseParticipant) HasTransferAcceptorAddress() bool {
	if o != nil && o.TransferAcceptorAddress != nil {
		return true
	}

	return false
}

// SetTransferAcceptorAddress gets a reference to the given ResponseTransferAcceptorAddress and assigns it to the TransferAcceptorAddress field.
func (o *ResponseParticipant) SetTransferAcceptorAddress(v ResponseTransferAcceptorAddress) {
	o.TransferAcceptorAddress = &v
}

func (o ResponseParticipant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantCategoryCode != nil {
		toSerialize["merchant_category_code"] = o.MerchantCategoryCode
	}
	if o.InterchangeRateDesignator != nil {
		toSerialize["interchange_rate_designator"] = o.InterchangeRateDesignator
	}
	if o.CardAcceptorId != nil {
		toSerialize["card_acceptor_id"] = o.CardAcceptorId
	}
	if o.CustomerServiceContactInfo != nil {
		toSerialize["customer_service_contact_info"] = o.CustomerServiceContactInfo
	}
	if o.TransferAcceptorAddress != nil {
		toSerialize["transfer_acceptor_address"] = o.TransferAcceptorAddress
	}
	return json.Marshal(toSerialize)
}

type NullableResponseParticipant struct {
	value *ResponseParticipant
	isSet bool
}

func (v NullableResponseParticipant) Get() *ResponseParticipant {
	return v.value
}

func (v *NullableResponseParticipant) Set(val *ResponseParticipant) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseParticipant(val *ResponseParticipant) *NullableResponseParticipant {
	return &NullableResponseParticipant{value: val, isSet: true}
}

func (v NullableResponseParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


