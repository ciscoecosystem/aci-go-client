package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateClouddomainprofile(description string, cloudDomPattr models.ClouddomainprofileAttributes) (*models.Clouddomainprofile, error) {	
	rn := fmt.Sprintf("clouddomp")
	parentDn := fmt.Sprintf("uni")
	cloudDomP := models.NewClouddomainprofile(rn, parentDn, description, cloudDomPattr)
	err := sm.Save(cloudDomP)
	return cloudDomP, err
}

func (sm *ServiceManager) ReadClouddomainprofile() (*models.Clouddomainprofile, error) {
	dn := fmt.Sprintf("uni/clouddomp")    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudDomP := models.ClouddomainprofileFromContainer(cont)
	return cloudDomP, nil
}

func (sm *ServiceManager) DeleteClouddomainprofile() error {
	dn := fmt.Sprintf("uni/clouddomp")
	return sm.DeleteByDn(dn, models.ClouddompClassName)
}

func (sm *ServiceManager) UpdateClouddomainprofile(description string, cloudDomPattr models.ClouddomainprofileAttributes) (*models.Clouddomainprofile, error) {
	rn := fmt.Sprintf("clouddomp")
	parentDn := fmt.Sprintf("uni")
	cloudDomP := models.NewClouddomainprofile(rn, parentDn, description, cloudDomPattr)

    cloudDomP.Status = "modified"
	err := sm.Save(cloudDomP)
	return cloudDomP, err

}

func (sm *ServiceManager) ListClouddomainprofile() ([]*models.Clouddomainprofile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/cloudDomP.json", baseurlStr )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ClouddomainprofileListFromContainer(cont)

	return list, err
}


