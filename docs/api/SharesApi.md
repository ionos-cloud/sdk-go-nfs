# \SharesApi

All URIs are relative to *https://nfs.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersSharesDelete**](SharesApi.md#ClustersSharesDelete) | **Delete** /clusters/{clusterId}/shares/{shareId} | Delete Share|
|[**ClustersSharesFindById**](SharesApi.md#ClustersSharesFindById) | **Get** /clusters/{clusterId}/shares/{shareId} | Retrieves a share|
|[**ClustersSharesGet**](SharesApi.md#ClustersSharesGet) | **Get** /clusters/{clusterId}/shares | Retrieves all shares|
|[**ClustersSharesPost**](SharesApi.md#ClustersSharesPost) | **Post** /clusters/{clusterId}/shares | Create Share|
|[**ClustersSharesPut**](SharesApi.md#ClustersSharesPut) | **Put** /clusters/{clusterId}/shares/{shareId} | Ensure Share|



## ClustersSharesDelete

```go
var result  = ClustersSharesDelete(ctx, clusterId, shareId)
                      .Execute()
```

Delete Share



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-nfs"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.
    shareId := "7b1ef56d-dfc6-51fe-aff0-7af2d6747868" // string | The identifier (UUID) of the share.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.SharesApi.ClustersSharesDelete(context.Background(), clusterId, shareId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.ClustersSharesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |
|**shareId** | **string** | The identifier (UUID) of the share. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSharesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"SharesApiService.ClustersSharesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "SharesApiService.ClustersSharesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "SharesApiService.ClustersSharesDelete": {
    "port": "8443",
},
})
```


## ClustersSharesFindById

```go
var result ShareRead = ClustersSharesFindById(ctx, clusterId, shareId)
                      .Execute()
```

Retrieves a share



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-nfs"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.
    shareId := "7b1ef56d-dfc6-51fe-aff0-7af2d6747868" // string | The share identifier.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SharesApi.ClustersSharesFindById(context.Background(), clusterId, shareId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.ClustersSharesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersSharesFindById`: ShareRead
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.ClustersSharesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |
|**shareId** | **string** | The share identifier. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSharesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ShareRead**](../models/ShareRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"SharesApiService.ClustersSharesFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "SharesApiService.ClustersSharesFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "SharesApiService.ClustersSharesFindById": {
    "port": "8443",
},
})
```


## ClustersSharesGet

```go
var result ShareReadList = ClustersSharesGet(ctx, clusterId)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieves all shares



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-nfs"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.
    offset := int32(0) // int32 | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SharesApi.ClustersSharesGet(context.Background(), clusterId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.ClustersSharesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersSharesGet`: ShareReadList
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.ClustersSharesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSharesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|

### Return type

[**ShareReadList**](../models/ShareReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"SharesApiService.ClustersSharesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "SharesApiService.ClustersSharesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "SharesApiService.ClustersSharesGet": {
    "port": "8443",
},
})
```


## ClustersSharesPost

```go
var result ShareRead = ClustersSharesPost(ctx, clusterId)
                      .ShareCreate(shareCreate)
                      .Execute()
```

Create Share



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-nfs"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.
    shareCreate := *openapiclient.NewShareCreate(*openapiclient.NewShare("Name_example", []openapiclient.ShareClientGroups{*openapiclient.NewShareClientGroups()})) // ShareCreate | Share to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SharesApi.ClustersSharesPost(context.Background(), clusterId).ShareCreate(shareCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.ClustersSharesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersSharesPost`: ShareRead
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.ClustersSharesPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSharesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **shareCreate** | [**ShareCreate**](../models/ShareCreate.md) | Share to create. | |

### Return type

[**ShareRead**](../models/ShareRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"SharesApiService.ClustersSharesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "SharesApiService.ClustersSharesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "SharesApiService.ClustersSharesPost": {
    "port": "8443",
},
})
```


## ClustersSharesPut

```go
var result ShareRead = ClustersSharesPut(ctx, clusterId, shareId)
                      .ShareEnsure(shareEnsure)
                      .Execute()
```

Ensure Share



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-nfs"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.
    shareId := "7b1ef56d-dfc6-51fe-aff0-7af2d6747868" // string | The identifier (UUID) of the share.
    shareEnsure := *openapiclient.NewShareEnsure("7b1ef56d-dfc6-51fe-aff0-7af2d6747868", *openapiclient.NewShare("Name_example", []openapiclient.ShareClientGroups{*openapiclient.NewShareClientGroups()})) // ShareEnsure | Update Share

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SharesApi.ClustersSharesPut(context.Background(), clusterId, shareId).ShareEnsure(shareEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesApi.ClustersSharesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersSharesPut`: ShareRead
    fmt.Fprintf(os.Stdout, "Response from `SharesApi.ClustersSharesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |
|**shareId** | **string** | The identifier (UUID) of the share. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSharesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **shareEnsure** | [**ShareEnsure**](../models/ShareEnsure.md) | Update Share | |

### Return type

[**ShareRead**](../models/ShareRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"SharesApiService.ClustersSharesPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "SharesApiService.ClustersSharesPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "SharesApiService.ClustersSharesPut": {
    "port": "8443",
},
})
```

