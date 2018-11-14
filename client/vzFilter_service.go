package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateFilter(name string, tenant string, description string, vzFilterattr models.FilterAttributes) (*models.Filter, error) {
	rn := fmt.Sprintf("flt-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant)
	vzFilter := models.NewFilter(rn, parentDn, description, vzFilterattr)
	err := sm.Save(vzFilter)
	return vzFilter, err
}

func (sm *ServiceManager) ReadFilter(name string, tenant string) (*models.Filter, error) {
	dn := fmt.Sprintf("uni/tn-%s/flt-%s", tenant, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzFilter := models.FilterFromContainer(cont)
	return vzFilter, nil
}

func (sm *ServiceManager) DeleteFilter(name string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/flt-%s", tenant, name)
	return sm.DeleteByDn(dn, models.VzfilterClassName)
}

func (sm *ServiceManager) UpdateFilter(name string, tenant string, description string, vzFilterattr models.FilterAttributes) (*models.Filter, error) {
	rn := fmt.Sprintf("flt-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant)
	vzFilter := models.NewFilter(rn, parentDn, description, vzFilterattr)

	vzFilter.Status = "modified"
	err := sm.Save(vzFilter)
	return vzFilter, err

}

func (sm *ServiceManager) ListFilter(tenant string) ([]*models.Filter, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/vzFilter.json", baseurlStr, tenant)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.FilterListFromContainer(cont)

	return list, err
}
