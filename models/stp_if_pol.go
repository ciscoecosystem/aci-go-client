package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const StpIfPolClassName = "stpIfPol"

type SpanningTreeInterfacePolicy struct {
	BaseAttributes
	SpanningTreeInterfacePolicyAttributes
}

type SpanningTreeInterfacePolicyAttributes struct {
	Name string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	Ctrl string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewSpanningTreeInterfacePolicy(stpIfPolRn, parentDn, description string, stpIfPolattr SpanningTreeInterfacePolicyAttributes) *SpanningTreeInterfacePolicy {
	dn := fmt.Sprintf("%s/%s", parentDn, stpIfPolRn)
	return &SpanningTreeInterfacePolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         StpIfPolClassName,
			Rn:                stpIfPolRn,
		},

		SpanningTreeInterfacePolicyAttributes: stpIfPolattr,
	}
}

func (stpIfPol *SpanningTreeInterfacePolicy) ToMap() (map[string]string, error) {
	stpIfPolMap, err := stpIfPol.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(stpIfPolMap, "name", stpIfPol.Name)

	A(stpIfPolMap, "annotation", stpIfPol.Annotation)

	A(stpIfPolMap, "ctrl", stpIfPol.Ctrl)

	A(stpIfPolMap, "nameAlias", stpIfPol.NameAlias)

	return stpIfPolMap, err
}

func SpanningTreeInterfacePolicyFromContainerList(cont *container.Container, index int) *SpanningTreeInterfacePolicy {

	SpanningTreeInterfacePolicyCont := cont.S("imdata").Index(index).S(StpIfPolClassName, "attributes")
	return &SpanningTreeInterfacePolicy{
		BaseAttributes{
			DistinguishedName: G(SpanningTreeInterfacePolicyCont, "dn"),
			Description:       G(SpanningTreeInterfacePolicyCont, "descr"),
			Status:            G(SpanningTreeInterfacePolicyCont, "status"),
			ClassName:         StpIfPolClassName,
			Rn:                G(SpanningTreeInterfacePolicyCont, "rn"),
		},

		SpanningTreeInterfacePolicyAttributes{

			Name: G(SpanningTreeInterfacePolicyCont, "name"),

			Annotation: G(SpanningTreeInterfacePolicyCont, "annotation"),

			Ctrl: G(SpanningTreeInterfacePolicyCont, "ctrl"),

			NameAlias: G(SpanningTreeInterfacePolicyCont, "nameAlias"),
		},
	}
}

func SpanningTreeInterfacePolicyFromContainer(cont *container.Container) *SpanningTreeInterfacePolicy {

	return SpanningTreeInterfacePolicyFromContainerList(cont, 0)
}

func SpanningTreeInterfacePolicyListFromContainer(cont *container.Container) []*SpanningTreeInterfacePolicy {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanningTreeInterfacePolicy, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanningTreeInterfacePolicyFromContainerList(cont, i)
	}

	return arr
}
