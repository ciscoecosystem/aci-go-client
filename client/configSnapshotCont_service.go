package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateConfigBackupSnapshotContainer(pDn string, description string, configSnapshotContattr models.ConfigBackupSnapshotContainerAttributes) (*models.ConfigBackupSnapshotContainer, error) {
	rn := fmt.Sprintf("backupst/snapshots-[%s]", pDn)
	parentDn := fmt.Sprintf("uni")
	configSnapshotCont := models.NewConfigBackupSnapshotContainer(rn, parentDn, description, configSnapshotContattr)
	err := sm.Save(configSnapshotCont)
	return configSnapshotCont, err
}

func (sm *ServiceManager) ReadConfigBackupSnapshotContainer(pDn string) (*models.ConfigBackupSnapshotContainer, error) {
	dn := fmt.Sprintf("uni/backupst/snapshots-[%s]", pDn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	configSnapshotCont := models.ConfigBackupSnapshotContainerFromContainer(cont)
	return configSnapshotCont, nil
}

func (sm *ServiceManager) DeleteConfigBackupSnapshotContainer(pDn string) error {
	dn := fmt.Sprintf("uni/backupst/snapshots-[%s]", pDn)
	return sm.DeleteByDn(dn, models.ConfigsnapshotcontClassName)
}

func (sm *ServiceManager) UpdateConfigBackupSnapshotContainer(pDn string, description string, configSnapshotContattr models.ConfigBackupSnapshotContainerAttributes) (*models.ConfigBackupSnapshotContainer, error) {
	rn := fmt.Sprintf("backupst/snapshots-[%s]", pDn)
	parentDn := fmt.Sprintf("uni")
	configSnapshotCont := models.NewConfigBackupSnapshotContainer(rn, parentDn, description, configSnapshotContattr)

	configSnapshotCont.Status = "modified"
	err := sm.Save(configSnapshotCont)
	return configSnapshotCont, err

}

func (sm *ServiceManager) ListConfigBackupSnapshotContainer() ([]*models.ConfigBackupSnapshotContainer, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/configSnapshotCont.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.ConfigBackupSnapshotContainerListFromContainer(cont)

	return list, err
}
