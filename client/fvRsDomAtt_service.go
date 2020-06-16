package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateDomain(tDn string, application_epg string, application_profile string, tenant string, description string, fvRsDomAttattr models.DomainAttributes) (*models.Domain, error) {
	rn := fmt.Sprintf("rsdomAtt-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant, application_profile, application_epg)
	fvRsDomAtt := models.NewDomain(rn, parentDn, description, fvRsDomAttattr)
	err := sm.Save(fvRsDomAtt)
	return fvRsDomAtt, err
}

func (sm *ServiceManager) ReadDomain(tDn string, application_epg string, application_profile string, tenant string) (*models.Domain, error) {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s/rsdomAtt-[%s]", tenant, application_profile, application_epg, tDn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvRsDomAtt := models.DomainFromContainer(cont)
	return fvRsDomAtt, nil
}

func (sm *ServiceManager) DeleteDomain(tDn string, application_epg string, application_profile string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s/rsdomAtt-[%s]", tenant, application_profile, application_epg, tDn)
	return sm.DeleteByDn(dn, models.FvrsdomattClassName)
}

func (sm *ServiceManager) UpdateDomain(tDn string, application_epg string, application_profile string, tenant string, description string, fvRsDomAttattr models.DomainAttributes) (*models.Domain, error) {
	rn := fmt.Sprintf("rsdomAtt-[%s]", tDn)
	parentDn := fmt.Sprintf("uni/tn-%s/ap-%s/epg-%s", tenant, application_profile, application_epg)
	fvRsDomAtt := models.NewDomain(rn, parentDn, description, fvRsDomAttattr)

	fvRsDomAtt.Status = "modified"
	err := sm.Save(fvRsDomAtt)
	return fvRsDomAtt, err

}

func (sm *ServiceManager) ListDomain(application_epg string, application_profile string, tenant string) ([]*models.Domain, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/ap-%s/epg-%s/fvRsDomAtt.json", baseurlStr, tenant, application_profile, application_epg)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.DomainListFromContainer(cont)

	return list, err
}
