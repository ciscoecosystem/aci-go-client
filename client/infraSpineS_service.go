package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"



	


)









func (sm *ServiceManager) CreateSwitchAssociation(switch_association_type string , name string ,spine_profile string , description string, infraSpineSattr models.SwitchAssociationAttributes) (*models.SwitchAssociation, error) {	
	rn := fmt.Sprintf("spines-%s-typ-%s",name, switch_association_type)
	parentDn := fmt.Sprintf("uni/infra/spprof-%s", spine_profile )
	infraSpineS := models.NewSwitchAssociation(rn, parentDn, description, infraSpineSattr)
	err := sm.Save(infraSpineS)
	return infraSpineS, err
}

func (sm *ServiceManager) ReadSwitchAssociation(switch_association_type string , name string ,spine_profile string ) (*models.SwitchAssociation, error) {
	dn := fmt.Sprintf("uni/infra/spprof-%s/spines-%s-typ-%s", spine_profile ,name , ,switch_association_type )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	infraSpineS := models.SwitchAssociationFromContainer(cont)
	return infraSpineS, nil
}

func (sm *ServiceManager) DeleteSwitchAssociation(switch_association_type string , name string ,spine_profile string ) error {
	dn := fmt.Sprintf("uni/infra/spprof-%s/spines-%s-typ-%s", spine_profile ,name , ,switch_association_type )
	return sm.DeleteByDn(dn, models.InfraspinesClassName)
}

func (sm *ServiceManager) UpdateSwitchAssociation(switch_association_type string , name string ,spine_profile string  ,description string, infraSpineSattr models.SwitchAssociationAttributes) (*models.SwitchAssociation, error) {
	rn := fmt.Sprintf("spines-%s-typ-%s",name, switch_association_type)
	parentDn := fmt.Sprintf("uni/infra/spprof-%s", spine_profile )
	infraSpineS := models.NewSwitchAssociation(rn, parentDn, description, infraSpineSattr)

    infraSpineS.Status = "modified"
	err := sm.Save(infraSpineS)
	return infraSpineS, err

}

func (sm *ServiceManager) ListSwitchAssociation(spine_profile string ) ([]*models.SwitchAssociation, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/infra/spprof-%s/infraSpineS.json", baseurlStr , spine_profile )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.SwitchAssociationListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationinfraRsSpineAccNodePGrpFromSwitchAssociation( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsspineAccNodePGrp", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tDn": "%s","annotation":"orchestrator:terraform"
								
			}
		}
	}`, "infraRsSpineAccNodePGrp", dn, tDn))

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

func (sm *ServiceManager) DeleteRelationinfraRsSpineAccNodePGrpFromSwitchAssociation(parentDn string) error{
	dn := fmt.Sprintf("%s/rsspineAccNodePGrp", parentDn)
	return sm.DeleteByDn(dn , "infraRsSpineAccNodePGrp")
}

func (sm *ServiceManager) ReadRelationinfraRsSpineAccNodePGrpFromSwitchAssociation( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"infraRsSpineAccNodePGrp")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"infraRsSpineAccNodePGrp")
	
	if len(contList) > 0 {
		dat := models.G(contList[0], "tnInfraSpineAccNodePGrpName")
		return dat, err
	} else {
		return nil,err
	}
		





}

