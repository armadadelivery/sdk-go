# EstimateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginFormat** | **string** |  | 
**Origin** | [**DeliveryRequestOrigin**](DeliveryRequestOrigin.md) |  | 
**DestinationFormat** | **string** |  | 
**Destination** | [**DeliveryRequestDestination**](DeliveryRequestDestination.md) |  | 

## Methods

### NewEstimateRequest

`func NewEstimateRequest(originFormat string, origin DeliveryRequestOrigin, destinationFormat string, destination DeliveryRequestDestination, ) *EstimateRequest`

NewEstimateRequest instantiates a new EstimateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateRequestWithDefaults

`func NewEstimateRequestWithDefaults() *EstimateRequest`

NewEstimateRequestWithDefaults instantiates a new EstimateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginFormat

`func (o *EstimateRequest) GetOriginFormat() string`

GetOriginFormat returns the OriginFormat field if non-nil, zero value otherwise.

### GetOriginFormatOk

`func (o *EstimateRequest) GetOriginFormatOk() (*string, bool)`

GetOriginFormatOk returns a tuple with the OriginFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFormat

`func (o *EstimateRequest) SetOriginFormat(v string)`

SetOriginFormat sets OriginFormat field to given value.


### GetOrigin

`func (o *EstimateRequest) GetOrigin() DeliveryRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EstimateRequest) GetOriginOk() (*DeliveryRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EstimateRequest) SetOrigin(v DeliveryRequestOrigin)`

SetOrigin sets Origin field to given value.


### GetDestinationFormat

`func (o *EstimateRequest) GetDestinationFormat() string`

GetDestinationFormat returns the DestinationFormat field if non-nil, zero value otherwise.

### GetDestinationFormatOk

`func (o *EstimateRequest) GetDestinationFormatOk() (*string, bool)`

GetDestinationFormatOk returns a tuple with the DestinationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFormat

`func (o *EstimateRequest) SetDestinationFormat(v string)`

SetDestinationFormat sets DestinationFormat field to given value.


### GetDestination

`func (o *EstimateRequest) GetDestination() DeliveryRequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *EstimateRequest) GetDestinationOk() (*DeliveryRequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *EstimateRequest) SetDestination(v DeliveryRequestDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


