package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudendpointselector(name string ,cloud_epg string ,cloud_application_container string ,tenant string , description string, cloudEPSelectorattr models.CloudendpointselectorAttributes) (*models.Cloudendpointselector, error) {	
	rn := fmt.Sprintf("epselector-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s", tenant ,cloud_application_container ,cloud_epg )
	cloudEPSelector := models.NewCloudendpointselector(rn, parentDn, description, cloudEPSelectorattr)
	err := sm.Save(cloudEPSelector)
	return cloudEPSelector, err
}

func (sm *ServiceManager) ReadCloudendpointselector(name string ,cloud_epg string ,cloud_application_container string ,tenant string ) (*models.Cloudendpointselector, error) {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s/epselector-%s", tenant ,cloud_application_container ,cloud_epg ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudEPSelector := models.CloudendpointselectorFromContainer(cont)
	return cloudEPSelector, nil
}

func (sm *ServiceManager) DeleteCloudendpointselector(name string ,cloud_epg string ,cloud_application_container string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s/epselector-%s", tenant ,cloud_application_container ,cloud_epg ,name )
	return sm.DeleteByDn(dn, models.CloudepselectorClassName)
}

func (sm *ServiceManager) UpdateCloudendpointselector(name string ,cloud_epg string ,cloud_application_container string ,tenant string  ,description string, cloudEPSelectorattr models.CloudendpointselectorAttributes) (*models.Cloudendpointselector, error) {
	rn := fmt.Sprintf("epselector-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s", tenant ,cloud_application_container ,cloud_epg )
	cloudEPSelector := models.NewCloudendpointselector(rn, parentDn, description, cloudEPSelectorattr)

    cloudEPSelector.Status = "modified"
	err := sm.Save(cloudEPSelector)
	return cloudEPSelector, err

}

func (sm *ServiceManager) ListCloudendpointselector(cloud_epg string ,cloud_application_container string ,tenant string ) ([]*models.Cloudendpointselector, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudapp-%s/cloudepg-%s/cloudEPSelector.json", baseurlStr , tenant ,cloud_application_container ,cloud_epg )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudendpointselectorListFromContainer(cont)

	return list, err
}


