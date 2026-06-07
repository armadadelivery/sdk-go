# DeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** | Partner-supplied unique id used for de-duplication. | 
**ScheduledDate** | Pointer to **time.Time** |  | [optional] 
**Payment** | [**Payment**](Payment.md) |  | 
**OriginFormat** | **string** |  | 
**Origin** | [**DeliveryRequestOrigin**](DeliveryRequestOrigin.md) |  | 
**DestinationFormat** | **string** |  | 
**Destination** | [**DeliveryRequestDestination**](DeliveryRequestDestination.md) |  | 

## Methods

### NewDeliveryRequest

`func NewDeliveryRequest(reference string, payment Payment, originFormat string, origin DeliveryRequestOrigin, destinationFormat string, destination DeliveryRequestDestination, ) *DeliveryRequest`

NewDeliveryRequest instantiates a new DeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryRequestWithDefaults

`func NewDeliveryRequestWithDefaults() *DeliveryRequest`

NewDeliveryRequestWithDefaults instantiates a new DeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *DeliveryRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DeliveryRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DeliveryRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetScheduledDate

`func (o *DeliveryRequest) GetScheduledDate() time.Time`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *DeliveryRequest) GetScheduledDateOk() (*time.Time, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *DeliveryRequest) SetScheduledDate(v time.Time)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *DeliveryRequest) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetPayment

`func (o *DeliveryRequest) GetPayment() Payment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *DeliveryRequest) GetPaymentOk() (*Payment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *DeliveryRequest) SetPayment(v Payment)`

SetPayment sets Payment field to given value.


### GetOriginFormat

`func (o *DeliveryRequest) GetOriginFormat() string`

GetOriginFormat returns the OriginFormat field if non-nil, zero value otherwise.

### GetOriginFormatOk

`func (o *DeliveryRequest) GetOriginFormatOk() (*string, bool)`

GetOriginFormatOk returns a tuple with the OriginFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFormat

`func (o *DeliveryRequest) SetOriginFormat(v string)`

SetOriginFormat sets OriginFormat field to given value.


### GetOrigin

`func (o *DeliveryRequest) GetOrigin() DeliveryRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DeliveryRequest) GetOriginOk() (*DeliveryRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DeliveryRequest) SetOrigin(v DeliveryRequestOrigin)`

SetOrigin sets Origin field to given value.


### GetDestinationFormat

`func (o *DeliveryRequest) GetDestinationFormat() string`

GetDestinationFormat returns the DestinationFormat field if non-nil, zero value otherwise.

### GetDestinationFormatOk

`func (o *DeliveryRequest) GetDestinationFormatOk() (*string, bool)`

GetDestinationFormatOk returns a tuple with the DestinationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFormat

`func (o *DeliveryRequest) SetDestinationFormat(v string)`

SetDestinationFormat sets DestinationFormat field to given value.


### GetDestination

`func (o *DeliveryRequest) GetDestination() DeliveryRequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DeliveryRequest) GetDestinationOk() (*DeliveryRequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DeliveryRequest) SetDestination(v DeliveryRequestDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


