package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const AaadomainrefClassName = "aaaDomainRef"

type TenantSecurityDomain struct {
	BaseAttributes
	TenantSecurityDomainAttributes
}

type TenantSecurityDomainAttributes struct {
	Name string `json:",omitempty"`

	Tenant string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewTenantSecurityDomain(aaaDomainRefRn, parentDn, description string, aaaDomainRefattr TenantSecurityDomainAttributes) *TenantSecurityDomain {
	dn := fmt.Sprintf("%s/%s", parentDn, aaaDomainRefRn)
	return &TenantSecurityDomain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         AaadomainrefClassName,
			Rn:                aaaDomainRefRn,
		},

		TenantSecurityDomainAttributes: aaaDomainRefattr,
	}
}

func (aaaDomainRef *TenantSecurityDomain) ToMap() (map[string]string, error) {
	aaaDomainRefMap, err := aaaDomainRef.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(aaaDomainRefMap, "name", aaaDomainRef.Name)

	A(aaaDomainRefMap, "tenant", aaaDomainRef.Tenant)

	A(aaaDomainRefMap, "annotation", aaaDomainRef.Annotation)

	A(aaaDomainRefMap, "nameAlias", aaaDomainRef.NameAlias)

	return aaaDomainRefMap, err
}

func TenantSecurityDomainFromContainerList(cont *container.Container, index int) *TenantSecurityDomain {

	TenantSecurityDomainCont := cont.S("imdata").Index(index).S(AaadomainrefClassName, "attributes")
	return &TenantSecurityDomain{
		BaseAttributes{
			DistinguishedName: G(TenantSecurityDomainCont, "dn"),
			Status:            G(TenantSecurityDomainCont, "status"),
			ClassName:         AaadomainrefClassName,
			Rn:                G(TenantSecurityDomainCont, "rn"),
		},

		TenantSecurityDomainAttributes{

			Name: G(TenantSecurityDomainCont, "name"),

			Tenant: G(TenantSecurityDomainCont, "tenant"),

			Annotation: G(TenantSecurityDomainCont, "annotation"),

			NameAlias: G(TenantSecurityDomainCont, "nameAlias"),
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
