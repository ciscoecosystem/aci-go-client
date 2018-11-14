package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateTenant(name string  ,description string, fvTenantattr models.TenantAttributes) (*models.Tenant, error) {	
	rn := fmt.Sprintf("tn-%s",name)
	parentDn := fmt.Sprintf("uni")
	fvTenant := models.NewTenant(rn, parentDn, description, fvTenantattr)
	err := sm.Save(fvTenant)
	return fvTenant, err
}

func (sm *ServiceManager) ReadTenant(name string ) (*models.Tenant, error) {
	dn := fmt.Sprintf("uni/tn-%s", name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvTenant := models.TenantFromContainer(cont)
	return fvTenant, nil
}

func (sm *ServiceManager) DeleteTenant(name string ) error {
	dn := fmt.Sprintf("uni/tn-%s", name )
	return sm.DeleteByDn(dn, models.FvtenantClassName)
}

func (sm *ServiceManager) UpdateTenant(name string  ,description string, fvTenantattr models.TenantAttributes) (*models.Tenant, error) {
	rn := fmt.Sprintf("tn-%s",name)
	parentDn := fmt.Sprintf("uni")
	fvTenant := models.NewTenant(rn, parentDn, description, fvTenantattr)

    fvTenant.Status = "modified"
	err := sm.Save(fvTenant)
	return fvTenant, err

}

func (sm *ServiceManager) ListTenant() ([]*models.Tenant, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/fvTenant.json", baseurlStr )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.TenantListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsTnDenyRule( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rstnDenyRule-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsTnDenyRule", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsTnDenyRule(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rstnDenyRule-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsTnDenyRule")
}


