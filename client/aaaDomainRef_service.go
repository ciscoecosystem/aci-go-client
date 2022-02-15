package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateTenantSecurityDomain(name string, tenant string, description string, aaaDomainRefattr models.TenantSecurityDomainAttributes) (*models.TenantSecurityDomain, error) {
	rn := fmt.Sprintf("domain-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant)
	aaaDomainRef := models.NewTenantSecurityDomain(rn, parentDn, description, aaaDomainRefattr)
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err
}

func (sm *ServiceManager) ReadTenantSecurityDomain(name string, tenant string) (*models.TenantSecurityDomain, error) {
	dn := fmt.Sprintf("uni/tn-%s/domain-%s", tenant, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	aaaDomainRef := models.TenantSecurityDomainFromContainer(cont)
	return aaaDomainRef, nil
}

func (sm *ServiceManager) DeleteTenantSecurityDomain(name string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/domain-%s", tenant, name)
	return sm.DeleteByDn(dn, models.AaadomainrefClassName)
}

func (sm *ServiceManager) UpdateTenantSecurityDomain(name string, tenant string, description string, aaaDomainRefattr models.TenantSecurityDomainAttributes) (*models.TenantSecurityDomain, error) {
	rn := fmt.Sprintf("domain-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant)

	aaaDomainRef := models.NewTenantSecurityDomain(rn, parentDn, description, aaaDomainRefattr)

	aaaDomainRef.Status = "modified"
	err := sm.Save(aaaDomainRef)
	return aaaDomainRef, err

}

func (sm *ServiceManager) ListTenantSecurityDomain(name string, tenant string) ([]*models.TenantSecurityDomain, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/domain-%s", baseurlStr, tenant, name)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.TenantSecurityDomainListFromContainer(cont)

	return list, err
}
