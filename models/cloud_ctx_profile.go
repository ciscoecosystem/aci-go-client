package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudctxprofileClassName = "cloudCtxProfile"

type Cloudcontextprofile struct {
	BaseAttributes
    CloudcontextprofileAttributes 
}
  
type CloudcontextprofileAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Type       string `json:",omitempty"`
    
}
   

func NewCloudcontextprofile(cloudCtxProfileRn, parentDn, description string, cloudCtxProfileattr CloudcontextprofileAttributes) *Cloudcontextprofile {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudCtxProfileRn)  
	return &Cloudcontextprofile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudctxprofileClassName,
			Rn:                cloudCtxProfileRn,
		},
        
		CloudcontextprofileAttributes: cloudCtxProfileattr,
         
	}
}

func (cloudCtxProfile *Cloudcontextprofile) ToMap() (map[string]string, error) {
	cloudCtxProfileMap, err := cloudCtxProfile.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudCtxProfileMap, "annotation",cloudCtxProfile.Annotation)
    A(cloudCtxProfileMap, "nameAlias",cloudCtxProfile.NameAlias)
    A(cloudCtxProfileMap, "type",cloudCtxProfile.Type)
    
	

	return cloudCtxProfileMap, err
}

func CloudcontextprofileFromContainerList(cont *container.Container, index int) *Cloudcontextprofile {

	CloudcontextprofileCont := cont.S("imdata").Index(index).S(CloudctxprofileClassName, "attributes")
	return &Cloudcontextprofile{
		BaseAttributes{
			DistinguishedName: G(CloudcontextprofileCont, "dn"),
			Description:       G(CloudcontextprofileCont, "descr"),
			Status:            G(CloudcontextprofileCont, "status"),
			ClassName:         CloudctxprofileClassName,
			Rn:                G(CloudcontextprofileCont, "rn"),
		},
        
		CloudcontextprofileAttributes{
        Annotation : G(CloudcontextprofileCont, "annotation"),
        NameAlias : G(CloudcontextprofileCont, "nameAlias"),
        Type : G(CloudcontextprofileCont, "type"),
        		
        },
        
	}
}

func CloudcontextprofileFromContainer(cont *container.Container) *Cloudcontextprofile {

	return CloudcontextprofileFromContainerList(cont, 0)
}

func CloudcontextprofileListFromContainer(cont *container.Container) []*Cloudcontextprofile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudcontextprofile, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudcontextprofileFromContainerList(cont, i)
	}

	return arr
}