package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const ConfigsnapshotClassName = "configSnapshot"

type JobInstance struct {
	BaseAttributes
	JobInstanceAttributes
}

type JobInstanceAttributes struct {
	Name string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	NameAlias string `json:",omitempty"`

	Retire string `json:",omitempty"`
}

func NewJobInstance(configSnapshotRn, parentDn, description string, configSnapshotattr JobInstanceAttributes) *JobInstance {
	dn := fmt.Sprintf("%s/%s", parentDn, configSnapshotRn)
	return &JobInstance{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         ConfigsnapshotClassName,
			Rn:                configSnapshotRn,
		},

		JobInstanceAttributes: configSnapshotattr,
	}
}

func (configSnapshot *JobInstance) ToMap() (map[string]string, error) {
	configSnapshotMap, err := configSnapshot.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(configSnapshotMap, "name", configSnapshot.Name)

	A(configSnapshotMap, "annotation", configSnapshot.Annotation)

	A(configSnapshotMap, "nameAlias", configSnapshot.NameAlias)

	A(configSnapshotMap, "retire", configSnapshot.Retire)

	return configSnapshotMap, err
}

func JobInstanceFromContainerList(cont *container.Container, index int) *JobInstance {

	JobInstanceCont := cont.S("imdata").Index(index).S(ConfigsnapshotClassName, "attributes")
	return &JobInstance{
		BaseAttributes{
			DistinguishedName: G(JobInstanceCont, "dn"),
			Description:       G(JobInstanceCont, "descr"),
			Status:            G(JobInstanceCont, "status"),
			ClassName:         ConfigsnapshotClassName,
			Rn:                G(JobInstanceCont, "rn"),
		},

		JobInstanceAttributes{

			Name: G(JobInstanceCont, "name"),

			Annotation: G(JobInstanceCont, "annotation"),

			NameAlias: G(JobInstanceCont, "nameAlias"),

			Retire: G(JobInstanceCont, "retire"),
		},
	}
}

func JobInstanceFromContainer(cont *container.Container) *JobInstance {

	return JobInstanceFromContainerList(cont, 0)
}

func JobInstanceListFromContainer(cont *container.Container) []*JobInstance {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*JobInstance, length)

	for i := 0; i < length; i++ {

		arr[i] = JobInstanceFromContainerList(cont, i)
	}

	return arr
}
