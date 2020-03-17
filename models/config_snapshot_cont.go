package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const ConfigsnapshotcontClassName = "configSnapshotCont"

type ConfigBackupSnapshotContainer struct {
	BaseAttributes
	ConfigBackupSnapshotContainerAttributes
}

type ConfigBackupSnapshotContainerAttributes struct {
	PDn string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewConfigBackupSnapshotContainer(configSnapshotContRn, parentDn, description string, configSnapshotContattr ConfigBackupSnapshotContainerAttributes) *ConfigBackupSnapshotContainer {
	dn := fmt.Sprintf("%s/%s", parentDn, configSnapshotContRn)
	return &ConfigBackupSnapshotContainer{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         ConfigsnapshotcontClassName,
			Rn:                configSnapshotContRn,
		},

		ConfigBackupSnapshotContainerAttributes: configSnapshotContattr,
	}
}

func (configSnapshotCont *ConfigBackupSnapshotContainer) ToMap() (map[string]string, error) {
	configSnapshotContMap, err := configSnapshotCont.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(configSnapshotContMap, "pDn", configSnapshotCont.PDn)

	A(configSnapshotContMap, "nameAlias", configSnapshotCont.NameAlias)

	return configSnapshotContMap, err
}

func ConfigBackupSnapshotContainerFromContainerList(cont *container.Container, index int) *ConfigBackupSnapshotContainer {

	ConfigBackupSnapshotContainerCont := cont.S("imdata").Index(index).S(ConfigsnapshotcontClassName, "attributes")
	return &ConfigBackupSnapshotContainer{
		BaseAttributes{
			DistinguishedName: G(ConfigBackupSnapshotContainerCont, "dn"),
			Description:       G(ConfigBackupSnapshotContainerCont, "descr"),
			Status:            G(ConfigBackupSnapshotContainerCont, "status"),
			ClassName:         ConfigsnapshotcontClassName,
			Rn:                G(ConfigBackupSnapshotContainerCont, "rn"),
		},

		ConfigBackupSnapshotContainerAttributes{

			PDn: G(ConfigBackupSnapshotContainerCont, "pDn"),

			NameAlias: G(ConfigBackupSnapshotContainerCont, "nameAlias"),
		},
	}
}

func ConfigBackupSnapshotContainerFromContainer(cont *container.Container) *ConfigBackupSnapshotContainer {

	return ConfigBackupSnapshotContainerFromContainerList(cont, 0)
}

func ConfigBackupSnapshotContainerListFromContainer(cont *container.Container) []*ConfigBackupSnapshotContainer {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*ConfigBackupSnapshotContainer, length)

	for i := 0; i < length; i++ {

		arr[i] = ConfigBackupSnapshotContainerFromContainerList(cont, i)
	}

	return arr
}
