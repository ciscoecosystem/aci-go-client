package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnfvRsCloudAccount        = "uni/tn-%s/rsCloudAccount"
	RnfvRsCloudAccount        = "rsCloudAccount"
	ParentDnfvRsCloudAccount  = "uni/tn-%s"
	FvrscloudaccountClassName = "fvRsCloudAccount"
)

type Tenanttoaccountassociation struct {
	BaseAttributes
	NameAliasAttribute
	TenanttoaccountassociationAttributes
}

type TenanttoaccountassociationAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewTenanttoaccountassociation(fvRsCloudAccountRn, parentDn, description, nameAlias string, fvRsCloudAccountAttr TenanttoaccountassociationAttributes) *Tenanttoaccountassociation {
	dn := fmt.Sprintf("%s/%s", parentDn, fvRsCloudAccountRn)
	return &Tenanttoaccountassociation{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvrscloudaccountClassName,
			Rn:                fvRsCloudAccountRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		TenanttoaccountassociationAttributes: fvRsCloudAccountAttr,
	}
}

func (fvRsCloudAccount *Tenanttoaccountassociation) ToMap() (map[string]string, error) {
	fvRsCloudAccountMap, err := fvRsCloudAccount.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := fvRsCloudAccount.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(fvRsCloudAccountMap, key, value)
	}

	A(fvRsCloudAccountMap, "annotation", fvRsCloudAccount.Annotation)
	A(fvRsCloudAccountMap, "tDn", fvRsCloudAccount.TDn)
	return fvRsCloudAccountMap, err
}

func TenanttoaccountassociationFromContainerList(cont *container.Container, index int) *Tenanttoaccountassociation {
	TenanttoaccountassociationCont := cont.S("imdata").Index(index).S(FvrscloudaccountClassName, "attributes")
	return &Tenanttoaccountassociation{
		BaseAttributes{
			DistinguishedName: G(TenanttoaccountassociationCont, "dn"),
			Description:       G(TenanttoaccountassociationCont, "descr"),
			Status:            G(TenanttoaccountassociationCont, "status"),
			ClassName:         FvrscloudaccountClassName,
			Rn:                G(TenanttoaccountassociationCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(TenanttoaccountassociationCont, "nameAlias"),
		},
		TenanttoaccountassociationAttributes{
			Annotation: G(TenanttoaccountassociationCont, "annotation"),
			TDn:        G(TenanttoaccountassociationCont, "tDn"),
		},
	}
}

func TenanttoaccountassociationFromContainer(cont *container.Container) *Tenanttoaccountassociation {
	return TenanttoaccountassociationFromContainerList(cont, 0)
}

func TenanttoaccountassociationListFromContainer(cont *container.Container) []*Tenanttoaccountassociation {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*Tenanttoaccountassociation, length)

	for i := 0; i < length; i++ {
		arr[i] = TenanttoaccountassociationFromContainerList(cont, i)
	}

	return arr
}
