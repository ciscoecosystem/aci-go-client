package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateContractSubject(name string ,contract string ,tenant string  ,description string, vzSubjattr models.ContractSubjectAttributes) (*models.ContractSubject, error) {	
	rn := fmt.Sprintf("subj-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant ,contract )
	vzSubj := models.NewContractSubject(rn, parentDn, description, vzSubjattr)
	err := sm.Save(vzSubj)
	return vzSubj, err
}

func (sm *ServiceManager) ReadContractSubject(name string ,contract string ,tenant string ) (*models.ContractSubject, error) {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant ,contract ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzSubj := models.ContractSubjectFromContainer(cont)
	return vzSubj, nil
}

func (sm *ServiceManager) DeleteContractSubject(name string ,contract string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant ,contract ,name )
	return sm.DeleteByDn(dn, models.VzsubjClassName)
}

func (sm *ServiceManager) UpdateContractSubject(name string ,contract string ,tenant string  ,description string, vzSubjattr models.ContractSubjectAttributes) (*models.ContractSubject, error) {
	rn := fmt.Sprintf("subj-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant ,contract )
	vzSubj := models.NewContractSubject(rn, parentDn, description, vzSubjattr)

    vzSubj.Status = "modified"
	err := sm.Save(vzSubj)
	return vzSubj, err

}

func (sm *ServiceManager) ListContractSubject(contract string ,tenant string ) ([]*models.ContractSubject, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/brc-%s/vzSubj.json", baseurlStr , tenant ,contract )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ContractSubjectListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationvzRsSubjGraphAtt( parentDn, tnVnsAbsGraphName string) error {
	dn := fmt.Sprintf("%s/rsSubjGraphAtt", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnVnsAbsGraphName": "%s"
								
			}
		}
	}`, "vzRsSubjGraphAtt", dn,tnVnsAbsGraphName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationvzRsSubjGraphAtt(parentDn string) error{
	dn := fmt.Sprintf("%s/rsSubjGraphAtt", parentDn)
	return sm.DeleteByDn(dn , "vzRsSubjGraphAtt")
}
func (sm *ServiceManager) CreateRelationvzRsSubjFiltAtt( parentDn, tnVzFilterName string) error {
	dn := fmt.Sprintf("%s/rssubjFiltAtt-%s", parentDn, tnVzFilterName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "vzRsSubjFiltAtt", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationvzRsSubjFiltAtt(parentDn , tnVzFilterName string) error{
	dn := fmt.Sprintf("%s/rssubjFiltAtt-%s", parentDn, tnVzFilterName)
	return sm.DeleteByDn(dn , "vzRsSubjFiltAtt")
}

