package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvrsdomattClassName = "fvRsDomAtt"

type Domain struct {
	BaseAttributes
	DomainAttributes
}

type DomainAttributes struct {
	TDn string `json:",omitempty"`

	Annotation string `json:",omitempty"`

	BindingType string `json:",omitempty"`

	ClassPref string `json:",omitempty"`

	Delimiter string `json:",omitempty"`

	Encap string `json:",omitempty"`

	EncapMode string `json:",omitempty"`

	EpgCos string `json:",omitempty"`

	EpgCosPref string `json:",omitempty"`

	InstrImedcy string `json:",omitempty"`

	LagPolicyName string `json:",omitempty"`

	NetflowDir string `json:",omitempty"`

	NetflowPref string `json:",omitempty"`

	NumPorts string `json:",omitempty"`

	PortAllocation string `json:",omitempty"`

	PrimaryEncap string `json:",omitempty"`

	PrimaryEncapInner string `json:",omitempty"`

	ResImedcy string `json:",omitempty"`

	SecondaryEncapInner string `json:",omitempty"`

	SwitchingMode string `json:",omitempty"`
}

func NewDomain(fvRsDomAttRn, parentDn, description string, fvRsDomAttattr DomainAttributes) *Domain {
	dn := fmt.Sprintf("%s/%s", parentDn, fvRsDomAttRn)
	return &Domain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvrsdomattClassName,
			Rn:                fvRsDomAttRn,
		},

		DomainAttributes: fvRsDomAttattr,
	}
}

func (fvRsDomAtt *Domain) ToMap() (map[string]string, error) {
	fvRsDomAttMap, err := fvRsDomAtt.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(fvRsDomAttMap, "tDn", fvRsDomAtt.TDn)

	A(fvRsDomAttMap, "annotation", fvRsDomAtt.Annotation)

	A(fvRsDomAttMap, "bindingType", fvRsDomAtt.BindingType)

	A(fvRsDomAttMap, "classPref", fvRsDomAtt.ClassPref)

	A(fvRsDomAttMap, "delimiter", fvRsDomAtt.Delimiter)

	A(fvRsDomAttMap, "encap", fvRsDomAtt.Encap)

	A(fvRsDomAttMap, "encapMode", fvRsDomAtt.EncapMode)

	A(fvRsDomAttMap, "epgCos", fvRsDomAtt.EpgCos)

	A(fvRsDomAttMap, "epgCosPref", fvRsDomAtt.EpgCosPref)

	A(fvRsDomAttMap, "instrImedcy", fvRsDomAtt.InstrImedcy)

	A(fvRsDomAttMap, "lagPolicyName", fvRsDomAtt.LagPolicyName)

	A(fvRsDomAttMap, "netflowDir", fvRsDomAtt.NetflowDir)

	A(fvRsDomAttMap, "netflowPref", fvRsDomAtt.NetflowPref)

	A(fvRsDomAttMap, "numPorts", fvRsDomAtt.NumPorts)

	A(fvRsDomAttMap, "portAllocation", fvRsDomAtt.PortAllocation)

	A(fvRsDomAttMap, "primaryEncap", fvRsDomAtt.PrimaryEncap)

	A(fvRsDomAttMap, "primaryEncapInner", fvRsDomAtt.PrimaryEncapInner)

	A(fvRsDomAttMap, "resImedcy", fvRsDomAtt.ResImedcy)

	A(fvRsDomAttMap, "secondaryEncapInner", fvRsDomAtt.SecondaryEncapInner)

	A(fvRsDomAttMap, "switchingMode", fvRsDomAtt.SwitchingMode)

	return fvRsDomAttMap, err
}

func DomainFromContainerList(cont *container.Container, index int) *Domain {

	DomainCont := cont.S("imdata").Index(index).S(FvrsdomattClassName, "attributes")
	return &Domain{
		BaseAttributes{
			DistinguishedName: G(DomainCont, "dn"),
			Description:       G(DomainCont, "descr"),
			Status:            G(DomainCont, "status"),
			ClassName:         FvrsdomattClassName,
			Rn:                G(DomainCont, "rn"),
		},

		DomainAttributes{

			TDn: G(DomainCont, "tDn"),

			Annotation: G(DomainCont, "annotation"),

			BindingType: G(DomainCont, "bindingType"),

			ClassPref: G(DomainCont, "classPref"),

			Delimiter: G(DomainCont, "delimiter"),

			Encap: G(DomainCont, "encap"),

			EncapMode: G(DomainCont, "encapMode"),

			EpgCos: G(DomainCont, "epgCos"),

			EpgCosPref: G(DomainCont, "epgCosPref"),

			InstrImedcy: G(DomainCont, "instrImedcy"),

			LagPolicyName: G(DomainCont, "lagPolicyName"),

			NetflowDir: G(DomainCont, "netflowDir"),

			NetflowPref: G(DomainCont, "netflowPref"),

			NumPorts: G(DomainCont, "numPorts"),

			PortAllocation: G(DomainCont, "portAllocation"),

			PrimaryEncap: G(DomainCont, "primaryEncap"),

			PrimaryEncapInner: G(DomainCont, "primaryEncapInner"),

			ResImedcy: G(DomainCont, "resImedcy"),

			SecondaryEncapInner: G(DomainCont, "secondaryEncapInner"),

			SwitchingMode: G(DomainCont, "switchingMode"),
		},
	}
}

func DomainFromContainer(cont *container.Container) *Domain {

	return DomainFromContainerList(cont, 0)
}

func DomainListFromContainer(cont *container.Container) []*Domain {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Domain, length)

	for i := 0; i < length; i++ {

		arr[i] = DomainFromContainerList(cont, i)
	}

	return arr
}
