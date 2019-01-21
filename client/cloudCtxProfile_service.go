package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateCloudcontextprofile(name string ,tenant string , description string, cloudCtxProfileattr models.CloudcontextprofileAttributes) (*models.Cloudcontextprofile, error) {	
	rn := fmt.Sprintf("ctxprofile-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudCtxProfile := models.NewCloudcontextprofile(rn, parentDn, description, cloudCtxProfileattr)
	err := sm.Save(cloudCtxProfile)
	return cloudCtxProfile, err
}

func (sm *ServiceManager) ReadCloudcontextprofile(name string ,tenant string ) (*models.Cloudcontextprofile, error) {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCtxProfile := models.CloudcontextprofileFromContainer(cont)
	return cloudCtxProfile, nil
}

func (sm *ServiceManager) DeleteCloudcontextprofile(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.CloudctxprofileClassName)
}

func (sm *ServiceManager) UpdateCloudcontextprofile(name string ,tenant string  ,description string, cloudCtxProfileattr models.CloudcontextprofileAttributes) (*models.Cloudcontextprofile, error) {
	rn := fmt.Sprintf("ctxprofile-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	cloudCtxProfile := models.NewCloudcontextprofile(rn, parentDn, description, cloudCtxProfileattr)

    cloudCtxProfile.Status = "modified"
	err := sm.Save(cloudCtxProfile)
	return cloudCtxProfile, err

}

func (sm *ServiceManager) ListCloudcontextprofile(tenant string ) ([]*models.Cloudcontextprofile, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudCtxProfile.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudcontextprofileListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationcloudRsCtxToFlowLogFromCloudcontextprofile( parentDn, tnCloudAwsFlowLogPolName string) error {
	dn := fmt.Sprintf("%s/rsctxToFlowLog", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudAwsFlowLogPolName": "%s"
								
			}
		}
	}`, "cloudRsCtxToFlowLog", dn,tnCloudAwsFlowLogPolName))

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

func (sm *ServiceManager) DeleteRelationcloudRsCtxToFlowLogFromCloudcontextprofile(parentDn string) error{
	dn := fmt.Sprintf("%s/rsctxToFlowLog", parentDn)
	return sm.DeleteByDn(dn , "cloudRsCtxToFlowLog")
}
func (sm *ServiceManager) CreateRelationcloudRsToCtxFromCloudcontextprofile( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rstoCtx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "cloudRsToCtx", dn,tnFvCtxName))

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
func (sm *ServiceManager) CreateRelationcloudRsCtxProfileToRegionFromCloudcontextprofile( parentDn, tnCloudRegionName string) error {
	dn := fmt.Sprintf("%s/rsctxProfileToRegion", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudRegionName": "%s"
								
			}
		}
	}`, "cloudRsCtxProfileToRegion", dn,tnCloudRegionName))

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

func (sm *ServiceManager) DeleteRelationcloudRsCtxProfileToRegionFromCloudcontextprofile(parentDn string) error{
	dn := fmt.Sprintf("%s/rsctxProfileToRegion", parentDn)
	return sm.DeleteByDn(dn , "cloudRsCtxProfileToRegion")
}

