package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudrouterpClassName = "cloudRouterP"

type CloudRouterProfile struct {
	BaseAttributes
	CloudRouterProfileAttributes
}

type CloudRouterProfileAttributes struct {
	Name string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`

	NumInstances string `json:",omitempty"`

	CloudRouterProfile_type string `json:",omitempty"`
}

func NewCloudRouterProfile(cloudRouterPRn, parentDn, description string, cloudRouterPattr CloudRouterProfileAttributes) *CloudRouterProfile {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudRouterPRn)
	return &CloudRouterProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudrouterpClassName,
			Rn:                cloudRouterPRn,
		},

		CloudRouterProfileAttributes: cloudRouterPattr,
	}
}

func (cloudRouterP *CloudRouterProfile) ToMap() (map[string]string, error) {
	cloudRouterPMap, err := cloudRouterP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudRouterPMap, "name", cloudRouterP.Name)

	A(cloudRouterPMap, "annotation", cloudRouterP.Annotation)

	A(cloudRouterPMap, "nameAlias", cloudRouterP.NameAlias)

	A(cloudRouterPMap, "numInstances", cloudRouterP.NumInstances)

	A(cloudRouterPMap, "type", cloudRouterP.CloudRouterProfile_type)

	return cloudRouterPMap, err
}

func CloudRouterProfileFromContainerList(cont *container.Container, index int) *CloudRouterProfile {

	CloudRouterProfileCont := cont.S("imdata").Index(index).S(CloudrouterpClassName, "attributes")
	return &CloudRouterProfile{
		BaseAttributes{
			DistinguishedName: G(CloudRouterProfileCont, "dn"),
			Description:       G(CloudRouterProfileCont, "descr"),
			Status:            G(CloudRouterProfileCont, "status"),
			ClassName:         CloudrouterpClassName,
			Rn:                G(CloudRouterProfileCont, "rn"),
		},

		CloudRouterProfileAttributes{

			Name: G(CloudRouterProfileCont, "name"),

			Annotation: G(CloudRouterProfileCont, "annotation"),

			NameAlias: G(CloudRouterProfileCont, "nameAlias"),

			NumInstances: G(CloudRouterProfileCont, "numInstances"),

			CloudRouterProfile_type: G(CloudRouterProfileCont, "type"),
		},
	}
}

func CloudRouterProfileFromContainer(cont *container.Container) *CloudRouterProfile {

	return CloudRouterProfileFromContainerList(cont, 0)
}

func CloudRouterProfileListFromContainer(cont *container.Container) []*CloudRouterProfile {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*CloudRouterProfile, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudRouterProfileFromContainerList(cont, i)
	}

	return arr
}
