package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudawsprovider(tenant string , description string, cloudAwsProviderattr models.CloudawsproviderAttributes) (*models.Cloudawsprovider, error) {	
	rn := fmt.Sprintf("awsprovider")
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudAwsProvider := models.NewCloudawsprovider(rn, parentDn, description, cloudAwsProviderattr)
	err := sm.Save(cloudAwsProvider)
	return cloudAwsProvider, err
}

func (sm *ServiceManager) ReadCloudawsprovider(tenant string ) (*models.Cloudawsprovider, error) {
	dn := fmt.Sprintf("uni/tn-%s/awsprovider", tenant )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudAwsProvider := models.CloudawsproviderFromContainer(cont)
	return cloudAwsProvider, nil
}

func (sm *ServiceManager) DeleteCloudawsprovider(tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/awsprovider", tenant )
	return sm.DeleteByDn(dn, models.CloudawsproviderClassName)
}

func (sm *ServiceManager) UpdateCloudawsprovider(tenant string  ,description string, cloudAwsProviderattr models.CloudawsproviderAttributes) (*models.Cloudawsprovider, error) {
	rn := fmt.Sprintf("awsprovider")
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudAwsProvider := models.NewCloudawsprovider(rn, parentDn, description, cloudAwsProviderattr)

    cloudAwsProvider.Status = "modified"
	err := sm.Save(cloudAwsProvider)
	return cloudAwsProvider, err

}

func (sm *ServiceManager) ListCloudawsprovider(tenant string ) ([]*models.Cloudawsprovider, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudAwsProvider.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudawsproviderListFromContainer(cont)

	return list, err
}


