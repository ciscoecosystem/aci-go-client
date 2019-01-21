package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const ClouddompClassName = "cloudDomP"

type Clouddomainprofile struct {
	BaseAttributes
    ClouddomainprofileAttributes 
}
  
type ClouddomainprofileAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    SiteId       string `json:",omitempty"`
    
}
   

func NewClouddomainprofile(cloudDomPRn, parentDn, description string, cloudDomPattr ClouddomainprofileAttributes) *Clouddomainprofile {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudDomPRn)  
	return &Clouddomainprofile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         ClouddompClassName,
			Rn:                cloudDomPRn,
		},
        
		ClouddomainprofileAttributes: cloudDomPattr,
         
	}
}

func (cloudDomP *Clouddomainprofile) ToMap() (map[string]string, error) {
	cloudDomPMap, err := cloudDomP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudDomPMap, "annotation",cloudDomP.Annotation)
    A(cloudDomPMap, "nameAlias",cloudDomP.NameAlias)
    A(cloudDomPMap, "siteId",cloudDomP.SiteId)
    
	

	return cloudDomPMap, err
}

func ClouddomainprofileFromContainerList(cont *container.Container, index int) *Clouddomainprofile {

	ClouddomainprofileCont := cont.S("imdata").Index(index).S(ClouddompClassName, "attributes")
	return &Clouddomainprofile{
		BaseAttributes{
			DistinguishedName: G(ClouddomainprofileCont, "dn"),
			Description:       G(ClouddomainprofileCont, "descr"),
			Status:            G(ClouddomainprofileCont, "status"),
			ClassName:         ClouddompClassName,
			Rn:                G(ClouddomainprofileCont, "rn"),
		},
        
		ClouddomainprofileAttributes{
        Annotation : G(ClouddomainprofileCont, "annotation"),
        NameAlias : G(ClouddomainprofileCont, "nameAlias"),
        SiteId : G(ClouddomainprofileCont, "siteId"),
        		
        },
        
	}
}

func ClouddomainprofileFromContainer(cont *container.Container) *Clouddomainprofile {

	return ClouddomainprofileFromContainerList(cont, 0)
}

func ClouddomainprofileListFromContainer(cont *container.Container) []*Clouddomainprofile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Clouddomainprofile, length)

	for i := 0; i < length; i++ {

		arr[i] = ClouddomainprofileFromContainerList(cont, i)
	}

	return arr
}