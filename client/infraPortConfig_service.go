package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateInfraPortConfiguration(subPort string, port string, card string, node string, description string, infraPortConfigAttr models.InfraPortConfigurationAttributes) (*models.InfraPortConfiguration, error) {
	rn := fmt.Sprintf(models.RninfraPortConfig, node, card, port, subPort)
	infraPortConfig := models.NewInfraPortConfiguration(rn, models.ParentDninfraPortConfig, description, infraPortConfigAttr)
	err := sm.Save(infraPortConfig)
	return infraPortConfig, err
}

func (sm *ServiceManager) ReadInfraPortConfiguration(subPort string, port string, card string, node string) (*models.InfraPortConfiguration, error) {
	rn := fmt.Sprintf(models.RninfraPortConfig, node, card, port, subPort)
	dn := fmt.Sprintf("%s/%s", models.ParentDninfraPortConfig, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	infraPortConfig := models.InfraPortConfigurationFromContainer(cont)
	return infraPortConfig, nil
}

func (sm *ServiceManager) DeleteInfraPortConfiguration(subPort string, port string, card string, node string) error {
	rn := fmt.Sprintf(models.RninfraPortConfig, node, card, port, subPort)
	dn := fmt.Sprintf("%s/%s", models.ParentDninfraPortConfig, rn)
	return sm.DeleteByDn(dn, models.InfraportconfigClassName)
}

func (sm *ServiceManager) UpdateInfraPortConfiguration(subPort string, port string, card string, node string, description string, infraPortConfigAttr models.InfraPortConfigurationAttributes) (*models.InfraPortConfiguration, error) {
	rn := fmt.Sprintf(models.RninfraPortConfig, node, card, port, subPort)
	infraPortConfig := models.NewInfraPortConfiguration(rn, models.ParentDninfraPortConfig, description, infraPortConfigAttr)
	infraPortConfig.Status = "modified"
	err := sm.Save(infraPortConfig)
	return infraPortConfig, err
}

func (sm *ServiceManager) ListInfraPortConfiguration() ([]*models.InfraPortConfiguration, error) {
	dnUrl := fmt.Sprintf("%s/uni/infra/infraPortConfig.json", models.BaseurlStr)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.InfraPortConfigurationListFromContainer(cont)
	return list, err
}
