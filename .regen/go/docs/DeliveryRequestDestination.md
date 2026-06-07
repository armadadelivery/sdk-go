# DeliveryRequestDestination

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
**Area** | **string** |  | 
**Block** | **string** |  | 
**Street** | **string** |  | 
**Building** | **string** |  | 
**City** | **string** |  | 
**District** | **string** |  | 
**ShortAddress** | **string** |  | 

## Methods

### NewDeliveryRequestDestination

`func NewDeliveryRequestDestination(contactName string, contactPhone string, latitude float32, longitude float32, area string, block string, street string, building string, city string, district string, shortAddress string, ) *DeliveryRequestDestination`

NewDeliveryRequestDestination instantiates a new DeliveryRequestDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryRequestDestinationWithDefaults

`func NewDeliveryRequestDestinationWithDefaults() *DeliveryRequestDestination`

NewDeliveryRequestDestinationWithDefaults instantiates a new DeliveryRequestDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactName

`func (o *DeliveryRequestDestination) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *DeliveryRequestDestination) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *DeliveryRequestDestination) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactPhone

`func (o *DeliveryRequestDestination) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *DeliveryRequestDestination) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *DeliveryRequestDestination) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.


### GetLatitude

`func (o *DeliveryRequestDestination) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DeliveryRequestDestination) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DeliveryRequestDestination) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *DeliveryRequestDestination) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DeliveryRequestDestination) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DeliveryRequestDestination) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetFirstLine

`func (o *DeliveryRequestDestination) GetFirstLine() string`

GetFirstLine returns the FirstLine field if non-nil, zero value otherwise.

### GetFirstLineOk

`func (o *DeliveryRequestDestination) GetFirstLineOk() (*string, bool)`

GetFirstLineOk returns a tuple with the FirstLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLine

`func (o *DeliveryRequestDestination) SetFirstLine(v string)`

SetFirstLine sets FirstLine field to given value.

### HasFirstLine

`func (o *DeliveryRequestDestination) HasFirstLine() bool`

HasFirstLine returns a boolean if a field has been set.

### GetFloor

`func (o *DeliveryRequestDestination) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *DeliveryRequestDestination) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *DeliveryRequestDestination) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *DeliveryRequestDestination) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetApartment

`func (o *DeliveryRequestDestination) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *DeliveryRequestDestination) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *DeliveryRequestDestination) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *DeliveryRequestDestination) HasApartment() bool`

HasApartment returns a boolean if a field has been set.

### GetInstructions

`func (o *DeliveryRequestDestination) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *DeliveryRequestDestination) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *DeliveryRequestDestination) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *DeliveryRequestDestination) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetArea

`func (o *DeliveryRequestDestination) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *DeliveryRequestDestination) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *DeliveryRequestDestination) SetArea(v string)`

SetArea sets Area field to given value.


### GetBlock

`func (o *DeliveryRequestDestination) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *DeliveryRequestDestination) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *DeliveryRequestDestination) SetBlock(v string)`

SetBlock sets Block field to given value.


### GetStreet

`func (o *DeliveryRequestDestination) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *DeliveryRequestDestination) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *DeliveryRequestDestination) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetBuilding

`func (o *DeliveryRequestDestination) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *DeliveryRequestDestination) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *DeliveryRequestDestination) SetBuilding(v string)`

SetBuilding sets Building field to given value.


### GetCity

`func (o *DeliveryRequestDestination) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeliveryRequestDestination) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeliveryRequestDestination) SetCity(v string)`

SetCity sets City field to given value.


### GetDistrict

`func (o *DeliveryRequestDestination) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *DeliveryRequestDestination) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *DeliveryRequestDestination) SetDistrict(v string)`

SetDistrict sets District field to given value.


### GetShortAddress

`func (o *DeliveryRequestDestination) GetShortAddress() string`

GetShortAddress returns the ShortAddress field if non-nil, zero value otherwise.

### GetShortAddressOk

`func (o *DeliveryRequestDestination) GetShortAddressOk() (*string, bool)`

GetShortAddressOk returns a tuple with the ShortAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortAddress

`func (o *DeliveryRequestDestination) SetShortAddress(v string)`

SetShortAddress sets ShortAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


