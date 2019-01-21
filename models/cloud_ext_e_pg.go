package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const CloudextepgClassName = "cloudExtEPg"

type Cloudexternalepg struct {
	BaseAttributes
    CloudexternalepgAttributes 
}
  
type CloudexternalepgAttributes struct {
    Annotation       string `json:",omitempty"`
    ExceptionTag       string `json:",omitempty"`
    FloodOnEncap       string `json:",omitempty"`
    MatchT       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    PrefGrMemb       string `json:",omitempty"`
    Prio       string `json:",omitempty"`
    RouteReachability       string `json:",omitempty"`
    
}
   

func NewCloudexternalepg(cloudExtEPgRn, parentDn, description string, cloudExtEPgattr CloudexternalepgAttributes) *Cloudexternalepg {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudExtEPgRn)  
	return &Cloudexternalepg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudextepgClassName,
			Rn:                cloudExtEPgRn,
		},
        
		CloudexternalepgAttributes: cloudExtEPgattr,
         
	}
}

func (cloudExtEPg *Cloudexternalepg) ToMap() (map[string]string, error) {
	cloudExtEPgMap, err := cloudExtEPg.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(cloudExtEPgMap, "annotation",cloudExtEPg.Annotation)
    A(cloudExtEPgMap, "exceptionTag",cloudExtEPg.ExceptionTag)
    A(cloudExtEPgMap, "floodOnEncap",cloudExtEPg.FloodOnEncap)
    A(cloudExtEPgMap, "matchT",cloudExtEPg.MatchT)
    A(cloudExtEPgMap, "nameAlias",cloudExtEPg.NameAlias)
    A(cloudExtEPgMap, "prefGrMemb",cloudExtEPg.PrefGrMemb)
    A(cloudExtEPgMap, "prio",cloudExtEPg.Prio)
    A(cloudExtEPgMap, "routeReachability",cloudExtEPg.RouteReachability)
    
	

	return cloudExtEPgMap, err
}

func CloudexternalepgFromContainerList(cont *container.Container, index int) *Cloudexternalepg {

	CloudexternalepgCont := cont.S("imdata").Index(index).S(CloudextepgClassName, "attributes")
	return &Cloudexternalepg{
		BaseAttributes{
			DistinguishedName: G(CloudexternalepgCont, "dn"),
			Description:       G(CloudexternalepgCont, "descr"),
			Status:            G(CloudexternalepgCont, "status"),
			ClassName:         CloudextepgClassName,
			Rn:                G(CloudexternalepgCont, "rn"),
		},
        
		CloudexternalepgAttributes{
        Annotation : G(CloudexternalepgCont, "annotation"),
        ExceptionTag : G(CloudexternalepgCont, "exceptionTag"),
        FloodOnEncap : G(CloudexternalepgCont, "floodOnEncap"),
        MatchT : G(CloudexternalepgCont, "matchT"),
        NameAlias : G(CloudexternalepgCont, "nameAlias"),
        PrefGrMemb : G(CloudexternalepgCont, "prefGrMemb"),
        Prio : G(CloudexternalepgCont, "prio"),
        RouteReachability : G(CloudexternalepgCont, "routeReachability"),
        		
        },
        
	}
}

func CloudexternalepgFromContainer(cont *container.Container) *Cloudexternalepg {

	return CloudexternalepgFromContainerList(cont, 0)
}

func CloudexternalepgListFromContainer(cont *container.Container) []*Cloudexternalepg {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Cloudexternalepg, length)

	for i := 0; i < length; i++ {

		arr[i] = CloudexternalepgFromContainerList(cont, i)
	}

	return arr
}