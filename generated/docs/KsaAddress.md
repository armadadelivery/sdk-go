# KsaAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactName** | **string** |  | 
**ContactPhone** | **string** |  | 
**City** | **string** |  | 
**Street** | **string** |  | 
**District** | **string** |  | 
**Building** | **string** |  | 
**Floor** | Pointer to **string** |  | [optional] 
**Apartment** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 

## Methods

### NewKsaAddress

`func NewKsaAddress(contactName string, contactPhone string, city string, street string, district string, building string, ) *KsaAddress`

NewKsaAddress instantiates a new KsaAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsaAddressWithDefaults

`func NewKsaAddressWithDefaults() *KsaAddress`

NewKsaAddressWithDefaults instantiates a new KsaAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactName

`func (o *KsaAddress) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *KsaAddress) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *KsaAddress) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactPhone

`func (o *KsaAddress) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *KsaAddress) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *KsaAddress) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.


### GetCity

`func (o *KsaAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *KsaAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *KsaAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *KsaAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *KsaAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *KsaAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetDistrict

`func (o *KsaAddress) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *KsaAddress) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *KsaAddress) SetDistrict(v string)`

SetDistrict sets District field to given value.


### GetBuilding

`func (o *KsaAddress) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *KsaAddress) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *KsaAddress) SetBuilding(v string)`

SetBuilding sets Building field to given value.


### GetFloor

`func (o *KsaAddress) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *KsaAddress) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *KsaAddress) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *KsaAddress) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetApartment

`func (o *KsaAddress) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *KsaAddress) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *KsaAddress) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *KsaAddress) HasApartment() bool`

HasApartment returns a boolean if a field has been set.

### GetInstructions

`func (o *KsaAddress) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *KsaAddress) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *KsaAddress) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *KsaAddress) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


