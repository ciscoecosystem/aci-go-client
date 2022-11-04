package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	DncloudtemplateBgpIpv4        = "uni/tn-%s/infranetwork-%s/extnetwork-%s/vpnnetwork-%s/ipsec-[%s]/bgpipv4-[%s]"
	RncloudtemplateBgpIpv4        = "bgpipv4-[%s]"
	ParentDncloudtemplateBgpIpv4  = "uni/tn-%s/infranetwork-%s/extnetwork-%s/vpnnetwork-%s/ipsec-[%s]"
	Cloudtemplatebgpipv4ClassName = "cloudtemplateBgpIpv4"
)

type BGPIPv4Peer struct {
	BaseAttributes
	NameAliasAttribute
	BGPIPv4PeerAttributes
}

type BGPIPv4PeerAttributes struct {
	Annotation string `json:",omitempty"`
	Peeraddr   string `json:",omitempty"`
	Peerasn    string `json:",omitempty"`
}

func NewBGPIPv4Peer(cloudtemplateBgpIpv4Rn, parentDn string, cloudtemplateBgpIpv4Attr BGPIPv4PeerAttributes) *BGPIPv4Peer {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudtemplateBgpIpv4Rn)
	return &BGPIPv4Peer{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         Cloudtemplatebgpipv4ClassName,
			Rn:                cloudtemplateBgpIpv4Rn,
		},
		BGPIPv4PeerAttributes: cloudtemplateBgpIpv4Attr,
	}
}

func (cloudtemplateBgpIpv4 *BGPIPv4Peer) ToMap() (map[string]string, error) {
	cloudtemplateBgpIpv4Map, err := cloudtemplateBgpIpv4.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudtemplateBgpIpv4Map, "annotation", cloudtemplateBgpIpv4.Annotation)
	A(cloudtemplateBgpIpv4Map, "peeraddr", cloudtemplateBgpIpv4.Peeraddr)
	A(cloudtemplateBgpIpv4Map, "peerasn", cloudtemplateBgpIpv4.Peerasn)
	return cloudtemplateBgpIpv4Map, err
}

func BGPIPv4PeerFromContainerList(cont *container.Container, index int) *BGPIPv4Peer {
	BGPIPv4PeerCont := cont.S("imdata").Index(index).S(Cloudtemplatebgpipv4ClassName, "attributes")
	return &BGPIPv4Peer{
		BaseAttributes{
			DistinguishedName: G(BGPIPv4PeerCont, "dn"),
			Status:            G(BGPIPv4PeerCont, "status"),
			ClassName:         Cloudtemplatebgpipv4ClassName,
			Rn:                G(BGPIPv4PeerCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(BGPIPv4PeerCont, "nameAlias"),
		},
		BGPIPv4PeerAttributes{
			Annotation: G(BGPIPv4PeerCont, "annotation"),
			Peeraddr:   G(BGPIPv4PeerCont, "peeraddr"),
			Peerasn:    G(BGPIPv4PeerCont, "peerasn"),
		},
	}
}

func BGPIPv4PeerFromContainer(cont *container.Container) *BGPIPv4Peer {
	return BGPIPv4PeerFromContainerList(cont, 0)
}

func BGPIPv4PeerListFromContainer(cont *container.Container) []*BGPIPv4Peer {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*BGPIPv4Peer, length)

	for i := 0; i < length; i++ {
		arr[i] = BGPIPv4PeerFromContainerList(cont, i)
	}

	return arr
}
