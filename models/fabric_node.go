package models

import (
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FabricNodeClassName = "fabricNode"

type FabricNode struct {
	BaseAttributes
	FabricNodeAttributes
}

type FabricNodeAttributes struct {
	AdSt             string `json:",omitempty"`
	Address          string `json:",omitempty"`
	Annotation       string `json:",omitempty"`
	ApicType         string `json:",omitempty"`
	DelayedHeartbeat string `json:",omitempty"`
	ExtMngdBy        string `json:",omitempty"`
	FabricSt         string `json:",omitempty"`
	Id               string `json:",omitempty"`
	LastStateModTs   string `json:",omitempty"`
	ModTs            string `json:",omitempty"`
	Model            string `json:",omitempty"`
	MonPolDn         string `json:",omitempty"`
	Name             string `json:",omitempty"`
	NameAlias        string `json:",omitempty"`
	NodeType         string `json:",omitempty"`
	Role             string `json:",omitempty"`
	Serial           string `json:",omitempty"`
	Uid              string `json:",omitempty"`
	Userdom          string `json:",omitempty"`
	Vendor           string `json:",omitempty"`
	Version          string `json:",omitempty"`
}

/*
 * No NewFabricNode as this is a non-configurable MO
func NewFabricNode(fabricNodeRn, parentDn, description string, fabricNodeattr FabricNodeAttributes) *FabricNode {
    dn := fmt.Sprintf("%s/%s", parentDn, fabricNodeRn)
    return &FabricNode{
        BaseAttributes: BaseAttributes{
            DistinguishedName: dn,
            Description:       description,
            Status:            "created, modified",
            ClassName:         FabricNodeClassName,
            Rn:                fabricNodeRn,
        },

        FabricNodeAttributes: fabricNodeattr,

    }
}
*/

func (fabricNode *FabricNode) ToMap() (map[string]string, error) {
	fabricNodeMap, err := fabricNode.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(fabricNodeMap, "adSt", fabricNode.AdSt)
	A(fabricNodeMap, "address", fabricNode.Address)
	A(fabricNodeMap, "annotation", fabricNode.Annotation)
	A(fabricNodeMap, "apicType", fabricNode.ApicType)
	A(fabricNodeMap, "delayedHeartbeat", fabricNode.DelayedHeartbeat)
	A(fabricNodeMap, "extMngdBy", fabricNode.ExtMngdBy)
	A(fabricNodeMap, "fabricSt", fabricNode.FabricSt)
	A(fabricNodeMap, "id", fabricNode.Id)
	A(fabricNodeMap, "lastStateModTs", fabricNode.LastStateModTs)
	A(fabricNodeMap, "modTs", fabricNode.ModTs)
	A(fabricNodeMap, "model", fabricNode.Model)
	A(fabricNodeMap, "monPolDn", fabricNode.MonPolDn)
	A(fabricNodeMap, "name", fabricNode.Name)
	A(fabricNodeMap, "nameAlias", fabricNode.NameAlias)
	A(fabricNodeMap, "nodeType", fabricNode.NodeType)
	A(fabricNodeMap, "role", fabricNode.Role)
	A(fabricNodeMap, "serial", fabricNode.Serial)
	A(fabricNodeMap, "uid", fabricNode.Uid)
	A(fabricNodeMap, "userdom", fabricNode.Userdom)
	A(fabricNodeMap, "vendor", fabricNode.Vendor)
	A(fabricNodeMap, "version", fabricNode.Version)

	return fabricNodeMap, err
}

func FabricNodeFromContainerList(cont *container.Container, index int) *FabricNode {

	FabricNodeCont := cont.S("imdata").Index(index).S(FabricNodeClassName, "attributes")
	return &FabricNode{
		BaseAttributes{
			DistinguishedName: G(FabricNodeCont, "dn"),
			Description:       G(FabricNodeCont, "descr"),
			Status:            G(FabricNodeCont, "status"),
			ClassName:         FabricNodeClassName,
			Rn:                G(FabricNodeCont, "rn"),
		},

		FabricNodeAttributes{
			AdSt:             G(FabricNodeCont, "adSt"),
			Address:          G(FabricNodeCont, "address"),
			Annotation:       G(FabricNodeCont, "annotation"),
			ApicType:         G(FabricNodeCont, "apicType"),
			DelayedHeartbeat: G(FabricNodeCont, "delayedHeartbeat"),
			ExtMngdBy:        G(FabricNodeCont, "extMngdBy"),
			FabricSt:         G(FabricNodeCont, "fabricSt"),
			Id:               G(FabricNodeCont, "id"),
			LastStateModTs:   G(FabricNodeCont, "lastStateModTs"),
			ModTs:            G(FabricNodeCont, "modTs"),
			Model:            G(FabricNodeCont, "model"),
			MonPolDn:         G(FabricNodeCont, "monPolDn"),
			Name:             G(FabricNodeCont, "name"),
			NameAlias:        G(FabricNodeCont, "nameAlias"),
			NodeType:         G(FabricNodeCont, "nodeType"),
			Role:             G(FabricNodeCont, "role"),
			Serial:           G(FabricNodeCont, "serial"),
			Uid:              G(FabricNodeCont, "uid"),
			Userdom:          G(FabricNodeCont, "userdom"),
			Vendor:           G(FabricNodeCont, "vendor"),
			Version:          G(FabricNodeCont, "version"),
		},
	}
}

func FabricNodeFromContainer(cont *container.Container) *FabricNode {

	return FabricNodeFromContainerList(cont, 0)
}

func FabricNodeListFromContainer(cont *container.Container) []*FabricNode {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*FabricNode, length)

	for i := 0; i < length; i++ {

		arr[i] = FabricNodeFromContainerList(cont, i)
	}

	return arr
}
