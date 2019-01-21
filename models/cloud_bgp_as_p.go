package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudbgpaspClassName = "cloudBgpAsP"

type Autonomoussystemprofile struct {
	BaseAttributes
    AutonomoussystemprofileAttributes 
}
  
type AutonomoussystemprofileAttributes struct {
    Annotation       string `json:",omitempty"`
    Asn       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewAutonomoussystemprofile(cloudBgpAsPRn, parentDn, description string, cloudBgpAsPattr AutonomoussystemprofileAttributes) *Autonomoussystemprofile {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudBgpAsPRn)  
	return &Autonomoussystemprofile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudbgpaspClassName,
			Rn:                cloudBgpAsPRn,
		},
        
		AutonomoussystemprofileAttributes: cloudBgpAsPattr,
         
	}
}

func (cloudBgpAsP *Autonomoussystemprofile) ToMap() (map[string]string, error) {
	cloudBgpAsPMap, err := cloudBgpAsP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudBgpAsPMap, "annotation",cloudBgpAsP.Annotation)
    A(cloudBgpAsPMap, "asn",cloudBgpAsP.Asn)
    A(cloudBgpAsPMap, "nameAlias",cloudBgpAsP.NameAlias)
    
	

	return cloudBgpAsPMap, err
}

func AutonomoussystemprofileFromContainerList(cont *container.Container, index int) *Autonomoussystemprofile {

	AutonomoussystemprofileCont := cont.S("imdata").Index(index).S(CloudbgpaspClassName, "attributes")
	return &Autonomoussystemprofile{
		BaseAttributes{
			DistinguishedName: G(AutonomoussystemprofileCont, "dn"),
			Description:       G(AutonomoussystemprofileCont, "descr"),
			Status:            G(AutonomoussystemprofileCont, "status"),
			ClassName:         CloudbgpaspClassName,
			Rn:                G(AutonomoussystemprofileCont, "rn"),
		},
        
		AutonomoussystemprofileAttributes{
        Annotation : G(AutonomoussystemprofileCont, "annotation"),
        Asn : G(AutonomoussystemprofileCont, "asn"),
        NameAlias : G(AutonomoussystemprofileCont, "nameAlias"),
        		
        },
        
	}
}

func AutonomoussystemprofileFromContainer(cont *container.Container) *Autonomoussystemprofile {

	return AutonomoussystemprofileFromContainerList(cont, 0)
}

func AutonomoussystemprofileListFromContainer(cont *container.Container) []*Autonomoussystemprofile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Autonomoussystemprofile, length)

	for i := 0; i < length; i++ {

		arr[i] = AutonomoussystemprofileFromContainerList(cont, i)
	}

	return arr
}