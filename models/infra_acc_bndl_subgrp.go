package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DninfraAccBndlSubgrp        = "uni/infra/funcprof/accbundle-%s/accsubbndl-%s"
	RninfraAccBndlSubgrp        = "accsubbndl-%s"
	ParentDninfraAccBndlSubgrp  = "uni/infra/funcprof/accbundle-%s"
	InfraaccbndlsubgrpClassName = "infraAccBndlSubgrp"
)

type OverridePolicyGroup struct {
	BaseAttributes
	OverridePolicyGroupAttributes
}

type OverridePolicyGroupAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewOverridePolicyGroup(infraAccBndlSubgrpRn, parentDn, description string, infraAccBndlSubgrpAttr OverridePolicyGroupAttributes) *OverridePolicyGroup {
	dn := fmt.Sprintf("%s/%s", parentDn, infraAccBndlSubgrpRn)
	return &OverridePolicyGroup{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         InfraaccbndlsubgrpClassName,
			Rn:                infraAccBndlSubgrpRn,
		},
		OverridePolicyGroupAttributes: infraAccBndlSubgrpAttr,
	}
}

func (infraAccBndlSubgrp *OverridePolicyGroup) ToMap() (map[string]string, error) {
	infraAccBndlSubgrpMap, err := infraAccBndlSubgrp.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(infraAccBndlSubgrpMap, "annotation", infraAccBndlSubgrp.Annotation)
	A(infraAccBndlSubgrpMap, "name", infraAccBndlSubgrp.Name)
	A(infraAccBndlSubgrpMap, "nameAlias", infraAccBndlSubgrp.NameAlias)
	return infraAccBndlSubgrpMap, err
}

func OverridePolicyGroupFromContainerList(cont *container.Container, index int) *OverridePolicyGroup {
	OverridePolicyGroupCont := cont.S("imdata").Index(index).S(InfraaccbndlsubgrpClassName, "attributes")
	return &OverridePolicyGroup{
		BaseAttributes{
			DistinguishedName: G(OverridePolicyGroupCont, "dn"),
			Description:       G(OverridePolicyGroupCont, "descr"),
			Status:            G(OverridePolicyGroupCont, "status"),
			ClassName:         InfraaccbndlsubgrpClassName,
			Rn:                G(OverridePolicyGroupCont, "rn"),
		},
		OverridePolicyGroupAttributes{
			Annotation: G(OverridePolicyGroupCont, "annotation"),
			Name:       G(OverridePolicyGroupCont, "name"),
			NameAlias:  G(OverridePolicyGroupCont, "nameAlias"),
		},
	}
}

func OverridePolicyGroupFromContainer(cont *container.Container) *OverridePolicyGroup {
	return OverridePolicyGroupFromContainerList(cont, 0)
}

func OverridePolicyGroupListFromContainer(cont *container.Container) []*OverridePolicyGroup {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*OverridePolicyGroup, length)

	for i := 0; i < length; i++ {
		arr[i] = OverridePolicyGroupFromContainerList(cont, i)
	}

	return arr
}
