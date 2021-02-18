package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateSpanningTreeInterfacePolicy(name string, description string, stpIfPolattr models.SpanningTreeInterfacePolicyAttributes) (*models.SpanningTreeInterfacePolicy, error) {
	rn := fmt.Sprintf("infra/ifPol-%s", name)
	parentDn := fmt.Sprintf("uni")
	stpIfPol := models.NewSpanningTreeInterfacePolicy(rn, parentDn, description, stpIfPolattr)
	err := sm.Save(stpIfPol)
	return stpIfPol, err
}

func (sm *ServiceManager) ReadSpanningTreeInterfacePolicy(name string) (*models.SpanningTreeInterfacePolicy, error) {
	dn := fmt.Sprintf("uni/infra/ifPol-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	stpIfPol := models.SpanningTreeInterfacePolicyFromContainer(cont)
	return stpIfPol, nil
}

func (sm *ServiceManager) DeleteSpanningTreeInterfacePolicy(name string) error {
	dn := fmt.Sprintf("uni/infra/ifPol-%s", name)
	return sm.DeleteByDn(dn, models.StpIfPolClassName)
}

func (sm *ServiceManager) UpdateSpanningTreeInterfacePolicy(name string, description string, stpIfPolattr models.SpanningTreeInterfacePolicyAttributes) (*models.SpanningTreeInterfacePolicy, error) {
	rn := fmt.Sprintf("infra/ifPol-%s", name)
	parentDn := fmt.Sprintf("uni")
	stpIfPol := models.NewSpanningTreeInterfacePolicy(rn, parentDn, description, stpIfPolattr)

	stpIfPol.Status = "modified"
	err := sm.Save(stpIfPol)
	return stpIfPol, err

}

func (sm *ServiceManager) ListSpanningTreeInterfacePolicy() ([]*models.SpanningTreeInterfacePolicy, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/stpIfPol.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanningTreeInterfacePolicyListFromContainer(cont)

	return list, err
}
