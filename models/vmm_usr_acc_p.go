package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VmmusraccpClassName = "vmmUsrAccP"

type VMMCredential struct {
	BaseAttributes
	VMMCredentialAttributes
}

type VMMCredentialAttributes struct {
	Name       string `json:",omitempty"`
	Annotation string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
	Pwd        string `json:",omitempty"`
	Usr        string `json:",omitempty"`
}

func NewVMMCredential(vmmUsrAccPRn, parentDn, description string, vmmUsrAccPattr VMMCredentialAttributes) *VMMCredential {
	dn := fmt.Sprintf("%s/%s", parentDn, vmmUsrAccPRn)
	return &VMMCredential{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VmmusraccpClassName,
			Rn:                vmmUsrAccPRn,
		},

		VMMCredentialAttributes: vmmUsrAccPattr,
	}
}

func (vmmUsrAccP *VMMCredential) ToMap() (map[string]string, error) {
	vmmUsrAccPMap, err := vmmUsrAccP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vmmUsrAccPMap, "name", vmmUsrAccP.Name)
	A(vmmUsrAccPMap, "annotation", vmmUsrAccP.Annotation)
	A(vmmUsrAccPMap, "nameAlias", vmmUsrAccP.NameAlias)
	A(vmmUsrAccPMap, "pwd", vmmUsrAccP.Pwd)
	A(vmmUsrAccPMap, "usr", vmmUsrAccP.Usr)

	return vmmUsrAccPMap, err
}

func VMMCredentialFromContainerList(cont *container.Container, index int) *VMMCredential {

	VMMCredentialCont := cont.S("imdata").Index(index).S(VmmusraccpClassName, "attributes")
	return &VMMCredential{
		BaseAttributes{
			DistinguishedName: G(VMMCredentialCont, "dn"),
			Description:       G(VMMCredentialCont, "descr"),
			Status:            G(VMMCredentialCont, "status"),
			ClassName:         VmmusraccpClassName,
			Rn:                G(VMMCredentialCont, "rn"),
		},

		VMMCredentialAttributes{
			Name:       G(VMMCredentialCont, "name"),
			Annotation: G(VMMCredentialCont, "annotation"),
			NameAlias:  G(VMMCredentialCont, "nameAlias"),
			Pwd:        G(VMMCredentialCont, "pwd"),
			Usr:        G(VMMCredentialCont, "usr"),
		},
	}
}

func VMMCredentialFromContainer(cont *container.Container) *VMMCredential {

	return VMMCredentialFromContainerList(cont, 0)
}

func VMMCredentialListFromContainer(cont *container.Container) []*VMMCredential {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*VMMCredential, length)

	for i := 0; i < length; i++ {

		arr[i] = VMMCredentialFromContainerList(cont, i)
	}

	return arr
}
