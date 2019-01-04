package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateVMMDomain(name string ,provider_profile_vendor string  ,description string, vmmDomPattr models.VMMDomainAttributes) (*models.VMMDomain, error) {	
	rn := fmt.Sprintf("dom-%s",name)
	parentDn := fmt.Sprintf("uni/vmmp-%s", provider_profile_vendor )
	vmmDomP := models.NewVMMDomain(rn, parentDn, description, vmmDomPattr)
	err := sm.Save(vmmDomP)
	return vmmDomP, err
}

func (sm *ServiceManager) ReadVMMDomain(name string ,provider_profile_vendor string ) (*models.VMMDomain, error) {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vmmDomP := models.VMMDomainFromContainer(cont)
	return vmmDomP, nil
}

func (sm *ServiceManager) DeleteVMMDomain(name string ,provider_profile_vendor string ) error {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor ,name )
	return sm.DeleteByDn(dn, models.VmmdompClassName)
}

func (sm *ServiceManager) UpdateVMMDomain(name string ,provider_profile_vendor string  ,description string, vmmDomPattr models.VMMDomainAttributes) (*models.VMMDomain, error) {
	rn := fmt.Sprintf("dom-%s",name)
	parentDn := fmt.Sprintf("uni/vmmp-%s", provider_profile_vendor )
	vmmDomP := models.NewVMMDomain(rn, parentDn, description, vmmDomPattr)

    vmmDomP.Status = "modified"
	err := sm.Save(vmmDomP)
	return vmmDomP, err

}

func (sm *ServiceManager) ListVMMDomain(provider_profile_vendor string ) ([]*models.VMMDomain, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/vmmp-%s/vmmDomP.json", baseurlStr , provider_profile_vendor )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.VMMDomainListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationinfraRsVlanNsFromVMMDomain( parentDn, tnFvnsVlanInstPName string) error {
	dn := fmt.Sprintf("%s/rsvlanNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsVlanInstPName": "%s"
								
			}
		}
	}`, "infraRsVlanNs", dn,tnFvnsVlanInstPName))

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

func (sm *ServiceManager) DeleteRelationinfraRsVlanNsFromVMMDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsvlanNs", parentDn)
	return sm.DeleteByDn(dn , "infraRsVlanNs")
}
func (sm *ServiceManager) CreateRelationvmmRsDomMcastAddrNsFromVMMDomain( parentDn, tnFvnsMcastAddrInstPName string) error {
	dn := fmt.Sprintf("%s/rsdomMcastAddrNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsMcastAddrInstPName": "%s"
								
			}
		}
	}`, "vmmRsDomMcastAddrNs", dn,tnFvnsMcastAddrInstPName))

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

func (sm *ServiceManager) DeleteRelationvmmRsDomMcastAddrNsFromVMMDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsdomMcastAddrNs", parentDn)
	return sm.DeleteByDn(dn , "vmmRsDomMcastAddrNs")
}
func (sm *ServiceManager) CreateRelationvmmRsDefaultCdpIfPolFromVMMDomain( parentDn, tnCdpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultCdpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCdpIfPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultCdpIfPol", dn,tnCdpIfPolName))

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
func (sm *ServiceManager) CreateRelationvmmRsDefaultLacpLagPolFromVMMDomain( parentDn, tnLacpLagPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultLacpLagPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnLacpLagPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultLacpLagPol", dn,tnLacpLagPolName))

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
func (sm *ServiceManager) CreateRelationinfraRsVlanNsDefFromVMMDomain( parentDn, tnFvnsAInstPName string) error {
	dn := fmt.Sprintf("%s/rsvlanNsDef", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAInstPName": "%s"
								
			}
		}
	}`, "infraRsVlanNsDef", dn,tnFvnsAInstPName))

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
func (sm *ServiceManager) CreateRelationinfraRsVipAddrNsFromVMMDomain( parentDn, tnFvnsAddrInstName string) error {
	dn := fmt.Sprintf("%s/rsvipAddrNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAddrInstName": "%s"
								
			}
		}
	}`, "infraRsVipAddrNs", dn,tnFvnsAddrInstName))

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

func (sm *ServiceManager) DeleteRelationinfraRsVipAddrNsFromVMMDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsvipAddrNs", parentDn)
	return sm.DeleteByDn(dn , "infraRsVipAddrNs")
}
func (sm *ServiceManager) CreateRelationvmmRsDefaultLldpIfPolFromVMMDomain( parentDn, tnLldpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultLldpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnLldpIfPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultLldpIfPol", dn,tnLldpIfPolName))

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
func (sm *ServiceManager) CreateRelationvmmRsDefaultL2InstPolFromVMMDomain( parentDn, tnL2InstPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultL2InstPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnL2InstPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultL2InstPol", dn,tnL2InstPolName))

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
func (sm *ServiceManager) CreateRelationvmmRsDefaultStpIfPolFromVMMDomain( parentDn, tnStpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultStpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnStpIfPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultStpIfPol", dn,tnStpIfPolName))

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
func (sm *ServiceManager) CreateRelationinfraRsDomVxlanNsDefFromVMMDomain( parentDn, tnFvnsAInstPName string) error {
	dn := fmt.Sprintf("%s/rsdomVxlanNsDef", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAInstPName": "%s"
								
			}
		}
	}`, "infraRsDomVxlanNsDef", dn,tnFvnsAInstPName))

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
func (sm *ServiceManager) CreateRelationvmmRsDefaultFwPolFromVMMDomain( parentDn, tnNwsFwPolName string) error {
	dn := fmt.Sprintf("%s/rsdefaultFwPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnNwsFwPolName": "%s"
								
			}
		}
	}`, "vmmRsDefaultFwPol", dn,tnNwsFwPolName))

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

