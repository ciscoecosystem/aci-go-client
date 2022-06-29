package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnfvnsMcastAddrBlk        = "uni/infra/maddrns-%s/fromaddr-[%s]-toaddr-[%s]"
	RnfvnsMcastAddrBlk        = "fromaddr-[%s]-toaddr-[%s]"
	ParentDnfvnsMcastAddrBlk  = "uni/infra/maddrns-%s"
	FvnsmcastaddrblkClassName = "fvnsMcastAddrBlk"
)

type AbstractionofIPAddressBlock struct {
	BaseAttributes
	AbstractionofIPAddressBlockAttributes
}

type AbstractionofIPAddressBlockAttributes struct {
	Annotation string `json:",omitempty"`
	From       string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
	To         string `json:",omitempty"`
}

func NewAbstractionofIPAddressBlock(fvnsMcastAddrBlkRn, parentDn, description string, fvnsMcastAddrBlkAttr AbstractionofIPAddressBlockAttributes) *AbstractionofIPAddressBlock {
	dn := fmt.Sprintf("%s/%s", parentDn, fvnsMcastAddrBlkRn)
	return &AbstractionofIPAddressBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvnsmcastaddrblkClassName,
			Rn:                fvnsMcastAddrBlkRn,
		},
		AbstractionofIPAddressBlockAttributes: fvnsMcastAddrBlkAttr,
	}
}

func (fvnsMcastAddrBlk *AbstractionofIPAddressBlock) ToMap() (map[string]string, error) {
	fvnsMcastAddrBlkMap, err := fvnsMcastAddrBlk.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(fvnsMcastAddrBlkMap, "annotation", fvnsMcastAddrBlk.Annotation)
	A(fvnsMcastAddrBlkMap, "from", fvnsMcastAddrBlk.From)
	A(fvnsMcastAddrBlkMap, "name", fvnsMcastAddrBlk.Name)
	A(fvnsMcastAddrBlkMap, "nameAlias", fvnsMcastAddrBlk.NameAlias)
	A(fvnsMcastAddrBlkMap, "to", fvnsMcastAddrBlk.To)
	return fvnsMcastAddrBlkMap, err
}

func AbstractionofIPAddressBlockFromContainerList(cont *container.Container, index int) *AbstractionofIPAddressBlock {
	AbstractionofIPAddressBlockCont := cont.S("imdata").Index(index).S(FvnsmcastaddrblkClassName, "attributes")
	return &AbstractionofIPAddressBlock{
		BaseAttributes{
			DistinguishedName: G(AbstractionofIPAddressBlockCont, "dn"),
			Description:       G(AbstractionofIPAddressBlockCont, "descr"),
			Status:            G(AbstractionofIPAddressBlockCont, "status"),
			ClassName:         FvnsmcastaddrblkClassName,
			Rn:                G(AbstractionofIPAddressBlockCont, "rn"),
		},
		AbstractionofIPAddressBlockAttributes{
			Annotation: G(AbstractionofIPAddressBlockCont, "annotation"),
			From:       G(AbstractionofIPAddressBlockCont, "from"),
			Name:       G(AbstractionofIPAddressBlockCont, "name"),
			NameAlias:  G(AbstractionofIPAddressBlockCont, "nameAlias"),
			To:         G(AbstractionofIPAddressBlockCont, "to"),
		},
	}
}

func AbstractionofIPAddressBlockFromContainer(cont *container.Container) *AbstractionofIPAddressBlock {
	return AbstractionofIPAddressBlockFromContainerList(cont, 0)
}

func AbstractionofIPAddressBlockListFromContainer(cont *container.Container) []*AbstractionofIPAddressBlock {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*AbstractionofIPAddressBlock, length)

	for i := 0; i < length; i++ {
		arr[i] = AbstractionofIPAddressBlockFromContainerList(cont, i)
	}

	return arr
}
