package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DncloudtemplateVpnNetwork        = "uni/tn-%s/infranetwork-%s/extnetwork-%s/vpnnetwork-%s"
	RncloudtemplateVpnNetwork        = "vpnnetwork-%s"
	ParentDncloudtemplateVpnNetwork  = "uni/tn-%s/infranetwork-%s/extnetwork-%s"
	CloudtemplatevpnnetworkClassName = "cloudtemplateVpnNetwork"
)

type TemplateforVPNNetwork struct {
	BaseAttributes
	NameAliasAttribute
	TemplateforVPNNetworkAttributes
}

type TemplateforVPNNetworkAttributes struct {
	Annotation     string `json:",omitempty"`
	Name           string `json:",omitempty"`
	RemoteSiteId   string `json:",omitempty"`
	RemoteSiteName string `json:",omitempty"`
}

func NewTemplateforVPNNetwork(cloudtemplateVpnNetworkRn, parentDn, nameAlias string, cloudtemplateVpnNetworkAttr TemplateforVPNNetworkAttributes) *TemplateforVPNNetwork {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudtemplateVpnNetworkRn)
	return &TemplateforVPNNetwork{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         CloudtemplatevpnnetworkClassName,
			Rn:                cloudtemplateVpnNetworkRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		TemplateforVPNNetworkAttributes: cloudtemplateVpnNetworkAttr,
	}
}

func (cloudtemplateVpnNetwork *TemplateforVPNNetwork) ToMap() (map[string]string, error) {
	cloudtemplateVpnNetworkMap, err := cloudtemplateVpnNetwork.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := cloudtemplateVpnNetwork.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(cloudtemplateVpnNetworkMap, key, value)
	}

	A(cloudtemplateVpnNetworkMap, "annotation", cloudtemplateVpnNetwork.Annotation)
	A(cloudtemplateVpnNetworkMap, "name", cloudtemplateVpnNetwork.Name)
	A(cloudtemplateVpnNetworkMap, "remoteSiteId", cloudtemplateVpnNetwork.RemoteSiteId)
	A(cloudtemplateVpnNetworkMap, "remoteSiteName", cloudtemplateVpnNetwork.RemoteSiteName)
	return cloudtemplateVpnNetworkMap, err
}

func TemplateforVPNNetworkFromContainerList(cont *container.Container, index int) *TemplateforVPNNetwork {
	TemplateforVPNNetworkCont := cont.S("imdata").Index(index).S(CloudtemplatevpnnetworkClassName, "attributes")
	return &TemplateforVPNNetwork{
		BaseAttributes{
			DistinguishedName: G(TemplateforVPNNetworkCont, "dn"),
			Status:            G(TemplateforVPNNetworkCont, "status"),
			ClassName:         CloudtemplatevpnnetworkClassName,
			Rn:                G(TemplateforVPNNetworkCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(TemplateforVPNNetworkCont, "nameAlias"),
		},
		TemplateforVPNNetworkAttributes{
			Annotation:     G(TemplateforVPNNetworkCont, "annotation"),
			Name:           G(TemplateforVPNNetworkCont, "name"),
			RemoteSiteId:   G(TemplateforVPNNetworkCont, "remoteSiteId"),
			RemoteSiteName: G(TemplateforVPNNetworkCont, "remoteSiteName"),
		},
	}
}

func TemplateforVPNNetworkFromContainer(cont *container.Container) *TemplateforVPNNetwork {
	return TemplateforVPNNetworkFromContainerList(cont, 0)
}

func TemplateforVPNNetworkListFromContainer(cont *container.Container) []*TemplateforVPNNetwork {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*TemplateforVPNNetwork, length)

	for i := 0; i < length; i++ {
		arr[i] = TemplateforVPNNetworkFromContainerList(cont, i)
	}

	return arr
}
