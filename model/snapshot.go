package model

type CreateVmSnapshotRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Params   CreateVmSnapshotParams `json:"params"`
}

type CreateVmSnapshotParams struct {
	CreateVmSnapshot CreateVmSnapshot `json:"createImage"`
}

type CreateVmSnapshot struct {
	Name             string `json:"name"`
	InstanceSnapshot string `json:"instance_snapshot"` //true,false
}

type RevertVmSnapshotRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Params   RevertVmSnapshotParams `json:"params"`
}

type RevertVmSnapshotParams struct {
	RevertVmSnapshot RevertVmSnapshot `json:"rebuild"`
}

type RevertVmSnapshot struct {
	Description string `json:"description,omitempty"`
	ImageRef    string `json:"imageRef"` //快照id
}

type CreateVolumeSnapshotRequest struct {
	Domain   string                     `json:"domain"`
	Token    string                     `json:"token"`
	TenantId string                     `json:"tenantId"`
	Params   CreateVolumeSnapshotParams `json:"params"`
}

type CreateVolumeSnapshotParams struct {
	CreateVolumeSnapshot CreateVolumeSnapshot `json:"snapshot"`
}

type CreateVolumeSnapshot struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	VolumeId    string `json:"volume_id"`
	Force       bool   `json:"force"`
}

type RevertVolumeSnapshotRequest struct {
	Domain     string                     `json:"domain"`
	Token      string                     `json:"token"`
	TenantId   string                     `json:"tenantId"`
	SnapshotId string                     `json:"snapshotId"`
	Params     RevertVolumeSnapshotParams `json:"params"`
}

type RevertVolumeSnapshotParams struct {
	RevertVolumeSnapshot RevertVolumeSnapshot `json:"rollback"`
}

type RevertVolumeSnapshot struct {
	Name     string `json:"name"`
	VolumeId string `json:"volume_id"`
}

type DeleteSnapshotRequest struct {
	Domain     string `json:"domain"`
	Token      string `json:"token"`
	TenantId   string `json:"tenantId"`
	SnapshotId string `json:"snapshotId"`
}

type QuerySnapshotRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	VolumeId string `json:"volume_id"`
}
