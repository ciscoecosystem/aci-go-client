package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnPimExtP        = "pimextp"
	DnPimExtP        = "uni/tn-%s/out-%s/pimextp"
	ParentDnPimExtP  = "uni/tn-%s/out-%s"
	PimExtPClassName = "pimExtP"
)

type ExternalProfile struct {
	BaseAttributes
	ExternalProfileAttributes
}

type ExternalProfileAttributes struct {
	Annotation string `json:",omitempty"`
	EnabledAf  string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewExternalProfile(pimExtPRn, parentDn, description string, pimExtPAttr ExternalProfileAttributes) *ExternalProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, pimExtPRn)
	return &ExternalProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         PimExtPClassName,
			Rn:                pimExtPRn,
		},
		ExternalProfileAttributes: pimExtPAttr,
	}
}

func (pimExtP *ExternalProfile) ToMap() (map[string]string, error) {
	pimExtPMap, err := pimExtP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(pimExtPMap, "annotation", pimExtP.Annotation)
	A(pimExtPMap, "enabledAf", pimExtP.EnabledAf)
	A(pimExtPMap, "name", pimExtP.Name)
	A(pimExtPMap, "nameAlias", pimExtP.NameAlias)
	return pimExtPMap, err
}

func ExternalProfileFromContainerList(cont *container.Container, index int) *ExternalProfile {
	ExternalProfileCont := cont.S("imdata").Index(index).S(PimExtPClassName, "attributes")
	return &ExternalProfile{
		BaseAttributes{
			DistinguishedName: G(ExternalProfileCont, "dn"),
			Description:       G(ExternalProfileCont, "descr"),
			Status:            G(ExternalProfileCont, "status"),
			ClassName:         PimExtPClassName,
			Rn:                G(ExternalProfileCont, "rn"),
		},
		ExternalProfileAttributes{
			Annotation: G(ExternalProfileCont, "annotation"),
			EnabledAf:  G(ExternalProfileCont, "enabledAf"),
			Name:       G(ExternalProfileCont, "name"),
			NameAlias:  G(ExternalProfileCont, "nameAlias"),
		},
	}
}

func ExternalProfileFromContainer(cont *container.Container) *ExternalProfile {
	return ExternalProfileFromContainerList(cont, 0)
}

func ExternalProfileListFromContainer(cont *container.Container) []*ExternalProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*ExternalProfile, length)

	for i := 0; i < length; i++ {
		arr[i] = ExternalProfileFromContainerList(cont, i)
	}

	return arr
}
