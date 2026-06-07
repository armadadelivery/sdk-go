# \DeliveriesAPI

All URIs are relative to *https://sandbox.api.armadadelivery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDelivery**](DeliveriesAPI.md#CancelDelivery) | **Post** /v2/deliveries/{id}/cancel | Cancel a delivery
[**CreateDelivery**](DeliveriesAPI.md#CreateDelivery) | **Post** /v2/deliveries | Create a delivery
[**EstimateDelivery**](DeliveriesAPI.md#EstimateDelivery) | **Post** /v2/deliveries/estimate | Estimate a delivery fee
[**EstimateDeliveryStatic**](DeliveriesAPI.md#EstimateDeliveryStatic) | **Post** /v2/deliveries/estimate/static | Estimate a delivery fee using static pricing (no live traffic).
[**GetDelivery**](DeliveriesAPI.md#GetDelivery) | **Get** /v2/deliveries/{id} | Retrieve a delivery
[**RetryDelivery**](DeliveriesAPI.md#RetryDelivery) | **Post** /v2/deliveries/{id}/retry | Retry a failed delivery



## CancelDelivery

> Delivery CancelDelivery(ctx, id).CancelDeliveryRequest(cancelDeliveryRequest).Execute()

Cancel a delivery

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	cancelDeliveryRequest := *openapiclient.NewCancelDeliveryRequest() // CancelDeliveryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.CancelDelivery(context.Background(), id).CancelDeliveryRequest(cancelDeliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.CancelDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelDelivery`: Delivery
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.CancelDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelDeliveryRequest** | [**CancelDeliveryRequest**](CancelDeliveryRequest.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDelivery

> Delivery CreateDelivery(ctx).DeliveryRequest(deliveryRequest).Execute()

Create a delivery

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deliveryRequest := *openapiclient.NewDeliveryRequest("Reference_example", *openapiclient.NewPayment(float32(123), "Type_example"), "OriginFormat_example", openapiclient.DeliveryRequest_origin{BranchOrigin: openapiclient.NewBranchOrigin("BranchId_example")}, "DestinationFormat_example", openapiclient.DeliveryRequest_destination{KsaAddress: openapiclient.NewKsaAddress("ContactName_example", "ContactPhone_example", "City_example", "Street_example", "District_example", "Building_example")}) // DeliveryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.CreateDelivery(context.Background()).DeliveryRequest(deliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.CreateDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDelivery`: Delivery
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.CreateDelivery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryRequest** | [**DeliveryRequest**](DeliveryRequest.md) |  | 

### Return type

[**Delivery**](Delivery.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateDelivery

> Estimate EstimateDelivery(ctx).EstimateRequest(estimateRequest).Execute()

Estimate a delivery fee

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	estimateRequest := *openapiclient.NewEstimateRequest("OriginFormat_example", openapiclient.DeliveryRequest_origin{BranchOrigin: openapiclient.NewBranchOrigin("BranchId_example")}, "DestinationFormat_example", openapiclient.DeliveryRequest_destination{KsaAddress: openapiclient.NewKsaAddress("ContactName_example", "ContactPhone_example", "City_example", "Street_example", "District_example", "Building_example")}) // EstimateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.EstimateDelivery(context.Background()).EstimateRequest(estimateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.EstimateDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateDelivery`: Estimate
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.EstimateDelivery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateRequest** | [**EstimateRequest**](EstimateRequest.md) |  | 

### Return type

[**Estimate**](Estimate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateDeliveryStatic

> Estimate EstimateDeliveryStatic(ctx).EstimateRequest(estimateRequest).Execute()

Estimate a delivery fee using static pricing (no live traffic).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	estimateRequest := *openapiclient.NewEstimateRequest("OriginFormat_example", openapiclient.DeliveryRequest_origin{BranchOrigin: openapiclient.NewBranchOrigin("BranchId_example")}, "DestinationFormat_example", openapiclient.DeliveryRequest_destination{KsaAddress: openapiclient.NewKsaAddress("ContactName_example", "ContactPhone_example", "City_example", "Street_example", "District_example", "Building_example")}) // EstimateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.EstimateDeliveryStatic(context.Background()).EstimateRequest(estimateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.EstimateDeliveryStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateDeliveryStatic`: Estimate
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.EstimateDeliveryStatic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateDeliveryStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateRequest** | [**EstimateRequest**](EstimateRequest.md) |  | 

### Return type

[**Estimate**](Estimate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelivery

> Delivery GetDelivery(ctx, id).Execute()

Retrieve a delivery

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.GetDelivery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.GetDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelivery`: Delivery
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.GetDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delivery**](Delivery.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryDelivery

> Delivery RetryDelivery(ctx, id).Execute()

Retry a failed delivery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveriesAPI.RetryDelivery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveriesAPI.RetryDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryDelivery`: Delivery
	fmt.Fprintf(os.Stdout, "Response from `DeliveriesAPI.RetryDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Delivery**](Delivery.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

