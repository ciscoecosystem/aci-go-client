package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const AaaUserClassName = "aaaUser"

type User struct {
	BaseAttributes
	UserAttributes
}

type UserAttributes struct {
	Name string `json:"omit empty"`

	AccountStatus string `json:"omit empty"`

	NameAlias string `json:"omit empty"`

	//OtpEnable bool `json:"omit empty"`

	OtpKey string `json:"omit empty"`

	Phone string `json:"omit empty"`

	Pwd string `json:"omit empty"`

	//PwdLifetime int `json:"omit empty"`

	//PwdUpdateRequired bool `json:"omit empty"`

	RbacString string `json:"omit empty"`

	//UnixUserId int `json:"omit empty"`

	Annotation string `json:"omit empty"`

	CertAttribute string `json:"omit empty"`

	ClearPwdHistory string `json:"omit empty"`

	Email string `json:"omit empty"`

	Expiration string `json:"omit empty"`

	Expires string `json:"omit empty"`

	FirstName string `json:"omit empty"`

	LastName string `json:"omit empty"`
}

func NewUser(aaaUserRn, parentDn, description string, aaaUserattr UserAttributes) *User {
	dn := fmt.Sprintf("%s/%s", parentDn, aaaUserRn)
	return &User{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         AaaUserClassName,
			Rn:                aaaUserRn,
		},

		UserAttributes: aaaUserattr,
	}
}

func (aaaUser *User) ToMap() (map[string]string, error) {
	aaaUserMap, err := aaaUser.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(aaaUserMap, "name", aaaUser.Name)
	A(aaaUserMap, "account_status", aaaUser.AccountStatus)
	A(aaaUserMap, "name_alias", aaaUser.NameAlias)
	//A(aaaUserMap, "otp_enable", aaaUser.OtpEnable)
	A(aaaUserMap, "otp_key", aaaUser.OtpKey)
	A(aaaUserMap, "phone", aaaUser.Phone)
	A(aaaUserMap, "pwd", aaaUser.Pwd)
	//A(aaaUserMap, "pwd_life_time", aaaUser.PwdLifetime)
	//A(aaaUserMap, "pwd_update_required", aaaUser.PwdUpdateRequired)
	A(aaaUserMap, "rbac_string", aaaUser.RbacString)
	//A(aaaUserMap, "unix_user_id", aaaUser.UnixUserId)
	A(aaaUserMap, "annotation", aaaUser.Annotation)
	A(aaaUserMap, "cert_attribute", aaaUser.CertAttribute)
	A(aaaUserMap, "clear_pwd_history", aaaUser.ClearPwdHistory)
	A(aaaUserMap, "email", aaaUser.Email)
	A(aaaUserMap, "expiration", aaaUser.Expiration)
	A(aaaUserMap, "expires", aaaUser.Expires)
	A(aaaUserMap, "first_name", aaaUser.FirstName)
	A(aaaUserMap, "last_name", aaaUser.LastName)

	return aaaUserMap, err
}

func UserFromContainerList(cont *container.Container, index int) *User {

	UserCont := cont.S("imdata").Index(index).S(AaaUserClassName, "attributes")
	return &User{
		BaseAttributes{
			DistinguishedName: G(UserCont, "dn"),
			Description:       G(UserCont, "descr"),
			Status:            G(UserCont, "status"),
			ClassName:         AaaUserClassName,
			Rn:                G(UserCont, "rn"),
		},

		UserAttributes{

			Name:          G(UserCont, "name"),
			AccountStatus: G(UserCont, "account_status"),
			NameAlias:     G(UserCont, "name_alias"),
			//OtpEnable:         G(UserCont, "otp_enable"),
			OtpKey: G(UserCont, "otp_key"),
			Phone:  G(UserCont, "phone"),
			Pwd:    G(UserCont, "pwd"),
			//PwdLifetime:       G(UserCont, "pwd_life_time"),
			//PwdUpdateRequired: G(UserCont, "pwd_update_required"),
			//UnixUserId:        G(UserCont, "unix_user_id"),
			RbacString:      G(UserCont, "rbac_string"),
			Annotation:      G(UserCont, "annotation"),
			CertAttribute:   G(UserCont, "cert_attribute"),
			ClearPwdHistory: G(UserCont, "clear_pwd_history"),
			Email:           G(UserCont, "email"),
			Expiration:      G(UserCont, "expiration"),
			Expires:         G(UserCont, "expires"),
			FirstName:       G(UserCont, "first_name"),
			LastName:        G(UserCont, "last_name"),
		},
	}
}

func UserFromContainer(cont *container.Container) *User {

	return UserFromContainerList(cont, 0)
}

func UserListFromContainer(cont *container.Container) []*User {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*User, length)

	for i := 0; i < length; i++ {

		arr[i] = UserFromContainerList(cont, i)
	}

	return arr
}
