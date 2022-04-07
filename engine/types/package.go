package types
//
//import (
//	//actionClient "tiops/common/action-client"
//	//tiopsConfigs "tiops/common/config"
//	actionClient "tiops/common/action-client"
//	tiopsConfigs "tiops/common/config"
//	"tiops/common/models"
//	//"workflow-engine/global"
//)
//
//type Package struct {
//	info    *models.ProjectInfo
//	actions map[string]Action
//	*actionClient.RemoteActionClient
//}
//
//func (p *Package) GetAction(name string) Action {
//	return p.actions[name]
//}
//
//func BuildPackage(info *models.ProjectInfo) *Package {
//	//fmt.Println(info)
//
//	//client :=
//	//for _, actionInfo := range info.Actions {
//	//	switch actionInfo.Type {
//	//	default:
//	//		pg.actions[actionInfo.Name] = &RemoteServiceAction{
//	//			info:   actionInfo,
//	//			client: client,
//	//		}
//	//	}
//	//}
//	//fmt.Println(pg.actions["add"])
//	return &Package{
//		info: info,
//		actions: map[string]Action{},
//		RemoteActionClient: actionClient.NewRemoteActionClient(tiopsConfigs.ActionServiceName(info.XId), tiopsConfigs.ActionServerPort),
//	}
//}
