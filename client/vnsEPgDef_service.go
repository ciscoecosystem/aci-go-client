package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateEPgDef(name string, legvnode string, node_instance string, l4l7_service_graph_instance_scopeDn string, l4l7_service_graph_instance_graphDn string, l4l7_service_graph_instance_ctrctDn string, tenant string, description string, nameAlias string, vnsEPgDefAttr models.EPgDefAttributes) (*models.EPgDef, error) {
	rn := fmt.Sprintf(models.RnvnsEPgDef, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsEPgDef, tenant, l4l7_service_graph_instance_ctrctDn, l4l7_service_graph_instance_graphDn, l4l7_service_graph_instance_scopeDn, node_instance, legvnode)
	vnsEPgDef := models.NewEPgDef(rn, parentDn, description, nameAlias, vnsEPgDefAttr)
	err := sm.Save(vnsEPgDef)
	return vnsEPgDef, err
}

func (sm *ServiceManager) ReadEPgDef(name string, legvnode string, node_instance string, l4l7_service_graph_instance_scopeDn string, l4l7_service_graph_instance_graphDn string, l4l7_service_graph_instance_ctrctDn string, tenant string) (*models.EPgDef, error) {
	dn := fmt.Sprintf(models.DnvnsEPgDef, tenant, l4l7_service_graph_instance_ctrctDn, l4l7_service_graph_instance_graphDn, l4l7_service_graph_instance_scopeDn, node_instance, legvnode, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vnsEPgDef := models.EPgDefFromContainer(cont)
	return vnsEPgDef, nil
}

func (sm *ServiceManager) DeleteEPgDef(name string, legvnode string, node_instance string, l4l7_service_graph_instance_scopeDn string, l4l7_service_graph_instance_graphDn string, l4l7_service_graph_instance_ctrctDn string, tenant string) error {
	dn := fmt.Sprintf(models.DnvnsEPgDef, tenant, l4l7_service_graph_instance_ctrctDn, l4l7_service_graph_instance_graphDn, l4l7_service_graph_instance_scopeDn, node_instance, legvnode, name)
	return sm.DeleteByDn(dn, models.VnsepgdefClassName)
}

func (sm *ServiceManager) UpdateEPgDef(name string, legvnode string, node_instance string, l4l7_service_graph_instance_scopeDn string, l4l7_service_graph_instance_graphDn string, l4l7_service_graph_instance_ctrctDn string, tenant string, description string, nameAlias string, vnsEPgDefAttr models.EPgDefAttributes) (*models.EPgDef, error) {
	rn := fmt.Sprintf(models.RnvnsEPgDef, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsEPgDef, tenant, l4l7_service_graph_instance_ctrctDn, l4l7_service_graph_instance_graphDn, l4l7_service_graph_instance_scopeDn, node_instance, legvnode)
	vnsEPgDef := models.NewEPgDef(rn, parentDn, description, nameAlias, vnsEPgDefAttr)
	vnsEPgDef.Status = "modified"
	err := sm.Save(vnsEPgDef)
	return vnsEPgDef, err
}

func (sm *ServiceManager) ListEPgDef(legvnode string, node_instance string, l4l7_service_graph_instance_scopeDn string, l4l7_service_graph_instance_graphDn string, l4l7_service_graph_instance_ctrctDn string, tenant string) ([]*models.EPgDef, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/GraphInst_C-[%s]-G-[%s]-S-[%s]/NodeInst-%s/LegVNode-%s/vnsEPgDef.json", models.BaseurlStr, tenant, l4l7_service_graph_instance_ctrctDn, l4l7_service_graph_instance_graphDn, l4l7_service_graph_instance_scopeDn, node_instance, legvnode)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.EPgDefListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationvnsRsEPgDefToConn(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsEPgDefToConn", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "vnsRsEPgDefToConn", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationvnsRsEPgDefToConn(parentDn string) error {
	dn := fmt.Sprintf("%s/rsEPgDefToConn", parentDn)
	return sm.DeleteByDn(dn, "vnsRsEPgDefToConn")
}

func (sm *ServiceManager) ReadRelationvnsRsEPgDefToConn(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsEPgDefToConn")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "vnsRsEPgDefToConn")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}

func (sm *ServiceManager) CreateRelationvnsRsEPgDefToLIf(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsEPgDefToLIf", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "vnsRsEPgDefToLIf", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationvnsRsEPgDefToLIf(parentDn string) error {
	dn := fmt.Sprintf("%s/rsEPgDefToLIf", parentDn)
	return sm.DeleteByDn(dn, "vnsRsEPgDefToLIf")
}

func (sm *ServiceManager) ReadRelationvnsRsEPgDefToLIf(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsEPgDefToLIf")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "vnsRsEPgDefToLIf")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}

func (sm *ServiceManager) CreateRelationvnsRsEPpInfoAtt(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsePpInfoAtt", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "vnsRsEPpInfoAtt", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationvnsRsEPpInfoAtt(parentDn string) error {
	dn := fmt.Sprintf("%s/rsePpInfoAtt", parentDn)
	return sm.DeleteByDn(dn, "vnsRsEPpInfoAtt")
}

func (sm *ServiceManager) ReadRelationvnsRsEPpInfoAtt(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsEPpInfoAtt")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "vnsRsEPpInfoAtt")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}

func (sm *ServiceManager) CreateRelationvnsRsSEPpInfoAtt(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsSEPpInfoAtt", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "vnsRsSEPpInfoAtt", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationvnsRsSEPpInfoAtt(parentDn string) error {
	dn := fmt.Sprintf("%s/rsSEPpInfoAtt", parentDn)
	return sm.DeleteByDn(dn, "vnsRsSEPpInfoAtt")
}

func (sm *ServiceManager) ReadRelationvnsRsSEPpInfoAtt(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsSEPpInfoAtt")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "vnsRsSEPpInfoAtt")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tDn")
		return dat, err
	} else {
		return nil, err
	}
}
