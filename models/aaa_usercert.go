package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const AaaUserCertClassName = "aaaUserCert"

type UserCert struct {
	BaseAttributes
	UserCertAttributes
}

type UserCertAttributes struct {
	Name string `json:"omit empty"`

	NameAlias string `json:"omit empty"`

	Annotation string `json:"omit empty"`

	Data string `json:"omit empty"`
}

func NewUserCert(aaaUserCertRn, parentDn, description string, aaaUserCertattr UserCertAttributes) *UserCert {
	dn := fmt.Sprintf("%s/%s", parentDn, aaaUserCertRn)
	return &UserCert{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         AaaUserCertClassName,
			Rn:                aaaUserCertRn,
		},

		UserCertAttributes: aaaUserCertattr,
	}
}

func (aaaUserCert *UserCert) ToMap() (map[string]string, error) {
	aaaUserCertMap, err := aaaUserCert.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(aaaUserCertMap, "name", aaaUserCert.Name)

	A(aaaUserCertMap, "name_alias", aaaUserCert.NameAlias)

	A(aaaUserCertMap, "annotation", aaaUserCert.Annotation)

	A(aaaUserCertMap, "data", aaaUserCert.Data)

	return aaaUserCertMap, err
}

func UserCertFromContainerList(cont *container.Container, index int) *UserCert {

	UserCertCont := cont.S("imdata").Index(index).S(AaaUserCertClassName, "attributes")
	return &UserCert{
		BaseAttributes{
			DistinguishedName: G(UserCertCont, "dn"),
			Description:       G(UserCertCont, "descr"),
			Status:            G(UserCertCont, "status"),
			ClassName:         AaaUserCertClassName,
			Rn:                G(UserCertCont, "rn"),
		},

		UserCertAttributes{

			Name: G(UserCertCont, "name"),

			NameAlias: G(UserCertCont, "name_alias"),

			Annotation: G(UserCertCont, "annotation"),
			Data:       G(UserCertCont, "data"),
		},
	}
}

func UserCertFromContainer(cont *container.Container) *UserCert {

	return UserCertFromContainerList(cont, 0)
}

func UserCertListFromContainer(cont *container.Container) []*UserCert {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*UserCert, length)

	for i := 0; i < length; i++ {

		arr[i] = UserCertFromContainerList(cont, i)
	}

	return arr
}
