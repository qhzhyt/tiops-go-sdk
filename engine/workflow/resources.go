package workflow

import (
	"path"
	"tiops/common/config"
	"tiops/common/models"
	"tiops/engine/types"
)

//func SetRuntimeConfigsToActionApp(runtimeConfig *models.RuntimeConfig, *models.A)  {
//
//}

func ResourcesPreProcess(resources *models.WorkflowResources, workflow *types.Workflow) *models.WorkflowResources {
	apps := resources.Apps

	for _, app := range apps {
		//projectInfo := workflow.Projects[app.ProjectId]

		actionInfo := workflow.GetActionInfo(app.ActionId)

		//fmt.Println(actionInfo)

		if actionInfo == nil {
			continue
		}

		//actionInfo := action.Info()

		runtimeConfig := actionInfo.RuntimeConfig

		if len(app.WorkContainers) < 1{
			app.WorkContainers = append(app.WorkContainers, &models.K8SContainer{})
		}

		mainContainer := app.WorkContainers[0]

		if mainContainer.Image == "" {
			switch actionInfo.Source {
			case models.ActionSource_FromImage:
				mainContainer.Image = actionInfo.Image
			case models.ActionSource_FromProject:
				mainContainer.Image = config.WorkflowActionServerImage(actionInfo.ProjectId)
			}
		}

		if mainContainer.ResourcesLimits == nil {
			mainContainer.ResourcesLimits = &models.ContainerResources{Extra: map[string]string{}}
		}

		if mainContainer.ResourcesRequests == nil {
			mainContainer.ResourcesRequests = &models.ContainerResources{Extra: map[string]string{}}
		}

		if runtimeConfig != nil {
			if runtimeConfig.UseAllDatasetVolumes {
				app.Volumes = map[string]string{
					"tiops-datasets": "/datasets",
				}
			} else if runtimeConfig.UseVolume && len(runtimeConfig.DatasetVolumes) > 0{
				app.Volumes = map[string]string{}
				for _, volume := range runtimeConfig.DatasetVolumes {
					app.Volumes[path.Join("tiops-datasets", volume.DatasetId)] = volume.MountPath
				}
			}

			if runtimeConfig.UseGPU {
				mainContainer.ResourcesLimits.Gpu = "1"
				mainContainer.ResourcesRequests.Gpu = "1"
			} else {
				mainContainer.ResourcesLimits.Gpu = "0"
				mainContainer.ResourcesRequests.Gpu = "0"
			}

			if len(runtimeConfig.Resources) > 0 {
				for _, resource := range runtimeConfig.Resources {
					switch resource.Name {
					case "cpu":
						mainContainer.ResourcesLimits.Cpu = resource.Limit
						mainContainer.ResourcesRequests.Cpu = resource.Request
					case "memory":
						mainContainer.ResourcesLimits.Memory = resource.Limit
						mainContainer.ResourcesRequests.Memory = resource.Request
					default:
						mainContainer.ResourcesLimits.Extra[resource.Name] = resource.Limit
						mainContainer.ResourcesRequests.Extra[resource.Name] = resource.Request
					}
				}
			}
		}
	}
	return resources
}
