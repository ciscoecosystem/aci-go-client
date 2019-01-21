package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudepgClassName = "cloudEPg"

type Cloudepg struct {
	BaseAttributes
    CloudepgAttributes 
}
  
type CloudepgAttributes struct {
    Annotation       string `json:",omitempty"`
    ExceptionTag       string `json:",omitempty"`
    FloodOnEncap       string `json:",omitempty"`
    MatchT       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    PrefGrMemb       string `json:",omitempty"`
    Prio       string `json:",omitempty"`
    
}
   

func NewCloudepg(cloudEPgRn, parentDn, description string, cloudEPgattr CloudepgAttributes) *Cloudepg {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudEPgRn)  
	return &Cloudepg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudepgClassName,
			Rn:                cloudEPgRn,
		},
        
		CloudepgAttributes: cloudEPgattr,
         
	}
}

func (cloudEPg *Cloudepg) ToMap() (map[string]string, error) {
	cloudEPgMap, err := cloudEPg.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudEPgMap, "annotation",cloudEPg.Annotation)
    A(cloudEPgMap, "exceptionTag",cloudEPg.ExceptionTag)
    A(cloudEPgMap, "floodOnEncap",cloudEPg.FloodOnEncap)
    A(cloudEPgMap, "matchT",cloudEPg.MatchT)
    A(cloudEPgMap, "nameAlias",cloudEPg.NameAlias)
    A(cloudEPgMap, "prefGrMemb",cloudEPg.PrefGrMemb)
    A(cloudEPgMap, "prio",cloudEPg.Prio)
    
	

	return cloudEPgMap, err
}

func CloudepgFromContainerList(cont *container.Container, index int) *Cloudepg {

	CloudepgCont := cont.S("imdata").Index(index).S(CloudepgClassName, "attributes")
	return &Cloudepg{
		BaseAttributes{
			DistinguishedName: G(CloudepgCont, "dn"),
			Description:       G(CloudepgCont, "descr"),
			Status:            G(CloudepgCont, "status"),
			ClassName:         CloudepgClassName,
			Rn:                G(CloudepgCont, "rn"),
		},
        
		CloudepgAttributes{
        Annotation : G(CloudepgCont, "annotation"),
        ExceptionTag : G(CloudepgCont, "exceptionTag"),
        FloodOnEncap : G(CloudepgCont, "floodOnEncap"),
        MatchT : G(CloudepgCont, "matchT"),
        NameAlias : G(CloudepgCont, "nameAlias"),
        PrefGrMemb : G(CloudepgCont, "prefGrMemb"),
        Prio : G(CloudepgCont, "prio"),
        		
        },
        
	}
}

func CloudepgFromContainer(cont *container.Container) *Cloudepg {

	return CloudepgFromContainerList(cont, 0)
}

func CloudepgListFromContainer(cont *container.Container) []*Cloudepg {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudepg, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudepgFromContainerList(cont, i)
	}

	return arr
}