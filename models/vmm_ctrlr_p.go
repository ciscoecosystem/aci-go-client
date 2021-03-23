package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VmmctrlrpClassName = "vmmCtrlrP"

type VMMController struct {
	BaseAttributes
	VMMControllerAttributes
}

type VMMControllerAttributes struct {
	Name             string `json:",omitempty"`
	Annotation       string `json:",omitempty"`
	DvsVersion       string `json:",omitempty"`
	HostOrIp         string `json:",omitempty"`
	InventoryTrigSt  string `json:",omitempty"`
	Mode             string `json:",omitempty"`
	MsftConfigErrMsg string `json:",omitempty"`
	MsftConfigIssues string `json:",omitempty"`
	N1kvStatsMode    string `json:",omitempty"`
	NameAlias        string `json:",omitempty"`
	Port             string `json:",omitempty"`
	RootContName     string `json:",omitempty"`
	Scope            string `json:",omitempty"`
	SeqNum           string `json:",omitempty"`
	StatsMode        string `json:",omitempty"`
	VxlanDeplPref    string `json:",omitempty"`
}

func NewVMMController(vmmCtrlrPRn, parentDn, description string, vmmCtrlrPattr VMMControllerAttributes) *VMMController {
	dn := fmt.Sprintf("%s/%s", parentDn, vmmCtrlrPRn)
	return &VMMController{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VmmctrlrpClassName,
			Rn:                vmmCtrlrPRn,
		},

		VMMControllerAttributes: vmmCtrlrPattr,
	}
}

func (vmmCtrlrP *VMMController) ToMap() (map[string]string, error) {
	vmmCtrlrPMap, err := vmmCtrlrP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vmmCtrlrPMap, "name", vmmCtrlrP.Name)
	A(vmmCtrlrPMap, "annotation", vmmCtrlrP.Annotation)
	A(vmmCtrlrPMap, "dvsVersion", vmmCtrlrP.DvsVersion)
	A(vmmCtrlrPMap, "hostOrIp", vmmCtrlrP.HostOrIp)
	A(vmmCtrlrPMap, "inventoryTrigSt", vmmCtrlrP.InventoryTrigSt)
	A(vmmCtrlrPMap, "mode", vmmCtrlrP.Mode)
	A(vmmCtrlrPMap, "msftConfigErrMsg", vmmCtrlrP.MsftConfigErrMsg)
	A(vmmCtrlrPMap, "msftConfigIssues", vmmCtrlrP.MsftConfigIssues)
	A(vmmCtrlrPMap, "n1kvStatsMode", vmmCtrlrP.N1kvStatsMode)
	A(vmmCtrlrPMap, "nameAlias", vmmCtrlrP.NameAlias)
	A(vmmCtrlrPMap, "port", vmmCtrlrP.Port)
	A(vmmCtrlrPMap, "rootContName", vmmCtrlrP.RootContName)
	A(vmmCtrlrPMap, "scope", vmmCtrlrP.Scope)
	A(vmmCtrlrPMap, "seqNum", vmmCtrlrP.SeqNum)
	A(vmmCtrlrPMap, "statsMode", vmmCtrlrP.StatsMode)
	A(vmmCtrlrPMap, "vxlanDeplPref", vmmCtrlrP.VxlanDeplPref)

	return vmmCtrlrPMap, err
}

func VMMControllerFromContainerList(cont *container.Container, index int) *VMMController {

	VMMControllerCont := cont.S("imdata").Index(index).S(VmmctrlrpClassName, "attributes")
	return &VMMController{
		BaseAttributes{
			DistinguishedName: G(VMMControllerCont, "dn"),
			Description:       G(VMMControllerCont, "descr"),
			Status:            G(VMMControllerCont, "status"),
			ClassName:         VmmctrlrpClassName,
			Rn:                G(VMMControllerCont, "rn"),
		},

		VMMControllerAttributes{
			Name:             G(VMMControllerCont, "name"),
			Annotation:       G(VMMControllerCont, "annotation"),
			DvsVersion:       G(VMMControllerCont, "dvsVersion"),
			HostOrIp:         G(VMMControllerCont, "hostOrIp"),
			InventoryTrigSt:  G(VMMControllerCont, "inventoryTrigSt"),
			Mode:             G(VMMControllerCont, "mode"),
			MsftConfigErrMsg: G(VMMControllerCont, "msftConfigErrMsg"),
			MsftConfigIssues: G(VMMControllerCont, "msftConfigIssues"),
			N1kvStatsMode:    G(VMMControllerCont, "n1kvStatsMode"),
			NameAlias:        G(VMMControllerCont, "nameAlias"),
			Port:             G(VMMControllerCont, "port"),
			RootContName:     G(VMMControllerCont, "rootContName"),
			Scope:            G(VMMControllerCont, "scope"),
			SeqNum:           G(VMMControllerCont, "seqNum"),
			StatsMode:        G(VMMControllerCont, "statsMode"),
			VxlanDeplPref:    G(VMMControllerCont, "vxlanDeplPref"),
		},
	}
}

func VMMControllerFromContainer(cont *container.Container) *VMMController {

	return VMMControllerFromContainerList(cont, 0)
}

func VMMControllerListFromContainer(cont *container.Container) []*VMMController {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*VMMController, length)

	for i := 0; i < length; i++ {

		arr[i] = VMMControllerFromContainerList(cont, i)
	}

	return arr
}
