package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (sm *ServiceManager) CreateLogicalInterface(name string, l4_l7_devices string, tenant string, nameAlias string, vnsLIfAttr models.LogicalInterfaceAttributes) (*models.LogicalInterface, error) {
	rn := fmt.Sprintf(models.RnvnsLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsLIf, tenant, l4_l7_devices)
	vnsLIf := models.NewLogicalInterface(rn, parentDn, nameAlias, vnsLIfAttr)
	err := sm.Save(vnsLIf)
	return vnsLIf, err
}

func (sm *ServiceManager) ReadLogicalInterface(name string, l4_l7_devices string, tenant string) (*models.LogicalInterface, error) {
	dn := fmt.Sprintf(models.DnvnsLIf, tenant, l4_l7_devices, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vnsLIf := models.LogicalInterfaceFromContainer(cont)
	return vnsLIf, nil
}

func (sm *ServiceManager) DeleteLogicalInterface(name string, l4_l7_devices string, tenant string) error {
	dn := fmt.Sprintf(models.DnvnsLIf, tenant, l4_l7_devices, name)
	return sm.DeleteByDn(dn, models.VnslifClassName)
}

func (sm *ServiceManager) UpdateLogicalInterface(name string, l4_l7_devices string, tenant string, nameAlias string, vnsLIfAttr models.LogicalInterfaceAttributes) (*models.LogicalInterface, error) {
	rn := fmt.Sprintf(models.RnvnsLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsLIf, tenant, l4_l7_devices)
	vnsLIf := models.NewLogicalInterface(rn, parentDn, nameAlias, vnsLIfAttr)
	vnsLIf.Status = "modified"
	err := sm.Save(vnsLIf)
	return vnsLIf, err
}

func (sm *ServiceManager) ListLogicalInterface(l4_l7_devices string, tenant string) ([]*models.LogicalInterface, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/lDevVip-%s/vnsLIf.json", models.BaseurlStr, tenant, l4_l7_devices)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.LogicalInterfaceListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationvnsRsCIfAttN(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rscIfAttN-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "vnsRsCIfAttN", dn, annotation, tDn))

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

func (sm *ServiceManager) DeleteRelationvnsRsCIfAttN(parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rscIfAttN-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn, "vnsRsCIfAttN")
}

func (sm *ServiceManager) ReadRelationvnsRsCIfAttN(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsCIfAttN")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "vnsRsCIfAttN")

	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList {
		dat := models.G(contItem, "tDn")
		st.Add(dat)
	}
	return st, err
}
