package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateAutonomoussystemprofile(description string, cloudBgpAsPattr models.AutonomoussystemprofileAttributes) (*models.Autonomoussystemprofile, error) {	
	rn := fmt.Sprintf("clouddomp/as")
	parentDn := fmt.Sprintf("uni")
	cloudBgpAsP := models.NewAutonomoussystemprofile(rn, parentDn, description, cloudBgpAsPattr)
	err := sm.Save(cloudBgpAsP)
	return cloudBgpAsP, err
}

func (sm *ServiceManager) ReadAutonomoussystemprofile() (*models.Autonomoussystemprofile, error) {
	dn := fmt.Sprintf("uni/clouddomp/as")    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudBgpAsP := models.AutonomoussystemprofileFromContainer(cont)
	return cloudBgpAsP, nil
}

func (sm *ServiceManager) DeleteAutonomoussystemprofile() error {
	dn := fmt.Sprintf("uni/clouddomp/as")
	return sm.DeleteByDn(dn, models.CloudbgpaspClassName)
}

func (sm *ServiceManager) UpdateAutonomoussystemprofile(description string, cloudBgpAsPattr models.AutonomoussystemprofileAttributes) (*models.Autonomoussystemprofile, error) {
	rn := fmt.Sprintf("clouddomp/as")
	parentDn := fmt.Sprintf("uni")
	cloudBgpAsP := models.NewAutonomoussystemprofile(rn, parentDn, description, cloudBgpAsPattr)

    cloudBgpAsP.Status = "modified"
	err := sm.Save(cloudBgpAsP)
	return cloudBgpAsP, err

}

func (sm *ServiceManager) ListAutonomoussystemprofile() ([]*models.Autonomoussystemprofile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/cloudBgpAsP.json", baseurlStr )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.AutonomoussystemprofileListFromContainer(cont)

	return list, err
}


