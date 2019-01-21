package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudsubnetClassName = "cloudSubnet"

type Cloudsubnet struct {
	BaseAttributes
    CloudsubnetAttributes 
}
  
type CloudsubnetAttributes struct {
    Annotation       string `json:",omitempty"`
    Ip       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Scope       string `json:",omitempty"`
    Usage       string `json:",omitempty"`
    
}
   

func NewCloudsubnet(cloudSubnetRn, parentDn, description string, cloudSubnetattr CloudsubnetAttributes) *Cloudsubnet {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudSubnetRn)  
	return &Cloudsubnet{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudsubnetClassName,
			Rn:                cloudSubnetRn,
		},
        
		CloudsubnetAttributes: cloudSubnetattr,
         
	}
}

func (cloudSubnet *Cloudsubnet) ToMap() (map[string]string, error) {
	cloudSubnetMap, err := cloudSubnet.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudSubnetMap, "annotation",cloudSubnet.Annotation)
    A(cloudSubnetMap, "ip",cloudSubnet.Ip)
    A(cloudSubnetMap, "nameAlias",cloudSubnet.NameAlias)
    A(cloudSubnetMap, "scope",cloudSubnet.Scope)
    A(cloudSubnetMap, "usage",cloudSubnet.Usage)
    
	

	return cloudSubnetMap, err
}

func CloudsubnetFromContainerList(cont *container.Container, index int) *Cloudsubnet {

	CloudsubnetCont := cont.S("imdata").Index(index).S(CloudsubnetClassName, "attributes")
	return &Cloudsubnet{
		BaseAttributes{
			DistinguishedName: G(CloudsubnetCont, "dn"),
			Description:       G(CloudsubnetCont, "descr"),
			Status:            G(CloudsubnetCont, "status"),
			ClassName:         CloudsubnetClassName,
			Rn:                G(CloudsubnetCont, "rn"),
		},
        
		CloudsubnetAttributes{
        Annotation : G(CloudsubnetCont, "annotation"),
        Ip : G(CloudsubnetCont, "ip"),
        NameAlias : G(CloudsubnetCont, "nameAlias"),
        Scope : G(CloudsubnetCont, "scope"),
        Usage : G(CloudsubnetCont, "usage"),
        		
        },
        
	}
}

func CloudsubnetFromContainer(cont *container.Container) *Cloudsubnet {

	return CloudsubnetFromContainerList(cont, 0)
}

func CloudsubnetListFromContainer(cont *container.Container) []*Cloudsubnet {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudsubnet, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudsubnetFromContainerList(cont, i)
	}

	return arr
}