package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateExternalProfile(l3_outside string, tenant string, description string, pimExtPAttr models.PIMExternalProfileAttributes) (*models.PIMExternalProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimExtP, tenant, l3_outside)
	pimExtP := models.NewPIMExternalProfile(models.RnPimExtP, parentDn, description, pimExtPAttr)

	err := sm.Save(pimExtP)
	return pimExtP, err
}

func (sm *ServiceManager) ReadExternalProfile(l3_outside string, tenant string) (*models.PIMExternalProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimExtP, tenant, l3_outside)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnPimExtP)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	pimExtP := models.PIMExternalProfileFromContainer(cont)
	return pimExtP, nil
}

func (sm *ServiceManager) DeleteExternalProfile(l3_outside string, tenant string) error {

	parentDn := fmt.Sprintf(models.ParentDnPimExtP, tenant, l3_outside)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnPimExtP)

	return sm.DeleteByDn(dn, models.PimExtPClassName)
}

func (sm *ServiceManager) UpdateExternalProfile(l3_outside string, tenant string, description string, pimExtPAttr models.PIMExternalProfileAttributes) (*models.PIMExternalProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimExtP, tenant, l3_outside)
	pimExtP := models.NewPIMExternalProfile(models.RnPimExtP, parentDn, description, pimExtPAttr)

	pimExtP.Status = "modified"
	err := sm.Save(pimExtP)
	return pimExtP, err
}

func (sm *ServiceManager) ListExternalProfile(l3_outside string, tenant string) ([]*models.PIMExternalProfile, error) {

	parentDn := fmt.Sprintf(models.ParentDnPimExtP, tenant, l3_outside)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.PimExtPClassName)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.PIMExternalProfileListFromContainer(cont)
	return list, err
}
