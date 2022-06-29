package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateAbstractionofIPAddressBlock(to string, _from string, multicast_address_pool string, description string, fvnsMcastAddrBlkAttr models.AbstractionofIPAddressBlockAttributes) (*models.AbstractionofIPAddressBlock, error) {
	rn := fmt.Sprintf(models.RnfvnsMcastAddrBlk, _from, to)
	parentDn := fmt.Sprintf(models.ParentDnfvnsMcastAddrBlk, multicast_address_pool)
	fvnsMcastAddrBlk := models.NewAbstractionofIPAddressBlock(rn, parentDn, description, fvnsMcastAddrBlkAttr)
	err := sm.Save(fvnsMcastAddrBlk)
	return fvnsMcastAddrBlk, err
}

func (sm *ServiceManager) ReadAbstractionofIPAddressBlock(to string, _from string, multicast_address_pool string) (*models.AbstractionofIPAddressBlock, error) {
	dn := fmt.Sprintf(models.DnfvnsMcastAddrBlk, multicast_address_pool, _from, to)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvnsMcastAddrBlk := models.AbstractionofIPAddressBlockFromContainer(cont)
	return fvnsMcastAddrBlk, nil
}

func (sm *ServiceManager) DeleteAbstractionofIPAddressBlock(to string, _from string, multicast_address_pool string) error {
	dn := fmt.Sprintf(models.DnfvnsMcastAddrBlk, multicast_address_pool, _from, to)
	return sm.DeleteByDn(dn, models.FvnsmcastaddrblkClassName)
}

func (sm *ServiceManager) UpdateAbstractionofIPAddressBlock(to string, _from string, multicast_address_pool string, description string, fvnsMcastAddrBlkAttr models.AbstractionofIPAddressBlockAttributes) (*models.AbstractionofIPAddressBlock, error) {
	rn := fmt.Sprintf(models.RnfvnsMcastAddrBlk, _from, to)
	parentDn := fmt.Sprintf(models.ParentDnfvnsMcastAddrBlk, multicast_address_pool)
	fvnsMcastAddrBlk := models.NewAbstractionofIPAddressBlock(rn, parentDn, description, fvnsMcastAddrBlkAttr)
	fvnsMcastAddrBlk.Status = "modified"
	err := sm.Save(fvnsMcastAddrBlk)
	return fvnsMcastAddrBlk, err
}

func (sm *ServiceManager) ListAbstractionofIPAddressBlock(multicast_address_pool string) ([]*models.AbstractionofIPAddressBlock, error) {
	dnUrl := fmt.Sprintf("%s/uni/infra/maddrns-%s/fvnsMcastAddrBlk.json", models.BaseurlStr, multicast_address_pool)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.AbstractionofIPAddressBlockListFromContainer(cont)
	return list, err
}
