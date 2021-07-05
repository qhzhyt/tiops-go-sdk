package config

import (
	"fmt"
	"os"
	"path"
)

var (
	NFSServer         = "172.16.3.21"
	NFSCustomPVBase   = "/export/deepops_nfs/custom_pvs"
	NFSDatasetsPVPath = path.Join(NFSCustomPVBase, "tiops-datasets")
	ApiServerHost     = fmt.Sprintf("tiops-api-server.%s", TiopsSystemNamespace)
)

func init() {
	if os.Getenv("NFS_SERVER") != "" {
		NFSServer = os.Getenv("NFS_HOST")
	}
	if TiopsSystemNamespace == "" {
		ApiServerHost = fmt.Sprintf("tiops-api-server.%s", DefaultTiopsSystemNamespace)
	} else {
		ApiServerHost = fmt.Sprintf("tiops-api-server.%s", TiopsSystemNamespace)
	}
}
