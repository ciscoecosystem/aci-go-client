package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudregionClassName = "cloudRegion"

type Cloudprovidersregion struct {
	BaseAttributes
    CloudprovidersregionAttributes 
}
  
type CloudprovidersregionAttributes struct {
    AdminSt       string `json:",omitempty"`
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewCloudprovidersregion(cloudRegionRn, parentDn, description string, cloudRegionattr CloudprovidersregionAttributes) *Cloudprovidersregion {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudRegionRn)  
	return &Cloudprovidersregion{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudregionClassName,
			Rn:                cloudRegionRn,
		},
        
		CloudprovidersregionAttributes: cloudRegionattr,
         
	}
}

func (cloudRegion *Cloudprovidersregion) ToMap() (map[string]string, error) {
	cloudRegionMap, err := cloudRegion.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudRegionMap, "adminSt",cloudRegion.AdminSt)
    A(cloudRegionMap, "annotation",cloudRegion.Annotation)
    A(cloudRegionMap, "nameAlias",cloudRegion.NameAlias)
    
	

	return cloudRegionMap, err
}

func CloudprovidersregionFromContainerList(cont *container.Container, index int) *Cloudprovidersregion {

	CloudprovidersregionCont := cont.S("imdata").Index(index).S(CloudregionClassName, "attributes")
	return &Cloudprovidersregion{
		BaseAttributes{
			DistinguishedName: G(CloudprovidersregionCont, "dn"),
			Description:       G(CloudprovidersregionCont, "descr"),
			Status:            G(CloudprovidersregionCont, "status"),
			ClassName:         CloudregionClassName,
			Rn:                G(CloudprovidersregionCont, "rn"),
		},
        
		CloudprovidersregionAttributes{
        AdminSt : G(CloudprovidersregionCont, "adminSt"),
        Annotation : G(CloudprovidersregionCont, "annotation"),
        NameAlias : G(CloudprovidersregionCont, "nameAlias"),
        		
        },
        
	}
}

func CloudprovidersregionFromContainer(cont *container.Container) *Cloudprovidersregion {

	return CloudprovidersregionFromContainerList(cont, 0)
}

func CloudprovidersregionListFromContainer(cont *container.Container) []*Cloudprovidersregion {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudprovidersregion, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudprovidersregionFromContainerList(cont, i)
	}

	return arr
}