package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateFabricVpcDomainPolicy(name string, description string, fabricVpcDomPolattr models.FabricVpcDomainPolicyAttributes) (*models.FabricVpcDomainPolicy, error) {
	rn := fmt.Sprintf("fabric/vpcInst-%s", name)
	parentDn := fmt.Sprintf("uni")
	fabricVpcDomPol := models.NewFabricVpcDomainPolicy(rn, parentDn, description, fabricVpcDomPolattr)
	err := sm.Save(fabricVpcDomPol)
	return fabricVpcDomPol, err
}

func (sm *ServiceManager) ReadFabricVpcDomainPolicy(name string) (*models.FabricVpcDomainPolicy, error) {
	dn := fmt.Sprintf("uni/fabric/vpcInst-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fabricVpcDomPol := models.FabricVpcDomainPolicyFromContainer(cont)
	return fabricVpcDomPol, nil
}

func (sm *ServiceManager) DeleteFabricVpcDomainPolicy(name string) error {
	dn := fmt.Sprintf("uni/fabric/vpcInst-%s", name)
	return sm.DeleteByDn(dn, models.VpcDomPolClassName)
}

func (sm *ServiceManager) UpdateFabricVpcDomainPolicy(name string, description string, fabricVpcDomPolattr models.FabricVpcDomainPolicyAttributes) (*models.FabricVpcDomainPolicy, error) {
	rn := fmt.Sprintf("fabric/vpcInst-%s", name)
	parentDn := fmt.Sprintf("uni")
	fabricVpcDomPol := models.NewFabricVpcDomainPolicy(rn, parentDn, description, fabricVpcDomPolattr)

	fabricVpcDomPol.Status = "modified"
	err := sm.Save(fabricVpcDomPol)
	return fabricVpcDomPol, err

}

func (sm *ServiceManager) ListFabricVpcDomainPolicy() ([]*models.FabricVpcDomainPolicy, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/vpcInstPol.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.FabricVpcDomainPolicyListFromContainer(cont)

	return list, err
}
