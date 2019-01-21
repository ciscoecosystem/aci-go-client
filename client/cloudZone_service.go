package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudavailabilityzone(name string ,cloud_providers_region string ,cloud_provider_profile_vendor string , description string, cloudZoneattr models.CloudavailabilityzoneAttributes) (*models.Cloudavailabilityzone, error) {	
	rn := fmt.Sprintf("zone-%s",name)
	parentDn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s", cloud_provider_profile_vendor ,cloud_providers_region )
	cloudZone := models.NewCloudavailabilityzone(rn, parentDn, description, cloudZoneattr)
	err := sm.Save(cloudZone)
	return cloudZone, err
}

func (sm *ServiceManager) ReadCloudavailabilityzone(name string ,cloud_providers_region string ,cloud_provider_profile_vendor string ) (*models.Cloudavailabilityzone, error) {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s/zone-%s", cloud_provider_profile_vendor ,cloud_providers_region ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudZone := models.CloudavailabilityzoneFromContainer(cont)
	return cloudZone, nil
}

func (sm *ServiceManager) DeleteCloudavailabilityzone(name string ,cloud_providers_region string ,cloud_provider_profile_vendor string ) error {
	dn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s/zone-%s", cloud_provider_profile_vendor ,cloud_providers_region ,name )
	return sm.DeleteByDn(dn, models.CloudzoneClassName)
}

func (sm *ServiceManager) UpdateCloudavailabilityzone(name string ,cloud_providers_region string ,cloud_provider_profile_vendor string  ,description string, cloudZoneattr models.CloudavailabilityzoneAttributes) (*models.Cloudavailabilityzone, error) {
	rn := fmt.Sprintf("zone-%s",name)
	parentDn := fmt.Sprintf("uni/clouddomp/provp-%s/region-%s", cloud_provider_profile_vendor ,cloud_providers_region )
	cloudZone := models.NewCloudavailabilityzone(rn, parentDn, description, cloudZoneattr)

    cloudZone.Status = "modified"
	err := sm.Save(cloudZone)
	return cloudZone, err

}

func (sm *ServiceManager) ListCloudavailabilityzone(cloud_providers_region string ,cloud_provider_profile_vendor string ) ([]*models.Cloudavailabilityzone, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/clouddomp/provp-%s/region-%s/cloudZone.json", baseurlStr , cloud_provider_profile_vendor ,cloud_providers_region )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudavailabilityzoneListFromContainer(cont)

	return list, err
}


