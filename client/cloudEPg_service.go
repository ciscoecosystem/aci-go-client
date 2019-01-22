package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateCloudEPg(name string ,cloud_application_container string ,tenant string , description string, cloudEPgattr models.CloudEPgAttributes) (*models.CloudEPg, error) {	
	rn := fmt.Sprintf("cloudepg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,cloud_application_container )
	cloudEPg := models.NewCloudEPg(rn, parentDn, description, cloudEPgattr)
	err := sm.Save(cloudEPg)
	return cloudEPg, err
}

func (sm *ServiceManager) ReadCloudEPg(name string ,cloud_application_container string ,tenant string ) (*models.CloudEPg, error) {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s", tenant ,cloud_application_container ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudEPg := models.CloudEPgFromContainer(cont)
	return cloudEPg, nil
}

func (sm *ServiceManager) DeleteCloudEPg(name string ,cloud_application_container string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudepg-%s", tenant ,cloud_application_container ,name )
	return sm.DeleteByDn(dn, models.CloudepgClassName)
}

func (sm *ServiceManager) UpdateCloudEPg(name string ,cloud_application_container string ,tenant string  ,description string, cloudEPgattr models.CloudEPgAttributes) (*models.CloudEPg, error) {
	rn := fmt.Sprintf("cloudepg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,cloud_application_container )
	cloudEPg := models.NewCloudEPg(rn, parentDn, description, cloudEPgattr)

    cloudEPg.Status = "modified"
	err := sm.Save(cloudEPg)
	return cloudEPg, err

}

func (sm *ServiceManager) ListCloudEPg(cloud_application_container string ,tenant string ) ([]*models.CloudEPg, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudapp-%s/cloudEPg.json", baseurlStr , tenant ,cloud_application_container )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudEPgListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsSecInheritedFromCloudEPg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsSecInherited", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsSecInheritedFromCloudEPg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsSecInherited")
}
func (sm *ServiceManager) CreateRelationfvRsProvFromCloudEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProv", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsProvFromCloudEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsProv")
}
func (sm *ServiceManager) CreateRelationfvRsConsIfFromCloudEPg( parentDn, tnVzCPIfName string) error {
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsConsIf", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsConsIfFromCloudEPg(parentDn , tnVzCPIfName string) error{
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	return sm.DeleteByDn(dn , "fvRsConsIf")
}
func (sm *ServiceManager) CreateRelationfvRsCustQosPolFromCloudEPg( parentDn, tnQosCustomPolName string) error {
	dn := fmt.Sprintf("%s/rscustQosPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnQosCustomPolName": "%s"
								
			}
		}
	}`, "fvRsCustQosPol", dn,tnQosCustomPolName))

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
func (sm *ServiceManager) CreateRelationfvRsConsFromCloudEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCons", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsConsFromCloudEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsCons")
}
func (sm *ServiceManager) CreateRelationcloudRsCloudEPgCtxFromCloudEPg( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rsCloudEPgCtx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "cloudRsCloudEPgCtx", dn,tnFvCtxName))

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
func (sm *ServiceManager) CreateRelationfvRsProtByFromCloudEPg( parentDn, tnVzTabooName string) error {
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProtBy", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsProtByFromCloudEPg(parentDn , tnVzTabooName string) error{
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	return sm.DeleteByDn(dn , "fvRsProtBy")
}
func (sm *ServiceManager) CreateRelationfvRsIntraEpgFromCloudEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsIntraEpg", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsIntraEpgFromCloudEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsIntraEpg")
}

