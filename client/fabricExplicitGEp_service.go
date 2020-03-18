package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateVPCExplicitProtectionGroup(name string, description string, fabricExplicitGEpattr models.VPCExplicitProtectionGroupAttributes) (*models.VPCExplicitProtectionGroup, error) {
	rn := fmt.Sprintf("fabric/protpol/expgep-%s", name)
	parentDn := fmt.Sprintf("uni")
	fabricExplicitGEp := models.NewVPCExplicitProtectionGroup(rn, parentDn, description, fabricExplicitGEpattr)
	err := sm.Save(fabricExplicitGEp)
	return fabricExplicitGEp, err
}

func (sm *ServiceManager) ReadVPCExplicitProtectionGroup(name string) (*models.VPCExplicitProtectionGroup, error) {
	dn := fmt.Sprintf("uni/fabric/protpol/expgep-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fabricExplicitGEp := models.VPCExplicitProtectionGroupFromContainer(cont)
	return fabricExplicitGEp, nil
}

func (sm *ServiceManager) DeleteVPCExplicitProtectionGroup(name string) error {
	dn := fmt.Sprintf("uni/fabric/protpol/expgep-%s", name)
	return sm.DeleteByDn(dn, models.FabricexplicitgepClassName)
}

func (sm *ServiceManager) UpdateVPCExplicitProtectionGroup(name string, description string, fabricExplicitGEpattr models.VPCExplicitProtectionGroupAttributes) (*models.VPCExplicitProtectionGroup, error) {
	rn := fmt.Sprintf("fabric/protpol/expgep-%s", name)
	parentDn := fmt.Sprintf("uni")
	fabricExplicitGEp := models.NewVPCExplicitProtectionGroup(rn, parentDn, description, fabricExplicitGEpattr)

	fabricExplicitGEp.Status = "modified"
	err := sm.Save(fabricExplicitGEp)
	return fabricExplicitGEp, err

}

func (sm *ServiceManager) ListVPCExplicitProtectionGroup() ([]*models.VPCExplicitProtectionGroup, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/fabricExplicitGEp.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.VPCExplicitProtectionGroupListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfabricRsVpcInstPolFromVPCExplicitProtectionGroup(parentDn, tnVpcInstPolName string) error {
	dn := fmt.Sprintf("%s/rsvpcInstPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnVpcInstPolName": "%s"
								
			}
		}
	}`, "fabricRsVpcInstPol", dn, tnVpcInstPolName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationfabricRsVpcInstPolFromVPCExplicitProtectionGroup(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "fabricRsVpcInstPol")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "fabricRsVpcInstPol")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnVpcInstPolName")
		return dat, err
	} else {
		return nil, err
	}

}
