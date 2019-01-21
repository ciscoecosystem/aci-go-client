package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvbdClassName = "fvBD"

type Bridgedomain struct {
	BaseAttributes
    BridgedomainAttributes 
}
  
type BridgedomainAttributes struct {
    OptimizeWanBandwidth       string `json:",omitempty"`
    Annotation       string `json:",omitempty"`
    ArpFlood       string `json:",omitempty"`
    EpClear       string `json:",omitempty"`
    EpMoveDetectMode       string `json:",omitempty"`
    HostBasedRouting       string `json:",omitempty"`
    IntersiteBumTrafficAllow       string `json:",omitempty"`
    IntersiteL2Stretch       string `json:",omitempty"`
    IpLearning       string `json:",omitempty"`
    Ipv6McastAllow       string `json:",omitempty"`
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
    V6unkMcastAct       string `json:",omitempty"`
    Vmac       string `json:",omitempty"`
    
}
   

func NewBridgedomain(fvBDRn, parentDn, description string, fvBDattr BridgedomainAttributes) *Bridgedomain {
	dn := fmt.Sprintf("%s/%s", parentDn, fvBDRn)  
	return &Bridgedomain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvbdClassName,
			Rn:                fvBDRn,
		},
        
		BridgedomainAttributes: fvBDattr,
         
	}
}

func (fvBD *Bridgedomain) ToMap() (map[string]string, error) {
	fvBDMap, err := fvBD.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(fvBDMap, "OptimizeWanBandwidth",fvBD.OptimizeWanBandwidth)
    A(fvBDMap, "annotation",fvBD.Annotation)
    A(fvBDMap, "arpFlood",fvBD.ArpFlood)
    A(fvBDMap, "epClear",fvBD.EpClear)
    A(fvBDMap, "epMoveDetectMode",fvBD.EpMoveDetectMode)
    A(fvBDMap, "hostBasedRouting",fvBD.HostBasedRouting)
    A(fvBDMap, "intersiteBumTrafficAllow",fvBD.IntersiteBumTrafficAllow)
    A(fvBDMap, "intersiteL2Stretch",fvBD.IntersiteL2Stretch)
    A(fvBDMap, "ipLearning",fvBD.IpLearning)
    A(fvBDMap, "ipv6McastAllow",fvBD.Ipv6McastAllow)
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
    A(fvBDMap, "v6unkMcastAct",fvBD.V6unkMcastAct)
    A(fvBDMap, "vmac",fvBD.Vmac)
    
	

	return fvBDMap, err
}

func BridgedomainFromContainerList(cont *container.Container, index int) *Bridgedomain {

	BridgedomainCont := cont.S("imdata").Index(index).S(FvbdClassName, "attributes")
	return &Bridgedomain{
		BaseAttributes{
			DistinguishedName: G(BridgedomainCont, "dn"),
			Description:       G(BridgedomainCont, "descr"),
			Status:            G(BridgedomainCont, "status"),
			ClassName:         FvbdClassName,
			Rn:                G(BridgedomainCont, "rn"),
		},
        
		BridgedomainAttributes{
        OptimizeWanBandwidth : G(BridgedomainCont, "OptimizeWanBandwidth"),
        Annotation : G(BridgedomainCont, "annotation"),
        ArpFlood : G(BridgedomainCont, "arpFlood"),
        EpClear : G(BridgedomainCont, "epClear"),
        EpMoveDetectMode : G(BridgedomainCont, "epMoveDetectMode"),
        HostBasedRouting : G(BridgedomainCont, "hostBasedRouting"),
        IntersiteBumTrafficAllow : G(BridgedomainCont, "intersiteBumTrafficAllow"),
        IntersiteL2Stretch : G(BridgedomainCont, "intersiteL2Stretch"),
        IpLearning : G(BridgedomainCont, "ipLearning"),
        Ipv6McastAllow : G(BridgedomainCont, "ipv6McastAllow"),
        LimitIpLearnToSubnets : G(BridgedomainCont, "limitIpLearnToSubnets"),
        LlAddr : G(BridgedomainCont, "llAddr"),
        Mac : G(BridgedomainCont, "mac"),
        McastAllow : G(BridgedomainCont, "mcastAllow"),
        MultiDstPktAct : G(BridgedomainCont, "multiDstPktAct"),
        NameAlias : G(BridgedomainCont, "nameAlias"),
        Type : G(BridgedomainCont, "type"),
        UnicastRoute : G(BridgedomainCont, "unicastRoute"),
        UnkMacUcastAct : G(BridgedomainCont, "unkMacUcastAct"),
        UnkMcastAct : G(BridgedomainCont, "unkMcastAct"),
        V6unkMcastAct : G(BridgedomainCont, "v6unkMcastAct"),
        Vmac : G(BridgedomainCont, "vmac"),
        		
        },
        
	}
}

func BridgedomainFromContainer(cont *container.Container) *Bridgedomain {

	return BridgedomainFromContainerList(cont, 0)
}

func BridgedomainListFromContainer(cont *container.Container) []*Bridgedomain {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Bridgedomain, length)

	for i := 0; i < length; i++ {

		arr[i] = BridgedomainFromContainerList(cont, i)
	}

	return arr
}