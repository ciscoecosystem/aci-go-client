package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateAdditionalconfigforregion(parentDn string, cloudtemplateRegionDetailAttr models.AdditionalconfigforregionAttributes) (*models.Additionalconfigforregion, error) {

	cloudtemplateRegionDetail := models.NewAdditionalconfigforregion(models.RnCloudtemplateRegionDetail, parentDn, cloudtemplateRegionDetailAttr)

	err := sm.Save(cloudtemplateRegionDetail)
	return cloudtemplateRegionDetail, err
}

func (sm *ServiceManager) ReadAdditionalconfigforregion(parentDn string) (*models.Additionalconfigforregion, error) {

	dn := fmt.Sprintf("%s/%s", parentDn, models.RnCloudtemplateRegionDetail)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudtemplateRegionDetail := models.AdditionalconfigforregionFromContainer(cont)
	return cloudtemplateRegionDetail, nil
}

func (sm *ServiceManager) DeleteAdditionalconfigforregion(parentDn string) error {

	dn := fmt.Sprintf("%s/%s", parentDn, models.RnCloudtemplateRegionDetail)

	return sm.DeleteByDn(dn, models.CloudtemplateRegionDetailClassName)
}

func (sm *ServiceManager) UpdateAdditionalconfigforregion(parentDn string, cloudtemplateRegionDetailAttr models.AdditionalconfigforregionAttributes) (*models.Additionalconfigforregion, error) {

	cloudtemplateRegionDetail := models.NewAdditionalconfigforregion(models.RnCloudtemplateRegionDetail, parentDn, cloudtemplateRegionDetailAttr)

	cloudtemplateRegionDetail.Status = "modified"
	err := sm.Save(cloudtemplateRegionDetail)
	return cloudtemplateRegionDetail, err
}

func (sm *ServiceManager) ListAdditionalconfigforregion(parentDn string) ([]*models.Additionalconfigforregion, error) {

	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudtemplateRegionDetailClassName)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.AdditionalconfigforregionListFromContainer(cont)
	return list, err
}
