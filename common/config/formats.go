package config

import (
	"fmt"
	"regexp"
	"strings"
)

func WorkflowActionServerLabel(name string) string {
	return fmt.Sprintf(WorkflowActionServerLabelBase, name)
}

func WorkflowTestJobId(workflowId string) string {
	return fmt.Sprintf("wf-job-%s-test", workflowId)
}

func ActionServiceName(name string) string {
	return fmt.Sprintf(WorkflowActionServerLabelBase, name)
}

func StandAloneActionServiceName(name string, nodeId string) string {
	if len(nodeId) > 8 {
		nodeId = nodeId[:8]
	}
	return CleanName(fmt.Sprintf(WorkflowActionServerLabelBase, name+"-"+nodeId))
}

func WorkflowActionServerImage(name string) string {
	return fmt.Sprintf(WorkflowActionServerImageBase, name)
}

func WorkflowEngineImage(project string, name string) string {
	return fmt.Sprintf(WorkflowActionServerImageBase, project)
}

var (
	nameCleanRe = regexp.MustCompile(`[^-a-z0-9]`)
)

func CleanName(name string) string {
	name = strings.ToLower(name)
	name = nameCleanRe.ReplaceAllString(name, "-")

	//runes.

	for len(name) > 0 && !(name[0] >= 'a' && name[0] <= 'z') {
		name = name[1:]
	}

	return name
}
