package workflow

import (
	"path"
	"tiops/common/config"
	"tiops/common/models"
	"tiops/engine/types"
)

func ResourcesPreProcess(resources *models.WorkflowResources, workflow *types.Workflow) *models.WorkflowResources {
	apps := resources.Apps

	for _, app := range apps {
		//projectInfo := workflow.Projects[app.ProjectId]

		action := workflow.GetAction(app.ActionId)

		if action == nil {
			continue
		}

		actionInfo := action.Info()

		runtimeConfig := action.Info().RuntimeConfig

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
			mainContainer.ResourcesLimits = &models.ContainerResources{}
		}

		if mainContainer.ResourcesRequests == nil {
			mainContainer.ResourcesRequests = &models.ContainerResources{}
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
		}
	}
	return resources
}
