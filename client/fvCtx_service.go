package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateVRF(name string ,tenant string  ,description string, fvCtxattr models.VRFAttributes) (*models.VRF, error) {	
	rn := fmt.Sprintf("ctx-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvCtx := models.NewVRF(rn, parentDn, description, fvCtxattr)
	err := sm.Save(fvCtx)
	return fvCtx, err
}

func (sm *ServiceManager) ReadVRF(name string ,tenant string ) (*models.VRF, error) {
	dn := fmt.Sprintf("uni/tn-%s/ctx-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvCtx := models.VRFFromContainer(cont)
	return fvCtx, nil
}

func (sm *ServiceManager) DeleteVRF(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/ctx-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvctxClassName)
}

func (sm *ServiceManager) UpdateVRF(name string ,tenant string  ,description string, fvCtxattr models.VRFAttributes) (*models.VRF, error) {
	rn := fmt.Sprintf("ctx-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvCtx := models.NewVRF(rn, parentDn, description, fvCtxattr)

    fvCtx.Status = "modified"
	err := sm.Save(fvCtx)
	return fvCtx, err

}

func (sm *ServiceManager) ListVRF(tenant string ) ([]*models.VRF, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvCtx.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.VRFListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsOspfCtxPolFromVRF( parentDn, tnOspfCtxPolName string) error {
	dn := fmt.Sprintf("%s/rsospfCtxPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnOspfCtxPolName": "%s"
								
			}
		}
	}`, "fvRsOspfCtxPol", dn,tnOspfCtxPolName))

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
func (sm *ServiceManager) CreateRelationfvRsVrfValidationPolFromVRF( parentDn, tnL3extVrfValidationPolName string) error {
	dn := fmt.Sprintf("%s/rsvrfValidationPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnL3extVrfValidationPolName": "%s"
								
			}
		}
	}`, "fvRsVrfValidationPol", dn,tnL3extVrfValidationPolName))

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
func (sm *ServiceManager) CreateRelationfvRsCtxMcastToFromVRF( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsctxMcastTo-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCtxMcastTo", dn))

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
func (sm *ServiceManager) CreateRelationfvRsBgpCtxPolFromVRF( parentDn, tnBgpCtxPolName string) error {
	dn := fmt.Sprintf("%s/rsbgpCtxPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnBgpCtxPolName": "%s"
								
			}
		}
	}`, "fvRsBgpCtxPol", dn,tnBgpCtxPolName))

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
func (sm *ServiceManager) CreateRelationfvRsCtxToExtRouteTagPolFromVRF( parentDn, tnL3extRouteTagPolName string) error {
	dn := fmt.Sprintf("%s/rsctxToExtRouteTagPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnL3extRouteTagPolName": "%s"
								
			}
		}
	}`, "fvRsCtxToExtRouteTagPol", dn,tnL3extRouteTagPolName))

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
func (sm *ServiceManager) CreateRelationfvRsCtxToEigrpCtxAfPolFromVRF( parentDn, tnEigrpCtxAfPolName,af string) error {
	dn := fmt.Sprintf("%s/rsctxToEigrpCtxAfPol-[%s]-%s", parentDn, tnEigrpCtxAfPolName,af)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCtxToEigrpCtxAfPol", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsCtxToEigrpCtxAfPolFromVRF(parentDn , tnEigrpCtxAfPolName,af string) error{
	dn := fmt.Sprintf("%s/rsctxToEigrpCtxAfPol-[%s]-%s", parentDn, tnEigrpCtxAfPolName,af)
	return sm.DeleteByDn(dn , "fvRsCtxToEigrpCtxAfPol")
}
func (sm *ServiceManager) CreateRelationfvRsCtxToOspfCtxPolFromVRF( parentDn, tnOspfCtxPolName,af string) error {
	dn := fmt.Sprintf("%s/rsctxToOspfCtxPol-[%s]-%s", parentDn, tnOspfCtxPolName,af)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCtxToOspfCtxPol", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsCtxToOspfCtxPolFromVRF(parentDn , tnOspfCtxPolName,af string) error{
	dn := fmt.Sprintf("%s/rsctxToOspfCtxPol-[%s]-%s", parentDn, tnOspfCtxPolName,af)
	return sm.DeleteByDn(dn , "fvRsCtxToOspfCtxPol")
}
func (sm *ServiceManager) CreateRelationfvRsCtxToEpRetFromVRF( parentDn, tnFvEpRetPolName string) error {
	dn := fmt.Sprintf("%s/rsctxToEpRet", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvEpRetPolName": "%s"
								
			}
		}
	}`, "fvRsCtxToEpRet", dn,tnFvEpRetPolName))

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
func (sm *ServiceManager) CreateRelationfvRsCtxMonPolFromVRF( parentDn, tnMonEPGPolName string) error {
	dn := fmt.Sprintf("%s/rsCtxMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonEPGPolName": "%s"
								
			}
		}
	}`, "fvRsCtxMonPol", dn,tnMonEPGPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsCtxMonPolFromVRF(parentDn string) error{
	dn := fmt.Sprintf("%s/rsCtxMonPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsCtxMonPol")
}
func (sm *ServiceManager) CreateRelationfvRsCtxToBgpCtxAfPolFromVRF( parentDn, tnBgpCtxAfPolName,af string) error {
	dn := fmt.Sprintf("%s/rsctxToBgpCtxAfPol-[%s]-%s", parentDn, tnBgpCtxAfPolName,af)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCtxToBgpCtxAfPol", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsCtxToBgpCtxAfPolFromVRF(parentDn , tnBgpCtxAfPolName,af string) error{
	dn := fmt.Sprintf("%s/rsctxToBgpCtxAfPol-[%s]-%s", parentDn, tnBgpCtxAfPolName,af)
	return sm.DeleteByDn(dn , "fvRsCtxToBgpCtxAfPol")
}

