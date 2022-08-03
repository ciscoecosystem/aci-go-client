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

type TenanttoCloudAccountAssociation struct {
	BaseAttributes
	NameAliasAttribute
	TenanttoCloudAccountAssociationAttributes
}

type TenanttoCloudAccountAssociationAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewTenanttoCloudAccountAssociation(fvRsCloudAccountRn, parentDn, nameAlias string, fvRsCloudAccountAttr TenanttoCloudAccountAssociationAttributes) *TenanttoCloudAccountAssociation {
	dn := fmt.Sprintf("%s/%s", parentDn, fvRsCloudAccountRn)
	return &TenanttoCloudAccountAssociation{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         FvrscloudaccountClassName,
			Rn:                fvRsCloudAccountRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		TenanttoCloudAccountAssociationAttributes: fvRsCloudAccountAttr,
	}
}

func (fvRsCloudAccount *TenanttoCloudAccountAssociation) ToMap() (map[string]string, error) {
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

func TenanttoCloudAccountAssociationFromContainerList(cont *container.Container, index int) *TenanttoCloudAccountAssociation {
	TenanttoCloudAccountAssociationCont := cont.S("imdata").Index(index).S(FvrscloudaccountClassName, "attributes")
	return &TenanttoCloudAccountAssociation{
		BaseAttributes{
			DistinguishedName: G(TenanttoCloudAccountAssociationCont, "dn"),
			Status:            G(TenanttoCloudAccountAssociationCont, "status"),
			ClassName:         FvrscloudaccountClassName,
			Rn:                G(TenanttoCloudAccountAssociationCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(TenanttoCloudAccountAssociationCont, "nameAlias"),
		},
		TenanttoCloudAccountAssociationAttributes{
			Annotation: G(TenanttoCloudAccountAssociationCont, "annotation"),
			TDn:        G(TenanttoCloudAccountAssociationCont, "tDn"),
		},
	}
}

func TenanttoCloudAccountAssociationFromContainer(cont *container.Container) *TenanttoCloudAccountAssociation {
	return TenanttoCloudAccountAssociationFromContainerList(cont, 0)
}

func TenanttoCloudAccountAssociationListFromContainer(cont *container.Container) []*TenanttoCloudAccountAssociation {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*TenanttoCloudAccountAssociation, length)

	for i := 0; i < length; i++ {
		arr[i] = TenanttoCloudAccountAssociationFromContainerList(cont, i)
	}

	return arr
}
