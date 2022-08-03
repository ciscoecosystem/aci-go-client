package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateActiveDirectory(active_directory_id string, tenant string, description string, nameAlias string, cloudADAttr models.ActiveDirectoryAttributes) (*models.ActiveDirectory, error) {
	rn := fmt.Sprintf(models.RncloudAD, active_directory_id)
	parentDn := fmt.Sprintf(models.ParentDncloudAD, tenant)
	cloudAD := models.NewActiveDirectory(rn, parentDn, description, nameAlias, cloudADAttr)
	err := sm.Save(cloudAD)
	return cloudAD, err
}

func (sm *ServiceManager) ReadActiveDirectory(active_directory_id string, tenant string) (*models.ActiveDirectory, error) {
	dn := fmt.Sprintf(models.DncloudAD, tenant, active_directory_id)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudAD := models.ActiveDirectoryFromContainer(cont)
	return cloudAD, nil
}

func (sm *ServiceManager) DeleteActiveDirectory(active_directory_id string, tenant string) error {
	dn := fmt.Sprintf(models.DncloudAD, tenant, active_directory_id)
	return sm.DeleteByDn(dn, models.CloudadClassName)
}

func (sm *ServiceManager) UpdateActiveDirectory(active_directory_id string, tenant string, description string, nameAlias string, cloudADAttr models.ActiveDirectoryAttributes) (*models.ActiveDirectory, error) {
	rn := fmt.Sprintf(models.RncloudAD, active_directory_id)
	parentDn := fmt.Sprintf(models.ParentDncloudAD, tenant)
	cloudAD := models.NewActiveDirectory(rn, parentDn, description, nameAlias, cloudADAttr)
	cloudAD.Status = "modified"
	err := sm.Save(cloudAD)
	return cloudAD, err
}

func (sm *ServiceManager) ListActiveDirectory(tenant string) ([]*models.ActiveDirectory, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudAD.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.ActiveDirectoryListFromContainer(cont)
	return list, err
}
