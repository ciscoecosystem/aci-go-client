package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudprovidersregion(name string ,cloud_provider_profile_vendor string , description string, cloudRegionattr models.CloudprovidersregionAttributes) (*models.Cloudprovidersregion, error) {	
	rn := fmt.Sprintf("region-%s",name)
	parentDn := fmt.Sprintf("uni/clouddomp/provp-%s", cloud_provider_profile_vendor )
	cloudRegion := models.NewCloudprovidersregion(rn, parentDn, description, cloudRegionattr)
	err := sm.Save(cloudRegion)
	return cloudRegion, err
}

func (sm *ServiceManager) ReadCloudprovidersregion(name string ,cloud_provider_profile_vendor string ) (*models.Cloudprovidersregion, error) {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s", cloud_provider_profile_vendor ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudRegion := models.CloudprovidersregionFromContainer(cont)
	return cloudRegion, nil
}

func (sm *ServiceManager) DeleteCloudprovidersregion(name string ,cloud_provider_profile_vendor string ) error {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s", cloud_provider_profile_vendor ,name )
	return sm.DeleteByDn(dn, models.CloudregionClassName)
}

func (sm *ServiceManager) UpdateCloudprovidersregion(name string ,cloud_provider_profile_vendor string  ,description string, cloudRegionattr models.CloudprovidersregionAttributes) (*models.Cloudprovidersregion, error) {
	rn := fmt.Sprintf("region-%s",name)
	parentDn := fmt.Sprintf("uni/clouddomp/provp-%s", cloud_provider_profile_vendor )
	cloudRegion := models.NewCloudprovidersregion(rn, parentDn, description, cloudRegionattr)

    cloudRegion.Status = "modified"
	err := sm.Save(cloudRegion)
	return cloudRegion, err

}

func (sm *ServiceManager) ListCloudprovidersregion(cloud_provider_profile_vendor string ) ([]*models.Cloudprovidersregion, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/clouddomp/provp-%s/cloudRegion.json", baseurlStr , cloud_provider_profile_vendor )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudprovidersregionListFromContainer(cont)

	return list, err
}


