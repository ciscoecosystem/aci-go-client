package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateEPGsUsingFunction(tDn string, attachable_access_entity_profile string, description string, infraRsFuncToEpgattr models.EPGsUsingFunctionAttributes) (*models.EPGsUsingFunction, error) {
	rn := fmt.Sprintf("provacc/rsfuncToEpg-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/infra/attentp-%s", attachable_access_entity_profile)
	infraRsFuncToEpg := models.NewEPGsUsingFunction(rn, parentDn, description, infraRsFuncToEpgattr)
	err := sm.Save(infraRsFuncToEpg)
	return infraRsFuncToEpg, err
}

func (sm *ServiceManager) ReadEPGsUsingFunction(tDn string, attachable_access_entity_profile string) (*models.EPGsUsingFunction, error) {
	dn := fmt.Sprintf("uni/infra/attentp-%s/provacc/rsfuncToEpg-[%s]", attachable_access_entity_profile, tDn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	infraRsFuncToEpg := models.EPGsUsingFunctionFromContainer(cont)
	return infraRsFuncToEpg, nil
}

func (sm *ServiceManager) DeleteEPGsUsingFunction(tDn string, attachable_access_entity_profile string) error {
	dn := fmt.Sprintf("uni/infra/attentp-%s/provacc/rsfuncToEpg-[%s]", attachable_access_entity_profile, tDn)
	return sm.DeleteByDn(dn, models.InfrarsfunctoepgClassName)
}

func (sm *ServiceManager) UpdateEPGsUsingFunction(tDn string, attachable_access_entity_profile string, description string, infraRsFuncToEpgattr models.EPGsUsingFunctionAttributes) (*models.EPGsUsingFunction, error) {
	rn := fmt.Sprintf("provacc/rsfuncToEpg-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/infra/attentp-%s", attachable_access_entity_profile)
	infraRsFuncToEpg := models.NewEPGsUsingFunction(rn, parentDn, description, infraRsFuncToEpgattr)

	infraRsFuncToEpg.Status = "modified"
	err := sm.Save(infraRsFuncToEpg)
	return infraRsFuncToEpg, err

}

func (sm *ServiceManager) ListEPGsUsingFunction(attachable_access_entity_profile string) ([]*models.EPGsUsingFunction, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/infra/attentp-%s/infraRsFuncToEpg.json", baseurlStr, attachable_access_entity_profile)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.EPGsUsingFunctionListFromContainer(cont)

	return list, err
}
