package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DnvzRsSubjFiltAtt        = "uni/tn-%s/brc-%s/subj-%s/rssubjFiltAtt-%s"
	RnvzRsSubjFiltAtt        = "rssubjFiltAtt-%s"
	ParentDnvzRsSubjFiltAtt  = "uni/tn-%s/brc-%s/subj-%s"
	VzrssubjfiltattClassName = "vzRsSubjFiltAtt"
)

type SubjectFilter struct {
	BaseAttributes
	NameAliasAttribute
	SubjectFilterAttributes
}

type SubjectFilterAttributes struct {
	Annotation       string `json:",omitempty"`
	Action           string `json:",omitempty"`
	Directives       string `json:",omitempty"`
	PriorityOverride string `json:",omitempty"`
	TnVzFilterName   string `json:",omitempty"`
}

func NewSubjectFilter(vzRsSubjFiltAttRn, parentDn, description, nameAlias string, vzRsSubjFiltAttAttr SubjectFilterAttributes) *SubjectFilter {
	dn := fmt.Sprintf("%s/%s", parentDn, vzRsSubjFiltAttRn)
	return &SubjectFilter{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VzrssubjfiltattClassName,
			Rn:                vzRsSubjFiltAttRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		SubjectFilterAttributes: vzRsSubjFiltAttAttr,
	}
}

func (vzRsSubjFiltAtt *SubjectFilter) ToMap() (map[string]string, error) {
	vzRsSubjFiltAttMap, err := vzRsSubjFiltAtt.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := vzRsSubjFiltAtt.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(vzRsSubjFiltAttMap, key, value)
	}

	A(vzRsSubjFiltAttMap, "action", vzRsSubjFiltAtt.Action)
	A(vzRsSubjFiltAttMap, "directives", vzRsSubjFiltAtt.Directives)
	A(vzRsSubjFiltAttMap, "priorityOverride", vzRsSubjFiltAtt.PriorityOverride)
	A(vzRsSubjFiltAttMap, "tnVzFilterName", vzRsSubjFiltAtt.TnVzFilterName)
	return vzRsSubjFiltAttMap, err
}

func SubjectFilterFromContainerList(cont *container.Container, index int) *SubjectFilter {
	SubjectFilterCont := cont.S("imdata").Index(index).S(VzrssubjfiltattClassName, "attributes")
	return &SubjectFilter{
		BaseAttributes{
			DistinguishedName: G(SubjectFilterCont, "dn"),
			Description:       G(SubjectFilterCont, "descr"),
			Status:            G(SubjectFilterCont, "status"),
			ClassName:         VzrssubjfiltattClassName,
			Rn:                G(SubjectFilterCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(SubjectFilterCont, "nameAlias"),
		},
		SubjectFilterAttributes{
			Action:           G(SubjectFilterCont, "action"),
			Directives:       G(SubjectFilterCont, "directives"),
			PriorityOverride: G(SubjectFilterCont, "priorityOverride"),
			TnVzFilterName:   G(SubjectFilterCont, "tnVzFilterName"),
		},
	}
}

func SubjectFilterFromContainer(cont *container.Container) *SubjectFilter {
	return SubjectFilterFromContainerList(cont, 0)
}

func SubjectFilterListFromContainer(cont *container.Container) []*SubjectFilter {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*SubjectFilter, length)

	for i := 0; i < length; i++ {
		arr[i] = SubjectFilterFromContainerList(cont, i)
	}

	return arr
}
