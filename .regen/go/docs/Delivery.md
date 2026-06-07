# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TestMode** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DeliveryFee** | Pointer to **float32** |  | [optional] 
**Customer** | Pointer to [**DeliveryCustomer**](DeliveryCustomer.md) |  | [optional] 
**Driver** | Pointer to [**DeliveryDriver**](DeliveryDriver.md) |  | [optional] 
**Logistics** | Pointer to [**DeliveryLogistics**](DeliveryLogistics.md) |  | [optional] 

## Methods

### NewDelivery

`func NewDelivery() *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Delivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Delivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Delivery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Delivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Delivery) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Delivery) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Delivery) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Delivery) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTestMode

`func (o *Delivery) GetTestMode() bool`

GetTestMode returns the TestMode field if non-nil, zero value otherwise.

### GetTestModeOk

`func (o *Delivery) GetTestModeOk() (*bool, bool)`

GetTestModeOk returns a tuple with the TestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMode

`func (o *Delivery) SetTestMode(v bool)`

SetTestMode sets TestMode field to given value.

### HasTestMode

`func (o *Delivery) HasTestMode() bool`

HasTestMode returns a boolean if a field has been set.

### GetAmount

`func (o *Delivery) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Delivery) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Delivery) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Delivery) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Delivery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Delivery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Delivery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Delivery) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Delivery) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Delivery) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Delivery) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Delivery) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeliveryFee

`func (o *Delivery) GetDeliveryFee() float32`

GetDeliveryFee returns the DeliveryFee field if non-nil, zero value otherwise.

### GetDeliveryFeeOk

`func (o *Delivery) GetDeliveryFeeOk() (*float32, bool)`

GetDeliveryFeeOk returns a tuple with the DeliveryFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFee

`func (o *Delivery) SetDeliveryFee(v float32)`

SetDeliveryFee sets DeliveryFee field to given value.

### HasDeliveryFee

`func (o *Delivery) HasDeliveryFee() bool`

HasDeliveryFee returns a boolean if a field has been set.

### GetCustomer

`func (o *Delivery) GetCustomer() DeliveryCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Delivery) GetCustomerOk() (*DeliveryCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Delivery) SetCustomer(v DeliveryCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Delivery) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetDriver

`func (o *Delivery) GetDriver() DeliveryDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *Delivery) GetDriverOk() (*DeliveryDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *Delivery) SetDriver(v DeliveryDriver)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *Delivery) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetLogistics

`func (o *Delivery) GetLogistics() DeliveryLogistics`

GetLogistics returns the Logistics field if non-nil, zero value otherwise.

### GetLogisticsOk

`func (o *Delivery) GetLogisticsOk() (*DeliveryLogistics, bool)`

GetLogisticsOk returns a tuple with the Logistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogistics

`func (o *Delivery) SetLogistics(v DeliveryLogistics)`

SetLogistics sets Logistics field to given value.

### HasLogistics

`func (o *Delivery) HasLogistics() bool`

HasLogistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


