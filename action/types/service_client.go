// @Title  client.go
// @Description  系统服务调用客户端
// @Create  heyitong  2022/6/24 16:36
// @Update  heyitong  2022/6/24 16:36

package types

import (
	"context"
	"github.com/google/uuid"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/services"
)

type ServiceActionClient struct {
	serviceName   string
	actionName    string
	nodeID        string
	actionClient  *actionClient.RemoteActionClient
	actionOptions ActionOptions
}

func (sac *ServiceActionClient) Call(ctx context.Context, batch ActionDataBatch) (ActionDataBatch, error) {

	actionDataMap := map[string]*services.ActionData{}

	for name, data := range batch.ToActionDataMap() {
		actionDataMap[name] = data.RawData()
	}

	response, err := sac.actionClient.CallAction(ctx, &services.ActionRequest{
		ActionName: sac.actionName,
		NodeId:     sac.nodeID,
		Inputs:     actionDataMap,
	})

	if err != nil {
		return nil, err
	}

	var names []string

	for n := range response.Outputs {
		names = append(names, n)
	}

	return ServicesActionDataMapToActionDataBatch(response.Outputs), nil

}

func NewServiceActionClient(serviceName, actionName string, options ActionOptions) (*ServiceActionClient, error) {

	client, err := actionClient.GetRemoteActionClient(tiopsConfigs.ServicesNamespace, serviceName)

	if err != nil {
		return nil, err
	}

	nodeID := uuid.New().String()

	_, err = client.RegisterActionNode(context.TODO(), &services.RegisterActionNodeRequest{
		ActionName:    actionName,
		NodeId:        nodeID,
		ActionInfo:    nil,
		ActionOptions: options,
	})

	if err != nil {
		return nil, err
	}

	sac := &ServiceActionClient{
		serviceName:   serviceName,
		actionName:    actionName,
		nodeID:        nodeID,
		actionClient:  client,
		actionOptions: options,
	}

	return sac, nil
}

type ServicesClient struct {
	serviceName            string
	actionName             string
	serviceActionClientMap map[string]*ServiceActionClient
}

func (c *ServicesClient) CallServiceAction(ctx context.Context, serviceName, actionName string, inputs ActionDataBatch) (ActionDataBatch, error) {
	return c.CallServiceActionWithOptions(ctx, serviceName, actionName, inputs, ActionOptions{})
}

func (c *ServicesClient) CallServiceActionWithOptions(ctx context.Context, serviceName, actionName string, inputs ActionDataBatch, options ActionOptions) (ActionDataBatch, error) {
	//clientID := actionName + "." + serviceName
	client, err := NewServiceActionClient(serviceName, actionName, options)
	if err != nil {
		return nil, err
	}
	return client.Call(ctx, inputs)
}

func (c *ServicesClient) NewServiceActionClient(serviceName, actionName string, options ActionOptions) (*ServiceActionClient, error) {
	return NewServiceActionClient(serviceName, actionName, options)
}
