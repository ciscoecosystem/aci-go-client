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

type PrivateLinkLabel struct {
	BaseAttributes
	PrivateLinkLabelAttributes
}

type PrivateLinkLabelAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewPrivateLinkLabel(cloudPrivateLinkLabelRn, parentDn, description string, cloudPrivateLinkLabelAttr PrivateLinkLabelAttributes) *PrivateLinkLabel {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudPrivateLinkLabelRn)
	return &PrivateLinkLabel{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudPrivateLinkLabelClassName,
			Rn:                cloudPrivateLinkLabelRn,
		},
		PrivateLinkLabelAttributes: cloudPrivateLinkLabelAttr,
	}
}

func (cloudPrivateLinkLabel *PrivateLinkLabel) ToMap() (map[string]string, error) {
	cloudPrivateLinkLabelMap, err := cloudPrivateLinkLabel.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudPrivateLinkLabelMap, "annotation", cloudPrivateLinkLabel.Annotation)
	A(cloudPrivateLinkLabelMap, "name", cloudPrivateLinkLabel.Name)
	A(cloudPrivateLinkLabelMap, "nameAlias", cloudPrivateLinkLabel.NameAlias)
	return cloudPrivateLinkLabelMap, err
}

func PrivateLinkLabelFromContainerList(cont *container.Container, index int) *PrivateLinkLabel {
	PrivateLinkLabelCont := cont.S("imdata").Index(index).S(CloudPrivateLinkLabelClassName, "attributes")
	return &PrivateLinkLabel{
		BaseAttributes{
			DistinguishedName: G(PrivateLinkLabelCont, "dn"),
			Description:       G(PrivateLinkLabelCont, "descr"),
			Status:            G(PrivateLinkLabelCont, "status"),
			ClassName:         CloudPrivateLinkLabelClassName,
			Rn:                G(PrivateLinkLabelCont, "rn"),
		},
		PrivateLinkLabelAttributes{
			Annotation: G(PrivateLinkLabelCont, "annotation"),
			Name:       G(PrivateLinkLabelCont, "name"),
			NameAlias:  G(PrivateLinkLabelCont, "nameAlias"),
		},
	}
}

func PrivateLinkLabelFromContainer(cont *container.Container) *PrivateLinkLabel {
	return PrivateLinkLabelFromContainerList(cont, 0)
}

func PrivateLinkLabelListFromContainer(cont *container.Container) []*PrivateLinkLabel {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*PrivateLinkLabel, length)

	for i := 0; i < length; i++ {
		arr[i] = PrivateLinkLabelFromContainerList(cont, i)
	}

	return arr
}
