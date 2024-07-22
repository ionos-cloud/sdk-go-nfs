# \ClustersApi

All URIs are relative to *https://nfs.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete Cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Retrieve Cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Retrieve Clusters|
|[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create Cluster|
|[**ClustersPut**](ClustersApi.md#ClustersPut) | **Put** /clusters/{clusterId} | Ensure Cluster|



## ClustersDelete

```go
var result  = ClustersDelete(ctx, clusterId)
                      .Execute()
```

Delete Cluster



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.ClustersApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersDelete": {
    "port": "8443",
},
})
```


## ClustersFindById

```go
var result ClusterRead = ClustersFindById(ctx, clusterId)
                      .Execute()
```

Retrieve Cluster



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersFindById`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersFindById": {
    "port": "8443",
},
})
```


## ClustersGet

```go
var result ClusterReadList = ClustersGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterDatacenterId(filterDatacenterId)
                      .Execute()
```

Retrieve Clusters



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
    offset := int32(0) // int32 | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    filterDatacenterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The datacenter identifier to filter by. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).Offset(offset).Limit(limit).FilterDatacenterId(filterDatacenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersGet`: ClusterReadList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **filterDatacenterId** | **string** | The datacenter identifier to filter by. | |

### Return type

[**ClusterReadList**](../models/ClusterReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersGet": {
    "port": "8443",
},
})
```


## ClustersPost

```go
var result ClusterRead = ClustersPost(ctx)
                      .ClusterCreate(clusterCreate)
                      .Execute()
```

Create Cluster



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
    clusterCreate := *openapiclient.NewClusterCreate(*openapiclient.NewCluster("Cluster 1", []openapiclient.ClusterConnections{*openapiclient.NewClusterConnections("123e4567-e89b-12d3-a456-426614174001", "1", "10.254.64.1/24")})) // ClusterCreate | Cluster to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPost(context.Background()).ClusterCreate(clusterCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPost`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **clusterCreate** | [**ClusterCreate**](../models/ClusterCreate.md) | Cluster to create. | |

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersPost": {
    "port": "8443",
},
})
```


## ClustersPut

```go
var result ClusterRead = ClustersPut(ctx, clusterId)
                      .ClusterEnsure(clusterEnsure)
                      .Execute()
```

Ensure Cluster



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
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    clusterEnsure := *openapiclient.NewClusterEnsure("e69b22a5-8fee-56b1-b6fb-4a07e4205ead", *openapiclient.NewCluster("Cluster 1", []openapiclient.ClusterConnections{*openapiclient.NewClusterConnections("123e4567-e89b-12d3-a456-426614174001", "1", "10.254.64.1/24")})) // ClusterEnsure | Update Cluster

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPut(context.Background(), clusterId).ClusterEnsure(clusterEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPut`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **clusterEnsure** | [**ClusterEnsure**](../models/ClusterEnsure.md) | Update Cluster | |

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersPut": {
    "port": "8443",
},
})
```

