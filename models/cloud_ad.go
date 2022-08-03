package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DncloudAD        = "uni/tn-%s/ad-%s"
	RncloudAD        = "ad-%s"
	ParentDncloudAD  = "uni/tn-%s"
	CloudadClassName = "cloudAD"
)

type ActiveDirectory struct {
	BaseAttributes
	NameAliasAttribute
	ActiveDirectoryAttributes
}

type ActiveDirectoryAttributes struct {
	Annotation         string `json:",omitempty"`
	ActiveDirectory_id string `json:",omitempty"`
	Name               string `json:",omitempty"`
}

func NewActiveDirectory(cloudADRn, parentDn, description, nameAlias string, cloudADAttr ActiveDirectoryAttributes) *ActiveDirectory {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudADRn)
	return &ActiveDirectory{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudadClassName,
			Rn:                cloudADRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		ActiveDirectoryAttributes: cloudADAttr,
	}
}

func (cloudAD *ActiveDirectory) ToMap() (map[string]string, error) {
	cloudADMap, err := cloudAD.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := cloudAD.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(cloudADMap, key, value)
	}

	A(cloudADMap, "ActiveDirectory_id", cloudAD.ActiveDirectory_id)
	A(cloudADMap, "name", cloudAD.Name)
	return cloudADMap, err
}

func ActiveDirectoryFromContainerList(cont *container.Container, index int) *ActiveDirectory {
	ActiveDirectoryCont := cont.S("imdata").Index(index).S(CloudadClassName, "attributes")
	return &ActiveDirectory{
		BaseAttributes{
			DistinguishedName: G(ActiveDirectoryCont, "dn"),
			Description:       G(ActiveDirectoryCont, "descr"),
			Status:            G(ActiveDirectoryCont, "status"),
			ClassName:         CloudadClassName,
			Rn:                G(ActiveDirectoryCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(ActiveDirectoryCont, "nameAlias"),
		},
		ActiveDirectoryAttributes{
			ActiveDirectory_id: G(ActiveDirectoryCont, "ActiveDirectory_id"),
			Name:               G(ActiveDirectoryCont, "name"),
		},
	}
}

func ActiveDirectoryFromContainer(cont *container.Container) *ActiveDirectory {
	return ActiveDirectoryFromContainerList(cont, 0)
}

func ActiveDirectoryListFromContainer(cont *container.Container) []*ActiveDirectory {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*ActiveDirectory, length)

	for i := 0; i < length; i++ {
		arr[i] = ActiveDirectoryFromContainerList(cont, i)
	}

	return arr
}
