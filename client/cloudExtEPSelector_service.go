package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudendpointselectorforexternalepgs(name string ,cloud_external_epg string ,cloud_application_container string ,tenant string , description string, cloudExtEPSelectorattr models.CloudendpointselectorforexternalepgsAttributes) (*models.Cloudendpointselectorforexternalepgs, error) {	
	rn := fmt.Sprintf("extepselector-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s", tenant ,cloud_application_container ,cloud_external_epg )
	cloudExtEPSelector := models.NewCloudendpointselectorforexternalepgs(rn, parentDn, description, cloudExtEPSelectorattr)
	err := sm.Save(cloudExtEPSelector)
	return cloudExtEPSelector, err
}

func (sm *ServiceManager) ReadCloudendpointselectorforexternalepgs(name string ,cloud_external_epg string ,cloud_application_container string ,tenant string ) (*models.Cloudendpointselectorforexternalepgs, error) {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s/extepselector-%s", tenant ,cloud_application_container ,cloud_external_epg ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPSelector := models.CloudendpointselectorforexternalepgsFromContainer(cont)
	return cloudExtEPSelector, nil
}

func (sm *ServiceManager) DeleteCloudendpointselectorforexternalepgs(name string ,cloud_external_epg string ,cloud_application_container string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s/extepselector-%s", tenant ,cloud_application_container ,cloud_external_epg ,name )
	return sm.DeleteByDn(dn, models.CloudextepselectorClassName)
}

func (sm *ServiceManager) UpdateCloudendpointselectorforexternalepgs(name string ,cloud_external_epg string ,cloud_application_container string ,tenant string  ,description string, cloudExtEPSelectorattr models.CloudendpointselectorforexternalepgsAttributes) (*models.Cloudendpointselectorforexternalepgs, error) {
	rn := fmt.Sprintf("extepselector-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s", tenant ,cloud_application_container ,cloud_external_epg )
	cloudExtEPSelector := models.NewCloudendpointselectorforexternalepgs(rn, parentDn, description, cloudExtEPSelectorattr)

    cloudExtEPSelector.Status = "modified"
	err := sm.Save(cloudExtEPSelector)
	return cloudExtEPSelector, err

}

func (sm *ServiceManager) ListCloudendpointselectorforexternalepgs(cloud_external_epg string ,cloud_application_container string ,tenant string ) ([]*models.Cloudendpointselectorforexternalepgs, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudapp-%s/cloudextepg-%s/cloudExtEPSelector.json", baseurlStr , tenant ,cloud_application_container ,cloud_external_epg )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudendpointselectorforexternalepgsListFromContainer(cont)

	return list, err
}


