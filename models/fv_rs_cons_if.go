package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnfvRsConsIf        = "uni/tn-%s/ap-%s/epg-%s/rsconsIf-%s"
	RnfvRsConsIf        = "rsconsIf-%s"
	ParentDnfvRsConsIf  = "uni/tn-%s/ap-%s/epg-%s"
	FvrsconsifClassName = "fvRsConsIf"
)

type ContractInterface struct {
	BaseAttributes
	NameAliasAttribute
	ContractInterfaceAttributes
}

type ContractInterfaceAttributes struct {
	Annotation   string `json:",omitempty"`
	Prio         string `json:",omitempty"`
	TnVzCPIfName string `json:",omitempty"`
}

func NewContractInterface(fvRsConsIfRn, parentDn, description, nameAlias string, fvRsConsIfAttr ContractInterfaceAttributes) *ContractInterface {
	dn := fmt.Sprintf("%s/%s", parentDn, fvRsConsIfRn)
	return &ContractInterface{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvrsconsifClassName,
			Rn:                fvRsConsIfRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		ContractInterfaceAttributes: fvRsConsIfAttr,
	}
}

func (fvRsConsIf *ContractInterface) ToMap() (map[string]string, error) {
	fvRsConsIfMap, err := fvRsConsIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := fvRsConsIf.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(fvRsConsIfMap, key, value)
	}

	A(fvRsConsIfMap, "prio", fvRsConsIf.Prio)
	A(fvRsConsIfMap, "tnVzCPIfName", fvRsConsIf.TnVzCPIfName)
	return fvRsConsIfMap, err
}

func ContractInterfaceFromContainerList(cont *container.Container, index int) *ContractInterface {
	ContractInterfaceCont := cont.S("imdata").Index(index).S(FvrsconsifClassName, "attributes")
	return &ContractInterface{
		BaseAttributes{
			DistinguishedName: G(ContractInterfaceCont, "dn"),
			Description:       G(ContractInterfaceCont, "descr"),
			Status:            G(ContractInterfaceCont, "status"),
			ClassName:         FvrsconsifClassName,
			Rn:                G(ContractInterfaceCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(ContractInterfaceCont, "nameAlias"),
		},
		ContractInterfaceAttributes{
			Prio:         G(ContractInterfaceCont, "prio"),
			TnVzCPIfName: G(ContractInterfaceCont, "tnVzCPIfName"),
		},
	}
}

func ContractInterfaceFromContainer(cont *container.Container) *ContractInterface {
	return ContractInterfaceFromContainerList(cont, 0)
}

func ContractInterfaceListFromContainer(cont *container.Container) []*ContractInterface {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*ContractInterface, length)

	for i := 0; i < length; i++ {
		arr[i] = ContractInterfaceFromContainerList(cont, i)
	}

	return arr
}
