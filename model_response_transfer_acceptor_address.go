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

// ResponseTransferAcceptorAddress Address of the transfer acceptor in response.
type ResponseTransferAcceptorAddress struct {
	// First line of the transfer acceptor's address. Alphanumeric Special [a-zA-Z0-9 !\"#$%&'()*+,-./&#92;:;<=>?@[]_`{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50.
	Line1 *string `json:"line1,omitempty"`
	// Second line of the transfer acceptor's address. Alphanumeric Special [a-zA-Z0-9 !\"#$%&'()*+,-./&#92;:;<=>?@[]_`{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50.
	Line2 *string `json:"line2,omitempty"`
	// The transfer acceptor's city. Alphanumeric Special [a-zA-Z0-9 !\"#$%&'()*+,-./&#92;:;<=>?@[]_`{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-25.
	City *string `json:"city,omitempty"`
	// The transfer acceptor's province, state or territory, as an ISO 3166-2 uppercase alpha 2 or 3 character country subdivision code. For example, Missouri is MO.
	CountrySubdivision *string `json:"country_subdivision,omitempty"`
	// The transfer acceptor's postal code. For USA, this is 5 digits or 5 digits hyphen 4 digits, for example '63368', '63368-5555'. For other regions, this is alphanumeric, length 1-10.
	PostalCode *string `json:"postal_code,omitempty"`
	// The transfer acceptor's country as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA.
	Country *string `json:"country,omitempty"`
}

// NewResponseTransferAcceptorAddress instantiates a new ResponseTransferAcceptorAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTransferAcceptorAddress() *ResponseTransferAcceptorAddress {
	this := ResponseTransferAcceptorAddress{}
	return &this
}

// NewResponseTransferAcceptorAddressWithDefaults instantiates a new ResponseTransferAcceptorAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTransferAcceptorAddressWithDefaults() *ResponseTransferAcceptorAddress {
	this := ResponseTransferAcceptorAddress{}
	return &this
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetLine1() string {
	if o == nil || o.Line1 == nil {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetLine1Ok() (*string, bool) {
	if o == nil || o.Line1 == nil {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasLine1() bool {
	if o != nil && o.Line1 != nil {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *ResponseTransferAcceptorAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetLine2() string {
	if o == nil || o.Line2 == nil {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetLine2Ok() (*string, bool) {
	if o == nil || o.Line2 == nil {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasLine2() bool {
	if o != nil && o.Line2 != nil {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *ResponseTransferAcceptorAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *ResponseTransferAcceptorAddress) SetCity(v string) {
	o.City = &v
}

// GetCountrySubdivision returns the CountrySubdivision field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetCountrySubdivision() string {
	if o == nil || o.CountrySubdivision == nil {
		var ret string
		return ret
	}
	return *o.CountrySubdivision
}

// GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetCountrySubdivisionOk() (*string, bool) {
	if o == nil || o.CountrySubdivision == nil {
		return nil, false
	}
	return o.CountrySubdivision, true
}

// HasCountrySubdivision returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasCountrySubdivision() bool {
	if o != nil && o.CountrySubdivision != nil {
		return true
	}

	return false
}

// SetCountrySubdivision gets a reference to the given string and assigns it to the CountrySubdivision field.
func (o *ResponseTransferAcceptorAddress) SetCountrySubdivision(v string) {
	o.CountrySubdivision = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *ResponseTransferAcceptorAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ResponseTransferAcceptorAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransferAcceptorAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ResponseTransferAcceptorAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ResponseTransferAcceptorAddress) SetCountry(v string) {
	o.Country = &v
}

func (o ResponseTransferAcceptorAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Line1 != nil {
		toSerialize["line1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line2"] = o.Line2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CountrySubdivision != nil {
		toSerialize["country_subdivision"] = o.CountrySubdivision
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableResponseTransferAcceptorAddress struct {
	value *ResponseTransferAcceptorAddress
	isSet bool
}

func (v NullableResponseTransferAcceptorAddress) Get() *ResponseTransferAcceptorAddress {
	return v.value
}

func (v *NullableResponseTransferAcceptorAddress) Set(val *ResponseTransferAcceptorAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTransferAcceptorAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTransferAcceptorAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTransferAcceptorAddress(val *ResponseTransferAcceptorAddress) *NullableResponseTransferAcceptorAddress {
	return &NullableResponseTransferAcceptorAddress{value: val, isSet: true}
}

func (v NullableResponseTransferAcceptorAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTransferAcceptorAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


