/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-31 23:21:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-21 00:04:34
 */

package reverseproxy

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/leiyudaye/gateway/discover"
	"github.com/leiyudaye/gateway/lib/comm"
	"github.com/leiyudaye/gateway/lib/util"
	"github.com/leiyudaye/gateway/log"
	"google.golang.org/grpc"
)

var (
	gFilenames []string               // proto文件列表名
	gFileDescs []*desc.FileDescriptor // protobuf文件描述
)

func init() {
	// 读取proto文件
	var parser protoparse.Parser
	gFilenames = util.GetFileNameByPath("./protobuf")
	if len(gFilenames) == 0 {
		log.Error("read protobuf file failed")
		return
	}

	var err error
	gFileDescs, err = parser.ParseFiles(gFilenames...)
	if err != nil {
		log.Error("parser failed, file=%v, err=%v", gFilenames, err)
	}
	// TODO， 启定时器循环读
}

func convertType(fieldDesc *desc.FieldDescriptor, v interface{}) interface{} {
	switch fieldDesc.GetType().String() {
	case "TYPE_DOUBLE":
	case "TYPE_FLOAT":
	case "TYPE_INT64":
		if reflect.TypeOf(v).Kind().String() == "float64" {
			return int64(v.(float64))
		}
	case "TYPE_UINT64":
		if reflect.TypeOf(v).Kind().String() == "float64" {
			return uint64(v.(float64))
		}
	case "TYPE_INT32":
		if reflect.TypeOf(v).Kind().String() == "float64" {
			return int32(v.(float64))
		}
	case "TYPE_FIXED64":
	case "TYPE_FIXED32":
	case "TYPE_BOOL":
	case "TYPE_STRING":
	case "TYPE_GROUP":
	case "TYPE_MESSAGE":
	case "TYPE_BYTES":
	case "TYPE_UINT32":
	case "TYPE_ENUM":
	case "TYPE_SFIXED32":
	case "TYPE_SFIXED64":
	case "TYPE_SINT32":
	case "TYPE_SINT64":
	default:
	}
	return v
}

func convertObject(object map[string]interface{}, srvReq *dynamic.Message) error {
	for k, v := range object {
		if reflect.TypeOf(v).Kind().String() != "map" ||
			reflect.TypeOf(v).Kind().String() != "list" {
			fieldDesc := srvReq.FindFieldDescriptorByName(k)
			if fieldDesc == nil {
				log.Error("no found this field, k=%v", k)
				return fmt.Errorf("no found field, [%v]", k)
			}
			val := convertType(fieldDesc, v)
			if err := srvReq.TrySetFieldByName(k, val); err != nil {
				log.Error("try set field failed, err=%v", err)
				return fmt.Errorf("try set field failed, [%v]", val)
			}
			srvReq.SetFieldByName(k, val)
		} else {
			return convertObject(v.(map[string]interface{}), srvReq)
		}
	}
	return nil
}

func NewGinForGrpcReverseProxy(gParams comm.GatewayParams) (comm.GatewayRsp, error) {
	var (
		srvReq         *dynamic.Message                              // grpc服务请求参数
		srvRsp         *dynamic.Message                              // grpc服务返回参数
		srvModule      = gParams.Module.Module                       // 服务方法名
		srvMethod      = gParams.Module.Method                       // 服务方法名
		srvFullMethod  = fmt.Sprintf("/%v/%v", srvModule, srvMethod) // 服务的全量名称
		srvReqName     = srvMethod + "Req"                           // 请求参数名称
		srvRspName     = srvMethod + "Rsp"                           // 返回参数名称
		srvFullReqName = strings.Split(gParams.Module.Module, ".")[0] + "." + srvReqName
		srvFullRspName = strings.Split(gParams.Module.Module, ".")[0] + "." + srvRspName
		gatewayRsp     comm.GatewayRsp // 返回参数
	)

	// 服务发现
	disConn, err := discover.NewDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		log.Error("discover connect failed. err=%v", err)
		gatewayRsp.Code = comm.ErrDiscoverFailed
		return gatewayRsp, fmt.Errorf("discover connect failed. err=%v", err)
	}

	tagAddr, err := disConn.Discover(gParams.Module.Module)
	if err != nil {
		log.Error(err.Error())
		gatewayRsp.Code = comm.ErrDiscoverFailed
		return gatewayRsp, err
	}

	for _, fileDesc := range gFileDescs {
		fmt.Println(fileDesc)

		// 请求Req参数
		msgDesc := fileDesc.FindMessage(srvFullReqName)
		if msgDesc == nil {
			log.Error("no found message, reqName=%v", srvFullReqName)
			continue
		}
		srvReq = dynamic.NewMessage(msgDesc)
		err := convertObject(gParams.Module.Param.(map[string]interface{}), srvReq)
		if err != nil {
			gatewayRsp.Code = comm.ErrCovertFailed
			return gatewayRsp, err
		}

		// 返回Rsp参数M
		msgDesc = fileDesc.FindMessage(srvFullRspName)
		if msgDesc == nil {
			log.Error("no found message, rspNmae=%v", srvFullRspName)
			return gatewayRsp, fmt.Errorf("no found message, rspNmae=%v", srvFullRspName)
		}
		srvRsp = dynamic.NewMessage(msgDesc)

	}

	if srvReq == nil || srvRsp == nil {
		gatewayRsp.Code = comm.ErrNoFoundReqField
		return gatewayRsp, fmt.Errorf("no found req or rsp")
	}

	// grpc 代理
	conn, err := grpc.Dial(tagAddr, grpc.WithInsecure())
	if err != nil {
		gatewayRsp.Code = comm.ErrNetworkConnectFailed
		return gatewayRsp, fmt.Errorf("grpc dial failed")
	}
	defer conn.Close()

	err = conn.Invoke(context.Background(), srvFullMethod, srvReq, srvRsp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(srvRsp)
	bt, _ := srvRsp.MarshalJSON()
	gatewayRsp.Code = 0
	gatewayRsp.Data = string(bt)
	return gatewayRsp, nil
}
