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

func (c *APIClient) GetWorkflowById(id string) *models.WorkflowInfo {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := c.APIServiceClient.GetWorkflowById(ctx, &services.QueryRequest{Id: id})
	if err != nil {
		log.Fatal(err)
	}
	return result
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
