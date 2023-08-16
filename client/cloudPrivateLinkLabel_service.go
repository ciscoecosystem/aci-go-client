package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreatePrivateLinkLabel(name string, parentDn string, description string, cloudPrivateLinkLabelAttr models.PrivateLinkLabelAttributes) (*models.PrivateLinkLabel, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	cloudPrivateLinkLabel := models.NewPrivateLinkLabel(rn, parentDn, description, cloudPrivateLinkLabelAttr)

	err := sm.Save(cloudPrivateLinkLabel)
	return cloudPrivateLinkLabel, err
}

func (sm *ServiceManager) ReadPrivateLinkLabel(name string, parentDn string) (*models.PrivateLinkLabel, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudPrivateLinkLabel := models.PrivateLinkLabelFromContainer(cont)
	return cloudPrivateLinkLabel, nil
}

func (sm *ServiceManager) DeletePrivateLinkLabel(name string, parentDn string) error {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	return sm.DeleteByDn(dn, models.CloudPrivateLinkLabelClassName)
}

func (sm *ServiceManager) UpdatePrivateLinkLabel(name string, parentDn string, description string, cloudPrivateLinkLabelAttr models.PrivateLinkLabelAttributes) (*models.PrivateLinkLabel, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	cloudPrivateLinkLabel := models.NewPrivateLinkLabel(rn, parentDn, description, cloudPrivateLinkLabelAttr)

	cloudPrivateLinkLabel.Status = "modified"
	err := sm.Save(cloudPrivateLinkLabel)
	return cloudPrivateLinkLabel, err
}

func (sm *ServiceManager) ListPrivateLinkLabel(parentDn string) ([]*models.PrivateLinkLabel, error) {

	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudPrivateLinkLabelClassName)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.PrivateLinkLabelListFromContainer(cont)
	return list, err
}
