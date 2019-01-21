package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateCloudsubnet(ip string ,cloud_cidr_pool_addr string ,cloud_context_profile string ,tenant string , description string, cloudSubnetattr models.CloudsubnetAttributes) (*models.Cloudsubnet, error) {	
	rn := fmt.Sprintf("subnet-[%s]",ip)
	parentDn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]", tenant ,cloud_context_profile ,cloud_cidr_pool_addr )
	cloudSubnet := models.NewCloudsubnet(rn, parentDn, description, cloudSubnetattr)
	err := sm.Save(cloudSubnet)
	return cloudSubnet, err
}

func (sm *ServiceManager) ReadCloudsubnet(ip string ,cloud_cidr_pool_addr string ,cloud_context_profile string ,tenant string ) (*models.Cloudsubnet, error) {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]/subnet-[%s]", tenant ,cloud_context_profile ,cloud_cidr_pool_addr ,ip )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudSubnet := models.CloudsubnetFromContainer(cont)
	return cloudSubnet, nil
}

func (sm *ServiceManager) DeleteCloudsubnet(ip string ,cloud_cidr_pool_addr string ,cloud_context_profile string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]/subnet-[%s]", tenant ,cloud_context_profile ,cloud_cidr_pool_addr ,ip )
	return sm.DeleteByDn(dn, models.CloudsubnetClassName)
}

func (sm *ServiceManager) UpdateCloudsubnet(ip string ,cloud_cidr_pool_addr string ,cloud_context_profile string ,tenant string  ,description string, cloudSubnetattr models.CloudsubnetAttributes) (*models.Cloudsubnet, error) {
	rn := fmt.Sprintf("subnet-[%s]",ip)
	parentDn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s/cidr-[%s]", tenant ,cloud_context_profile ,cloud_cidr_pool_addr )
	cloudSubnet := models.NewCloudsubnet(rn, parentDn, description, cloudSubnetattr)

    cloudSubnet.Status = "modified"
	err := sm.Save(cloudSubnet)
	return cloudSubnet, err

}

func (sm *ServiceManager) ListCloudsubnet(cloud_cidr_pool_addr string ,cloud_context_profile string ,tenant string ) ([]*models.Cloudsubnet, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ctxprofile-%s/cidr-[%s]/cloudSubnet.json", baseurlStr , tenant ,cloud_context_profile ,cloud_cidr_pool_addr )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudsubnetListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationcloudRsZoneAttachFromCloudsubnet( parentDn, tnCloudZoneName string) error {
	dn := fmt.Sprintf("%s/rszoneAttach", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudZoneName": "%s"
								
			}
		}
	}`, "cloudRsZoneAttach", dn,tnCloudZoneName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationcloudRsZoneAttachFromCloudsubnet(parentDn string) error{
	dn := fmt.Sprintf("%s/rszoneAttach", parentDn)
	return sm.DeleteByDn(dn , "cloudRsZoneAttach")
}
func (sm *ServiceManager) CreateRelationcloudRsSubnetToFlowLogFromCloudsubnet( parentDn, tnCloudAwsFlowLogPolName string) error {
	dn := fmt.Sprintf("%s/rssubnetToFlowLog", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudAwsFlowLogPolName": "%s"
								
			}
		}
	}`, "cloudRsSubnetToFlowLog", dn,tnCloudAwsFlowLogPolName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationcloudRsSubnetToFlowLogFromCloudsubnet(parentDn string) error{
	dn := fmt.Sprintf("%s/rssubnetToFlowLog", parentDn)
	return sm.DeleteByDn(dn , "cloudRsSubnetToFlowLog")
}

