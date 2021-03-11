package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateAutonomousSystemProfile(peer_connectivity_profile_addr string, logical_node_profile string, l3_outside string, tenant string, description string, bgpAsPattr models.AutonomousSystemProfileAttributes) (*models.AutonomousSystemProfile, error) {
	rn := fmt.Sprintf("as")
	parentDn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/peerP-[%s]", tenant, l3_outside, logical_node_profile, peer_connectivity_profile_addr)
	bgpAsP := models.NewAutonomousSystemProfile(rn, parentDn, description, bgpAsPattr)
	err := sm.Save(bgpAsP)
	return bgpAsP, err
}

func (sm *ServiceManager) ReadAutonomousSystemProfile(peer_connectivity_profile_addr string, logical_node_profile string, l3_outside string, tenant string) (*models.AutonomousSystemProfile, error) {
	dn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/peerP-[%s]/as", tenant, l3_outside, logical_node_profile, peer_connectivity_profile_addr)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	bgpAsP := models.AutonomousSystemProfileFromContainer(cont)
	return bgpAsP, nil
}

func (sm *ServiceManager) DeleteAutonomousSystemProfile(peer_connectivity_profile_addr string, logical_node_profile string, l3_outside string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/peerP-[%s]/as", tenant, l3_outside, logical_node_profile, peer_connectivity_profile_addr)
	return sm.DeleteByDn(dn, models.BgpaspClassName)
}

func (sm *ServiceManager) UpdateAutonomousSystemProfile(peer_connectivity_profile_addr string, logical_node_profile string, l3_outside string, tenant string, description string, bgpAsPattr models.AutonomousSystemProfileAttributes) (*models.AutonomousSystemProfile, error) {
	rn := fmt.Sprintf("as")
	parentDn := fmt.Sprintf("uni/tn-%s/out-%s/lnodep-%s/peerP-[%s]", tenant, l3_outside, logical_node_profile, peer_connectivity_profile_addr)
	bgpAsP := models.NewAutonomousSystemProfile(rn, parentDn, description, bgpAsPattr)

	bgpAsP.Status = "modified"
	err := sm.Save(bgpAsP)
	return bgpAsP, err

}

func (sm *ServiceManager) ListAutonomousSystemProfile(peer_connectivity_profile_addr string, logical_node_profile string, l3_outside string, tenant string) ([]*models.AutonomousSystemProfile, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/out-%s/lnodep-%s/peerP-[%s]/bgpAsP.json", baseurlStr, tenant, l3_outside, logical_node_profile, peer_connectivity_profile_addr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.AutonomousSystemProfileListFromContainer(cont)

	return list, err
}
