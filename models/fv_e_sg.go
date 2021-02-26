package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvesgClassName = "fvESg"

type EndpointSecurityGroup struct {
	BaseAttributes
    EndpointSecurityGroupAttributes 
}
  
type EndpointSecurityGroupAttributes struct {
	
	
	Name string `json:",omitempty"`
	
	
    
	Annotation       string `json:",omitempty"`
	
    
	ExceptionTag       string `json:",omitempty"`
	
    
	FloodOnEncap       string `json:",omitempty"`
	
    
	MatchT       string `json:",omitempty"`
	
    
	NameAlias       string `json:",omitempty"`
	
    
	PcEnfPref       string `json:",omitempty"`
	
    
	PrefGrMemb       string `json:",omitempty"`
	
    
	Prio       string `json:",omitempty"`
	
    
	Userdom       string `json:",omitempty"`
	
    
}
   

func NewEndpointSecurityGroup(fvESgRn, parentDn, description string, fvESgattr EndpointSecurityGroupAttributes) *EndpointSecurityGroup {
	dn := fmt.Sprintf("%s/%s", parentDn, fvESgRn)  
	return &EndpointSecurityGroup{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvesgClassName,
			Rn:                fvESgRn,
		},
        
		EndpointSecurityGroupAttributes: fvESgattr,
         
	}
}

func (fvESg *EndpointSecurityGroup) ToMap() (map[string]string, error) {
	fvESgMap, err := fvESg.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	
	
	A(fvESgMap, "name",fvESg.Name)
	
	
    
	A(fvESgMap, "annotation",fvESg.Annotation)
	
    
	A(fvESgMap, "exceptionTag",fvESg.ExceptionTag)
	
    
	A(fvESgMap, "floodOnEncap",fvESg.FloodOnEncap)
	
    
	A(fvESgMap, "matchT",fvESg.MatchT)
	
    
	A(fvESgMap, "nameAlias",fvESg.NameAlias)
	
    
	A(fvESgMap, "pcEnfPref",fvESg.PcEnfPref)
	
    
	A(fvESgMap, "prefGrMemb",fvESg.PrefGrMemb)
	
    
	A(fvESgMap, "prio",fvESg.Prio)
	
    
	A(fvESgMap, "userdom",fvESg.Userdom)
	
    
	

	return fvESgMap, err
}

func EndpointSecurityGroupFromContainerList(cont *container.Container, index int) *EndpointSecurityGroup {

	EndpointSecurityGroupCont := cont.S("imdata").Index(index).S(FvesgClassName, "attributes")
	return &EndpointSecurityGroup{
		BaseAttributes{
			DistinguishedName: G(EndpointSecurityGroupCont, "dn"),
			Description:       G(EndpointSecurityGroupCont, "descr"),
			Status:            G(EndpointSecurityGroupCont, "status"),
			ClassName:         FvesgClassName,
			Rn:                G(EndpointSecurityGroupCont, "rn"),
		},
        
		EndpointSecurityGroupAttributes{
		
		
			Name : G(EndpointSecurityGroupCont, "name"),
		
		
        
	        Annotation : G(EndpointSecurityGroupCont, "annotation"),
		
        
	        ExceptionTag : G(EndpointSecurityGroupCont, "exceptionTag"),
		
        
	        FloodOnEncap : G(EndpointSecurityGroupCont, "floodOnEncap"),
		
        
	        MatchT : G(EndpointSecurityGroupCont, "matchT"),
		
        
	        NameAlias : G(EndpointSecurityGroupCont, "nameAlias"),
		
        
	        PcEnfPref : G(EndpointSecurityGroupCont, "pcEnfPref"),
		
        
	        PrefGrMemb : G(EndpointSecurityGroupCont, "prefGrMemb"),
		
        
	        Prio : G(EndpointSecurityGroupCont, "prio"),
		
        
	        Userdom : G(EndpointSecurityGroupCont, "userdom"),
		
        		
        },
        
	}
}

func EndpointSecurityGroupFromContainer(cont *container.Container) *EndpointSecurityGroup {

	return EndpointSecurityGroupFromContainerList(cont, 0)
}

func EndpointSecurityGroupListFromContainer(cont *container.Container) []*EndpointSecurityGroup {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*EndpointSecurityGroup, length)

	for i := 0; i < length; i++ {

		arr[i] = EndpointSecurityGroupFromContainerList(cont, i)
	}

	return arr
}