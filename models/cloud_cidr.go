package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudcidrClassName = "cloudCidr"

type Cloudcidrpool struct {
	BaseAttributes
    CloudcidrpoolAttributes 
}
  
type CloudcidrpoolAttributes struct {
    Addr       string `json:",omitempty"`
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Primary       string `json:",omitempty"`
    
}
   

func NewCloudcidrpool(cloudCidrRn, parentDn, description string, cloudCidrattr CloudcidrpoolAttributes) *Cloudcidrpool {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudCidrRn)  
	return &Cloudcidrpool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudcidrClassName,
			Rn:                cloudCidrRn,
		},
        
		CloudcidrpoolAttributes: cloudCidrattr,
         
	}
}

func (cloudCidr *Cloudcidrpool) ToMap() (map[string]string, error) {
	cloudCidrMap, err := cloudCidr.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudCidrMap, "addr",cloudCidr.Addr)
    A(cloudCidrMap, "annotation",cloudCidr.Annotation)
    A(cloudCidrMap, "nameAlias",cloudCidr.NameAlias)
    A(cloudCidrMap, "primary",cloudCidr.Primary)
    
	

	return cloudCidrMap, err
}

func CloudcidrpoolFromContainerList(cont *container.Container, index int) *Cloudcidrpool {

	CloudcidrpoolCont := cont.S("imdata").Index(index).S(CloudcidrClassName, "attributes")
	return &Cloudcidrpool{
		BaseAttributes{
			DistinguishedName: G(CloudcidrpoolCont, "dn"),
			Description:       G(CloudcidrpoolCont, "descr"),
			Status:            G(CloudcidrpoolCont, "status"),
			ClassName:         CloudcidrClassName,
			Rn:                G(CloudcidrpoolCont, "rn"),
		},
        
		CloudcidrpoolAttributes{
        Addr : G(CloudcidrpoolCont, "addr"),
        Annotation : G(CloudcidrpoolCont, "annotation"),
        NameAlias : G(CloudcidrpoolCont, "nameAlias"),
        Primary : G(CloudcidrpoolCont, "primary"),
        		
        },
        
	}
}

func CloudcidrpoolFromContainer(cont *container.Container) *Cloudcidrpool {

	return CloudcidrpoolFromContainerList(cont, 0)
}

func CloudcidrpoolListFromContainer(cont *container.Container) []*Cloudcidrpool {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudcidrpool, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudcidrpoolFromContainerList(cont, i)
	}

	return arr
}