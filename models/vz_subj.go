package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VzsubjClassName = "vzSubj"

type Contractsubject struct {
	BaseAttributes
    ContractsubjectAttributes 
}
  
type ContractsubjectAttributes struct {
    Annotation       string `json:",omitempty"`
    ConsMatchT       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Prio       string `json:",omitempty"`
    ProvMatchT       string `json:",omitempty"`
    RevFltPorts       string `json:",omitempty"`
    TargetDscp       string `json:",omitempty"`
    
}
   

func NewContractsubject(vzSubjRn, parentDn, description string, vzSubjattr ContractsubjectAttributes) *Contractsubject {
	dn := fmt.Sprintf("%s/%s", parentDn, vzSubjRn)  
	return &Contractsubject{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VzsubjClassName,
			Rn:                vzSubjRn,
		},
        
		ContractsubjectAttributes: vzSubjattr,
         
	}
}

func (vzSubj *Contractsubject) ToMap() (map[string]string, error) {
	vzSubjMap, err := vzSubj.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(vzSubjMap, "annotation",vzSubj.Annotation)
    A(vzSubjMap, "consMatchT",vzSubj.ConsMatchT)
    A(vzSubjMap, "nameAlias",vzSubj.NameAlias)
    A(vzSubjMap, "prio",vzSubj.Prio)
    A(vzSubjMap, "provMatchT",vzSubj.ProvMatchT)
    A(vzSubjMap, "revFltPorts",vzSubj.RevFltPorts)
    A(vzSubjMap, "targetDscp",vzSubj.TargetDscp)
    
	

	return vzSubjMap, err
}

func ContractsubjectFromContainerList(cont *container.Container, index int) *Contractsubject {

	ContractsubjectCont := cont.S("imdata").Index(index).S(VzsubjClassName, "attributes")
	return &Contractsubject{
		BaseAttributes{
			DistinguishedName: G(ContractsubjectCont, "dn"),
			Description:       G(ContractsubjectCont, "descr"),
			Status:            G(ContractsubjectCont, "status"),
			ClassName:         VzsubjClassName,
			Rn:                G(ContractsubjectCont, "rn"),
		},
        
		ContractsubjectAttributes{
        Annotation : G(ContractsubjectCont, "annotation"),
        ConsMatchT : G(ContractsubjectCont, "consMatchT"),
        NameAlias : G(ContractsubjectCont, "nameAlias"),
        Prio : G(ContractsubjectCont, "prio"),
        ProvMatchT : G(ContractsubjectCont, "provMatchT"),
        RevFltPorts : G(ContractsubjectCont, "revFltPorts"),
        TargetDscp : G(ContractsubjectCont, "targetDscp"),
        		
        },
        
	}
}

func ContractsubjectFromContainer(cont *container.Container) *Contractsubject {

	return ContractsubjectFromContainerList(cont, 0)
}

func ContractsubjectListFromContainer(cont *container.Container) []*Contractsubject {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Contractsubject, length)

	for i := 0; i < length; i++ {

		arr[i] = ContractsubjectFromContainerList(cont, i)
	}

	return arr
}