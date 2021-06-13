/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-31 23:21:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-14 03:33:30
 */

package reverseproxy

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/leiyudaye/gateway/discover"
	lib "github.com/leiyudaye/gateway/lib/comm"
	"github.com/leiyudaye/gateway/log"

	"google.golang.org/grpc"
)

func NewGinForGrpcReverseProxy(gParams lib.GatewayParams) (string, error) {
	var (
		req           *dynamic.Message
		rsp           *dynamic.Message
		srvMethod     string = gParams.Module.Method
		srvFullMethod string = "/" + gParams.Module.Module + "/" + gParams.Module.Method
		reqName       string = srvMethod + "Req"
		rspName       string = srvMethod + "Rsp"
	)

	// 服务发现
	disConn, err := discover.NewDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		log.Error("discover connect failed. err=%v", err)
		return "", nil
	}

	tagAddr, err := disConn.Discover(gParams.Module.Module)
	if err != nil {
		log.Error(err.Error())
		return "", nil
	}

	// grpc 代理
	conn, err := grpc.Dial(tagAddr, grpc.WithInsecure())
	if err != nil {
		return "", nil
	}
	defer conn.Close()

	// 读取proto文件
	var parser protoparse.Parser
	fileDescriptors, err := parser.ParseFiles("./protobuf/user.proto")
	if err != nil {
		fmt.Printf("err=%v", err)
	}
	fileDescriptor := fileDescriptors[0]
	for _, msgDescriptor := range fileDescriptor.GetMessageTypes() {
		fmt.Println(msgDescriptor.GetName())
		if msgDescriptor.GetName() == reqName {
			req = dynamic.NewMessage(msgDescriptor)
			for k, v := range gParams.Module.Param.(map[string]interface{}) {
				if reflect.TypeOf(v).Kind().String() == "float64" {
					req.SetFieldByName(k, int32(v.(float64)))
					continue
				}
				if err := req.TrySetFieldByName(k, v); err != nil {
					fmt.Println(err)
					return "", nil
				}
				req.SetFieldByName(k, v)
			}
		}

		if msgDescriptor.GetName() == rspName {
			rsp = dynamic.NewMessage(msgDescriptor)
		}
		// for _, fieldDesc := range msgDescriptor.GetFields() {
		// 	if fieldDesc.GetName() == "userID" {
		// 		req.TrySetFieldByName("userID", fieldDesc.GetType())
		// 		req.SetFieldByName("userID", interface{}(1))
		// 	}
		// }
	}

	err = conn.Invoke(context.Background(), srvFullMethod, req, rsp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
	bt, _ := rsp.MarshalJSON()
	return string(bt), nil
}
