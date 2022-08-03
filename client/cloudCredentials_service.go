package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateAccessCredentialtomanagethecloudresources(name string, tenant string, description string, nameAlias string, cloudCredentialsAttr models.AccessCredentialtomanagethecloudresourcesAttributes) (*models.AccessCredentialtomanagethecloudresources, error) {
	rn := fmt.Sprintf(models.RncloudCredentials, name)
	parentDn := fmt.Sprintf(models.ParentDncloudCredentials, tenant)
	cloudCredentials := models.NewAccessCredentialtomanagethecloudresources(rn, parentDn, description, nameAlias, cloudCredentialsAttr)
	err := sm.Save(cloudCredentials)
	return cloudCredentials, err
}

func (sm *ServiceManager) ReadAccessCredentialtomanagethecloudresources(name string, tenant string) (*models.AccessCredentialtomanagethecloudresources, error) {
	dn := fmt.Sprintf(models.DncloudCredentials, tenant, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCredentials := models.AccessCredentialtomanagethecloudresourcesFromContainer(cont)
	return cloudCredentials, nil
}

func (sm *ServiceManager) DeleteAccessCredentialtomanagethecloudresources(name string, tenant string) error {
	dn := fmt.Sprintf(models.DncloudCredentials, tenant, name)
	return sm.DeleteByDn(dn, models.CloudcredentialsClassName)
}

func (sm *ServiceManager) UpdateAccessCredentialtomanagethecloudresources(name string, tenant string, description string, nameAlias string, cloudCredentialsAttr models.AccessCredentialtomanagethecloudresourcesAttributes) (*models.AccessCredentialtomanagethecloudresources, error) {
	rn := fmt.Sprintf(models.RncloudCredentials, name)
	parentDn := fmt.Sprintf(models.ParentDncloudCredentials, tenant)
	cloudCredentials := models.NewAccessCredentialtomanagethecloudresources(rn, parentDn, description, nameAlias, cloudCredentialsAttr)
	cloudCredentials.Status = "modified"
	err := sm.Save(cloudCredentials)
	return cloudCredentials, err
}

func (sm *ServiceManager) ListAccessCredentialtomanagethecloudresources(tenant string) ([]*models.AccessCredentialtomanagethecloudresources, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudCredentials.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.AccessCredentialtomanagethecloudresourcesListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationcloudRsAD(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsAD", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "cloudRsAD", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationcloudRsAD(parentDn string) error {
	dn := fmt.Sprintf("%s/rsAD", parentDn)
	return sm.DeleteByDn(dn, "cloudRsAD")
}

func (sm *ServiceManager) ReadRelationcloudRsAD(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsAD")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "cloudRsAD")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}
