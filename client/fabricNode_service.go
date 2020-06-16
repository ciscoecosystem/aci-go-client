package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) ReadFabricNode(pod int, node int) (*models.FabricNode, error) {
	dn := fmt.Sprintf("topology/pod-%d/node-%d", pod, node)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	fabricNode := models.FabricNodeFromContainer(cont)
	return fabricNode, nil
}

func (sm *ServiceManager) ListFabricNode() ([]*models.FabricNode, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/fabricNode.json", baseurlStr)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.FabricNodeListFromContainer(cont)
	return list, err
}
