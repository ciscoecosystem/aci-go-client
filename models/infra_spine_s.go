package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const InfraspinesClassName = "infraSpineS"

type SwitchAssociation struct {
	BaseAttributes
	SwitchAssociationAttributes
}

type SwitchAssociationAttributes struct {
	Name string `json:",omitempty"`

	SwitchAssociationType string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewSwitchAssociation(infraSpineSRn, parentDn, description string, infraSpineSattr SwitchAssociationAttributes) *SwitchAssociation {
	dn := fmt.Sprintf("%s/%s", parentDn, infraSpineSRn)
	return &SwitchAssociation{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         InfraspinesClassName,
			Rn:                infraSpineSRn,
		},

		SwitchAssociationAttributes: infraSpineSattr,
	}
}

func (infraSpineS *SwitchAssociation) ToMap() (map[string]string, error) {
	infraSpineSMap, err := infraSpineS.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(infraSpineSMap, "name", infraSpineS.Name)

	A(infraSpineSMap, "type", infraSpineS.Switch_association_type)

	A(infraSpineSMap, "annotation", infraSpineS.Annotation)

	A(infraSpineSMap, "nameAlias", infraSpineS.NameAlias)

	return infraSpineSMap, err
}

func SwitchAssociationFromContainerList(cont *container.Container, index int) *SwitchAssociation {

	SwitchAssociationCont := cont.S("imdata").Index(index).S(InfraspinesClassName, "attributes")
	return &SwitchAssociation{
		BaseAttributes{
			DistinguishedName: G(SwitchAssociationCont, "dn"),
			Description:       G(SwitchAssociationCont, "descr"),
			Status:            G(SwitchAssociationCont, "status"),
			ClassName:         InfraspinesClassName,
			Rn:                G(SwitchAssociationCont, "rn"),
		},

		SwitchAssociationAttributes{

			Name: G(SwitchAssociationCont, "name"),

			Switch_association_type: G(SwitchAssociationCont, "type"),

			Annotation: G(SwitchAssociationCont, "annotation"),

			NameAlias: G(SwitchAssociationCont, "nameAlias"),
		},
	}
}

func SwitchAssociationFromContainer(cont *container.Container) *SwitchAssociation {

	return SwitchAssociationFromContainerList(cont, 0)
}

func SwitchAssociationListFromContainer(cont *container.Container) []*SwitchAssociation {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SwitchAssociation, length)

	for i := 0; i < length; i++ {

		arr[i] = SwitchAssociationFromContainerList(cont, i)
	}

	return arr
}
