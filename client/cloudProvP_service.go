package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudproviderprofile(vendor string , description string, cloudProvPattr models.CloudproviderprofileAttributes) (*models.Cloudproviderprofile, error) {	
	rn := fmt.Sprintf("clouddomp/provp-%s",vendor)
	parentDn := fmt.Sprintf("uni")
	cloudProvP := models.NewCloudproviderprofile(rn, parentDn, description, cloudProvPattr)
	err := sm.Save(cloudProvP)
	return cloudProvP, err
}

func (sm *ServiceManager) ReadCloudproviderprofile(vendor string ) (*models.Cloudproviderprofile, error) {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s", vendor )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudProvP := models.CloudproviderprofileFromContainer(cont)
	return cloudProvP, nil
}

func (sm *ServiceManager) DeleteCloudproviderprofile(vendor string ) error {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s", vendor )
	return sm.DeleteByDn(dn, models.CloudprovpClassName)
}

func (sm *ServiceManager) UpdateCloudproviderprofile(vendor string  ,description string, cloudProvPattr models.CloudproviderprofileAttributes) (*models.Cloudproviderprofile, error) {
	rn := fmt.Sprintf("clouddomp/provp-%s",vendor)
	parentDn := fmt.Sprintf("uni")
	cloudProvP := models.NewCloudproviderprofile(rn, parentDn, description, cloudProvPattr)

    cloudProvP.Status = "modified"
	err := sm.Save(cloudProvP)
	return cloudProvP, err

}

func (sm *ServiceManager) ListCloudproviderprofile() ([]*models.Cloudproviderprofile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/cloudProvP.json", baseurlStr )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudproviderprofileListFromContainer(cont)

	return list, err
}


