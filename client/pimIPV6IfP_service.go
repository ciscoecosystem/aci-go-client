package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreatePimIPv6InterfaceProfile(logical_interface_profile string, logical_node_profile string, l3_outside string, tenant string, description string, pimIPV6IfPAttr models.PimIPv6InterfaceProfileAttributes) (*models.PimIPv6InterfaceProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimIPV6IfP, tenant, l3_outside, logical_node_profile, logical_interface_profile)
	pimIPV6IfP := models.NewPimIPv6InterfaceProfile(models.RnPimIPV6IfP, parentDn, description, pimIPV6IfPAttr)

	err := sm.Save(pimIPV6IfP)
	return pimIPV6IfP, err
}

func (sm *ServiceManager) ReadPimIPv6InterfaceProfile(logical_interface_profile string, logical_node_profile string, l3_outside string, tenant string) (*models.PimIPv6InterfaceProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimIPV6IfP, tenant, l3_outside, logical_node_profile, logical_interface_profile)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnPimIPV6IfP)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	pimIPV6IfP := models.PimIPv6InterfaceProfileFromContainer(cont)
	return pimIPV6IfP, nil
}

func (sm *ServiceManager) DeletePimIPv6InterfaceProfile(logical_interface_profile string, logical_node_profile string, l3_outside string, tenant string) error {

	parentDn := fmt.Sprintf(models.ParentDnPimIPV6IfP, tenant, l3_outside, logical_node_profile, logical_interface_profile)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnPimIPV6IfP)

	return sm.DeleteByDn(dn, models.PimIPV6IfPClassName)
}

func (sm *ServiceManager) UpdatePimIPv6InterfaceProfile(logical_interface_profile string, logical_node_profile string, l3_outside string, tenant string, description string, pimIPV6IfPAttr models.PimIPv6InterfaceProfileAttributes) (*models.PimIPv6InterfaceProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimIPV6IfP, tenant, l3_outside, logical_node_profile, logical_interface_profile)
	pimIPV6IfP := models.NewPimIPv6InterfaceProfile(models.RnPimIPV6IfP, parentDn, description, pimIPV6IfPAttr)

	pimIPV6IfP.Status = "modified"
	err := sm.Save(pimIPV6IfP)
	return pimIPV6IfP, err
}

func (sm *ServiceManager) ListPimIPv6InterfaceProfile(logical_interface_profile string, logical_node_profile string, l3_outside string, tenant string) ([]*models.PimIPv6InterfaceProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimIPV6IfP, tenant, l3_outside, logical_node_profile, logical_interface_profile)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.PimIPV6IfPClassName)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.PimIPv6InterfaceProfileListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationPimIPv6RsIfPol(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsV6IfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"	
			}
		}
	}`, "pimRsV6IfPol", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationPimIPv6RsIfPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsV6IfPol", parentDn)
	return sm.DeleteByDn(dn, "pimRsV6IfPol")
}

func (sm *ServiceManager) ReadRelationPimIPv6RsIfPol(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "pimRsV6IfPol")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "pimRsV6IfPol")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}
