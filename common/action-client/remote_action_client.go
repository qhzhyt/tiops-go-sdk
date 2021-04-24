package action_client

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"tiops/common/services"
)

type RemoteActionClient struct {
	services.ActionsServiceClient
	conn *grpc.ClientConn
}

func NewRemoteActionClient(ip string, port int) *RemoteActionClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	// 如果要增加 Recv 可以接受的一个消息的数据量，必须增加 grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000))
	//conn, err := grpc.Dial("unix:///var/lib/test.socket", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000)))
	if err != nil {
		log.Fatalf("fail to connect server: %v", err)
	}

	client := &RemoteActionClient{ActionsServiceClient: services.NewActionsServiceClient(conn), conn: conn}
	//defer conn.Close()
	//client :=

	//ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	//defer cancel()
	//startTime := time.Now().Unix()
	//n := 0
	//for i := 0 ; i < 10000; i++ {
	//
	//	go func() {
	//		_, err := client.CallAction(ctx, &ActionRequest{Name: "square", Args: &ActionArguments{Type: ArgumentsType_JSON, Value: "[1]"}})
	//
	//		if err != nil {
	//			fmt.Printf("error %v", err)
	//			return
	//		}
	//		n++
	//	}()
	//
	//	//fmt.Printf("%v\n", stream)
	//	if i % 1000 == 999 {
	//		fmt.Println(i, "/", 10000)
	//	}
	//}
	//
	//for true {
	//	time.Sleep(time.Millisecond * 500)
	//	if n > 9990 {
	//		break
	//	}
	//}
	//
	//fmt.Println("used time", time.Now().Unix() - startTime)

	return client
}
