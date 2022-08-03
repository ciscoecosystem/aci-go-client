package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateTenanttoaccountassociation(tenant string, nameAlias string, fvRsCloudAccountAttr models.TenanttoCloudAccountAssociationAttributes) (*models.TenanttoCloudAccountAssociation, error) {
	rn := fmt.Sprintf(models.RnfvRsCloudAccount)
	parentDn := fmt.Sprintf(models.ParentDnfvRsCloudAccount, tenant)
	fvRsCloudAccount := models.NewTenanttoCloudAccountAssociation(rn, parentDn, nameAlias, fvRsCloudAccountAttr)
	err := sm.Save(fvRsCloudAccount)
	return fvRsCloudAccount, err
}

func (sm *ServiceManager) ReadTenanttoaccountassociation(tenant string) (*models.TenanttoCloudAccountAssociation, error) {
	dn := fmt.Sprintf(models.DnfvRsCloudAccount, tenant)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvRsCloudAccount := models.TenanttoCloudAccountAssociationFromContainer(cont)
	return fvRsCloudAccount, nil
}

func (sm *ServiceManager) DeleteTenanttoaccountassociation(tenant string) error {
	dn := fmt.Sprintf(models.DnfvRsCloudAccount, tenant)
	return sm.DeleteByDn(dn, models.FvrscloudaccountClassName)
}

func (sm *ServiceManager) UpdateTenanttoaccountassociation(tenant string, nameAlias string, fvRsCloudAccountAttr models.TenanttoCloudAccountAssociationAttributes) (*models.TenanttoCloudAccountAssociation, error) {
	rn := fmt.Sprintf(models.RnfvRsCloudAccount)
	parentDn := fmt.Sprintf(models.ParentDnfvRsCloudAccount, tenant)
	fvRsCloudAccount := models.NewTenanttoCloudAccountAssociation(rn, parentDn, nameAlias, fvRsCloudAccountAttr)
	fvRsCloudAccount.Status = "modified"
	err := sm.Save(fvRsCloudAccount)
	return fvRsCloudAccount, err
}

func (sm *ServiceManager) ListTenanttoaccountassociation(tenant string) ([]*models.TenanttoCloudAccountAssociation, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvRsCloudAccount.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.TenanttoCloudAccountAssociationListFromContainer(cont)
	return list, err
}
