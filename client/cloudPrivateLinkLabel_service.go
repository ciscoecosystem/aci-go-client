package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreatePrivateLinkLabelfortheserviceEPg(name string, parentDn string, description string, cloudPrivateLinkLabelAttr models.PrivateLinkLabelfortheserviceEPgAttributes) (*models.PrivateLinkLabelfortheserviceEPg, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	cloudPrivateLinkLabel := models.NewPrivateLinkLabelfortheserviceEPg(rn, parentDn, description, cloudPrivateLinkLabelAttr)

	err := sm.Save(cloudPrivateLinkLabel)
	return cloudPrivateLinkLabel, err
}

func (sm *ServiceManager) ReadPrivateLinkLabelfortheserviceEPg(name string, parentDn string) (*models.PrivateLinkLabelfortheserviceEPg, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudPrivateLinkLabel := models.PrivateLinkLabelfortheserviceEPgFromContainer(cont)
	return cloudPrivateLinkLabel, nil
}

func (sm *ServiceManager) DeletePrivateLinkLabelfortheserviceEPg(name string, parentDn string) error {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	return sm.DeleteByDn(dn, models.CloudPrivateLinkLabelClassName)
}

func (sm *ServiceManager) UpdatePrivateLinkLabelfortheserviceEPg(name string, parentDn string, description string, cloudPrivateLinkLabelAttr models.PrivateLinkLabelfortheserviceEPgAttributes) (*models.PrivateLinkLabelfortheserviceEPg, error) {

	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	cloudPrivateLinkLabel := models.NewPrivateLinkLabelfortheserviceEPg(rn, parentDn, description, cloudPrivateLinkLabelAttr)

	cloudPrivateLinkLabel.Status = "modified"
	err := sm.Save(cloudPrivateLinkLabel)
	return cloudPrivateLinkLabel, err
}

func (sm *ServiceManager) ListPrivateLinkLabelfortheserviceEPg(parentDn string) (*models.PrivateLinkLabelfortheserviceEPg, error) {

	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudPrivateLinkLabelClassName)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.PrivateLinkLabelfortheserviceEPgListFromContainer(cont)
	return list, err
}
