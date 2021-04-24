package config

import (
	"os"
	"path"
)

var (
	NFSServer         = "172.16.3.21"
	NFSCustomPVBase   = "/export/deepops_nfs/custom_pvs"
	NFSDatasetsPVPath = path.Join(NFSCustomPVBase, "tiops-datasets")
)

func init() {
	if os.Getenv("NFS_SERVER") != "" {
		NFSServer = os.Getenv("NFS_HOST")
	}
}