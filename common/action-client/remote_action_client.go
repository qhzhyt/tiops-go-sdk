package action_client

import (
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"math"
	"time"
	"tiops/common/services"
)

type RemoteActionClient struct {
	services.ActionsServiceClient
	conn *grpc.ClientConn
}


func NewRemoteActionClient(ip string, port int) *RemoteActionClient {

	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	retryOps := []grpc_retry.CallOption{
		grpc_retry.WithMax(2),
		grpc_retry.WithPerRetryTimeout(time.Minute * 20),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinearWithJitter(time.Second/2, 0.2)),
	}
	retryInterceptor := grpc_retry.UnaryClientInterceptor(retryOps...)


	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32), grpc.MaxCallSendMsgSize(math.MaxInt32)),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithChainUnaryInterceptor(retryInterceptor),
	)

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
