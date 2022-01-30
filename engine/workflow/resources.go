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
		projectInfo := workflow.Projects[app.ProjectId]

		projectConfig := projectInfo.Config

		if app.MainContainer == nil {
			app.MainContainer = &models.K8SContainer{}
		}

		if app.MainContainer.Image == "" {
			app.MainContainer.Image = config.WorkflowActionServerImage(app.ProjectId)
		}

		if app.MainContainer.ResourcesLimits == nil {
			app.MainContainer.ResourcesLimits = &models.ContainerResources{}
		}

		if app.MainContainer.ResourcesRequests == nil {
			app.MainContainer.ResourcesRequests = &models.ContainerResources{}
		}

		if projectConfig != nil {
			if projectConfig.UseAllDatasetVolumes {
				app.Volumes = map[string]string{
					"tiops-datasets": "/datasets",
				}
			} else if projectConfig.UseVolume && len(projectConfig.DatasetVolumes) > 0{
				app.Volumes = map[string]string{}
				for _, volume := range projectConfig.DatasetVolumes {
					app.Volumes[path.Join("tiops-datasets", volume.DatasetId)] = volume.MountPath
				}
			}

			if projectConfig.UseGPU {
				app.MainContainer.ResourcesLimits.Gpu = "1"
				app.MainContainer.ResourcesRequests.Gpu = "1"
			} else {
				app.MainContainer.ResourcesLimits.Gpu = "0"
				app.MainContainer.ResourcesRequests.Gpu = "0"
			}
		}
	}
	return resources
}
