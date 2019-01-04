# aci-go-client
 This repository contains the golang client SDK to interact with CISCO APIC using REST API calls. This SDK is used by [terraform-provider-aci](https://github.com/ciscoecosystem/terraform-provider-aci).

## Installation ##

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.


```sh
$go get github.com/ciscoecosystem/aci-go-client
```

There are no additional dependancies needed to be installed.

## Overview ##
  
* <strong>client</strong> :- This package contains the HTTP Client configuration as well as service methods which serves the CRUD operations on the Model Objects in CISCO ACI.

* <strong>models</strong> :- This package contains all the models structs and utility methods for the same.

* <strong>tests</strong> :- This package contains the unit tests for the CRUD operations that can be performed on the Model Objects.

## How to Use ##

import the client in your go application and retrive the client object by calling client.GetClient() method.
```golang
import github.com/ciscoecosystem/aci-go-client/client
client.GetClient("URL", "Username", client.Password("Password"), client.Insecure(true/false))
```

aci-go-client supports signature based authentication also. To use signature based authentication call the GetClient method as follows.  
  

```golang
client.GetClient("URL", "Username", client.PrivateKey("PrivateKey path"),client.AdminCert("Certificate name"), client.Insecure(true/false))

```

Use that client object to call the service methods to perform the CRUD operations on the model objects.

Example,

```golang
    client.CreateTenant("tenant_name","description",tenantAttributesStruct)
    # tenantAttributesStruct is struct present in models/fv_tenant.go
```