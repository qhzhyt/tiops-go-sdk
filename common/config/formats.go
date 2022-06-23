// @Title  formats.go
// @Description  全局字符串格式化函数
// @Create  heyitong  2022/6/23 15:54
// @Update  heyitong  2022/6/23 15:54

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
		nodeId = nodeId[:6]
	}
	return CleanName(fmt.Sprintf(WorkflowActionServerLabelBase, name+"-"+nodeId))
}

func SystemActionServiceName(name string) string {

	return fmt.Sprintf(SystemActionServerLabelBase, CleanName(name))
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
