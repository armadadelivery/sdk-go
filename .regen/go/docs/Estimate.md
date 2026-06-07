# Estimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **float32** | Meters. | [optional] 
**Duration** | Pointer to **float32** | Seconds. | [optional] 

## Methods

### NewEstimate

`func NewEstimate() *Estimate`

NewEstimate instantiates a new Estimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateWithDefaults

`func NewEstimateWithDefaults() *Estimate`

NewEstimateWithDefaults instantiates a new Estimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *Estimate) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Estimate) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Estimate) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Estimate) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetCurrency

`func (o *Estimate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Estimate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Estimate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Estimate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDistance

`func (o *Estimate) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Estimate) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Estimate) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Estimate) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetDuration

`func (o *Estimate) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Estimate) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Estimate) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Estimate) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


