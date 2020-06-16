package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const L3extrsnodel3outattClassName = "l3extRsNodeL3OutAtt"

type L3extRsNodeL3OutAtt struct {
	BaseAttributes
	L3extRsNodeL3OutAttAttributes
}

type L3extRsNodeL3OutAttAttributes struct {
	TDn string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	ConfigIssues string `json:",omitempty"`

	RtrId string `json:",omitempty"`

	RtrIdLoopBack string `json:",omitempty"`
}

func NewL3extRsNodeL3OutAtt(l3extRsNodeL3OutAttRn, parentDn, description string, l3extRsNodeL3OutAttattr L3extRsNodeL3OutAttAttributes) *L3extRsNodeL3OutAtt {
	dn := fmt.Sprintf("%s/%s", parentDn, l3extRsNodeL3OutAttRn)
	return &L3extRsNodeL3OutAtt{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         L3extrsnodel3outattClassName,
			Rn:                l3extRsNodeL3OutAttRn,
		},

		L3extRsNodeL3OutAttAttributes: l3extRsNodeL3OutAttattr,
	}
}

func (l3extRsNodeL3OutAtt *L3extRsNodeL3OutAtt) ToMap() (map[string]string, error) {
	l3extRsNodeL3OutAttMap, err := l3extRsNodeL3OutAtt.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(l3extRsNodeL3OutAttMap, "tDn", l3extRsNodeL3OutAtt.TDn)

	A(l3extRsNodeL3OutAttMap, "annotation", l3extRsNodeL3OutAtt.Annotation)

	A(l3extRsNodeL3OutAttMap, "configIssues", l3extRsNodeL3OutAtt.ConfigIssues)

	A(l3extRsNodeL3OutAttMap, "rtrId", l3extRsNodeL3OutAtt.RtrId)

	A(l3extRsNodeL3OutAttMap, "rtrIdLoopBack", l3extRsNodeL3OutAtt.RtrIdLoopBack)

	return l3extRsNodeL3OutAttMap, err
}

func L3extRsNodeL3OutAttFromContainerList(cont *container.Container, index int) *L3extRsNodeL3OutAtt {

	L3extRsNodeL3OutAttCont := cont.S("imdata").Index(index).S(L3extrsnodel3outattClassName, "attributes")
	return &L3extRsNodeL3OutAtt{
		BaseAttributes{
			DistinguishedName: G(L3extRsNodeL3OutAttCont, "dn"),
			Description:       G(L3extRsNodeL3OutAttCont, "descr"),
			Status:            G(L3extRsNodeL3OutAttCont, "status"),
			ClassName:         L3extrsnodel3outattClassName,
			Rn:                G(L3extRsNodeL3OutAttCont, "rn"),
		},

		L3extRsNodeL3OutAttAttributes{

			TDn: G(L3extRsNodeL3OutAttCont, "tDn"),

			Annotation: G(L3extRsNodeL3OutAttCont, "annotation"),

			ConfigIssues: G(L3extRsNodeL3OutAttCont, "configIssues"),

			RtrId: G(L3extRsNodeL3OutAttCont, "rtrId"),

			RtrIdLoopBack: G(L3extRsNodeL3OutAttCont, "rtrIdLoopBack"),
		},
	}
}

func L3extRsNodeL3OutAttFromContainer(cont *container.Container) *L3extRsNodeL3OutAtt {

	return L3extRsNodeL3OutAttFromContainerList(cont, 0)
}

func L3extRsNodeL3OutAttListFromContainer(cont *container.Container) []*L3extRsNodeL3OutAtt {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*L3extRsNodeL3OutAtt, length)

	for i := 0; i < length; i++ {

		arr[i] = L3extRsNodeL3OutAttFromContainerList(cont, i)
	}

	return arr
}
