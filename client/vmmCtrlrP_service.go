package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateVMMController(name string ,vmm_domain string ,provider_profile_vendor string  ,description string, vmmCtrlrPattr models.VMMControllerAttributes) (*models.VMMController, error) {	
	rn := fmt.Sprintf("ctrlr-%s",name)
	parentDn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor ,vmm_domain )
	vmmCtrlrP := models.NewVMMController(rn, parentDn, description, vmmCtrlrPattr)
	err := sm.Save(vmmCtrlrP)
	return vmmCtrlrP, err
}

func (sm *ServiceManager) ReadVMMController(name string ,vmm_domain string ,provider_profile_vendor string ) (*models.VMMController, error) {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s/ctrlr-%s", provider_profile_vendor ,vmm_domain ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vmmCtrlrP := models.VMMControllerFromContainer(cont)
	return vmmCtrlrP, nil
}

func (sm *ServiceManager) DeleteVMMController(name string ,vmm_domain string ,provider_profile_vendor string ) error {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s/ctrlr-%s", provider_profile_vendor ,vmm_domain ,name )
	return sm.DeleteByDn(dn, models.VmmctrlrpClassName)
}

func (sm *ServiceManager) UpdateVMMController(name string ,vmm_domain string ,provider_profile_vendor string  ,description string, vmmCtrlrPattr models.VMMControllerAttributes) (*models.VMMController, error) {
	rn := fmt.Sprintf("ctrlr-%s",name)
	parentDn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor ,vmm_domain )
	vmmCtrlrP := models.NewVMMController(rn, parentDn, description, vmmCtrlrPattr)

    vmmCtrlrP.Status = "modified"
	err := sm.Save(vmmCtrlrP)
	return vmmCtrlrP, err

}

func (sm *ServiceManager) ListVMMController(vmm_domain string ,provider_profile_vendor string ) ([]*models.VMMController, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/vmmp-%s/dom-%s/vmmCtrlrP.json", baseurlStr , provider_profile_vendor ,vmm_domain )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.VMMControllerListFromContainer(cont)

	return list, err
}


func (sm *ServiceManager) CreateRelationvmmRsMcastAddrNs( parentDn, tnFvnsMcastAddrInstPName string) error {
	dn := fmt.Sprintf("%s/rsmcastAddrNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsMcastAddrInstPName": "%s"
								
			}
		}
	}`, "vmmRsMcastAddrNs", dn,tnFvnsMcastAddrInstPName))

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

func (sm *ServiceManager) DeleteRelationvmmRsMcastAddrNs(parentDn string) error{
	dn := fmt.Sprintf("%s/rsmcastAddrNs", parentDn)
	return sm.DeleteByDn(dn , "vmmRsMcastAddrNs")
}
func (sm *ServiceManager) CreateRelationvmmRsAcc( parentDn, tnVmmUsrAccPName string) error {
	dn := fmt.Sprintf("%s/rsacc", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnVmmUsrAccPName": "%s"
								
			}
		}
	}`, "vmmRsAcc", dn,tnVmmUsrAccPName))

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

func (sm *ServiceManager) DeleteRelationvmmRsAcc(parentDn string) error{
	dn := fmt.Sprintf("%s/rsacc", parentDn)
	return sm.DeleteByDn(dn , "vmmRsAcc")
}
func (sm *ServiceManager) CreateRelationvmmRsVmmCtrlrP( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsvmmCtrlrP-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "vmmRsVmmCtrlrP", dn))

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

func (sm *ServiceManager) DeleteRelationvmmRsVmmCtrlrP(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsvmmCtrlrP-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "vmmRsVmmCtrlrP")
}
func (sm *ServiceManager) CreateRelationvmmRsVxlanNs( parentDn, tnFvnsVxlanInstPName string) error {
	dn := fmt.Sprintf("%s/rsvxlanNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsVxlanInstPName": "%s"
								
			}
		}
	}`, "vmmRsVxlanNs", dn,tnFvnsVxlanInstPName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVxlanNs(parentDn string) error{
	dn := fmt.Sprintf("%s/rsvxlanNs", parentDn)
	return sm.DeleteByDn(dn , "vmmRsVxlanNs")
}
func (sm *ServiceManager) CreateRelationvmmRsCtrlrPMonPol( parentDn, tnMonInfraPolName string) error {
	dn := fmt.Sprintf("%s/rsctrlrPMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonInfraPolName": "%s"
								
			}
		}
	}`, "vmmRsCtrlrPMonPol", dn,tnMonInfraPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsCtrlrPMonPol(parentDn string) error{
	dn := fmt.Sprintf("%s/rsctrlrPMonPol", parentDn)
	return sm.DeleteByDn(dn , "vmmRsCtrlrPMonPol")
}
func (sm *ServiceManager) CreateRelationvmmRsMgmtEPg( parentDn, tnFvEPgName string) error {
	dn := fmt.Sprintf("%s/rsmgmtEPg", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvEPgName": "%s"
								
			}
		}
	}`, "vmmRsMgmtEPg", dn,tnFvEPgName))

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

func (sm *ServiceManager) DeleteRelationvmmRsMgmtEPg(parentDn string) error{
	dn := fmt.Sprintf("%s/rsmgmtEPg", parentDn)
	return sm.DeleteByDn(dn , "vmmRsMgmtEPg")
}

