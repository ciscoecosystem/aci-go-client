package models

import (
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VzrsifClassName = "vzRsIf"

type ContractExp struct {
	BaseAttributes
	ContractExpAttributes
}

type ContractExpAttributes struct {
	Annotation string `json:",omitempty"`

	Prio string `json:",omitempty"`

	TDn string `json:",omitempty"`
}

func NewContractExp(vzRsIfRn, parentDn, description string, vzRsIfattr ContractExpAttributes) *ContractExp {
	return &ContractExp{
		BaseAttributes: BaseAttributes{
			Description: description,
			Status:      "created, modified",
			ClassName:   VzrsifClassName,
			Rn:          vzRsIfRn,
		},

		ContractExpAttributes: vzRsIfattr,
	}
}

func (vzRsIf *ContractExp) ToMap() (map[string]string, error) {
	vzRsIfMap, err := vzRsIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vzRsIfMap, "annotation", vzRsIf.Annotation)

	A(vzRsIfMap, "prio", vzRsIf.Prio)

	A(vzRsIfMap, "tDn", vzRsIf.TDn)

	return vzRsIfMap, err
}

func ContractExpFromContainerList(cont *container.Container, index int) *ContractExp {

	ContractCont := cont.S("imdata").Index(index).S(VzrsifClassName, "attributes")
	return &ContractExp{
		BaseAttributes{
			DistinguishedName: G(ContractCont, "dn"),
			Description:       G(ContractCont, "descr"),
			Status:            G(ContractCont, "status"),
			ClassName:         VzrsifClassName,
			Rn:                G(ContractCont, "rn"),
		},

		ContractExpAttributes{

			Annotation: G(ContractCont, "annotation"),

			Prio: G(ContractCont, "prio"),

			TDn: G(ContractCont, "tDn"),
		},
	}
}

func ContractExpFromContainer(cont *container.Container) *ContractExp {

	return ContractExpFromContainerList(cont, 0)
}

func ContractExpListFromContainer(cont *container.Container) []*ContractExp {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*ContractExp, length)

	for i := 0; i < length; i++ {

		arr[i] = ContractExpFromContainerList(cont, i)
	}

	return arr
}
