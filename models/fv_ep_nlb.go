package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	RnfvEpNlb        = "epnlb"
	FvepnlbClassName = "fvEpNlb"
)

type NLBendpoint struct {
	BaseAttributes
	NameAliasAttribute
	NLBendpointAttributes
}

type NLBendpointAttributes struct {
	Annotation string `json:",omitempty"`
	Group      string `json:",omitempty"`
	Mac        string `json:",omitempty"`
	Mode       string `json:",omitempty"`
	Name       string `json:",omitempty"`
}

func NewNLBendpoint(fvEpNlbRn, parentDn, description, nameAlias string, fvEpNlbAttr NLBendpointAttributes) *NLBendpoint {
	dn := fmt.Sprintf("%s/%s", parentDn, fvEpNlbRn)
	return &NLBendpoint{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvepnlbClassName,
			Rn:                fvEpNlbRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		NLBendpointAttributes: fvEpNlbAttr,
	}
}

func (fvEpNlb *NLBendpoint) ToMap() (map[string]string, error) {
	fvEpNlbMap, err := fvEpNlb.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := fvEpNlb.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(fvEpNlbMap, key, value)
	}

	A(fvEpNlbMap, "group", fvEpNlb.Group)
	A(fvEpNlbMap, "mac", fvEpNlb.Mac)
	A(fvEpNlbMap, "mode", fvEpNlb.Mode)
	A(fvEpNlbMap, "name", fvEpNlb.Name)
	return fvEpNlbMap, err
}

func NLBendpointFromContainerList(cont *container.Container, index int) *NLBendpoint {
	NLBendpointCont := cont.S("imdata").Index(index).S(FvepnlbClassName, "attributes")
	return &NLBendpoint{
		BaseAttributes{
			DistinguishedName: G(NLBendpointCont, "dn"),
			Description:       G(NLBendpointCont, "descr"),
			Status:            G(NLBendpointCont, "status"),
			ClassName:         FvepnlbClassName,
			Rn:                G(NLBendpointCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(NLBendpointCont, "nameAlias"),
		},
		NLBendpointAttributes{
			Group: G(NLBendpointCont, "group"),
			Mac:   G(NLBendpointCont, "mac"),
			Mode:  G(NLBendpointCont, "mode"),
			Name:  G(NLBendpointCont, "name"),
		},
	}
}

func NLBendpointFromContainer(cont *container.Container) *NLBendpoint {
	return NLBendpointFromContainerList(cont, 0)
}

func NLBendpointListFromContainer(cont *container.Container) []*NLBendpoint {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*NLBendpoint, length)

	for i := 0; i < length; i++ {
		arr[i] = NLBendpointFromContainerList(cont, i)
	}

	return arr
}
