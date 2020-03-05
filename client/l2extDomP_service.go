package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func (sm *ServiceManager) CreateL2DomainProfile(name string, description string, l2extDomPattr models.L2DomainProfileAttributes) (*models.L2DomainProfile, error) {
	rn := fmt.Sprintf("l2dom-%s", name)
	parentDn := fmt.Sprintf("uni")
	l2extDomP := models.NewL2DomainProfile(rn, parentDn, description, l2extDomPattr)
	err := sm.Save(l2extDomP)
	return l2extDomP, err
}

func (sm *ServiceManager) ReadL2DomainProfile(name string) (*models.L2DomainProfile, error) {
	dn := fmt.Sprintf("uni/l2dom-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	l2extDomP := models.L2DomainProfileFromContainer(cont)
	return l2extDomP, nil
}

func (sm *ServiceManager) DeleteL2DomainProfile(name string) error {
	dn := fmt.Sprintf("uni/l2dom-%s", name)
	return sm.DeleteByDn(dn, models.L2extdompClassName)
}

func (sm *ServiceManager) UpdateL2DomainProfile(name string, description string, l2extDomPattr models.L2DomainProfileAttributes) (*models.L2DomainProfile, error) {
	rn := fmt.Sprintf("l2dom-%s", name)
	parentDn := fmt.Sprintf("uni")
	l2extDomP := models.NewL2DomainProfile(rn, parentDn, description, l2extDomPattr)

	l2extDomP.Status = "modified"
	err := sm.Save(l2extDomP)
	return l2extDomP, err

}

func (sm *ServiceManager) ListL2DomainProfile() ([]*models.L2DomainProfile, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/l2extDomP.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.L2DomainProfileListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationinfraRsVlanNsFromL2DomainProfile(parentDn, tnFvnsVlanInstPName string) error {
	dn := fmt.Sprintf("%s/rsvlanNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsVlanInstPName": "%s"
								
			}
		}
	}`, "infraRsVlanNs", dn, tnFvnsVlanInstPName))

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

func (sm *ServiceManager) DeleteRelationinfraRsVlanNsFromL2DomainProfile(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvlanNs", parentDn)
	return sm.DeleteByDn(dn, "infraRsVlanNs")
}

func (sm *ServiceManager) ReadRelationinfraRsVlanNsFromL2DomainProfile(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "infraRsVlanNs")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "infraRsVlanNs")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvnsVlanInstPName")
		return dat, err
	} else {
		return nil, err
	}

}
func (sm *ServiceManager) CreateRelationinfraRsVlanNsDefFromL2DomainProfile(parentDn, tnFvnsAInstPName string) error {
	dn := fmt.Sprintf("%s/rsvlanNsDef", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAInstPName": "%s"
								
			}
		}
	}`, "infraRsVlanNsDef", dn, tnFvnsAInstPName))

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

func (sm *ServiceManager) ReadRelationinfraRsVlanNsDefFromL2DomainProfile(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "infraRsVlanNsDef")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "infraRsVlanNsDef")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvnsAInstPName")
		return dat, err
	} else {
		return nil, err
	}

}
func (sm *ServiceManager) CreateRelationinfraRsVipAddrNsFromL2DomainProfile(parentDn, tnFvnsAddrInstName string) error {
	dn := fmt.Sprintf("%s/rsvipAddrNs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAddrInstName": "%s"
								
			}
		}
	}`, "infraRsVipAddrNs", dn, tnFvnsAddrInstName))

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

func (sm *ServiceManager) DeleteRelationinfraRsVipAddrNsFromL2DomainProfile(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvipAddrNs", parentDn)
	return sm.DeleteByDn(dn, "infraRsVipAddrNs")
}

func (sm *ServiceManager) ReadRelationinfraRsVipAddrNsFromL2DomainProfile(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "infraRsVipAddrNs")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "infraRsVipAddrNs")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvnsAddrInstName")
		return dat, err
	} else {
		return nil, err
	}

}
func (sm *ServiceManager) CreateRelationextnwRsOutFromL2DomainProfile(parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsout-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "extnwRsOut", dn))

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

func (sm *ServiceManager) ReadRelationextnwRsOutFromL2DomainProfile(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "extnwRsOut")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "extnwRsOut")

	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList {
		dat := models.G(contItem, "tDn")
		st.Add(dat)
	}
	return st, err

}
func (sm *ServiceManager) CreateRelationinfraRsDomVxlanNsDefFromL2DomainProfile(parentDn, tnFvnsAInstPName string) error {
	dn := fmt.Sprintf("%s/rsdomVxlanNsDef", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvnsAInstPName": "%s"
								
			}
		}
	}`, "infraRsDomVxlanNsDef", dn, tnFvnsAInstPName))

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

func (sm *ServiceManager) ReadRelationinfraRsDomVxlanNsDefFromL2DomainProfile(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "infraRsDomVxlanNsDef")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "infraRsDomVxlanNsDef")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvnsAInstPName")
		return dat, err
	} else {
		return nil, err
	}

}
