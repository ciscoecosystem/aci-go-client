package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateApplicationProfile(name string ,tenant string  ,description string, fvApattr models.ApplicationProfileAttributes) (*models.ApplicationProfile, error) {	
	rn := fmt.Sprintf("ap-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvAp := models.NewApplicationProfile(rn, parentDn, description, fvApattr)
	err := sm.Save(fvAp)
	return fvAp, err
}

func (sm *ServiceManager) ReadApplicationProfile(name string ,tenant string ) (*models.ApplicationProfile, error) {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAp := models.ApplicationProfileFromContainer(cont)
	return fvAp, nil
}

func (sm *ServiceManager) DeleteApplicationProfile(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvapClassName)
}

func (sm *ServiceManager) UpdateApplicationProfile(name string ,tenant string  ,description string, fvApattr models.ApplicationProfileAttributes) (*models.ApplicationProfile, error) {
	rn := fmt.Sprintf("ap-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvAp := models.NewApplicationProfile(rn, parentDn, description, fvApattr)

    fvAp.Status = "modified"
	err := sm.Save(fvAp)
	return fvAp, err

}

func (sm *ServiceManager) ListApplicationProfile(tenant string ) ([]*models.ApplicationProfile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvAp.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ApplicationProfileListFromContainer(cont)

	return list, err
}