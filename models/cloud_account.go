package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DncloudAccount        = "uni/tn-%s/act-[%s]-vendor-%s"
	RncloudAccount        = "act-[%s]-vendor-%s"
	ParentDncloudAccount  = "uni/tn-%s"
	CloudaccountClassName = "cloudAccount"
)

type Account struct {
	BaseAttributes
	NameAliasAttribute
	AccountAttributes
}

type AccountAttributes struct {
	Annotation string `json:",omitempty"`
	AccessType string `json:",omitempty"`
	Account_id string `json:",omitempty"`
	Name       string `json:",omitempty"`
	Vendor     string `json:",omitempty"`
}

func NewAccount(cloudAccountRn, parentDn, nameAlias string, cloudAccountAttr AccountAttributes) *Account {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudAccountRn)
	return &Account{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         CloudaccountClassName,
			Rn:                cloudAccountRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		AccountAttributes: cloudAccountAttr,
	}
}

func (cloudAccount *Account) ToMap() (map[string]string, error) {
	cloudAccountMap, err := cloudAccount.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := cloudAccount.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(cloudAccountMap, key, value)
	}

	A(cloudAccountMap, "accessType", cloudAccount.AccessType)
	A(cloudAccountMap, "id", cloudAccount.Account_id)
	A(cloudAccountMap, "name", cloudAccount.Name)
	A(cloudAccountMap, "vendor", cloudAccount.Vendor)
	return cloudAccountMap, err
}

func AccountFromContainerList(cont *container.Container, index int) *Account {
	AccountCont := cont.S("imdata").Index(index).S(CloudaccountClassName, "attributes")
	return &Account{
		BaseAttributes{
			DistinguishedName: G(AccountCont, "dn"),
			Status:            G(AccountCont, "status"),
			ClassName:         CloudaccountClassName,
			Rn:                G(AccountCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(AccountCont, "nameAlias"),
		},
		AccountAttributes{
			AccessType: G(AccountCont, "accessType"),
			Account_id: G(AccountCont, "id"),
			Name:       G(AccountCont, "name"),
			Vendor:     G(AccountCont, "vendor"),
		},
	}
}

func AccountFromContainer(cont *container.Container) *Account {
	return AccountFromContainerList(cont, 0)
}

func AccountListFromContainer(cont *container.Container) []*Account {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*Account, length)

	for i := 0; i < length; i++ {
		arr[i] = AccountFromContainerList(cont, i)
	}

	return arr
}
