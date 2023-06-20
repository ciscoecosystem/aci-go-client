package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnPimIPV6IfP        = "pimipv6ifp"
	DnPimIPV6IfP        = "uni/tn-%s/out-%s/lnodep-%s/lifp-%s/pimipv6ifp"
	ParentDnPimIPV6IfP  = "uni/tn-%s/out-%s/lnodep-%s/lifp-%s"
	PimIPV6IfPClassName = "pimIPV6IfP"
)

type PimIPv6InterfaceProfile struct {
	BaseAttributes
	PimIPv6InterfaceProfileAttributes
}

type PimIPv6InterfaceProfileAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewPimIPv6InterfaceProfile(pimIPV6IfPRn, parentDn, description string, pimIPV6IfPAttr PimIPv6InterfaceProfileAttributes) *PimIPv6InterfaceProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, pimIPV6IfPRn)
	return &PimIPv6InterfaceProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         PimIPV6IfPClassName,
			Rn:                pimIPV6IfPRn,
		},
		PimIPv6InterfaceProfileAttributes: pimIPV6IfPAttr,
	}
}

func (pimIPV6IfP *PimIPv6InterfaceProfile) ToMap() (map[string]string, error) {
	pimIPV6IfPMap, err := pimIPV6IfP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(pimIPV6IfPMap, "annotation", pimIPV6IfP.Annotation)
	A(pimIPV6IfPMap, "name", pimIPV6IfP.Name)
	A(pimIPV6IfPMap, "nameAlias", pimIPV6IfP.NameAlias)
	return pimIPV6IfPMap, err
}

func PimIPv6InterfaceProfileFromContainerList(cont *container.Container, index int) *PimIPv6InterfaceProfile {
	InterfaceProfileCont := cont.S("imdata").Index(index).S(PimIPV6IfPClassName, "attributes")
	return &PimIPv6InterfaceProfile{
		BaseAttributes{
			DistinguishedName: G(InterfaceProfileCont, "dn"),
			Description:       G(InterfaceProfileCont, "descr"),
			Status:            G(InterfaceProfileCont, "status"),
			ClassName:         PimIPV6IfPClassName,
			Rn:                G(InterfaceProfileCont, "rn"),
		},
		PimIPv6InterfaceProfileAttributes{
			Annotation: G(InterfaceProfileCont, "annotation"),
			Name:       G(InterfaceProfileCont, "name"),
			NameAlias:  G(InterfaceProfileCont, "nameAlias"),
		},
	}
}

func PimIPv6InterfaceProfileFromContainer(cont *container.Container) *PimIPv6InterfaceProfile {
	return PimIPv6InterfaceProfileFromContainerList(cont, 0)
}

func PimIPv6InterfaceProfileListFromContainer(cont *container.Container) []*PimIPv6InterfaceProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*PimIPv6InterfaceProfile, length)

	for i := 0; i < length; i++ {
		arr[i] = PimIPv6InterfaceProfileFromContainerList(cont, i)
	}

	return arr
}
