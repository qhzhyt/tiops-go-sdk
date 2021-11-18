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
	result, err := c.APIServiceClient.CreateExecutionRecord(ctx, record)
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

//func (c *APIClient) GetWorkflowExecutionResource(executionId string) (*models.K8SResources, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//	result, err := c.APIServiceClient.GetProjectByID(ctx, &services.QueryRequest{Id: executionId})
//	if err != nil {
//		log.Fatal(err)
//	}
//	return result
//}

//func (c *APIClient) GetPackageByID(id string) *models.PackageInfo {
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//
//	pg, err := c.APIServiceClient.GetPackageByID(ctx, &services.QueryRequest{Id: id})
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	if pg.Code != 1 {
//		fmt.Println("get package fail:", pg.Message, pg.Code)
//		return nil
//	}
//	result := &models.PackageInfo{}
//	ptypes.UnmarshalAny(pg.Data, result)
//	return result
//}

func NewAPIClient(ip string, port int) *APIClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := APIClient{APIServiceClient: services.NewAPIServiceClient(conn), grpcConn: conn}
	return &client
}
