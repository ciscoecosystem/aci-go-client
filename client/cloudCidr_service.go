package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"

)









func (sm *ServiceManager) CreateCloudcidrpool(addr string ,cloud_context_profile string ,tenant string , description string, cloudCidrattr models.CloudcidrpoolAttributes) (*models.Cloudcidrpool, error) {	
	rn := fmt.Sprintf("cidr-[%s]",addr)
	parentDn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant ,cloud_context_profile )
	cloudCidr := models.NewCloudcidrpool(rn, parentDn, description, cloudCidrattr)
	err := sm.Save(cloudCidr)
	return cloudCidr, err
}

func (sm *ServiceManager) ReadCloudcidrpool(addr string ,cloud_context_profile string ,tenant string ) (*models.Cloudcidrpool, error) {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]", tenant ,cloud_context_profile ,addr )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCidr := models.CloudcidrpoolFromContainer(cont)
	return cloudCidr, nil
}

func (sm *ServiceManager) DeleteCloudcidrpool(addr string ,cloud_context_profile string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]", tenant ,cloud_context_profile ,addr )
	return sm.DeleteByDn(dn, models.CloudcidrClassName)
}

func (sm *ServiceManager) UpdateCloudcidrpool(addr string ,cloud_context_profile string ,tenant string  ,description string, cloudCidrattr models.CloudcidrpoolAttributes) (*models.Cloudcidrpool, error) {
	rn := fmt.Sprintf("cidr-[%s]",addr)
	parentDn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant ,cloud_context_profile )
	cloudCidr := models.NewCloudcidrpool(rn, parentDn, description, cloudCidrattr)

    cloudCidr.Status = "modified"
	err := sm.Save(cloudCidr)
	return cloudCidr, err

}

func (sm *ServiceManager) ListCloudcidrpool(cloud_context_profile string ,tenant string ) ([]*models.Cloudcidrpool, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ctxprofile-%s/cloudCidr.json", baseurlStr , tenant ,cloud_context_profile )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudcidrpoolListFromContainer(cont)

	return list, err
}


