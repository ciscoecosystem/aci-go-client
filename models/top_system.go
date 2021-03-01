package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const TopsystemClassName = "topSystem"

type System struct {
	BaseAttributes
	SystemAttributes
}

type SystemAttributes struct {
	Address string `json:",omitempty"`

	EtepAddr string `json:",omitempty"`

	System_id string `json:",omitempty"`

	NameAlias string `json:",omitempty"`

	NodeType string `json:",omitempty"`

	RemoteNetworkId string `json:",omitempty"`

	RemoteNode string `json:",omitempty"`

	RldirectMode string `json:",omitempty"`

	Role string `json:",omitempty"`

	ServerType string `json:",omitempty"`
}

func NewSystem(topSystemRn, parentDn, description string, topSystemattr SystemAttributes) *System {
	dn := fmt.Sprintf("%s/%s", parentDn, topSystemRn)
	return &System{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         TopsystemClassName,
			Rn:                topSystemRn,
		},

		SystemAttributes: topSystemattr,
	}
}

func (topSystem *System) ToMap() (map[string]string, error) {
	topSystemMap, err := topSystem.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(topSystemMap, "address", topSystem.Address)

	A(topSystemMap, "etepAddr", topSystem.EtepAddr)

	A(topSystemMap, "id", topSystem.System_id)

	A(topSystemMap, "nameAlias", topSystem.NameAlias)

	A(topSystemMap, "nodeType", topSystem.NodeType)

	A(topSystemMap, "remoteNetworkId", topSystem.RemoteNetworkId)

	A(topSystemMap, "remoteNode", topSystem.RemoteNode)

	A(topSystemMap, "rldirectMode", topSystem.RldirectMode)

	A(topSystemMap, "role", topSystem.Role)

	A(topSystemMap, "serverType", topSystem.ServerType)

	return topSystemMap, err
}

func SystemFromContainerList(cont *container.Container, index int) *System {

	SystemCont := cont.S("imdata").Index(index).S(TopsystemClassName, "attributes")
	return &System{
		BaseAttributes{
			DistinguishedName: G(SystemCont, "dn"),
			Description:       G(SystemCont, "descr"),
			Status:            G(SystemCont, "status"),
			ClassName:         TopsystemClassName,
			Rn:                G(SystemCont, "rn"),
		},

		SystemAttributes{

			Address: G(SystemCont, "address"),

			EtepAddr: G(SystemCont, "etepAddr"),

			System_id: G(SystemCont, "id"),

			NameAlias: G(SystemCont, "nameAlias"),

			NodeType: G(SystemCont, "nodeType"),

			RemoteNetworkId: G(SystemCont, "remoteNetworkId"),

			RemoteNode: G(SystemCont, "remoteNode"),

			RldirectMode: G(SystemCont, "rldirectMode"),

			Role: G(SystemCont, "role"),

			ServerType: G(SystemCont, "serverType"),
		},
	}
}

func SystemFromContainer(cont *container.Container) *System {

	return SystemFromContainerList(cont, 0)
}

func SystemListFromContainer(cont *container.Container) []*System {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*System, length)

	for i := 0; i < length; i++ {

		arr[i] = SystemFromContainerList(cont, i)
	}

	return arr
}
