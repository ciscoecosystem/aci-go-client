package models

import (
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VzcpifClassName = "vzCPIf"

type ImportedContract struct {
	BaseAttributes
	ImportedContractAttributes
}

type ImportedContractAttributes struct {
	Name string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewImportedContract(vzCPIfRn, parentDn, description string, vzCPIfattr ImportedContractAttributes) *ImportedContract {
	return &ImportedContract{
		BaseAttributes: BaseAttributes{
			Description: description,
			Status:      "created, modified",
			ClassName:   VzcpifClassName,
			Rn:          vzCPIfRn,
		},

		ImportedContractAttributes: vzCPIfattr,
	}
}

func (vzCPIf *ImportedContract) ToMap() (map[string]string, error) {
	vzCPIfMap, err := vzCPIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vzCPIfMap, "name", vzCPIf.Name)

	A(vzCPIfMap, "annotation", vzCPIf.Annotation)

	A(vzCPIfMap, "nameAlias", vzCPIf.NameAlias)

	return vzCPIfMap, err
}

func ImportedContractFromContainerList(cont *container.Container, index int) *ImportedContract {

	ImportedContractCont := cont.S("imdata").Index(index).S(VzcpifClassName, "attributes")
	return &ImportedContract{
		BaseAttributes{
			DistinguishedName: G(ImportedContractCont, "dn"),
			Description:       G(ImportedContractCont, "descr"),
			Status:            G(ImportedContractCont, "status"),
			ClassName:         VzcpifClassName,
			Rn:                G(ImportedContractCont, "rn"),
		},

		ImportedContractAttributes{

			Name: G(ImportedContractCont, "name"),

			Annotation: G(ImportedContractCont, "annotation"),

			NameAlias: G(ImportedContractCont, "nameAlias"),
		},
	}
}

func ImportedContractFromContainer(cont *container.Container) *ImportedContract {

	return ImportedContractFromContainerList(cont, 0)
}

func ImportedContractListFromContainer(cont *container.Container) []*ImportedContract {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*ImportedContract, length)

	for i := 0; i < length; i++ {

		arr[i] = ImportedContractFromContainerList(cont, i)
	}

	return arr
}
