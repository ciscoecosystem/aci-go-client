package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateTenantSecurityDomain(name string, tenant string, description string, nameAlias string, aaaDomainRefAttr models.TenantSecurityDomainAttributes) (*models.TenantSecurityDomain, error) {
	rn := fmt.Sprintf(models.RnaaaDomainRef, name)
	parentDn := fmt.Sprintf(models.ParentDnaaaDomainRef, tenant)
	aaaDomainRef := models.NewTenantSecurityDomain(rn, parentDn, description, nameAlias, aaaDomainRefAttr)
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err
}

func (sm *ServiceManager) ReadTenantSecurityDomain(name string, tenant string) (*models.TenantSecurityDomain, error) {
	dn := fmt.Sprintf(models.DnaaaDomainRef, tenant, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	aaaDomainRef := models.TenantSecurityDomainFromContainer(cont)
	return aaaDomainRef, nil
}

func (sm *ServiceManager) DeleteTenantSecurityDomain(name string, tenant string) error {
	dn := fmt.Sprintf(models.DnaaaDomainRef, tenant, name)
	return sm.DeleteByDn(dn, models.AaadomainrefClassName)
}

func (sm *ServiceManager) UpdateTenantSecurityDomain(name string, tenant string, description string, nameAlias string, aaaDomainRefAttr models.TenantSecurityDomainAttributes) (*models.TenantSecurityDomain, error) {
	rn := fmt.Sprintf(models.RnaaaDomainRef, name)
	parentDn := fmt.Sprintf(models.ParentDnaaaDomainRef, tenant)
	aaaDomainRef := models.NewTenantSecurityDomain(rn, parentDn, description, nameAlias, aaaDomainRefAttr)
	aaaDomainRef.Status = "modified"
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err
}

func (sm *ServiceManager) ListTenantSecurityDomain(tenant string) ([]*models.TenantSecurityDomain, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/aaaDomainRef.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.TenantSecurityDomainListFromContainer(cont)
	return list, err
}
