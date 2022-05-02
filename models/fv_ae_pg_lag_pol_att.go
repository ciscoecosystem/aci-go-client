package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnfvAEPgLagPolAtt        = "uni/tn-%s/ap-%s/epg-%s/rsdomAtt-[%s]/epglagpolatt"
	RnfvAEPgLagPolAtt        = "epglagpolatt"
	ParentDnfvAEPgLagPolAtt  = "uni/tn-%s/ap-%s/epg-%s/rsdomAtt-[%s]"
	FvaepglagpolattClassName = "fvAEPgLagPolAtt"
)

type ApplicationEPGLagPolicy struct {
	BaseAttributes
}

type ApplicationEPGLagPolicyAttributes struct {
	Annotation string `json:",omitempty"`
}

func NewApplicationEPGLagPolicy(fvAEPgLagPolAttRn, parentDn string) *ApplicationEPGLagPolicy {
	dn := fmt.Sprintf("%s/%s", parentDn, fvAEPgLagPolAttRn)
	return &ApplicationEPGLagPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         FvaepglagpolattClassName,
			Rn:                fvAEPgLagPolAttRn,
		},
	}
}

func (fvAEPgLagPolAtt *ApplicationEPGLagPolicy) ToMap() (map[string]string, error) {
	fvAEPgLagPolAttMap, err := fvAEPgLagPolAtt.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
	return fvAEPgLagPolAttMap, err
}

func ApplicationEPGLagPolicyFromContainerList(cont *container.Container, index int) *ApplicationEPGLagPolicy {
	ApplicationEPGLagPolicyCont := cont.S("imdata").Index(index).S(FvaepglagpolattClassName, "attributes")
	return &ApplicationEPGLagPolicy{
		BaseAttributes{
			DistinguishedName: G(ApplicationEPGLagPolicyCont, "dn"),
			Status:            G(ApplicationEPGLagPolicyCont, "status"),
			ClassName:         FvaepglagpolattClassName,
			Rn:                G(ApplicationEPGLagPolicyCont, "rn"),
		},
	}
}

func ApplicationEPGLagPolicyFromContainer(cont *container.Container) *ApplicationEPGLagPolicy {
	return ApplicationEPGLagPolicyFromContainerList(cont, 0)
}

func ApplicationEPGLagPolicyListFromContainer(cont *container.Container) []*ApplicationEPGLagPolicy {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*ApplicationEPGLagPolicy, length)

	for i := 0; i < length; i++ {
		arr[i] = ApplicationEPGLagPolicyFromContainerList(cont, i)
	}

	return arr
}
