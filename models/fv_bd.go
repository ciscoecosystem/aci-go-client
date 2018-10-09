package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvbdClassName = "fvBD"

type BridgeDomain struct {
	BaseAttributes
    BridgeDomainAttributes 
}
  
type BridgeDomainAttributes struct {
    OptimizeWanBandwidth       string `json:",omitempty"`
    ArpFlood       string `json:",omitempty"`
    EpClear       string `json:",omitempty"`
    EpMoveDetectMode       string `json:",omitempty"`
    IntersiteBumTrafficAllow       string `json:",omitempty"`
    IntersiteL2Stretch       string `json:",omitempty"`
    IpLearning       string `json:",omitempty"`
    LimitIpLearnToSubnets       string `json:",omitempty"`
    LlAddr       string `json:",omitempty"`
    Mac       string `json:",omitempty"`
    McastAllow       string `json:",omitempty"`
    MultiDstPktAct       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Type       string `json:",omitempty"`
    UnicastRoute       string `json:",omitempty"`
    UnkMacUcastAct       string `json:",omitempty"`
    UnkMcastAct       string `json:",omitempty"`
    Vmac       string `json:",omitempty"`
    
}
   

func NewBridgeDomain(fvBDRn, parentDn, description string, fvBDattr BridgeDomainAttributes) *BridgeDomain {
	dn := fmt.Sprintf("%s/%s", parentDn, fvBDRn)  
	return &BridgeDomain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvbdClassName,
			Rn:                fvBDRn,
		},
        
		BridgeDomainAttributes: fvBDattr,
         
	}
}

func (fvBD *BridgeDomain) ToMap() (map[string]string, error) {
	fvBDMap, err := fvBD.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(fvBDMap, "OptimizeWanBandwidth",fvBD.OptimizeWanBandwidth)
    A(fvBDMap, "arpFlood",fvBD.ArpFlood)
    A(fvBDMap, "epClear",fvBD.EpClear)
    A(fvBDMap, "epMoveDetectMode",fvBD.EpMoveDetectMode)
    A(fvBDMap, "intersiteBumTrafficAllow",fvBD.IntersiteBumTrafficAllow)
    A(fvBDMap, "intersiteL2Stretch",fvBD.IntersiteL2Stretch)
    A(fvBDMap, "ipLearning",fvBD.IpLearning)
    A(fvBDMap, "limitIpLearnToSubnets",fvBD.LimitIpLearnToSubnets)
    A(fvBDMap, "llAddr",fvBD.LlAddr)
    A(fvBDMap, "mac",fvBD.Mac)
    A(fvBDMap, "mcastAllow",fvBD.McastAllow)
    A(fvBDMap, "multiDstPktAct",fvBD.MultiDstPktAct)
    A(fvBDMap, "nameAlias",fvBD.NameAlias)
    A(fvBDMap, "type",fvBD.Type)
    A(fvBDMap, "unicastRoute",fvBD.UnicastRoute)
    A(fvBDMap, "unkMacUcastAct",fvBD.UnkMacUcastAct)
    A(fvBDMap, "unkMcastAct",fvBD.UnkMcastAct)
    A(fvBDMap, "vmac",fvBD.Vmac)
    
	

	return fvBDMap, err
}

func BridgeDomainFromContainerList(cont *container.Container, index int) *BridgeDomain {

	BridgeDomainCont := cont.S("imdata").Index(index).S(FvbdClassName, "attributes")
	return &BridgeDomain{
		BaseAttributes{
			DistinguishedName: G(BridgeDomainCont, "dn"),
			Description:       G(BridgeDomainCont, "descr"),
			Status:            G(BridgeDomainCont, "status"),
			ClassName:         FvbdClassName,
			Rn:                G(BridgeDomainCont, "rn"),
		},
        
		BridgeDomainAttributes{
        OptimizeWanBandwidth : G(BridgeDomainCont, "OptimizeWanBandwidth"),
        ArpFlood : G(BridgeDomainCont, "arpFlood"),
        EpClear : G(BridgeDomainCont, "epClear"),
        EpMoveDetectMode : G(BridgeDomainCont, "epMoveDetectMode"),
        IntersiteBumTrafficAllow : G(BridgeDomainCont, "intersiteBumTrafficAllow"),
        IntersiteL2Stretch : G(BridgeDomainCont, "intersiteL2Stretch"),
        IpLearning : G(BridgeDomainCont, "ipLearning"),
        LimitIpLearnToSubnets : G(BridgeDomainCont, "limitIpLearnToSubnets"),
        LlAddr : G(BridgeDomainCont, "llAddr"),
        Mac : G(BridgeDomainCont, "mac"),
        McastAllow : G(BridgeDomainCont, "mcastAllow"),
        MultiDstPktAct : G(BridgeDomainCont, "multiDstPktAct"),
        NameAlias : G(BridgeDomainCont, "nameAlias"),
        Type : G(BridgeDomainCont, "type"),
        UnicastRoute : G(BridgeDomainCont, "unicastRoute"),
        UnkMacUcastAct : G(BridgeDomainCont, "unkMacUcastAct"),
        UnkMcastAct : G(BridgeDomainCont, "unkMcastAct"),
        Vmac : G(BridgeDomainCont, "vmac"),
        		
        },
        
	}
}

func BridgeDomainFromContainer(cont *container.Container) *BridgeDomain {

	return BridgeDomainFromContainerList(cont, 0)
}

func BridgeDomainListFromContainer(cont *container.Container) []*BridgeDomain {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*BridgeDomain, length)

	for i := 0; i < length; i++ {

		arr[i] = BridgeDomainFromContainerList(cont, i)
	}

	return arr
}