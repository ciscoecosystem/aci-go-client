package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateFabricPortConfiguration(subPort string, port string, card string, node string, description string, fabricPortConfigAttr models.FabricPortConfigurationAttributes) (*models.FabricPortConfiguration, error) {
	rn := fmt.Sprintf(models.RnfabricPortConfig, node, card, port, subPort)
	fabricPortConfig := models.NewFabricPortConfiguration(rn, models.ParentDnfabricPortConfig, description, fabricPortConfigAttr)
	err := sm.Save(fabricPortConfig)
	return fabricPortConfig, err
}

func (sm *ServiceManager) ReadFabricPortConfiguration(subPort string, port string, card string, node string) (*models.FabricPortConfiguration, error) {
	rn := fmt.Sprintf(models.RnfabricPortConfig, node, card, port, subPort)
	dn := fmt.Sprintf("%s/%s", models.ParentDnfabricPortConfig, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	fabricPortConfig := models.FabricPortConfigurationFromContainer(cont)
	return fabricPortConfig, nil
}

func (sm *ServiceManager) DeleteFabricPortConfiguration(subPort string, port string, card string, node string) error {
	rn := fmt.Sprintf(models.RnfabricPortConfig, node, card, port, subPort)
	dn := fmt.Sprintf("%s/%s", models.ParentDnfabricPortConfig, rn)
	return sm.DeleteByDn(dn, models.FabricportconfigClassName)
}

func (sm *ServiceManager) UpdateFabricPortConfiguration(subPort string, port string, card string, node string, description string, fabricPortConfigAttr models.FabricPortConfigurationAttributes) (*models.FabricPortConfiguration, error) {
	rn := fmt.Sprintf(models.RnfabricPortConfig, node, card, port, subPort)
	fabricPortConfig := models.NewFabricPortConfiguration(rn, models.ParentDnfabricPortConfig, description, fabricPortConfigAttr)
	fabricPortConfig.Status = "modified"
	err := sm.Save(fabricPortConfig)
	return fabricPortConfig, err
}

func (sm *ServiceManager) ListFabricPortConfiguration() ([]*models.FabricPortConfiguration, error) {
	dnUrl := fmt.Sprintf("%s/uni/fabric/fabricPortConfig.json", models.BaseurlStr)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.FabricPortConfigurationListFromContainer(cont)
	return list, err
}
