package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateBridgedomain(name string ,tenant string , description string, fvBDattr models.BridgedomainAttributes) (*models.Bridgedomain, error) {	
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgedomain(rn, parentDn, description, fvBDattr)
	err := sm.Save(fvBD)
	return fvBD, err
}

func (sm *ServiceManager) ReadBridgedomain(name string ,tenant string ) (*models.Bridgedomain, error) {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvBD := models.BridgedomainFromContainer(cont)
	return fvBD, nil
}

func (sm *ServiceManager) DeleteBridgedomain(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvbdClassName)
}

func (sm *ServiceManager) UpdateBridgedomain(name string ,tenant string  ,description string, fvBDattr models.BridgedomainAttributes) (*models.Bridgedomain, error) {
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgedomain(rn, parentDn, description, fvBDattr)

    fvBD.Status = "modified"
	err := sm.Save(fvBD)
	return fvBD, err

}

func (sm *ServiceManager) ListBridgedomain(tenant string ) ([]*models.Bridgedomain, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvBD.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.BridgedomainListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsBDToProfileFromBridgedomain( parentDn, tnRtctrlProfileName string) error {
	dn := fmt.Sprintf("%s/rsBDToProfile", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnRtctrlProfileName": "%s"
								
			}
		}
	}`, "fvRsBDToProfile", dn,tnRtctrlProfileName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToProfileFromBridgedomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToProfile", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToProfile")
}
func (sm *ServiceManager) CreateRelationfvRsMldsnFromBridgedomain( parentDn, tnMldSnoopPolName string) error {
	dn := fmt.Sprintf("%s/rsmldsn", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMldSnoopPolName": "%s"
								
			}
		}
	}`, "fvRsMldsn", dn,tnMldSnoopPolName))

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
func (sm *ServiceManager) CreateRelationfvRsABDPolMonPolFromBridgedomain( parentDn, tnMonEPGPolName string) error {
	dn := fmt.Sprintf("%s/rsABDPolMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonEPGPolName": "%s"
								
			}
		}
	}`, "fvRsABDPolMonPol", dn,tnMonEPGPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsABDPolMonPolFromBridgedomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsABDPolMonPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsABDPolMonPol")
}
func (sm *ServiceManager) CreateRelationfvRsBDToNdPFromBridgedomain( parentDn, tnNdIfPolName string) error {
	dn := fmt.Sprintf("%s/rsBDToNdP", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnNdIfPolName": "%s"
								
			}
		}
	}`, "fvRsBDToNdP", dn,tnNdIfPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBdFloodToFromBridgedomain( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsbdFloodTo-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBdFloodTo", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBdFloodToFromBridgedomain(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsbdFloodTo-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsBdFloodTo")
}
func (sm *ServiceManager) CreateRelationfvRsBDToFhsFromBridgedomain( parentDn, tnFhsBDPolName string) error {
	dn := fmt.Sprintf("%s/rsBDToFhs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFhsBDPolName": "%s"
								
			}
		}
	}`, "fvRsBDToFhs", dn,tnFhsBDPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToFhsFromBridgedomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToFhs", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToFhs")
}
func (sm *ServiceManager) CreateRelationfvRsBDToRelayPFromBridgedomain( parentDn, tnDhcpRelayPName string) error {
	dn := fmt.Sprintf("%s/rsBDToRelayP", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnDhcpRelayPName": "%s"
								
			}
		}
	}`, "fvRsBDToRelayP", dn,tnDhcpRelayPName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToRelayPFromBridgedomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToRelayP", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToRelayP")
}
func (sm *ServiceManager) CreateRelationfvRsCtxFromBridgedomain( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rsctx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "fvRsCtx", dn,tnFvCtxName))

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
func (sm *ServiceManager) CreateRelationfvRsBDToNetflowMonitorPolFromBridgedomain( parentDn, tnNetflowMonitorPolName,fltType string) error {
	dn := fmt.Sprintf("%s/rsBDToNetflowMonitorPol-[%s]-%s", parentDn, tnNetflowMonitorPolName,fltType)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBDToNetflowMonitorPol", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToNetflowMonitorPolFromBridgedomain(parentDn , tnNetflowMonitorPolName,fltType string) error{
	dn := fmt.Sprintf("%s/rsBDToNetflowMonitorPol-[%s]-%s", parentDn, tnNetflowMonitorPolName,fltType)
	return sm.DeleteByDn(dn , "fvRsBDToNetflowMonitorPol")
}
func (sm *ServiceManager) CreateRelationfvRsIgmpsnFromBridgedomain( parentDn, tnIgmpSnoopPolName string) error {
	dn := fmt.Sprintf("%s/rsigmpsn", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnIgmpSnoopPolName": "%s"
								
			}
		}
	}`, "fvRsIgmpsn", dn,tnIgmpSnoopPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBdToEpRetFromBridgedomain( parentDn, tnFvEpRetPolName string) error {
	dn := fmt.Sprintf("%s/rsbdToEpRet", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvEpRetPolName": "%s"
								
			}
		}
	}`, "fvRsBdToEpRet", dn,tnFvEpRetPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBDToOutFromBridgedomain( parentDn, tnL3extOutName string) error {
	dn := fmt.Sprintf("%s/rsBDToOut-%s", parentDn, tnL3extOutName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBDToOut", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToOutFromBridgedomain(parentDn , tnL3extOutName string) error{
	dn := fmt.Sprintf("%s/rsBDToOut-%s", parentDn, tnL3extOutName)
	return sm.DeleteByDn(dn , "fvRsBDToOut")
}

