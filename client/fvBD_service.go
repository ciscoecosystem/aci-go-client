package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateBridgeDomain(name string ,tenant string  ,description string, fvBDattr models.BridgeDomainAttributes) (*models.BridgeDomain, error) {	
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgeDomain(rn, parentDn, description, fvBDattr)
	err := sm.Save(fvBD)
	return fvBD, err
}

func (sm *ServiceManager) ReadBridgeDomain(name string ,tenant string ) (*models.BridgeDomain, error) {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvBD := models.BridgeDomainFromContainer(cont)
	return fvBD, nil
}

func (sm *ServiceManager) DeleteBridgeDomain(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvbdClassName)
}

func (sm *ServiceManager) UpdateBridgeDomain(name string ,tenant string  ,description string, fvBDattr models.BridgeDomainAttributes) (*models.BridgeDomain, error) {
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgeDomain(rn, parentDn, description, fvBDattr)

    fvBD.Status = "modified"
	err := sm.Save(fvBD)
	return fvBD, err

}

func (sm *ServiceManager) ListBridgeDomain(tenant string ) ([]*models.BridgeDomain, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvBD.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.BridgeDomainListFromContainer(cont)

	return list, err
}