package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func (sm *ServiceManager) CreateEndpointSecurityGroup(name string ,application_profile string ,tenant string , description string, fvESgattr models.EndpointSecurityGroupAttributes) (*models.EndpointSecurityGroup, error) {	
	rn := fmt.Sprintf(models.Rn,name)
	parentDn := fmt.Sprintf(models.ParentDn, tenant ,application_profile )
	fvESg := models.NewEndpointSecurityGroup(rn, parentDn, description, fvESgattr)
	err := sm.Save(fvESg)
	return fvESg, err
}

func (sm *ServiceManager) ReadEndpointSecurityGroup(name string ,application_profile string ,tenant string ) (*models.EndpointSecurityGroup, error) {
	dn := fmt.Sprintf(models.Dn, tenant ,application_profile ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	fvESg := models.EndpointSecurityGroupFromContainer(cont)
	return fvESg, nil
}

func (sm *ServiceManager) DeleteEndpointSecurityGroup(name string ,application_profile string ,tenant string ) error {
	dn := fmt.Sprintf(models.Dn, tenant ,application_profile ,name )
	return sm.DeleteByDn(dn, models.FvesgClassName)
}

func (sm *ServiceManager) UpdateEndpointSecurityGroup(name string ,application_profile string ,tenant string  ,description string, fvESgattr models.EndpointSecurityGroupAttributes) (*models.EndpointSecurityGroup, error) {
	rn := fmt.Sprintf(models.Rn,name)
	parentDn := fmt.Sprintf(models.ParentDn, tenant ,application_profile )
	fvESg := models.NewEndpointSecurityGroup(rn, parentDn, description, fvESgattr)
    fvESg.Status = "modified"
	err := sm.Save(fvESg)
	return fvESg, err
}

func (sm *ServiceManager) ListEndpointSecurityGroup(application_profile string ,tenant string ) ([]*models.EndpointSecurityGroup, error) {
	dnUrl := fmt.Sprintf("%s/models.ParentDn/fvESg.json", models.BaseurlStr , tenant ,application_profile )
    cont, err := sm.GetViaURL(dnUrl)
	list := models.EndpointSecurityGroupListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsSecInheritedFromEndpointSecurityGroup( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsSecInherited", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsSecInheritedFromEndpointSecurityGroup(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsSecInherited")
}

func (sm *ServiceManager) ReadRelationfvRsSecInheritedFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsSecInherited")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsSecInherited")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tDn")
		st.Add(dat)
	}
	return st, err
}

func (sm *ServiceManager) CreateRelationfvRsProvFromEndpointSecurityGroup( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProv", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsProvFromEndpointSecurityGroup(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsProv")
}

func (sm *ServiceManager) ReadRelationfvRsProvFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsProv")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsProv")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
}

func (sm *ServiceManager) CreateRelationfvRsConsIfFromEndpointSecurityGroup( parentDn, tnVzCPIfName string) error {
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsConsIf", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsConsIfFromEndpointSecurityGroup(parentDn , tnVzCPIfName string) error{
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	return sm.DeleteByDn(dn , "fvRsConsIf")
}

func (sm *ServiceManager) ReadRelationfvRsConsIfFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsConsIf")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsConsIf")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzCPIfName")
		st.Add(dat)
	}
	return st, err
}

func (sm *ServiceManager) CreateRelationfvRsCustQosPolFromEndpointSecurityGroup( parentDn, tnQosCustomPolName string) error {
	dn := fmt.Sprintf("%s/rscustQosPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnQosCustomPolName": "%s"
								
			}
		}
	}`, "fvRsCustQosPol", dn,tnQosCustomPolName))
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

func (sm *ServiceManager) ReadRelationfvRsCustQosPolFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsCustQosPol")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsCustQosPol")
	if len(contList) > 0 {
		dat := models.G(contList[0], "tnQosCustomPolName")
		return dat, err
	} else {
		return nil,err
	}
}

func (sm *ServiceManager) CreateRelationfvRsConsFromEndpointSecurityGroup( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCons", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsConsFromEndpointSecurityGroup(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsCons")
}

func (sm *ServiceManager) ReadRelationfvRsConsFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsCons")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsCons")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
}

func (sm *ServiceManager) CreateRelationfvRsScopeFromEndpointSecurityGroup( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rsscope", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "fvRsScope", dn,tnFvCtxName))
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

func (sm *ServiceManager) ReadRelationfvRsScopeFromEndpointSecurityGroup( parentDn string) (interface{},error) {	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsScope")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsScope")
	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvCtxName")
		return dat, err
	} else {
		return nil,err
	}
}

func (sm *ServiceManager) CreateRelationfvRsProtByFromEndpointSecurityGroup( parentDn, tnVzTabooName string) error {
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProtBy", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsProtByFromEndpointSecurityGroup(parentDn , tnVzTabooName string) error{
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	return sm.DeleteByDn(dn , "fvRsProtBy")
}

func (sm *ServiceManager) ReadRelationfvRsProtByFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsProtBy")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsProtBy")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzTabooName")
		st.Add(dat)
	}
	return st, err
}

func (sm *ServiceManager) CreateRelationfvRsIntraEpgFromEndpointSecurityGroup( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsIntraEpg", dn))
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

func (sm *ServiceManager) DeleteRelationfvRsIntraEpgFromEndpointSecurityGroup(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsIntraEpg")
}

func (sm *ServiceManager) ReadRelationfvRsIntraEpgFromEndpointSecurityGroup( parentDn string) (interface{},error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json",models.BaseurlStr,parentDn,"fvRsIntraEpg")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont,"fvRsIntraEpg")
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
}

