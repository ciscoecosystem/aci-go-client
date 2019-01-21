package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudawsproviderClassName = "cloudAwsProvider"

type Cloudawsprovider struct {
	BaseAttributes
    CloudawsproviderAttributes 
}
  
type CloudawsproviderAttributes struct {
    AccessKeyId       string `json:",omitempty"`
    AccountId       string `json:",omitempty"`
    Annotation       string `json:",omitempty"`
    Email       string `json:",omitempty"`
    HttpProxy       string `json:",omitempty"`
    IsAccountInOrg       string `json:",omitempty"`
    IsTrusted       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    ProviderId       string `json:",omitempty"`
    Region       string `json:",omitempty"`
    SecretAccessKey       string `json:",omitempty"`
    
}
   

func NewCloudawsprovider(cloudAwsProviderRn, parentDn, description string, cloudAwsProviderattr CloudawsproviderAttributes) *Cloudawsprovider {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudAwsProviderRn)  
	return &Cloudawsprovider{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudawsproviderClassName,
			Rn:                cloudAwsProviderRn,
		},
        
		CloudawsproviderAttributes: cloudAwsProviderattr,
         
	}
}

func (cloudAwsProvider *Cloudawsprovider) ToMap() (map[string]string, error) {
	cloudAwsProviderMap, err := cloudAwsProvider.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudAwsProviderMap, "accessKeyId",cloudAwsProvider.AccessKeyId)
    A(cloudAwsProviderMap, "accountId",cloudAwsProvider.AccountId)
    A(cloudAwsProviderMap, "annotation",cloudAwsProvider.Annotation)
    A(cloudAwsProviderMap, "email",cloudAwsProvider.Email)
    A(cloudAwsProviderMap, "httpProxy",cloudAwsProvider.HttpProxy)
    A(cloudAwsProviderMap, "isAccountInOrg",cloudAwsProvider.IsAccountInOrg)
    A(cloudAwsProviderMap, "isTrusted",cloudAwsProvider.IsTrusted)
    A(cloudAwsProviderMap, "nameAlias",cloudAwsProvider.NameAlias)
    A(cloudAwsProviderMap, "providerId",cloudAwsProvider.ProviderId)
    A(cloudAwsProviderMap, "region",cloudAwsProvider.Region)
    A(cloudAwsProviderMap, "secretAccessKey",cloudAwsProvider.SecretAccessKey)
    
	

	return cloudAwsProviderMap, err
}

func CloudawsproviderFromContainerList(cont *container.Container, index int) *Cloudawsprovider {

	CloudawsproviderCont := cont.S("imdata").Index(index).S(CloudawsproviderClassName, "attributes")
	return &Cloudawsprovider{
		BaseAttributes{
			DistinguishedName: G(CloudawsproviderCont, "dn"),
			Description:       G(CloudawsproviderCont, "descr"),
			Status:            G(CloudawsproviderCont, "status"),
			ClassName:         CloudawsproviderClassName,
			Rn:                G(CloudawsproviderCont, "rn"),
		},
        
		CloudawsproviderAttributes{
        AccessKeyId : G(CloudawsproviderCont, "accessKeyId"),
        AccountId : G(CloudawsproviderCont, "accountId"),
        Annotation : G(CloudawsproviderCont, "annotation"),
        Email : G(CloudawsproviderCont, "email"),
        HttpProxy : G(CloudawsproviderCont, "httpProxy"),
        IsAccountInOrg : G(CloudawsproviderCont, "isAccountInOrg"),
        IsTrusted : G(CloudawsproviderCont, "isTrusted"),
        NameAlias : G(CloudawsproviderCont, "nameAlias"),
        ProviderId : G(CloudawsproviderCont, "providerId"),
        Region : G(CloudawsproviderCont, "region"),
        SecretAccessKey : G(CloudawsproviderCont, "secretAccessKey"),
        		
        },
        
	}
}

func CloudawsproviderFromContainer(cont *container.Container) *Cloudawsprovider {

	return CloudawsproviderFromContainerList(cont, 0)
}

func CloudawsproviderListFromContainer(cont *container.Container) []*Cloudawsprovider {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudawsprovider, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudawsproviderFromContainerList(cont, i)
	}

	return arr
}