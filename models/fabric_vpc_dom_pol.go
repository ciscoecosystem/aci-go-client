package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VpcDomPolClassName = "vpcInstPol"

type FabricVpcDomainPolicy struct {
	BaseAttributes
	FabricVpcDomainPolicyAttributes
}

type FabricVpcDomainPolicyAttributes struct {
	Name string `json:",omitempty"`

	PeerDeadInterval string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`
}

func NewFabricVpcDomainPolicy(vpcDomPolRn, parentDn, description string, vpcDomPolattr FabricVpcDomainPolicyAttributes) *FabricVpcDomainPolicy {
	dn := fmt.Sprintf("%s/%s", parentDn, vpcDomPolRn)
	return &FabricVpcDomainPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VpcDomPolClassName,
			Rn:                vpcDomPolRn,
		},

		FabricVpcDomainPolicyAttributes: vpcDomPolattr,
	}
}

func (vpcDomPol *FabricVpcDomainPolicy) ToMap() (map[string]string, error) {
	vpcDomPolMap, err := vpcDomPol.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vpcDomPolMap, "name", vpcDomPol.Name)

	A(vpcDomPolMap, "deadIntvl", vpcDomPol.PeerDeadInterval)

	A(vpcDomPolMap, "annotation", vpcDomPol.Annotation)

	A(vpcDomPolMap, "nameAlias", vpcDomPol.NameAlias)

	return vpcDomPolMap, err
}

func FabricVpcDomainPolicyFromContainerList(cont *container.Container, index int) *FabricVpcDomainPolicy {

	FabricVpcDomainPolicyCont := cont.S("imdata").Index(index).S(VpcDomPolClassName, "attributes")
	return &FabricVpcDomainPolicy{
		BaseAttributes{
			DistinguishedName: G(FabricVpcDomainPolicyCont, "dn"),
			Description:       G(FabricVpcDomainPolicyCont, "descr"),
			Status:            G(FabricVpcDomainPolicyCont, "status"),
			ClassName:         VpcDomPolClassName,
			Rn:                G(FabricVpcDomainPolicyCont, "rn"),
		},

		FabricVpcDomainPolicyAttributes{

			Name: G(FabricVpcDomainPolicyCont, "name"),

			PeerDeadInterval: G(FabricVpcDomainPolicyCont, "deadIntvl"),

			Annotation: G(FabricVpcDomainPolicyCont, "annotation"),

			NameAlias: G(FabricVpcDomainPolicyCont, "nameAlias"),
		},
	}
}

func FabricVpcDomainPolicyFromContainer(cont *container.Container) *FabricVpcDomainPolicy {

	return FabricVpcDomainPolicyFromContainerList(cont, 0)
}

func FabricVpcDomainPolicyListFromContainer(cont *container.Container) []*FabricVpcDomainPolicy {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*FabricVpcDomainPolicy, length)

	for i := 0; i < length; i++ {

		arr[i] = FabricVpcDomainPolicyFromContainerList(cont, i)
	}

	return arr
}
