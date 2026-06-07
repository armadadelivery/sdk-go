# LocationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactName** | **string** |  | 
**ContactPhone** | **string** |  | 
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**FirstLine** | Pointer to **string** |  | [optional] 
**Floor** | Pointer to **string** |  | [optional] 
**Apartment** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationAddress

`func NewLocationAddress(contactName string, contactPhone string, latitude float32, longitude float32, ) *LocationAddress`

NewLocationAddress instantiates a new LocationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAddressWithDefaults

`func NewLocationAddressWithDefaults() *LocationAddress`

NewLocationAddressWithDefaults instantiates a new LocationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactName

`func (o *LocationAddress) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *LocationAddress) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *LocationAddress) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactPhone

`func (o *LocationAddress) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *LocationAddress) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *LocationAddress) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.


### GetLatitude

`func (o *LocationAddress) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationAddress) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationAddress) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *LocationAddress) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationAddress) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationAddress) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetFirstLine

`func (o *LocationAddress) GetFirstLine() string`

GetFirstLine returns the FirstLine field if non-nil, zero value otherwise.

### GetFirstLineOk

`func (o *LocationAddress) GetFirstLineOk() (*string, bool)`

GetFirstLineOk returns a tuple with the FirstLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLine

`func (o *LocationAddress) SetFirstLine(v string)`

SetFirstLine sets FirstLine field to given value.

### HasFirstLine

`func (o *LocationAddress) HasFirstLine() bool`

HasFirstLine returns a boolean if a field has been set.

### GetFloor

`func (o *LocationAddress) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *LocationAddress) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *LocationAddress) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *LocationAddress) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetApartment

`func (o *LocationAddress) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *LocationAddress) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *LocationAddress) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *LocationAddress) HasApartment() bool`

HasApartment returns a boolean if a field has been set.

### GetInstructions

`func (o *LocationAddress) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *LocationAddress) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *LocationAddress) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *LocationAddress) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


