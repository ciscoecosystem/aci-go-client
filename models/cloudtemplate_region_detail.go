package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudtemplateRegionDetail        = "regiondetail"
	CloudtemplateRegionDetailClassName = "cloudtemplateRegionDetail"
)

type Additionalconfigforregion struct {
	BaseAttributes
	AdditionalconfigforregionAttributes
}

type AdditionalconfigforregionAttributes struct {
	Annotation           string `json:",omitempty"`
	HubNetworkingEnabled string `json:",omitempty"`
}

func NewAdditionalconfigforregion(cloudtemplateRegionDetailRn, parentDn string, cloudtemplateRegionDetailAttr AdditionalconfigforregionAttributes) *Additionalconfigforregion {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudtemplateRegionDetailRn)
	return &Additionalconfigforregion{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "modified",
			ClassName:         CloudtemplateRegionDetailClassName,
			Rn:                cloudtemplateRegionDetailRn,
		},
		AdditionalconfigforregionAttributes: cloudtemplateRegionDetailAttr,
	}
}

func (cloudtemplateRegionDetail *Additionalconfigforregion) ToMap() (map[string]string, error) {
	cloudtemplateRegionDetailMap, err := cloudtemplateRegionDetail.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudtemplateRegionDetailMap, "annotation", cloudtemplateRegionDetail.Annotation)
	A(cloudtemplateRegionDetailMap, "hubNetworkingEnabled", cloudtemplateRegionDetail.HubNetworkingEnabled)
	return cloudtemplateRegionDetailMap, err
}

func AdditionalconfigforregionFromContainerList(cont *container.Container, index int) *Additionalconfigforregion {
	AdditionalconfigforregionCont := cont.S("imdata").Index(index).S(CloudtemplateRegionDetailClassName, "attributes")
	return &Additionalconfigforregion{
		BaseAttributes{
			DistinguishedName: G(AdditionalconfigforregionCont, "dn"),
			Status:            G(AdditionalconfigforregionCont, "status"),
			ClassName:         CloudtemplateRegionDetailClassName,
			Rn:                G(AdditionalconfigforregionCont, "rn"),
		},
		AdditionalconfigforregionAttributes{
			Annotation:           G(AdditionalconfigforregionCont, "annotation"),
			HubNetworkingEnabled: G(AdditionalconfigforregionCont, "hubNetworkingEnabled"),
		},
	}
}

func AdditionalconfigforregionFromContainer(cont *container.Container) *Additionalconfigforregion {
	return AdditionalconfigforregionFromContainerList(cont, 0)
}

func AdditionalconfigforregionListFromContainer(cont *container.Container) []*Additionalconfigforregion {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*Additionalconfigforregion, length)

	for i := 0; i < length; i++ {
		arr[i] = AdditionalconfigforregionFromContainerList(cont, i)
	}

	return arr
}
