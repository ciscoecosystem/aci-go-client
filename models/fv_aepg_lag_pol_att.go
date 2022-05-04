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
	FvAEPgLagPolAttClassName = "fvAEPgLagPolAtt"
)

type ApplicationEPGLagPolicy struct {
	BaseAttributes
	NameAliasAttribute
	ApplicationEPGLagPolicyAttributes
}

type ApplicationEPGLagPolicyAttributes struct {
	Annotation string `json:",omitempty"`
}

func NewApplicationEPGLagPolicy(fvAEPgLagPolAttRn, parentDn, description, nameAlias string, fvAEPgLagPolAttAttr ApplicationEPGLagPolicyAttributes) *ApplicationEPGLagPolicy {
	dn := fmt.Sprintf("%s/%s", parentDn, fvAEPgLagPolAttRn)
	return &ApplicationEPGLagPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvAEPgLagPolAttClassName,
			Rn:                fvAEPgLagPolAttRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		ApplicationEPGLagPolicyAttributes: fvAEPgLagPolAttAttr,
	}
}

func (fvAEPgLagPolAtt *ApplicationEPGLagPolicy) ToMap() (map[string]string, error) {
	fvAEPgLagPolAttMap, err := fvAEPgLagPolAtt.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := fvAEPgLagPolAtt.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(fvAEPgLagPolAttMap, key, value)
	}

	A(fvAEPgLagPolAttMap, "annotation", fvAEPgLagPolAtt.Annotation)
	return fvAEPgLagPolAttMap, err
}

func ApplicationEPGLagPolicyFromContainerList(cont *container.Container, index int) *ApplicationEPGLagPolicy {
	ApplicationEPGLagPolicyCont := cont.S("imdata").Index(index).S(FvAEPgLagPolAttClassName, "attributes")
	return &ApplicationEPGLagPolicy{
		BaseAttributes{
			DistinguishedName: G(ApplicationEPGLagPolicyCont, "dn"),
			Description:       G(ApplicationEPGLagPolicyCont, "descr"),
			Status:            G(ApplicationEPGLagPolicyCont, "status"),
			ClassName:         FvAEPgLagPolAttClassName,
			Rn:                G(ApplicationEPGLagPolicyCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(ApplicationEPGLagPolicyCont, "nameAlias"),
		},
		ApplicationEPGLagPolicyAttributes{
			Annotation: G(ApplicationEPGLagPolicyCont, "annotation"),
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
