package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudepselectorClassName = "cloudEPSelector"

type Cloudendpointselector struct {
	BaseAttributes
    CloudendpointselectorAttributes 
}
  
type CloudendpointselectorAttributes struct {
    Annotation       string `json:",omitempty"`
    MatchExpression       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewCloudendpointselector(cloudEPSelectorRn, parentDn, description string, cloudEPSelectorattr CloudendpointselectorAttributes) *Cloudendpointselector {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudEPSelectorRn)  
	return &Cloudendpointselector{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudepselectorClassName,
			Rn:                cloudEPSelectorRn,
		},
        
		CloudendpointselectorAttributes: cloudEPSelectorattr,
         
	}
}

func (cloudEPSelector *Cloudendpointselector) ToMap() (map[string]string, error) {
	cloudEPSelectorMap, err := cloudEPSelector.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudEPSelectorMap, "annotation",cloudEPSelector.Annotation)
    A(cloudEPSelectorMap, "matchExpression",cloudEPSelector.MatchExpression)
    A(cloudEPSelectorMap, "nameAlias",cloudEPSelector.NameAlias)
    
	

	return cloudEPSelectorMap, err
}

func CloudendpointselectorFromContainerList(cont *container.Container, index int) *Cloudendpointselector {

	CloudendpointselectorCont := cont.S("imdata").Index(index).S(CloudepselectorClassName, "attributes")
	return &Cloudendpointselector{
		BaseAttributes{
			DistinguishedName: G(CloudendpointselectorCont, "dn"),
			Description:       G(CloudendpointselectorCont, "descr"),
			Status:            G(CloudendpointselectorCont, "status"),
			ClassName:         CloudepselectorClassName,
			Rn:                G(CloudendpointselectorCont, "rn"),
		},
        
		CloudendpointselectorAttributes{
        Annotation : G(CloudendpointselectorCont, "annotation"),
        MatchExpression : G(CloudendpointselectorCont, "matchExpression"),
        NameAlias : G(CloudendpointselectorCont, "nameAlias"),
        		
        },
        
	}
}

func CloudendpointselectorFromContainer(cont *container.Container) *Cloudendpointselector {

	return CloudendpointselectorFromContainerList(cont, 0)
}

func CloudendpointselectorListFromContainer(cont *container.Container) []*Cloudendpointselector {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudendpointselector, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudendpointselectorFromContainerList(cont, i)
	}

	return arr
}