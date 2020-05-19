package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"



	


)









func (sm *ServiceManager) CreateContract(imported_contract string ,tenant string , description string, vzRsIfattr models.ContractAttributes) (*models.Contract, error) {	
	rn := fmt.Sprintf("rsif")
	parentDn := fmt.Sprintf("uni/tn-%s/cif-%s", tenant ,imported_contract )
	vzRsIf := models.NewContract(rn, parentDn, description, vzRsIfattr)
	err := sm.Save(vzRsIf)
	return vzRsIf, err
}

func (sm *ServiceManager) ReadContract(imported_contract string ,tenant string ) (*models.Contract, error) {
	dn := fmt.Sprintf("uni/tn-%s/cif-%s/rsif", tenant ,imported_contract )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzRsIf := models.ContractFromContainer(cont)
	return vzRsIf, nil
}

func (sm *ServiceManager) DeleteContract(imported_contract string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cif-%s/rsif", tenant ,imported_contract )
	return sm.DeleteByDn(dn, models.VzrsifClassName)
}

func (sm *ServiceManager) UpdateContract(imported_contract string ,tenant string  ,description string, vzRsIfattr models.ContractAttributes) (*models.Contract, error) {
	rn := fmt.Sprintf("rsif")
	parentDn := fmt.Sprintf("uni/tn-%s/cif-%s", tenant ,imported_contract )
	vzRsIf := models.NewContract(rn, parentDn, description, vzRsIfattr)

    vzRsIf.Status = "modified"
	err := sm.Save(vzRsIf)
	return vzRsIf, err

}

func (sm *ServiceManager) ListContract(imported_contract string ,tenant string ) ([]*models.Contract, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cif-%s/vzRsIf.json", baseurlStr , tenant ,imported_contract )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ContractListFromContainer(cont)

	return list, err
}


