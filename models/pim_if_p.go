package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnPimIfP        = "pimifp"
	DnPimIfP        = "uni/tn-%s/out-%s/lnodep-%s/lifp-%s/pimifp"
	ParentDnPimIfP  = "uni/tn-%s/out-%s/lnodep-%s/lifp-%s"
	PimIfPClassName = "pimIfP"
)

type PimInterfaceProfile struct {
	BaseAttributes
	PimInterfaceProfileAttributes
}

type PimInterfaceProfileAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewPimInterfaceProfile(pimIfPRn, parentDn, description string, pimIfPAttr PimInterfaceProfileAttributes) *PimInterfaceProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, pimIfPRn)
	return &PimInterfaceProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         PimIfPClassName,
			Rn:                pimIfPRn,
		},
		PimInterfaceProfileAttributes: pimIfPAttr,
	}
}

func (pimIfP *PimInterfaceProfile) ToMap() (map[string]string, error) {
	pimIfPMap, err := pimIfP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(pimIfPMap, "annotation", pimIfP.Annotation)
	A(pimIfPMap, "name", pimIfP.Name)
	A(pimIfPMap, "nameAlias", pimIfP.NameAlias)
	return pimIfPMap, err
}

func PimInterfaceProfileFromContainerList(cont *container.Container, index int) *PimInterfaceProfile {
	InterfaceProfileCont := cont.S("imdata").Index(index).S(PimIfPClassName, "attributes")
	return &PimInterfaceProfile{
		BaseAttributes{
			DistinguishedName: G(InterfaceProfileCont, "dn"),
			Description:       G(InterfaceProfileCont, "descr"),
			Status:            G(InterfaceProfileCont, "status"),
			ClassName:         PimIfPClassName,
			Rn:                G(InterfaceProfileCont, "rn"),
		},
		PimInterfaceProfileAttributes{
			Annotation: G(InterfaceProfileCont, "annotation"),
			Name:       G(InterfaceProfileCont, "name"),
			NameAlias:  G(InterfaceProfileCont, "nameAlias"),
		},
	}
}

func PimInterfaceProfileFromContainer(cont *container.Container) *PimInterfaceProfile {
	return PimInterfaceProfileFromContainerList(cont, 0)
}

func PimInterfaceProfileListFromContainer(cont *container.Container) []*PimInterfaceProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*PimInterfaceProfile, length)

	for i := 0; i < length; i++ {
		arr[i] = PimInterfaceProfileFromContainerList(cont, i)
	}

	return arr
}
