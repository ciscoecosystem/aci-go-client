package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateApplicationprofile(name string ,tenant string , description string, fvApattr models.ApplicationprofileAttributes) (*models.Applicationprofile, error) {	
	rn := fmt.Sprintf("ap-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvAp := models.NewApplicationprofile(rn, parentDn, description, fvApattr)
	err := sm.Save(fvAp)
	return fvAp, err
}

func (sm *ServiceManager) ReadApplicationprofile(name string ,tenant string ) (*models.Applicationprofile, error) {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAp := models.ApplicationprofileFromContainer(cont)
	return fvAp, nil
}

func (sm *ServiceManager) DeleteApplicationprofile(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvapClassName)
}

func (sm *ServiceManager) UpdateApplicationprofile(name string ,tenant string  ,description string, fvApattr models.ApplicationprofileAttributes) (*models.Applicationprofile, error) {
	rn := fmt.Sprintf("ap-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvAp := models.NewApplicationprofile(rn, parentDn, description, fvApattr)

    fvAp.Status = "modified"
	err := sm.Save(fvAp)
	return fvAp, err

}

func (sm *ServiceManager) ListApplicationprofile(tenant string ) ([]*models.Applicationprofile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvAp.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ApplicationprofileListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsApMonPolFromApplicationprofile( parentDn, tnMonEPGPolName string) error {
	dn := fmt.Sprintf("%s/rsApMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonEPGPolName": "%s"
								
			}
		}
	}`, "fvRsApMonPol", dn,tnMonEPGPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsApMonPolFromApplicationprofile(parentDn string) error{
	dn := fmt.Sprintf("%s/rsApMonPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsApMonPol")
}

