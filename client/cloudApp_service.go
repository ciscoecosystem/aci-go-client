package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudapplicationcontainer(name string ,tenant string , description string, cloudAppattr models.CloudapplicationcontainerAttributes) (*models.Cloudapplicationcontainer, error) {	
	rn := fmt.Sprintf("cloudapp-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudApp := models.NewCloudapplicationcontainer(rn, parentDn, description, cloudAppattr)
	err := sm.Save(cloudApp)
	return cloudApp, err
}

func (sm *ServiceManager) ReadCloudapplicationcontainer(name string ,tenant string ) (*models.Cloudapplicationcontainer, error) {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudApp := models.CloudapplicationcontainerFromContainer(cont)
	return cloudApp, nil
}

func (sm *ServiceManager) DeleteCloudapplicationcontainer(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.CloudappClassName)
}

func (sm *ServiceManager) UpdateCloudapplicationcontainer(name string ,tenant string  ,description string, cloudAppattr models.CloudapplicationcontainerAttributes) (*models.Cloudapplicationcontainer, error) {
	rn := fmt.Sprintf("cloudapp-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudApp := models.NewCloudapplicationcontainer(rn, parentDn, description, cloudAppattr)

    cloudApp.Status = "modified"
	err := sm.Save(cloudApp)
	return cloudApp, err

}

func (sm *ServiceManager) ListCloudapplicationcontainer(tenant string ) ([]*models.Cloudapplicationcontainer, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudApp.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudapplicationcontainerListFromContainer(cont)

	return list, err
}


