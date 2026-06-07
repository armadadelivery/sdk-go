# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**InvoiceNo** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**OrdersCount** | Pointer to **int32** |  | [optional] 
**DeliveryTripsCount** | Pointer to **int32** |  | [optional] 
**PeriodBegin** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**DueOn** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNo

`func (o *Invoice) GetInvoiceNo() string`

GetInvoiceNo returns the InvoiceNo field if non-nil, zero value otherwise.

### GetInvoiceNoOk

`func (o *Invoice) GetInvoiceNoOk() (*string, bool)`

GetInvoiceNoOk returns a tuple with the InvoiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNo

`func (o *Invoice) SetInvoiceNo(v string)`

SetInvoiceNo sets InvoiceNo field to given value.

### HasInvoiceNo

`func (o *Invoice) HasInvoiceNo() bool`

HasInvoiceNo returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Invoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmount

`func (o *Invoice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Invoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrdersCount

`func (o *Invoice) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *Invoice) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *Invoice) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *Invoice) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### GetDeliveryTripsCount

`func (o *Invoice) GetDeliveryTripsCount() int32`

GetDeliveryTripsCount returns the DeliveryTripsCount field if non-nil, zero value otherwise.

### GetDeliveryTripsCountOk

`func (o *Invoice) GetDeliveryTripsCountOk() (*int32, bool)`

GetDeliveryTripsCountOk returns a tuple with the DeliveryTripsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTripsCount

`func (o *Invoice) SetDeliveryTripsCount(v int32)`

SetDeliveryTripsCount sets DeliveryTripsCount field to given value.

### HasDeliveryTripsCount

`func (o *Invoice) HasDeliveryTripsCount() bool`

HasDeliveryTripsCount returns a boolean if a field has been set.

### GetPeriodBegin

`func (o *Invoice) GetPeriodBegin() string`

GetPeriodBegin returns the PeriodBegin field if non-nil, zero value otherwise.

### GetPeriodBeginOk

`func (o *Invoice) GetPeriodBeginOk() (*string, bool)`

GetPeriodBeginOk returns a tuple with the PeriodBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBegin

`func (o *Invoice) SetPeriodBegin(v string)`

SetPeriodBegin sets PeriodBegin field to given value.

### HasPeriodBegin

`func (o *Invoice) HasPeriodBegin() bool`

HasPeriodBegin returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *Invoice) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *Invoice) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *Invoice) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *Invoice) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetDueOn

`func (o *Invoice) GetDueOn() string`

GetDueOn returns the DueOn field if non-nil, zero value otherwise.

### GetDueOnOk

`func (o *Invoice) GetDueOnOk() (*string, bool)`

GetDueOnOk returns a tuple with the DueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueOn

`func (o *Invoice) SetDueOn(v string)`

SetDueOn sets DueOn field to given value.

### HasDueOn

`func (o *Invoice) HasDueOn() bool`

HasDueOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Invoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


