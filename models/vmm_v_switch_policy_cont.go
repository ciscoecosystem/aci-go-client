package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VmmvswitchpolicycontClassName = "vmmVSwitchPolicyCont"

type VSwitchPolicyGroup struct {
	BaseAttributes
    VSwitchPolicyGroupAttributes 
}
  
type VSwitchPolicyGroupAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    
}
   

func NewVSwitchPolicyGroup(vmmVSwitchPolicyContRn, parentDn, description string, vmmVSwitchPolicyContattr VSwitchPolicyGroupAttributes) *VSwitchPolicyGroup {
	dn := fmt.Sprintf("%s/%s", parentDn, vmmVSwitchPolicyContRn)  
	return &VSwitchPolicyGroup{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VmmvswitchpolicycontClassName,
			Rn:                vmmVSwitchPolicyContRn,
		},
        
		VSwitchPolicyGroupAttributes: vmmVSwitchPolicyContattr,
         
	}
}

func (vmmVSwitchPolicyCont *VSwitchPolicyGroup) ToMap() (map[string]string, error) {
	vmmVSwitchPolicyContMap, err := vmmVSwitchPolicyCont.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(vmmVSwitchPolicyContMap, "annotation",vmmVSwitchPolicyCont.Annotation)
    A(vmmVSwitchPolicyContMap, "nameAlias",vmmVSwitchPolicyCont.NameAlias)
    
	

	return vmmVSwitchPolicyContMap, err
}

func VSwitchPolicyGroupFromContainerList(cont *container.Container, index int) *VSwitchPolicyGroup {

	VSwitchPolicyGroupCont := cont.S("imdata").Index(index).S(VmmvswitchpolicycontClassName, "attributes")
	return &VSwitchPolicyGroup{
		BaseAttributes{
			DistinguishedName: G(VSwitchPolicyGroupCont, "dn"),
			Description:       G(VSwitchPolicyGroupCont, "descr"),
			Status:            G(VSwitchPolicyGroupCont, "status"),
			ClassName:         VmmvswitchpolicycontClassName,
			Rn:                G(VSwitchPolicyGroupCont, "rn"),
		},
        
		VSwitchPolicyGroupAttributes{
        Annotation : G(VSwitchPolicyGroupCont, "annotation"),
        NameAlias : G(VSwitchPolicyGroupCont, "nameAlias"),
        		
        },
        
	}
}

func VSwitchPolicyGroupFromContainer(cont *container.Container) *VSwitchPolicyGroup {

	return VSwitchPolicyGroupFromContainerList(cont, 0)
}

func VSwitchPolicyGroupListFromContainer(cont *container.Container) []*VSwitchPolicyGroup {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*VSwitchPolicyGroup, length)

	for i := 0; i < length; i++ {

		arr[i] = VSwitchPolicyGroupFromContainerList(cont, i)
	}

	return arr
}