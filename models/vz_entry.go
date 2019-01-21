package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VzentryClassName = "vzEntry"

type Filterentry struct {
	BaseAttributes
    FilterentryAttributes 
}
  
type FilterentryAttributes struct {
    Annotation       string `json:",omitempty"`
    ApplyToFrag       string `json:",omitempty"`
    ArpOpc       string `json:",omitempty"`
    DFromPort       string `json:",omitempty"`
    DToPort       string `json:",omitempty"`
    EtherT       string `json:",omitempty"`
    Icmpv4T       string `json:",omitempty"`
    Icmpv6T       string `json:",omitempty"`
    MatchDscp       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Prot       string `json:",omitempty"`
    SFromPort       string `json:",omitempty"`
    SToPort       string `json:",omitempty"`
    Stateful       string `json:",omitempty"`
    TcpRules       string `json:",omitempty"`
    
}
   

func NewFilterentry(vzEntryRn, parentDn, description string, vzEntryattr FilterentryAttributes) *Filterentry {
	dn := fmt.Sprintf("%s/%s", parentDn, vzEntryRn)  
	return &Filterentry{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VzentryClassName,
			Rn:                vzEntryRn,
		},
        
		FilterentryAttributes: vzEntryattr,
         
	}
}

func (vzEntry *Filterentry) ToMap() (map[string]string, error) {
	vzEntryMap, err := vzEntry.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(vzEntryMap, "annotation",vzEntry.Annotation)
    A(vzEntryMap, "applyToFrag",vzEntry.ApplyToFrag)
    A(vzEntryMap, "arpOpc",vzEntry.ArpOpc)
    A(vzEntryMap, "dFromPort",vzEntry.DFromPort)
    A(vzEntryMap, "dToPort",vzEntry.DToPort)
    A(vzEntryMap, "etherT",vzEntry.EtherT)
    A(vzEntryMap, "icmpv4T",vzEntry.Icmpv4T)
    A(vzEntryMap, "icmpv6T",vzEntry.Icmpv6T)
    A(vzEntryMap, "matchDscp",vzEntry.MatchDscp)
    A(vzEntryMap, "nameAlias",vzEntry.NameAlias)
    A(vzEntryMap, "prot",vzEntry.Prot)
    A(vzEntryMap, "sFromPort",vzEntry.SFromPort)
    A(vzEntryMap, "sToPort",vzEntry.SToPort)
    A(vzEntryMap, "stateful",vzEntry.Stateful)
    A(vzEntryMap, "tcpRules",vzEntry.TcpRules)
    
	

	return vzEntryMap, err
}

func FilterentryFromContainerList(cont *container.Container, index int) *Filterentry {

	FilterentryCont := cont.S("imdata").Index(index).S(VzentryClassName, "attributes")
	return &Filterentry{
		BaseAttributes{
			DistinguishedName: G(FilterentryCont, "dn"),
			Description:       G(FilterentryCont, "descr"),
			Status:            G(FilterentryCont, "status"),
			ClassName:         VzentryClassName,
			Rn:                G(FilterentryCont, "rn"),
		},
        
		FilterentryAttributes{
        Annotation : G(FilterentryCont, "annotation"),
        ApplyToFrag : G(FilterentryCont, "applyToFrag"),
        ArpOpc : G(FilterentryCont, "arpOpc"),
        DFromPort : G(FilterentryCont, "dFromPort"),
        DToPort : G(FilterentryCont, "dToPort"),
        EtherT : G(FilterentryCont, "etherT"),
        Icmpv4T : G(FilterentryCont, "icmpv4T"),
        Icmpv6T : G(FilterentryCont, "icmpv6T"),
        MatchDscp : G(FilterentryCont, "matchDscp"),
        NameAlias : G(FilterentryCont, "nameAlias"),
        Prot : G(FilterentryCont, "prot"),
        SFromPort : G(FilterentryCont, "sFromPort"),
        SToPort : G(FilterentryCont, "sToPort"),
        Stateful : G(FilterentryCont, "stateful"),
        TcpRules : G(FilterentryCont, "tcpRules"),
        		
        },
        
	}
}

func FilterentryFromContainer(cont *container.Container) *Filterentry {

	return FilterentryFromContainerList(cont, 0)
}

func FilterentryListFromContainer(cont *container.Container) []*Filterentry {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Filterentry, length)

	for i := 0; i < length; i++ {

		arr[i] = FilterentryFromContainerList(cont, i)
	}

	return arr
}