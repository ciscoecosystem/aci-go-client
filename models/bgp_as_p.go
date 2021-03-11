package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const BgpaspClassName = "bgpAsP"

type AutonomousSystemProfile struct {
	BaseAttributes
	AutonomousSystemProfileAttributes
}

type AutonomousSystemProfileAttributes struct {
	Annotation string `json:",omitempty"`

	Asn string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewAutonomousSystemProfile(bgpAsPRn, parentDn, description string, bgpAsPattr AutonomousSystemProfileAttributes) *AutonomousSystemProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, bgpAsPRn)
	return &AutonomousSystemProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         BgpaspClassName,
			Rn:                bgpAsPRn,
		},

		AutonomousSystemProfileAttributes: bgpAsPattr,
	}
}

func (bgpAsP *AutonomousSystemProfile) ToMap() (map[string]string, error) {
	bgpAsPMap, err := bgpAsP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(bgpAsPMap, "annotation", bgpAsP.Annotation)

	A(bgpAsPMap, "asn", bgpAsP.Asn)

	A(bgpAsPMap, "nameAlias", bgpAsP.NameAlias)

	return bgpAsPMap, err
}

func AutonomousSystemProfileFromContainerList(cont *container.Container, index int) *AutonomousSystemProfile {

	AutonomousSystemProfileCont := cont.S("imdata").Index(index).S(BgpaspClassName, "attributes")
	return &AutonomousSystemProfile{
		BaseAttributes{
			DistinguishedName: G(AutonomousSystemProfileCont, "dn"),
			Description:       G(AutonomousSystemProfileCont, "descr"),
			Status:            G(AutonomousSystemProfileCont, "status"),
			ClassName:         BgpaspClassName,
			Rn:                G(AutonomousSystemProfileCont, "rn"),
		},

		AutonomousSystemProfileAttributes{

			Annotation: G(AutonomousSystemProfileCont, "annotation"),

			Asn: G(AutonomousSystemProfileCont, "asn"),

			NameAlias: G(AutonomousSystemProfileCont, "nameAlias"),
		},
	}
}

func AutonomousSystemProfileFromContainer(cont *container.Container) *AutonomousSystemProfile {

	return AutonomousSystemProfileFromContainerList(cont, 0)
}

func AutonomousSystemProfileListFromContainer(cont *container.Container) []*AutonomousSystemProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*AutonomousSystemProfile, length)

	for i := 0; i < length; i++ {

		arr[i] = AutonomousSystemProfileFromContainerList(cont, i)
	}

	return arr
}
