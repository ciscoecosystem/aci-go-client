package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateFilterRelationship(tnVzFilterName string, contract_subject string, contract string, tenant string, vzRsFiltAttAttr models.FilterRelationshipAttributes) (*models.FilterRelationship, error) {
	rn := fmt.Sprintf(models.RnvzRsFiltAtt, tnVzFilterName)
	parentDn := fmt.Sprintf(models.ParentDnvzRsFiltAtt, tenant, contract, contract_subject)
	vzRsFiltAtt := models.NewFilterRelationship(rn, parentDn, vzRsFiltAttAttr)
	err := sm.Save(vzRsFiltAtt)
	return vzRsFiltAtt, err
}

func (sm *ServiceManager) ReadFilterRelationship(tnVzFilterName string, contract_subject string, contract string, tenant string) (*models.FilterRelationship, error) {
	dn := fmt.Sprintf(models.DnvzRsFiltAtt, contract_subject, tnVzFilterName)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzRsFiltAtt := models.FilterRelationshipFromContainer(cont)
	return vzRsFiltAtt, nil
}

func (sm *ServiceManager) DeleteFilterRelationship(tnVzFilterName string, contract_subject string, contract string, tenant string) error {
	dn := fmt.Sprintf(models.DnvzRsFiltAtt, contract_subject, tnVzFilterName)
	return sm.DeleteByDn(dn, models.VzrsfiltattClassName)
}

func (sm *ServiceManager) UpdateFilterRelationship(tnVzFilterName string, contract_subject string, contract string, tenant string, vzRsFiltAttAttr models.FilterRelationshipAttributes) (*models.FilterRelationship, error) {
	rn := fmt.Sprintf(models.RnvzRsFiltAtt, tnVzFilterName)
	parentDn := fmt.Sprintf(models.ParentDnvzRsFiltAtt, tenant, contract, contract_subject)
	vzRsFiltAtt := models.NewFilterRelationship(rn, parentDn, vzRsFiltAttAttr)
	vzRsFiltAtt.Status = "modified"
	err := sm.Save(vzRsFiltAtt)
	return vzRsFiltAtt, err
}

func (sm *ServiceManager) ListFilterRelationship(contract_subject string, contract string, tenant string) ([]*models.FilterRelationship, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/brc-%s/subj-%s/intmnl/vzRsFiltAtt.json", models.BaseurlStr, tenant, contract, contract_subject)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.FilterRelationshipListFromContainer(cont)
	return list, err
}
