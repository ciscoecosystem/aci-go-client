package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnaaaDomainRef        = "uni/tn-%s/domain-%s"
	RnaaaDomainRef        = "domain-%s"
	ParentDnaaaDomainRef  = "uni/tn-%s"
	AaadomainrefClassName = "aaaDomainRef"
)

type TenantSecurityDomain struct {
	BaseAttributes
	NameAliasAttribute
	TenantSecurityDomainAttributes
}

type TenantSecurityDomainAttributes struct {
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
}

func NewTenantSecurityDomain(aaaDomainRefRn, parentDn, description, nameAlias string, aaaDomainRefAttr TenantSecurityDomainAttributes) *TenantSecurityDomain {
	dn := fmt.Sprintf("%s/%s", parentDn, aaaDomainRefRn)
	return &TenantSecurityDomain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         AaadomainrefClassName,
			Rn:                aaaDomainRefRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		TenantSecurityDomainAttributes: aaaDomainRefAttr,
	}
}

func (aaaDomainRef *TenantSecurityDomain) ToMap() (map[string]string, error) {
	aaaDomainRefMap, err := aaaDomainRef.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := aaaDomainRef.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(aaaDomainRefMap, key, value)
	}

	A(aaaDomainRefMap, "name", aaaDomainRef.Name)
	return aaaDomainRefMap, err
}

func TenantSecurityDomainFromContainerList(cont *container.Container, index int) *TenantSecurityDomain {
	TenantSecurityDomainCont := cont.S("imdata").Index(index).S(AaadomainrefClassName, "attributes")
	return &TenantSecurityDomain{
		BaseAttributes{
			DistinguishedName: G(TenantSecurityDomainCont, "dn"),
			Description:       G(TenantSecurityDomainCont, "descr"),
			Status:            G(TenantSecurityDomainCont, "status"),
			ClassName:         AaadomainrefClassName,
			Rn:                G(TenantSecurityDomainCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(TenantSecurityDomainCont, "nameAlias"),
		},
		TenantSecurityDomainAttributes{
			Name: G(TenantSecurityDomainCont, "name"),
		},
	}
}

func TenantSecurityDomainFromContainer(cont *container.Container) *TenantSecurityDomain {
	return TenantSecurityDomainFromContainerList(cont, 0)
}

func TenantSecurityDomainListFromContainer(cont *container.Container) []*TenantSecurityDomain {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*TenantSecurityDomain, length)

	for i := 0; i < length; i++ {
		arr[i] = TenantSecurityDomainFromContainerList(cont, i)
	}

	return arr
}
