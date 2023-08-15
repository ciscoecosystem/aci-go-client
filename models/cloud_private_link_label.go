package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudPrivateLinkLabel        = "privatelinklabel-%s"
	CloudPrivateLinkLabelClassName = "cloudPrivateLinkLabel"
)

type PrivateLinkLabelfortheserviceEPg struct {
	BaseAttributes
	PrivateLinkLabelfortheserviceEPgAttributes
}

type PrivateLinkLabelfortheserviceEPgAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewPrivateLinkLabelfortheserviceEPg(cloudPrivateLinkLabelRn, parentDn, description string, cloudPrivateLinkLabelAttr PrivateLinkLabelfortheserviceEPgAttributes) *PrivateLinkLabelfortheserviceEPg {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudPrivateLinkLabelRn)
	return &PrivateLinkLabelfortheserviceEPg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudPrivateLinkLabelClassName,
			Rn:                cloudPrivateLinkLabelRn,
		},
		PrivateLinkLabelfortheserviceEPgAttributes: cloudPrivateLinkLabelAttr,
	}
}

func (cloudPrivateLinkLabel *PrivateLinkLabelfortheserviceEPg) ToMap() (map[string]string, error) {
	cloudPrivateLinkLabelMap, err := cloudPrivateLinkLabel.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudPrivateLinkLabelMap, "annotation", cloudPrivateLinkLabel.Annotation)
	A(cloudPrivateLinkLabelMap, "name", cloudPrivateLinkLabel.Name)
	A(cloudPrivateLinkLabelMap, "nameAlias", cloudPrivateLinkLabel.NameAlias)
	return cloudPrivateLinkLabelMap, err
}

func PrivateLinkLabelfortheserviceEPgFromContainerList(cont *container.Container, index int) *PrivateLinkLabelfortheserviceEPg {
	PrivateLinkLabelfortheserviceEPgCont := cont.S("imdata").Index(index).S(CloudPrivateLinkLabelClassName, "attributes")
	return &PrivateLinkLabelfortheserviceEPg{
		BaseAttributes{
			DistinguishedName: G(PrivateLinkLabelfortheserviceEPgCont, "dn"),
			Description:       G(PrivateLinkLabelfortheserviceEPgCont, "descr"),
			Status:            G(PrivateLinkLabelfortheserviceEPgCont, "status"),
			ClassName:         CloudPrivateLinkLabelClassName,
			Rn:                G(PrivateLinkLabelfortheserviceEPgCont, "rn"),
		},
		PrivateLinkLabelfortheserviceEPgAttributes{
			Annotation: G(PrivateLinkLabelfortheserviceEPgCont, "annotation"),
			Name:       G(PrivateLinkLabelfortheserviceEPgCont, "name"),
			NameAlias:  G(PrivateLinkLabelfortheserviceEPgCont, "nameAlias"),
		},
	}
}

func PrivateLinkLabelfortheserviceEPgFromContainer(cont *container.Container) *PrivateLinkLabelfortheserviceEPg {
	return PrivateLinkLabelfortheserviceEPgFromContainerList(cont, 0)
}

func PrivateLinkLabelfortheserviceEPgListFromContainer(cont *container.Container) []*PrivateLinkLabelfortheserviceEPg {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*PrivateLinkLabelfortheserviceEPg, length)

	for i := 0; i < length; i++ {
		arr[i] = PrivateLinkLabelfortheserviceEPgFromContainerList(cont, i)
	}

	return arr
}
