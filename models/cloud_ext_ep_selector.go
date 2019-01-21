package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudextepselectorClassName = "cloudExtEPSelector"

type Cloudendpointselectorforexternalepgs struct {
	BaseAttributes
    CloudendpointselectorforexternalepgsAttributes 
}
  
type CloudendpointselectorforexternalepgsAttributes struct {
    Annotation       string `json:",omitempty"`
    IsShared       string `json:",omitempty"`
    MatchExpression       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Subnet       string `json:",omitempty"`
    
}
   

func NewCloudendpointselectorforexternalepgs(cloudExtEPSelectorRn, parentDn, description string, cloudExtEPSelectorattr CloudendpointselectorforexternalepgsAttributes) *Cloudendpointselectorforexternalepgs {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudExtEPSelectorRn)  
	return &Cloudendpointselectorforexternalepgs{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudextepselectorClassName,
			Rn:                cloudExtEPSelectorRn,
		},
        
		CloudendpointselectorforexternalepgsAttributes: cloudExtEPSelectorattr,
         
	}
}

func (cloudExtEPSelector *Cloudendpointselectorforexternalepgs) ToMap() (map[string]string, error) {
	cloudExtEPSelectorMap, err := cloudExtEPSelector.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudExtEPSelectorMap, "annotation",cloudExtEPSelector.Annotation)
    A(cloudExtEPSelectorMap, "isShared",cloudExtEPSelector.IsShared)
    A(cloudExtEPSelectorMap, "matchExpression",cloudExtEPSelector.MatchExpression)
    A(cloudExtEPSelectorMap, "nameAlias",cloudExtEPSelector.NameAlias)
    A(cloudExtEPSelectorMap, "subnet",cloudExtEPSelector.Subnet)
    
	

	return cloudExtEPSelectorMap, err
}

func CloudendpointselectorforexternalepgsFromContainerList(cont *container.Container, index int) *Cloudendpointselectorforexternalepgs {

	CloudendpointselectorforexternalepgsCont := cont.S("imdata").Index(index).S(CloudextepselectorClassName, "attributes")
	return &Cloudendpointselectorforexternalepgs{
		BaseAttributes{
			DistinguishedName: G(CloudendpointselectorforexternalepgsCont, "dn"),
			Description:       G(CloudendpointselectorforexternalepgsCont, "descr"),
			Status:            G(CloudendpointselectorforexternalepgsCont, "status"),
			ClassName:         CloudextepselectorClassName,
			Rn:                G(CloudendpointselectorforexternalepgsCont, "rn"),
		},
        
		CloudendpointselectorforexternalepgsAttributes{
        Annotation : G(CloudendpointselectorforexternalepgsCont, "annotation"),
        IsShared : G(CloudendpointselectorforexternalepgsCont, "isShared"),
        MatchExpression : G(CloudendpointselectorforexternalepgsCont, "matchExpression"),
        NameAlias : G(CloudendpointselectorforexternalepgsCont, "nameAlias"),
        Subnet : G(CloudendpointselectorforexternalepgsCont, "subnet"),
        		
        },
        
	}
}

func CloudendpointselectorforexternalepgsFromContainer(cont *container.Container) *Cloudendpointselectorforexternalepgs {

	return CloudendpointselectorforexternalepgsFromContainerList(cont, 0)
}

func CloudendpointselectorforexternalepgsListFromContainer(cont *container.Container) []*Cloudendpointselectorforexternalepgs {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudendpointselectorforexternalepgs, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudendpointselectorforexternalepgsFromContainerList(cont, i)
	}

	return arr
}