package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateApplicationepg(name string ,application_profile string ,tenant string , description string, fvAEPgattr models.ApplicationepgAttributes) (*models.Applicationepg, error) {	
	rn := fmt.Sprintf("epg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,application_profile )
	fvAEPg := models.NewApplicationepg(rn, parentDn, description, fvAEPgattr)
	err := sm.Save(fvAEPg)
	return fvAEPg, err
}

func (sm *ServiceManager) ReadApplicationepg(name string ,application_profile string ,tenant string ) (*models.Applicationepg, error) {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant ,application_profile ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAEPg := models.ApplicationepgFromContainer(cont)
	return fvAEPg, nil
}

func (sm *ServiceManager) DeleteApplicationepg(name string ,application_profile string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant ,application_profile ,name )
	return sm.DeleteByDn(dn, models.FvaepgClassName)
}

func (sm *ServiceManager) UpdateApplicationepg(name string ,application_profile string ,tenant string  ,description string, fvAEPgattr models.ApplicationepgAttributes) (*models.Applicationepg, error) {
	rn := fmt.Sprintf("epg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s", tenant ,application_profile )
	fvAEPg := models.NewApplicationepg(rn, parentDn, description, fvAEPgattr)

    fvAEPg.Status = "modified"
	err := sm.Save(fvAEPg)
	return fvAEPg, err

}

func (sm *ServiceManager) ListApplicationepg(application_profile string ,tenant string ) ([]*models.Applicationepg, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ap-%s/fvAEPg.json", baseurlStr , tenant ,application_profile )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ApplicationepgListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsBdFromApplicationepg( parentDn, tnFvBDName string) error {
	dn := fmt.Sprintf("%s/rsbd", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvBDName": "%s"
								
			}
		}
	}`, "fvRsBd", dn,tnFvBDName))

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
func (sm *ServiceManager) CreateRelationfvRsCustQosPolFromApplicationepg( parentDn, tnQosCustomPolName string) error {
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
func (sm *ServiceManager) CreateRelationfvRsDomAttFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsdomAtt-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsDomAtt", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsDomAttFromApplicationepg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsdomAtt-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsDomAtt")
}
func (sm *ServiceManager) CreateRelationfvRsFcPathAttFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsfcPathAtt-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsFcPathAtt", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsFcPathAttFromApplicationepg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsfcPathAtt-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsFcPathAtt")
}
func (sm *ServiceManager) CreateRelationfvRsProvFromApplicationepg( parentDn, tnVzBrCPName string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsProvFromApplicationepg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsProv")
}
func (sm *ServiceManager) CreateRelationfvRsGraphDefFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsgraphDef-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsGraphDef", dn))

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
func (sm *ServiceManager) CreateRelationfvRsConsIfFromApplicationepg( parentDn, tnVzCPIfName string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsConsIfFromApplicationepg(parentDn , tnVzCPIfName string) error{
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	return sm.DeleteByDn(dn , "fvRsConsIf")
}
func (sm *ServiceManager) CreateRelationfvRsSecInheritedFromApplicationepg( parentDn, tDn string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsSecInheritedFromApplicationepg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsSecInherited")
}
func (sm *ServiceManager) CreateRelationfvRsNodeAttFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsnodeAtt-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsNodeAtt", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsNodeAttFromApplicationepg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsnodeAtt-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsNodeAtt")
}
func (sm *ServiceManager) CreateRelationfvRsDppPolFromApplicationepg( parentDn, tnQosDppPolName string) error {
	dn := fmt.Sprintf("%s/rsdppPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnQosDppPolName": "%s"
								
			}
		}
	}`, "fvRsDppPol", dn,tnQosDppPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsDppPolFromApplicationepg(parentDn string) error{
	dn := fmt.Sprintf("%s/rsdppPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsDppPol")
}
func (sm *ServiceManager) CreateRelationfvRsConsFromApplicationepg( parentDn, tnVzBrCPName string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsConsFromApplicationepg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsCons")
}
func (sm *ServiceManager) CreateRelationfvRsProvDefFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsprovDef-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProvDef", dn))

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
func (sm *ServiceManager) CreateRelationfvRsTrustCtrlFromApplicationepg( parentDn, tnFhsTrustCtrlPolName string) error {
	dn := fmt.Sprintf("%s/rstrustCtrl", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFhsTrustCtrlPolName": "%s"
								
			}
		}
	}`, "fvRsTrustCtrl", dn,tnFhsTrustCtrlPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsTrustCtrlFromApplicationepg(parentDn string) error{
	dn := fmt.Sprintf("%s/rstrustCtrl", parentDn)
	return sm.DeleteByDn(dn , "fvRsTrustCtrl")
}
func (sm *ServiceManager) CreateRelationfvRsPathAttFromApplicationepg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rspathAtt-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsPathAtt", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsPathAttFromApplicationepg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rspathAtt-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsPathAtt")
}
func (sm *ServiceManager) CreateRelationfvRsProtByFromApplicationepg( parentDn, tnVzTabooName string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsProtByFromApplicationepg(parentDn , tnVzTabooName string) error{
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	return sm.DeleteByDn(dn , "fvRsProtBy")
}
func (sm *ServiceManager) CreateRelationfvRsAEPgMonPolFromApplicationepg( parentDn, tnMonEPGPolName string) error {
	dn := fmt.Sprintf("%s/rsAEPgMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonEPGPolName": "%s"
								
			}
		}
	}`, "fvRsAEPgMonPol", dn,tnMonEPGPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsAEPgMonPolFromApplicationepg(parentDn string) error{
	dn := fmt.Sprintf("%s/rsAEPgMonPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsAEPgMonPol")
}
func (sm *ServiceManager) CreateRelationfvRsIntraEpgFromApplicationepg( parentDn, tnVzBrCPName string) error {
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

func (sm *ServiceManager) DeleteRelationfvRsIntraEpgFromApplicationepg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsIntraEpg")
}

