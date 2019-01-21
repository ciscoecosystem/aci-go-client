package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateContractsubject(name string ,contract string ,tenant string , description string, vzSubjattr models.ContractsubjectAttributes) (*models.Contractsubject, error) {	
	rn := fmt.Sprintf("subj-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant ,contract )
	vzSubj := models.NewContractsubject(rn, parentDn, description, vzSubjattr)
	err := sm.Save(vzSubj)
	return vzSubj, err
}

func (sm *ServiceManager) ReadContractsubject(name string ,contract string ,tenant string ) (*models.Contractsubject, error) {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant ,contract ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vzSubj := models.ContractsubjectFromContainer(cont)
	return vzSubj, nil
}

func (sm *ServiceManager) DeleteContractsubject(name string ,contract string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/brc-%s/subj-%s", tenant ,contract ,name )
	return sm.DeleteByDn(dn, models.VzsubjClassName)
}

func (sm *ServiceManager) UpdateContractsubject(name string ,contract string ,tenant string  ,description string, vzSubjattr models.ContractsubjectAttributes) (*models.Contractsubject, error) {
	rn := fmt.Sprintf("subj-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/brc-%s", tenant ,contract )
	vzSubj := models.NewContractsubject(rn, parentDn, description, vzSubjattr)

    vzSubj.Status = "modified"
	err := sm.Save(vzSubj)
	return vzSubj, err

}

func (sm *ServiceManager) ListContractsubject(contract string ,tenant string ) ([]*models.Contractsubject, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/brc-%s/vzSubj.json", baseurlStr , tenant ,contract )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.ContractsubjectListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationvzRsSubjGraphAttFromContractsubject( parentDn, tnVnsAbsGraphName string) error {
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

func (sm *ServiceManager) DeleteRelationvzRsSubjGraphAttFromContractsubject(parentDn string) error{
	dn := fmt.Sprintf("%s/rsSubjGraphAtt", parentDn)
	return sm.DeleteByDn(dn , "vzRsSubjGraphAtt")
}
func (sm *ServiceManager) CreateRelationvzRsSdwanPolFromContractsubject( parentDn, tnExtdevSDWanSlaPolName string) error {
	dn := fmt.Sprintf("%s/rsSdwanPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnExtdevSDWanSlaPolName": "%s"
								
			}
		}
	}`, "vzRsSdwanPol", dn,tnExtdevSDWanSlaPolName))

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

func (sm *ServiceManager) DeleteRelationvzRsSdwanPolFromContractsubject(parentDn string) error{
	dn := fmt.Sprintf("%s/rsSdwanPol", parentDn)
	return sm.DeleteByDn(dn , "vzRsSdwanPol")
}
func (sm *ServiceManager) CreateRelationvzRsSubjFiltAttFromContractsubject( parentDn, tnVzFilterName string) error {
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

func (sm *ServiceManager) DeleteRelationvzRsSubjFiltAttFromContractsubject(parentDn , tnVzFilterName string) error{
	dn := fmt.Sprintf("%s/rssubjFiltAtt-%s", parentDn, tnVzFilterName)
	return sm.DeleteByDn(dn , "vzRsSubjFiltAtt")
}

