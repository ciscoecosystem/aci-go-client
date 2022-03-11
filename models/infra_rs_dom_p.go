package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DninfraRsDomP        = "%s/%s"
	RninfraRsDomP        = "rsdomP-[%s]"
	ParentDninfraRsDomP  = "uni/infra/attentp-%s"
	InfrarsdompClassName = "infraRsDomP"
)

type InfraRsDomP struct {
	BaseAttributes
	NameAliasAttribute
	InfraRsDomPAttributes
}

type InfraRsDomPAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewInfraRsDomP(infraRsDomPRn, parentDn, nameAlias string, infraRsDomPAttr InfraRsDomPAttributes) *InfraRsDomP {
	dn := fmt.Sprintf("%s/%s", parentDn, infraRsDomPRn)
	return &InfraRsDomP{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         InfrarsdompClassName,
			Rn:                infraRsDomPRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		InfraRsDomPAttributes: infraRsDomPAttr,
	}
}

func (infraRsDomP *InfraRsDomP) ToMap() (map[string]string, error) {
	infraRsDomPMap, err := infraRsDomP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := infraRsDomP.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(infraRsDomPMap, key, value)
	}

	A(infraRsDomPMap, "annotation", infraRsDomP.Annotation)
	A(infraRsDomPMap, "tDn", infraRsDomP.TDn)
	return infraRsDomPMap, err
}

func InfraRsDomPFromContainerList(cont *container.Container, index int) *InfraRsDomP {
	InfraRsDomPCont := cont.S("imdata").Index(index).S(InfrarsdompClassName, "attributes")
	return &InfraRsDomP{
		BaseAttributes{
			DistinguishedName: G(InfraRsDomPCont, "dn"),
			Status:            G(InfraRsDomPCont, "status"),
			ClassName:         InfrarsdompClassName,
			Rn:                G(InfraRsDomPCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(InfraRsDomPCont, "nameAlias"),
		},
		InfraRsDomPAttributes{
			Annotation: G(InfraRsDomPCont, "annotation"),
			TDn:        G(InfraRsDomPCont, "tDn"),
		},
	}
}

func InfraRsDomPFromContainer(cont *container.Container) *InfraRsDomP {
	return InfraRsDomPFromContainerList(cont, 0)
}

func InfraRsDomPListFromContainer(cont *container.Container) []*InfraRsDomP {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*InfraRsDomP, length)

	for i := 0; i < length; i++ {
		arr[i] = InfraRsDomPFromContainerList(cont, i)
	}

	return arr
}
