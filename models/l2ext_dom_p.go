package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const L2extdompClassName = "l2extDomP"

type L2DomainProfile struct {
	BaseAttributes
    L2DomainProfileAttributes 
}
  
type L2DomainProfileAttributes struct {
	
	
	Name string `json:",omitempty"`
	
	
    
	Annotation       string `json:",omitempty"`
	
    
	NameAlias       string `json:",omitempty"`
	
    
}
   

func NewL2DomainProfile(l2extDomPRn, parentDn, description string, l2extDomPattr L2DomainProfileAttributes) *L2DomainProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, l2extDomPRn)  
	return &L2DomainProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         L2extdompClassName,
			Rn:                l2extDomPRn,
		},
        
		L2DomainProfileAttributes: l2extDomPattr,
         
	}
}

func (l2extDomP *L2DomainProfile) ToMap() (map[string]string, error) {
	l2extDomPMap, err := l2extDomP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	
	
	A(l2extDomPMap, "name",l2extDomP.Name)
	
	
    
	A(l2extDomPMap, "annotation",l2extDomP.Annotation)
	
    
	A(l2extDomPMap, "nameAlias",l2extDomP.NameAlias)
	
    
	

	return l2extDomPMap, err
}

func L2DomainProfileFromContainerList(cont *container.Container, index int) *L2DomainProfile {

	L2DomainProfileCont := cont.S("imdata").Index(index).S(L2extdompClassName, "attributes")
	return &L2DomainProfile{
		BaseAttributes{
			DistinguishedName: G(L2DomainProfileCont, "dn"),
			Description:       G(L2DomainProfileCont, "descr"),
			Status:            G(L2DomainProfileCont, "status"),
			ClassName:         L2extdompClassName,
			Rn:                G(L2DomainProfileCont, "rn"),
		},
        
		L2DomainProfileAttributes{
		
		
			Name : G(L2DomainProfileCont, "name"),
		
		
        
	        Annotation : G(L2DomainProfileCont, "annotation"),
		
        
	        NameAlias : G(L2DomainProfileCont, "nameAlias"),
		
        		
        },
        
	}
}

func L2DomainProfileFromContainer(cont *container.Container) *L2DomainProfile {

	return L2DomainProfileFromContainerList(cont, 0)
}

func L2DomainProfileListFromContainer(cont *container.Container) []*L2DomainProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*L2DomainProfile, length)

	for i := 0; i < length; i++ {

		arr[i] = L2DomainProfileFromContainerList(cont, i)
	}

	return arr
}