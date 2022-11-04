package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateBGPIPv4Peer(peeraddr string, template_for_ipsec_tunnel_peeraddr string, template_for_vpn_network string, template_for_external_network string, infra_network_template string, tenant string, cloudtemplateBgpIpv4Attr models.BGPIPv4PeerAttributes) (*models.BGPIPv4Peer, error) {
	rn := fmt.Sprintf(models.RncloudtemplateBgpIpv4, peeraddr)
	parentDn := fmt.Sprintf(models.ParentDncloudtemplateBgpIpv4, tenant, infra_network_template, template_for_external_network, template_for_vpn_network, template_for_ipsec_tunnel_peeraddr)
	cloudtemplateBgpIpv4 := models.NewBGPIPv4Peer(rn, parentDn, cloudtemplateBgpIpv4Attr)
	err := sm.Save(cloudtemplateBgpIpv4)
	return cloudtemplateBgpIpv4, err
}

func (sm *ServiceManager) ReadBGPIPv4Peer(peeraddr string, template_for_ipsec_tunnel_peeraddr string, template_for_vpn_network string, template_for_external_network string, infra_network_template string, tenant string) (*models.BGPIPv4Peer, error) {
	dn := fmt.Sprintf(models.DncloudtemplateBgpIpv4, tenant, infra_network_template, template_for_external_network, template_for_vpn_network, template_for_ipsec_tunnel_peeraddr, peeraddr)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudtemplateBgpIpv4 := models.BGPIPv4PeerFromContainer(cont)
	return cloudtemplateBgpIpv4, nil
}

func (sm *ServiceManager) DeleteBGPIPv4Peer(peeraddr string, template_for_ipsec_tunnel_peeraddr string, template_for_vpn_network string, template_for_external_network string, infra_network_template string, tenant string) error {
	dn := fmt.Sprintf(models.DncloudtemplateBgpIpv4, tenant, infra_network_template, template_for_external_network, template_for_vpn_network, template_for_ipsec_tunnel_peeraddr, peeraddr)
	return sm.DeleteByDn(dn, models.Cloudtemplatebgpipv4ClassName)
}

func (sm *ServiceManager) UpdateBGPIPv4Peer(peeraddr string, template_for_ipsec_tunnel_peeraddr string, template_for_vpn_network string, template_for_external_network string, infra_network_template string, tenant string, cloudtemplateBgpIpv4Attr models.BGPIPv4PeerAttributes) (*models.BGPIPv4Peer, error) {
	rn := fmt.Sprintf(models.RncloudtemplateBgpIpv4, peeraddr)
	parentDn := fmt.Sprintf(models.ParentDncloudtemplateBgpIpv4, tenant, infra_network_template, template_for_external_network, template_for_vpn_network, template_for_ipsec_tunnel_peeraddr)
	cloudtemplateBgpIpv4 := models.NewBGPIPv4Peer(rn, parentDn, cloudtemplateBgpIpv4Attr)
	cloudtemplateBgpIpv4.Status = "modified"
	err := sm.Save(cloudtemplateBgpIpv4)
	return cloudtemplateBgpIpv4, err
}

func (sm *ServiceManager) ListBGPIPv4Peer(parentDn string) ([]*models.BGPIPv4Peer, error) {
	dnUrl := fmt.Sprintf("%s/%s/cloudtemplateBgpIpv4.json", models.BaseurlStr, parentDn)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.BGPIPv4PeerListFromContainer(cont)
	return list, err
}
