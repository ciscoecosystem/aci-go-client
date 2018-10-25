package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateContractSubject(name string, contract string, tenant string, description string, vzSubjattr models.ContractSubjectAttributes) (*models.ContractSubject, error) {
	rn := fmt.Sprintf("subj-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant, contract)
	vzSubj := models.NewContractSubject(rn, parentDn, description, vzSubjattr)
	err := sm.Save(vzSubj)
	return vzSubj, err
}

func (sm *ServiceManager) ReadContractSubject(name string, contract string, tenant string) (*models.ContractSubject, error) {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant, contract, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzSubj := models.ContractSubjectFromContainer(cont)
	return vzSubj, nil
}

func (sm *ServiceManager) DeleteContractSubject(name string, contract string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant, contract, name)
	return sm.DeleteByDn(dn, models.VzsubjClassName)
}

func (sm *ServiceManager) UpdateContractSubject(name string, contract string, tenant string, description string, vzSubjattr models.ContractSubjectAttributes) (*models.ContractSubject, error) {
	rn := fmt.Sprintf("subj-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant, contract)
	vzSubj := models.NewContractSubject(rn, parentDn, description, vzSubjattr)

	vzSubj.Status = "modified"
	err := sm.Save(vzSubj)
	return vzSubj, err

}

func (sm *ServiceManager) ListContractSubject(contract string, tenant string) ([]*models.ContractSubject, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/brc-%s/vzSubj.json", baseurlStr, tenant, contract)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.ContractSubjectListFromContainer(cont)

	return list, err
}
