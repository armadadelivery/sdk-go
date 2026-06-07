# BranchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Phone** | **string** |  | 
**Address** | **string** |  | 
**Location** | Pointer to [**BranchInputLocation**](BranchInputLocation.md) |  | [optional] 

## Methods

### NewBranchInput

`func NewBranchInput(name string, phone string, address string, ) *BranchInput`

NewBranchInput instantiates a new BranchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchInputWithDefaults

`func NewBranchInputWithDefaults() *BranchInput`

NewBranchInputWithDefaults instantiates a new BranchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BranchInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchInput) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *BranchInput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BranchInput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BranchInput) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetAddress

`func (o *BranchInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BranchInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BranchInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLocation

`func (o *BranchInput) GetLocation() BranchInputLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BranchInput) GetLocationOk() (*BranchInputLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BranchInput) SetLocation(v BranchInputLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BranchInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


