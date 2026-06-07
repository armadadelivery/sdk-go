# ListInvoices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Invoices** | Pointer to [**[]Invoice**](Invoice.md) |  | [optional] 

## Methods

### NewListInvoices200Response

`func NewListInvoices200Response() *ListInvoices200Response`

NewListInvoices200Response instantiates a new ListInvoices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoices200ResponseWithDefaults

`func NewListInvoices200ResponseWithDefaults() *ListInvoices200Response`

NewListInvoices200ResponseWithDefaults instantiates a new ListInvoices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListInvoices200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListInvoices200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListInvoices200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListInvoices200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ListInvoices200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListInvoices200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListInvoices200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListInvoices200Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotal

`func (o *ListInvoices200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListInvoices200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListInvoices200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListInvoices200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetInvoices

`func (o *ListInvoices200Response) GetInvoices() []Invoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *ListInvoices200Response) GetInvoicesOk() (*[]Invoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *ListInvoices200Response) SetInvoices(v []Invoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *ListInvoices200Response) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


