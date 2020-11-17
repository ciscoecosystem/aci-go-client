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

aci-go-client supports concurrent connections to different targets by calling client.NewClient() instead of client.GetClient().

```golang
client.NewClient("URL", "Username", client.PrivateKey("PrivateKey path"),client.AdminCert("Certificate name"), client.Insecure(true/false))
```

When making PyQuery calls (API calls which start with /mqapi2/), ensure that the APIC-Cookie is populated in the Requests
as PyQuery may not function using only the Certificate-Based Request.

If using Username/Password, no extra steps are necessary, as the APIC-Cookie will be obtained at the time of the API call.
If using Certificate based authentication, call client.Authenticate() first, to obtain a recent APIC-Cookie authorization token
PyQuery APIs currently do not support Certificate + Username authentication.

```golang
client.NewClient("URL", "Username", client.AppUserName("AppUserName"), client.PrivateKey("PrivateKey path"), client.AdminCert("Certificate name"), client.Insecure(true/false))
```
