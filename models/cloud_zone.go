package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudzoneClassName = "cloudZone"

type Cloudavailabilityzone struct {
	BaseAttributes
    CloudavailabilityzoneAttributes 
}
  
type CloudavailabilityzoneAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewCloudavailabilityzone(cloudZoneRn, parentDn, description string, cloudZoneattr CloudavailabilityzoneAttributes) *Cloudavailabilityzone {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudZoneRn)  
	return &Cloudavailabilityzone{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudzoneClassName,
			Rn:                cloudZoneRn,
		},
        
		CloudavailabilityzoneAttributes: cloudZoneattr,
         
	}
}

func (cloudZone *Cloudavailabilityzone) ToMap() (map[string]string, error) {
	cloudZoneMap, err := cloudZone.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudZoneMap, "annotation",cloudZone.Annotation)
    A(cloudZoneMap, "nameAlias",cloudZone.NameAlias)
    
	

	return cloudZoneMap, err
}

func CloudavailabilityzoneFromContainerList(cont *container.Container, index int) *Cloudavailabilityzone {

	CloudavailabilityzoneCont := cont.S("imdata").Index(index).S(CloudzoneClassName, "attributes")
	return &Cloudavailabilityzone{
		BaseAttributes{
			DistinguishedName: G(CloudavailabilityzoneCont, "dn"),
			Description:       G(CloudavailabilityzoneCont, "descr"),
			Status:            G(CloudavailabilityzoneCont, "status"),
			ClassName:         CloudzoneClassName,
			Rn:                G(CloudavailabilityzoneCont, "rn"),
		},
        
		CloudavailabilityzoneAttributes{
        Annotation : G(CloudavailabilityzoneCont, "annotation"),
        NameAlias : G(CloudavailabilityzoneCont, "nameAlias"),
        		
        },
        
	}
}

func CloudavailabilityzoneFromContainer(cont *container.Container) *Cloudavailabilityzone {

	return CloudavailabilityzoneFromContainerList(cont, 0)
}

func CloudavailabilityzoneListFromContainer(cont *container.Container) []*Cloudavailabilityzone {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudavailabilityzone, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudavailabilityzoneFromContainerList(cont, i)
	}

	return arr
}