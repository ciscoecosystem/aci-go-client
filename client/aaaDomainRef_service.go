package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateAaaDomainRef(name string, tenant string, description string, nameAlias string, aaaDomainRefAttr models.AaaDomainRefAttributes) (*models.AaaDomainRef, error) {
	rn := fmt.Sprintf(models.RnaaaDomainRef, name)
	parentDn := fmt.Sprintf(models.ParentDnaaaDomainRef, tenant)
	aaaDomainRef := models.NewAaaDomainRef(rn, parentDn, description, nameAlias, aaaDomainRefAttr)
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err
}

func (sm *ServiceManager) ReadAaaDomainRef(name string, tenant string) (*models.AaaDomainRef, error) {
	dn := fmt.Sprintf(models.DnaaaDomainRef, tenant, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	aaaDomainRef := models.AaaDomainRefFromContainer(cont)
	return aaaDomainRef, nil
}

func (sm *ServiceManager) DeleteAaaDomainRef(name string, tenant string) error {
	dn := fmt.Sprintf(models.DnaaaDomainRef, tenant, name)
	return sm.DeleteByDn(dn, models.AaadomainrefClassName)
}

func (sm *ServiceManager) UpdateAaaDomainRef(name string, tenant string, description string, nameAlias string, aaaDomainRefAttr models.AaaDomainRefAttributes) (*models.AaaDomainRef, error) {
	rn := fmt.Sprintf(models.RnaaaDomainRef, name)
	parentDn := fmt.Sprintf(models.ParentDnaaaDomainRef, tenant)
	aaaDomainRef := models.NewAaaDomainRef(rn, parentDn, description, nameAlias, aaaDomainRefAttr)
	aaaDomainRef.Status = "modified"
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err
}

func (sm *ServiceManager) ListAaaDomainRef(tenant string) ([]*models.AaaDomainRef, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/aaaDomainRef.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.AaaDomainRefListFromContainer(cont)
	return list, err
}
