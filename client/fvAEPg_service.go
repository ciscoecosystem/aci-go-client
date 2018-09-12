package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateApplicationEPG(name string ,application_profile string ,tenant string  ,description string, fvAEPgattr models.ApplicationEPGAttributes) (*models.ApplicationEPG, error) {	
	rn := fmt.Sprintf("epg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,application_profile )
	fvAEPg := models.NewApplicationEPG(rn, parentDn, description, fvAEPgattr)
	err := sm.Save(fvAEPg)
	return fvAEPg, err
}

func (sm *ServiceManager) ReadApplicationEPG(name string ,application_profile string ,tenant string ) (*models.ApplicationEPG, error) {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant ,application_profile ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAEPg := models.ApplicationEPGFromContainer(cont)
	return fvAEPg, nil
}

func (sm *ServiceManager) DeleteApplicationEPG(name string ,application_profile string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant ,application_profile ,name )
	return sm.DeleteByDn(dn, models.FvaepgClassName)
}

func (sm *ServiceManager) UpdateApplicationEPG(name string ,application_profile string ,tenant string  ,description string, fvAEPgattr models.ApplicationEPGAttributes) (*models.ApplicationEPG, error) {
	rn := fmt.Sprintf("epg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,application_profile )
	fvAEPg := models.NewApplicationEPG(rn, parentDn, description, fvAEPgattr)

    fvAEPg.Status = "modified"
	err := sm.Save(fvAEPg)
	return fvAEPg, err

}

func (sm *ServiceManager) ListApplicationEPG(application_profile string ,tenant string ) ([]*models.ApplicationEPG, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ap-%s/fvAEPg.json", baseurlStr , tenant ,application_profile )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ApplicationEPGListFromContainer(cont)

	return list, err
}