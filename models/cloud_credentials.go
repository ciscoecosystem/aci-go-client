package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const (
	DncloudCredentials        = "uni/tn-%s/credentials-%s"
	RncloudCredentials        = "credentials-%s"
	ParentDncloudCredentials  = "uni/tn-%s"
	CloudcredentialsClassName = "cloudCredentials"
)

type AccessCredentialtomanagethecloudresources struct {
	BaseAttributes
	NameAliasAttribute
	AccessCredentialtomanagethecloudresourcesAttributes
}

type AccessCredentialtomanagethecloudresourcesAttributes struct {
	Annotation    string `json:",omitempty"`
	ClientId      string `json:",omitempty"`
	Email         string `json:",omitempty"`
	HttpProxy     string `json:",omitempty"`
	Key           string `json:",omitempty"`
	KeyId         string `json:",omitempty"`
	Name          string `json:",omitempty"`
	RsaPrivateKey string `json:",omitempty"`
}

func NewAccessCredentialtomanagethecloudresources(cloudCredentialsRn, parentDn, description, nameAlias string, cloudCredentialsAttr AccessCredentialtomanagethecloudresourcesAttributes) *AccessCredentialtomanagethecloudresources {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudCredentialsRn)
	return &AccessCredentialtomanagethecloudresources{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudcredentialsClassName,
			Rn:                cloudCredentialsRn,
		},
		NameAliasAttribute: NameAliasAttribute{
			NameAlias: nameAlias,
		},
		AccessCredentialtomanagethecloudresourcesAttributes: cloudCredentialsAttr,
	}
}

func (cloudCredentials *AccessCredentialtomanagethecloudresources) ToMap() (map[string]string, error) {
	cloudCredentialsMap, err := cloudCredentials.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	alias, err := cloudCredentials.NameAliasAttribute.ToMap()
	if err != nil {
		return nil, err
	}

	for key, value := range alias {
		A(cloudCredentialsMap, key, value)
	}

	A(cloudCredentialsMap, "clientId", cloudCredentials.ClientId)
	A(cloudCredentialsMap, "email", cloudCredentials.Email)
	A(cloudCredentialsMap, "httpProxy", cloudCredentials.HttpProxy)
	A(cloudCredentialsMap, "key", cloudCredentials.Key)
	A(cloudCredentialsMap, "keyId", cloudCredentials.KeyId)
	A(cloudCredentialsMap, "name", cloudCredentials.Name)
	A(cloudCredentialsMap, "rsaPrivateKey", cloudCredentials.RsaPrivateKey)
	return cloudCredentialsMap, err
}

func AccessCredentialtomanagethecloudresourcesFromContainerList(cont *container.Container, index int) *AccessCredentialtomanagethecloudresources {
	AccessCredentialtomanagethecloudresourcesCont := cont.S("imdata").Index(index).S(CloudcredentialsClassName, "attributes")
	return &AccessCredentialtomanagethecloudresources{
		BaseAttributes{
			DistinguishedName: G(AccessCredentialtomanagethecloudresourcesCont, "dn"),
			Description:       G(AccessCredentialtomanagethecloudresourcesCont, "descr"),
			Status:            G(AccessCredentialtomanagethecloudresourcesCont, "status"),
			ClassName:         CloudcredentialsClassName,
			Rn:                G(AccessCredentialtomanagethecloudresourcesCont, "rn"),
		},
		NameAliasAttribute{
			NameAlias: G(AccessCredentialtomanagethecloudresourcesCont, "nameAlias"),
		},
		AccessCredentialtomanagethecloudresourcesAttributes{
			ClientId:      G(AccessCredentialtomanagethecloudresourcesCont, "clientId"),
			Email:         G(AccessCredentialtomanagethecloudresourcesCont, "email"),
			HttpProxy:     G(AccessCredentialtomanagethecloudresourcesCont, "httpProxy"),
			Key:           G(AccessCredentialtomanagethecloudresourcesCont, "key"),
			KeyId:         G(AccessCredentialtomanagethecloudresourcesCont, "keyId"),
			Name:          G(AccessCredentialtomanagethecloudresourcesCont, "name"),
			RsaPrivateKey: G(AccessCredentialtomanagethecloudresourcesCont, "rsaPrivateKey"),
		},
	}
}

func AccessCredentialtomanagethecloudresourcesFromContainer(cont *container.Container) *AccessCredentialtomanagethecloudresources {
	return AccessCredentialtomanagethecloudresourcesFromContainerList(cont, 0)
}

func AccessCredentialtomanagethecloudresourcesListFromContainer(cont *container.Container) []*AccessCredentialtomanagethecloudresources {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*AccessCredentialtomanagethecloudresources, length)

	for i := 0; i < length; i++ {
		arr[i] = AccessCredentialtomanagethecloudresourcesFromContainerList(cont, i)
	}

	return arr
}
