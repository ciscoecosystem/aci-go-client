package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudappClassName = "cloudApp"

type Cloudapplicationcontainer struct {
	BaseAttributes
    CloudapplicationcontainerAttributes 
}
  
type CloudapplicationcontainerAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewCloudapplicationcontainer(cloudAppRn, parentDn, description string, cloudAppattr CloudapplicationcontainerAttributes) *Cloudapplicationcontainer {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudAppRn)  
	return &Cloudapplicationcontainer{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudappClassName,
			Rn:                cloudAppRn,
		},
        
		CloudapplicationcontainerAttributes: cloudAppattr,
         
	}
}

func (cloudApp *Cloudapplicationcontainer) ToMap() (map[string]string, error) {
	cloudAppMap, err := cloudApp.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudAppMap, "annotation",cloudApp.Annotation)
    A(cloudAppMap, "nameAlias",cloudApp.NameAlias)
    
	

	return cloudAppMap, err
}

func CloudapplicationcontainerFromContainerList(cont *container.Container, index int) *Cloudapplicationcontainer {

	CloudapplicationcontainerCont := cont.S("imdata").Index(index).S(CloudappClassName, "attributes")
	return &Cloudapplicationcontainer{
		BaseAttributes{
			DistinguishedName: G(CloudapplicationcontainerCont, "dn"),
			Description:       G(CloudapplicationcontainerCont, "descr"),
			Status:            G(CloudapplicationcontainerCont, "status"),
			ClassName:         CloudappClassName,
			Rn:                G(CloudapplicationcontainerCont, "rn"),
		},
        
		CloudapplicationcontainerAttributes{
        Annotation : G(CloudapplicationcontainerCont, "annotation"),
        NameAlias : G(CloudapplicationcontainerCont, "nameAlias"),
        		
        },
        
	}
}

func CloudapplicationcontainerFromContainer(cont *container.Container) *Cloudapplicationcontainer {

	return CloudapplicationcontainerFromContainerList(cont, 0)
}

func CloudapplicationcontainerListFromContainer(cont *container.Container) []*Cloudapplicationcontainer {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudapplicationcontainer, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudapplicationcontainerFromContainerList(cont, i)
	}

	return arr
}