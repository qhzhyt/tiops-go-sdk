// @Title  client.go
// @Description  APIService客户端封装
// @Create  heyitong  2022/6/23 15:52
// @Update  heyitong  2022/6/23 15:52

package api_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	"tiops/common/models"
	"tiops/common/services"
)

// APIClient APIService客户端封装
type APIClient struct {
	services.APIServiceClient
	grpcConn *grpc.ClientConn
}

func (c *APIClient) Close() error {
	return c.grpcConn.Close()
}

func (c *APIClient) GetWorkflowById(id string) (*models.WorkflowInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetWorkflowById(ctx, &services.QueryRequest{Id: id})
	return result, err
}

func (c *APIClient) GetActionById(id string) (*models.ActionInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetActionById(ctx, &services.QueryRequest{Id: id})
	return result, err
}

func (c *APIClient) GetWorkflowJobById(id string) *models.WorkflowJob {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetWorkflowJobById(ctx, &services.QueryRequest{Id: id})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (c *APIClient) GetActionListByIds(ids []string) ([]*models.ActionInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetActionListByIds(ctx, &services.QueryRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	return result.List, nil
}

func (c *APIClient) GetProjectListByIds(ids []string) ([]*models.ProjectInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetProjectList(ctx, &services.QueryRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	return result.List, nil
}

func (c *APIClient) CreateOrUpdateWorkflowExecution(execution *models.WorkflowExecution) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.CreateOrUpdateWorkflowExecution(ctx, execution)
	if err != nil {
		return false, err
	}
	return result.Status == services.Status_Ok, nil
}

func (c *APIClient) CreateExecutionRecord(record *models.ExecutionRecord) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.CreateOrUpdateExecutionRecord(ctx, record)
	if err != nil {
		return false, err
	}
	return result.Status == services.Status_Ok, nil
}

func (c *APIClient) AddProcessRecord(record *models.ProcessRecord) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.AddProcessRecord(ctx, record)
	if err != nil {
		return false, err
	}
	return result.Status == services.Status_Ok, nil
}

func (c *APIClient) GetProjectByID(id string) *models.ProjectInfo {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetProjectByID(ctx, &services.QueryRequest{Id: id})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (c *APIClient) GetWorkflowExecutionById(id string) (*models.WorkflowExecution, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.APIServiceClient.GetWorkflowExecutionById(ctx, &services.QueryRequest{Id: id})
}

func NewAPIClient(ip string, port int) *APIClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := APIClient{APIServiceClient: services.NewAPIServiceClient(conn), grpcConn: conn}
	return &client
}
