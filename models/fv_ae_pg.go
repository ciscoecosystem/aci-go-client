package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const FvaepgClassName = "fvAEPg"

type Applicationepg struct {
	BaseAttributes
    ApplicationepgAttributes 
}
  
type ApplicationepgAttributes struct {
    Annotation       string `json:",omitempty"`
    ExceptionTag       string `json:",omitempty"`
    FloodOnEncap       string `json:",omitempty"`
    FwdCtrl       string `json:",omitempty"`
    HasMcastSource       string `json:",omitempty"`
    IsAttrBasedEPg       string `json:",omitempty"`
    MatchT       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    PcEnfPref       string `json:",omitempty"`
    PrefGrMemb       string `json:",omitempty"`
    Prio       string `json:",omitempty"`
    Shutdown       string `json:",omitempty"`
    
}
   

func NewApplicationepg(fvAEPgRn, parentDn, description string, fvAEPgattr ApplicationepgAttributes) *Applicationepg {
	dn := fmt.Sprintf("%s/%s", parentDn, fvAEPgRn)  
	return &Applicationepg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         FvaepgClassName,
			Rn:                fvAEPgRn,
		},
        
		ApplicationepgAttributes: fvAEPgattr,
         
	}
}

func (fvAEPg *Applicationepg) ToMap() (map[string]string, error) {
	fvAEPgMap, err := fvAEPg.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(fvAEPgMap, "annotation",fvAEPg.Annotation)
    A(fvAEPgMap, "exceptionTag",fvAEPg.ExceptionTag)
    A(fvAEPgMap, "floodOnEncap",fvAEPg.FloodOnEncap)
    A(fvAEPgMap, "fwdCtrl",fvAEPg.FwdCtrl)
    A(fvAEPgMap, "hasMcastSource",fvAEPg.HasMcastSource)
    A(fvAEPgMap, "isAttrBasedEPg",fvAEPg.IsAttrBasedEPg)
    A(fvAEPgMap, "matchT",fvAEPg.MatchT)
    A(fvAEPgMap, "nameAlias",fvAEPg.NameAlias)
    A(fvAEPgMap, "pcEnfPref",fvAEPg.PcEnfPref)
    A(fvAEPgMap, "prefGrMemb",fvAEPg.PrefGrMemb)
    A(fvAEPgMap, "prio",fvAEPg.Prio)
    A(fvAEPgMap, "shutdown",fvAEPg.Shutdown)
    
	

	return fvAEPgMap, err
}

func ApplicationepgFromContainerList(cont *container.Container, index int) *Applicationepg {

	ApplicationepgCont := cont.S("imdata").Index(index).S(FvaepgClassName, "attributes")
	return &Applicationepg{
		BaseAttributes{
			DistinguishedName: G(ApplicationepgCont, "dn"),
			Description:       G(ApplicationepgCont, "descr"),
			Status:            G(ApplicationepgCont, "status"),
			ClassName:         FvaepgClassName,
			Rn:                G(ApplicationepgCont, "rn"),
		},
        
		ApplicationepgAttributes{
        Annotation : G(ApplicationepgCont, "annotation"),
        ExceptionTag : G(ApplicationepgCont, "exceptionTag"),
        FloodOnEncap : G(ApplicationepgCont, "floodOnEncap"),
        FwdCtrl : G(ApplicationepgCont, "fwdCtrl"),
        HasMcastSource : G(ApplicationepgCont, "hasMcastSource"),
        IsAttrBasedEPg : G(ApplicationepgCont, "isAttrBasedEPg"),
        MatchT : G(ApplicationepgCont, "matchT"),
        NameAlias : G(ApplicationepgCont, "nameAlias"),
        PcEnfPref : G(ApplicationepgCont, "pcEnfPref"),
        PrefGrMemb : G(ApplicationepgCont, "prefGrMemb"),
        Prio : G(ApplicationepgCont, "prio"),
        Shutdown : G(ApplicationepgCont, "shutdown"),
        		
        },
        
	}
}

func ApplicationepgFromContainer(cont *container.Container) *Applicationepg {

	return ApplicationepgFromContainerList(cont, 0)
}

func ApplicationepgListFromContainer(cont *container.Container) []*Applicationepg {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Applicationepg, length)

	for i := 0; i < length; i++ {

		arr[i] = ApplicationepgFromContainerList(cont, i)
	}

	return arr
}