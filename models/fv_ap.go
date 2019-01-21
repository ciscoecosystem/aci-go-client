package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvapClassName = "fvAp"

type Applicationprofile struct {
	BaseAttributes
    ApplicationprofileAttributes 
}
  
type ApplicationprofileAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Prio       string `json:",omitempty"`
    
}
   

func NewApplicationprofile(fvApRn, parentDn, description string, fvApattr ApplicationprofileAttributes) *Applicationprofile {
	dn := fmt.Sprintf("%s/%s", parentDn, fvApRn)  
	return &Applicationprofile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvapClassName,
			Rn:                fvApRn,
		},
        
		ApplicationprofileAttributes: fvApattr,
         
	}
}

func (fvAp *Applicationprofile) ToMap() (map[string]string, error) {
	fvApMap, err := fvAp.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(fvApMap, "annotation",fvAp.Annotation)
    A(fvApMap, "nameAlias",fvAp.NameAlias)
    A(fvApMap, "prio",fvAp.Prio)
    
	

	return fvApMap, err
}

func ApplicationprofileFromContainerList(cont *container.Container, index int) *Applicationprofile {

	ApplicationprofileCont := cont.S("imdata").Index(index).S(FvapClassName, "attributes")
	return &Applicationprofile{
		BaseAttributes{
			DistinguishedName: G(ApplicationprofileCont, "dn"),
			Description:       G(ApplicationprofileCont, "descr"),
			Status:            G(ApplicationprofileCont, "status"),
			ClassName:         FvapClassName,
			Rn:                G(ApplicationprofileCont, "rn"),
		},
        
		ApplicationprofileAttributes{
        Annotation : G(ApplicationprofileCont, "annotation"),
        NameAlias : G(ApplicationprofileCont, "nameAlias"),
        Prio : G(ApplicationprofileCont, "prio"),
        		
        },
        
	}
}

func ApplicationprofileFromContainer(cont *container.Container) *Applicationprofile {

	return ApplicationprofileFromContainerList(cont, 0)
}

func ApplicationprofileListFromContainer(cont *container.Container) []*Applicationprofile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Applicationprofile, length)

	for i := 0; i < length; i++ {

		arr[i] = ApplicationprofileFromContainerList(cont, i)
	}

	return arr
}