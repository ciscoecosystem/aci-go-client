package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateNLBendpoint(parent_dn string, description string, nameAlias string, fvEpNlbAttr models.NLBendpointAttributes) (*models.NLBendpoint, error) {
	rn := fmt.Sprintf(models.RnfvEpNlb)
	fvEpNlb := models.NewNLBendpoint(rn, parent_dn, description, nameAlias, fvEpNlbAttr)
	err := sm.Save(fvEpNlb)
	return fvEpNlb, err
}

func (sm *ServiceManager) ReadNLBendpoint(parent_dn string) (*models.NLBendpoint, error) {
	dn := fmt.Sprintf("%s/%s", parent_dn, models.RnfvEpNlb)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvEpNlb := models.NLBendpointFromContainer(cont)
	return fvEpNlb, nil
}

func (sm *ServiceManager) DeleteNLBendpoint(parent_dn string) error {
	dn := fmt.Sprintf("%s/%s", parent_dn, models.RnfvEpNlb)
	return sm.DeleteByDn(dn, models.FvepnlbClassName)
}

func (sm *ServiceManager) UpdateNLBendpoint(parent_dn string, description string, nameAlias string, fvEpNlbAttr models.NLBendpointAttributes) (*models.NLBendpoint, error) {
	rn := fmt.Sprintf(models.RnfvEpNlb)
	fvEpNlb := models.NewNLBendpoint(rn, parent_dn, description, nameAlias, fvEpNlbAttr)
	fvEpNlb.Status = "modified"
	err := sm.Save(fvEpNlb)
	return fvEpNlb, err
}

func (sm *ServiceManager) ListNLBendpoint(parent_dn string) ([]*models.NLBendpoint, error) {
	dnUrl := fmt.Sprintf("%s/%s/fvEpNlb.json", models.BaseurlStr, parent_dn)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.NLBendpointListFromContainer(cont)
	return list, err
}
