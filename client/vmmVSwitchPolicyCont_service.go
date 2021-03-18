package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateVSwitchPolicyGroup(vmm_domain string, provider_profile_vendor string, description string, vmmVSwitchPolicyContattr models.VSwitchPolicyGroupAttributes) (*models.VSwitchPolicyGroup, error) {
	rn := fmt.Sprintf("vswitchpolcont")
	parentDn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor, vmm_domain)
	vmmVSwitchPolicyCont := models.NewVSwitchPolicyGroup(rn, parentDn, description, vmmVSwitchPolicyContattr)
	err := sm.Save(vmmVSwitchPolicyCont)
	return vmmVSwitchPolicyCont, err
}

func (sm *ServiceManager) ReadVSwitchPolicyGroup(vmm_domain string, provider_profile_vendor string) (*models.VSwitchPolicyGroup, error) {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s/vswitchpolcont", provider_profile_vendor, vmm_domain)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vmmVSwitchPolicyCont := models.VSwitchPolicyGroupFromContainer(cont)
	return vmmVSwitchPolicyCont, nil
}

func (sm *ServiceManager) DeleteVSwitchPolicyGroup(vmm_domain string, provider_profile_vendor string) error {
	dn := fmt.Sprintf("uni/vmmp-%s/dom-%s/vswitchpolcont", provider_profile_vendor, vmm_domain)
	return sm.DeleteByDn(dn, models.VmmvswitchpolicycontClassName)
}

func (sm *ServiceManager) UpdateVSwitchPolicyGroup(vmm_domain string, provider_profile_vendor string, description string, vmmVSwitchPolicyContattr models.VSwitchPolicyGroupAttributes) (*models.VSwitchPolicyGroup, error) {
	rn := fmt.Sprintf("vswitchpolcont")
	parentDn := fmt.Sprintf("uni/vmmp-%s/dom-%s", provider_profile_vendor, vmm_domain)
	vmmVSwitchPolicyCont := models.NewVSwitchPolicyGroup(rn, parentDn, description, vmmVSwitchPolicyContattr)

	vmmVSwitchPolicyCont.Status = "modified"
	err := sm.Save(vmmVSwitchPolicyCont)
	return vmmVSwitchPolicyCont, err

}

func (sm *ServiceManager) ListVSwitchPolicyGroup(vmm_domain string, provider_profile_vendor string) ([]*models.VSwitchPolicyGroup, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/vmmp-%s/dom-%s/vmmVSwitchPolicyCont.json", baseurlStr, provider_profile_vendor, vmm_domain)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.VSwitchPolicyGroupListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationvmmRsVswitchExporterPol(parentDn, tDn, activeFlowTimeout, idleFlowTimeout, samplingRate string) (string, error) {
	dn := fmt.Sprintf("%s/rsvswitchExporterPol-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"activeFlowTimeOut": "%s",
				"idleFlowTimeOut": "%s",
				"samplingRate": "%s"
			}
		}
	}`, "vmmRsVswitchExporterPol", dn, activeFlowTimeout, idleFlowTimeout, samplingRate))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return dn, err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return dn, err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return dn, err
	}
	fmt.Printf("%+v", cont)

	return dn, nil
}

func (sm *ServiceManager) DeleteRelationvmmRsVswitchExporterPol(parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchExporterPol-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchExporterPol")
}

func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideFwPol(parentDn, tnNwsFwPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideFwPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tnNwsFwPolName": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideFwPol", dn, tnNwsFwPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideFwPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideFwPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideFwPol")
}
func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideStpPol(parentDn, tnStpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideStpPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tnStpIfPolName": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideStpPol", dn, tnStpIfPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideStpPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideStpPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideStpPol")
}
func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideLldpIfPol(parentDn, LldpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideLldpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tDn": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideLldpIfPol", dn, LldpIfPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideLldpIfPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideLldpIfPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideLldpIfPol")
}
func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideMcpIfPol(parentDn, tnMcpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideMcpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tnMcpIfPolName": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideMcpIfPol", dn, tnMcpIfPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideMcpIfPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideMcpIfPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideMcpIfPol")
}
func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideCdpIfPol(parentDn, CdpIfPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideCdpIfPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tDn": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideCdpIfPol", dn, CdpIfPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideCdpIfPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideCdpIfPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideCdpIfPol")
}
func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideLacpPol(parentDn, LacpLagPolName string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideLacpPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tDn": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideLacpPol", dn, LacpLagPolName))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideLacpPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideLacpPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideLacpPol")
}

func (sm *ServiceManager) CreateRelationvmmRsVswitchOverrideMtuPol(parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideMtuPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tDn": "%s"
			}
		}
	}`, "vmmRsVswitchOverrideMtuPol", dn, tDn))

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

func (sm *ServiceManager) DeleteRelationvmmRsVswitchOverrideMtuPol(parentDn string) error {
	dn := fmt.Sprintf("%s/rsvswitchOverrideMtuPol", parentDn)
	return sm.DeleteByDn(dn, "vmmRsVswitchOverrideMtuPol")
}