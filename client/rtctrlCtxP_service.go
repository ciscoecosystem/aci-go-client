package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (sm *ServiceManager) CreateRouteControlContext(name string, route_control_profile string, tenant string, description string, nameAlias string, rtctrlCtxPAttr models.RouteControlContextAttributes) (*models.RouteControlContext, error) {
	rn := fmt.Sprintf(models.RnrtctrlCtxP, name)
	parentDn := fmt.Sprintf(models.ParentDnrtctrlCtxP, tenant, route_control_profile)
	rtctrlCtxP := models.NewRouteControlContext(rn, parentDn, description, nameAlias, rtctrlCtxPAttr)
	err := sm.Save(rtctrlCtxP)
	return rtctrlCtxP, err
}

func (sm *ServiceManager) ReadRouteControlContext(name string, route_control_profile string, tenant string) (*models.RouteControlContext, error) {
	dn := fmt.Sprintf(models.DnrtctrlCtxP, tenant, route_control_profile, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	rtctrlCtxP := models.RouteControlContextFromContainer(cont)
	return rtctrlCtxP, nil
}

func (sm *ServiceManager) DeleteRouteControlContext(name string, route_control_profile string, tenant string) error {
	dn := fmt.Sprintf(models.DnrtctrlCtxP, tenant, route_control_profile, name)
	return sm.DeleteByDn(dn, models.RtctrlctxpClassName)
}

func (sm *ServiceManager) UpdateRouteControlContext(name string, route_control_profile string, tenant string, description string, nameAlias string, rtctrlCtxPAttr models.RouteControlContextAttributes) (*models.RouteControlContext, error) {
	rn := fmt.Sprintf(models.RnrtctrlCtxP, name)
	parentDn := fmt.Sprintf(models.ParentDnrtctrlCtxP, tenant, route_control_profile)
	rtctrlCtxP := models.NewRouteControlContext(rn, parentDn, description, nameAlias, rtctrlCtxPAttr)
	rtctrlCtxP.Status = "modified"
	err := sm.Save(rtctrlCtxP)
	return rtctrlCtxP, err
}

func (sm *ServiceManager) ListRouteControlContext(route_control_profile string, tenant string) ([]*models.RouteControlContext, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/prof-%s/rtctrlCtxP.json", models.BaseurlStr, tenant, route_control_profile)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.RouteControlContextListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationrtctrlRsCtxPToSubjP(parentDn, annotation, tnRtctrlSubjPName string) error {
	dn := fmt.Sprintf("%s/rsctxPToSubjP-%s", parentDn, tnRtctrlSubjPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tnRtctrlSubjPName": "%s"
			}
		}
	}`, "rtctrlRsCtxPToSubjP", dn, annotation, tnRtctrlSubjPName))

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

func (sm *ServiceManager) DeleteRelationrtctrlRsCtxPToSubjP(parentDn, tnRtctrlSubjPName string) error {
	dn := fmt.Sprintf("%s/rsctxPToSubjP-%s", parentDn, tnRtctrlSubjPName)
	return sm.DeleteByDn(dn, "rtctrlRsCtxPToSubjP")
}

func (sm *ServiceManager) ReadRelationrtctrlRsCtxPToSubjP(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "rtctrlRsCtxPToSubjP")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "rtctrlRsCtxPToSubjP")

	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList {
		dat := models.G(contItem, "tnRtctrlSubjPName")
		st.Add(dat)
	}
	return st, err
}
