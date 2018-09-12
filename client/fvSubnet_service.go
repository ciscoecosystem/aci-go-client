package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSubnet(ip string ,bridge_domain string ,tenant string  ,description string, fvSubnetattr models.SubnetAttributes) (*models.Subnet, error) {	
	rn := fmt.Sprintf("subnet-[%s]",ip)
	parentDn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,bridge_domain )
	fvSubnet := models.NewSubnet(rn, parentDn, description, fvSubnetattr)
	err := sm.Save(fvSubnet)
	return fvSubnet, err
}

func (sm *ServiceManager) ReadSubnet(ip string ,bridge_domain string ,tenant string ) (*models.Subnet, error) {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s/subnet-[%s]", tenant ,bridge_domain ,ip )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvSubnet := models.SubnetFromContainer(cont)
	return fvSubnet, nil
}

func (sm *ServiceManager) DeleteSubnet(ip string ,bridge_domain string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s/subnet-[%s]", tenant ,bridge_domain ,ip )
	return sm.DeleteByDn(dn, models.FvsubnetClassName)
}

func (sm *ServiceManager) UpdateSubnet(ip string ,bridge_domain string ,tenant string  ,description string, fvSubnetattr models.SubnetAttributes) (*models.Subnet, error) {
	rn := fmt.Sprintf("subnet-[%s]",ip)
	parentDn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,bridge_domain )
	fvSubnet := models.NewSubnet(rn, parentDn, description, fvSubnetattr)

    fvSubnet.Status = "modified"
	err := sm.Save(fvSubnet)
	return fvSubnet, err

}

func (sm *ServiceManager) ListSubnet(bridge_domain string ,tenant string ) ([]*models.Subnet, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/BD-%s/fvSubnet.json", baseurlStr , tenant ,bridge_domain )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.SubnetListFromContainer(cont)

	return list, err
}