package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateContractInterface(tnVzCPIfName string, application_epg string, application_profile string, tenant string, description string, nameAlias string, fvRsConsIfAttr models.ContractInterfaceAttributes) (*models.ContractInterface, error) {
	rn := fmt.Sprintf(models.RnfvRsConsIf, tnVzCPIfName)
	parentDn := fmt.Sprintf(models.ParentDnfvRsConsIf, tenant, application_profile, application_epg)
	fvRsConsIf := models.NewContractInterface(rn, parentDn, description, nameAlias, fvRsConsIfAttr)
	err := sm.Save(fvRsConsIf)
	return fvRsConsIf, err
}

func (sm *ServiceManager) ReadContractInterface(tnVzCPIfName string, application_epg string, application_profile string, tenant string) (*models.ContractInterface, error) {
	dn := fmt.Sprintf(models.DnfvRsConsIf, tenant, application_profile, application_epg, tnVzCPIfName)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvRsConsIf := models.ContractInterfaceFromContainer(cont)
	return fvRsConsIf, nil
}

func (sm *ServiceManager) DeleteContractInterface(tnVzCPIfName string, application_epg string, application_profile string, tenant string) error {
	dn := fmt.Sprintf(models.DnfvRsConsIf, tenant, application_profile, application_epg, tnVzCPIfName)
	return sm.DeleteByDn(dn, models.FvrsconsifClassName)
}

func (sm *ServiceManager) UpdateContractInterface(tnVzCPIfName string, application_epg string, application_profile string, tenant string, description string, nameAlias string, fvRsConsIfAttr models.ContractInterfaceAttributes) (*models.ContractInterface, error) {
	rn := fmt.Sprintf(models.RnfvRsConsIf, tnVzCPIfName)
	parentDn := fmt.Sprintf(models.ParentDnfvRsConsIf, tenant, application_profile, application_epg)
	fvRsConsIf := models.NewContractInterface(rn, parentDn, description, nameAlias, fvRsConsIfAttr)
	fvRsConsIf.Status = "modified"
	err := sm.Save(fvRsConsIf)
	return fvRsConsIf, err
}

func (sm *ServiceManager) ListContractInterface(application_epg string, application_profile string, tenant string) ([]*models.ContractInterface, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ap-%s/epg-%s/fvRsConsIf.json", models.BaseurlStr, tenant, application_profile, application_epg)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.ContractInterfaceListFromContainer(cont)
	return list, err
}
