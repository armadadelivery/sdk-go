# Wallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**WarningBalanceLevel** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewWallet

`func NewWallet() *Wallet`

NewWallet instantiates a new Wallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithDefaults

`func NewWalletWithDefaults() *Wallet`

NewWalletWithDefaults instantiates a new Wallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Wallet) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Wallet) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Wallet) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Wallet) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *Wallet) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Wallet) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Wallet) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Wallet) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetWarningBalanceLevel

`func (o *Wallet) GetWarningBalanceLevel() float32`

GetWarningBalanceLevel returns the WarningBalanceLevel field if non-nil, zero value otherwise.

### GetWarningBalanceLevelOk

`func (o *Wallet) GetWarningBalanceLevelOk() (*float32, bool)`

GetWarningBalanceLevelOk returns a tuple with the WarningBalanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningBalanceLevel

`func (o *Wallet) SetWarningBalanceLevel(v float32)`

SetWarningBalanceLevel sets WarningBalanceLevel field to given value.

### HasWarningBalanceLevel

`func (o *Wallet) HasWarningBalanceLevel() bool`

HasWarningBalanceLevel returns a boolean if a field has been set.

### SetWarningBalanceLevelNil

`func (o *Wallet) SetWarningBalanceLevelNil(b bool)`

 SetWarningBalanceLevelNil sets the value for WarningBalanceLevel to be an explicit nil

### UnsetWarningBalanceLevel
`func (o *Wallet) UnsetWarningBalanceLevel()`

UnsetWarningBalanceLevel ensures that no value is present for WarningBalanceLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


