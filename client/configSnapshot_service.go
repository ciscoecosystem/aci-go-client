package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateJobInstance(name string, config_backup_snapshot_container_pDn string, description string, configSnapshotattr models.JobInstanceAttributes) (*models.JobInstance, error) {
	rn := fmt.Sprintf("snapshot-%s", name)
	parentDn := fmt.Sprintf("uni/backupst/snapshots-[%s]", config_backup_snapshot_container_pDn)
	configSnapshot := models.NewJobInstance(rn, parentDn, description, configSnapshotattr)
	err := sm.Save(configSnapshot)
	return configSnapshot, err
}

func (sm *ServiceManager) ReadJobInstance(name string, config_backup_snapshot_container_pDn string) (*models.JobInstance, error) {
	dn := fmt.Sprintf("uni/backupst/snapshots-[%s]/snapshot-%s", config_backup_snapshot_container_pDn, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	configSnapshot := models.JobInstanceFromContainer(cont)
	return configSnapshot, nil
}

func (sm *ServiceManager) DeleteJobInstance(name string, config_backup_snapshot_container_pDn string) error {
	dn := fmt.Sprintf("uni/backupst/snapshots-[%s]/snapshot-%s", config_backup_snapshot_container_pDn, name)
	return sm.DeleteByDn(dn, models.ConfigsnapshotClassName)
}

func (sm *ServiceManager) UpdateJobInstance(name string, config_backup_snapshot_container_pDn string, description string, configSnapshotattr models.JobInstanceAttributes) (*models.JobInstance, error) {
	rn := fmt.Sprintf("snapshot-%s", name)
	parentDn := fmt.Sprintf("uni/backupst/snapshots-[%s]", config_backup_snapshot_container_pDn)
	configSnapshot := models.NewJobInstance(rn, parentDn, description, configSnapshotattr)

	configSnapshot.Status = "modified"
	err := sm.Save(configSnapshot)
	return configSnapshot, err

}

func (sm *ServiceManager) ListJobInstance(config_backup_snapshot_container_pDn string) ([]*models.JobInstance, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/backupst/snapshots-[%s]/configSnapshot.json", baseurlStr, config_backup_snapshot_container_pDn)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.JobInstanceListFromContainer(cont)

	return list, err
}
