package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudprovpClassName = "cloudProvP"

type Cloudproviderprofile struct {
	BaseAttributes
    CloudproviderprofileAttributes 
}
  
type CloudproviderprofileAttributes struct {
    Annotation       string `json:",omitempty"`
    Vendor       string `json:",omitempty"`
    
}
   

func NewCloudproviderprofile(cloudProvPRn, parentDn, description string, cloudProvPattr CloudproviderprofileAttributes) *Cloudproviderprofile {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudProvPRn)  
	return &Cloudproviderprofile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudprovpClassName,
			Rn:                cloudProvPRn,
		},
        
		CloudproviderprofileAttributes: cloudProvPattr,
         
	}
}

func (cloudProvP *Cloudproviderprofile) ToMap() (map[string]string, error) {
	cloudProvPMap, err := cloudProvP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudProvPMap, "annotation",cloudProvP.Annotation)
    A(cloudProvPMap, "vendor",cloudProvP.Vendor)
    
	

	return cloudProvPMap, err
}

func CloudproviderprofileFromContainerList(cont *container.Container, index int) *Cloudproviderprofile {

	CloudproviderprofileCont := cont.S("imdata").Index(index).S(CloudprovpClassName, "attributes")
	return &Cloudproviderprofile{
		BaseAttributes{
			DistinguishedName: G(CloudproviderprofileCont, "dn"),
			Description:       G(CloudproviderprofileCont, "descr"),
			Status:            G(CloudproviderprofileCont, "status"),
			ClassName:         CloudprovpClassName,
			Rn:                G(CloudproviderprofileCont, "rn"),
		},
        
		CloudproviderprofileAttributes{
        Annotation : G(CloudproviderprofileCont, "annotation"),
        Vendor : G(CloudproviderprofileCont, "vendor"),
        		
        },
        
	}
}

func CloudproviderprofileFromContainer(cont *container.Container) *Cloudproviderprofile {

	return CloudproviderprofileFromContainerList(cont, 0)
}

func CloudproviderprofileListFromContainer(cont *container.Container) []*Cloudproviderprofile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudproviderprofile, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudproviderprofileFromContainerList(cont, i)
	}

	return arr
}