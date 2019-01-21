package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateFilterentry(name string ,filter string ,tenant string , description string, vzEntryattr models.FilterentryAttributes) (*models.Filterentry, error) {	
	rn := fmt.Sprintf("e-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/flt-%s", tenant ,filter )
	vzEntry := models.NewFilterentry(rn, parentDn, description, vzEntryattr)
	err := sm.Save(vzEntry)
	return vzEntry, err
}

func (sm *ServiceManager) ReadFilterentry(name string ,filter string ,tenant string ) (*models.Filterentry, error) {
	dn := fmt.Sprintf("uni/tn-%s/flt-%s/e-%s", tenant ,filter ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzEntry := models.FilterentryFromContainer(cont)
	return vzEntry, nil
}

func (sm *ServiceManager) DeleteFilterentry(name string ,filter string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/flt-%s/e-%s", tenant ,filter ,name )
	return sm.DeleteByDn(dn, models.VzentryClassName)
}

func (sm *ServiceManager) UpdateFilterentry(name string ,filter string ,tenant string  ,description string, vzEntryattr models.FilterentryAttributes) (*models.Filterentry, error) {
	rn := fmt.Sprintf("e-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/flt-%s", tenant ,filter )
	vzEntry := models.NewFilterentry(rn, parentDn, description, vzEntryattr)

    vzEntry.Status = "modified"
	err := sm.Save(vzEntry)
	return vzEntry, err

}

func (sm *ServiceManager) ListFilterentry(filter string ,tenant string ) ([]*models.Filterentry, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/flt-%s/vzEntry.json", baseurlStr , tenant ,filter )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.FilterentryListFromContainer(cont)

	return list, err
}


