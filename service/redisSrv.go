package service

import (
	pb "github.com/zhuCheer/neptune/grpc"
	"golang.org/x/net/context"
	"fmt"
	"github.com/zhuCheer/neptune/helper"
	"time"
	"log"
	"github.com/zhuCheer/neptune/connect"
	"encoding/json"
)

type RedisSrv struct{}
type doValueInfo []interface{}

func (this *RedisSrv) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Message: "PONG " + in.Name}, nil
}

// 连接池注册
func (this *RedisSrv) DoRegist(ctx context.Context, in *pb.DoRegistRequest) (*pb.DoResponse, error){
	address := in.GetAddress()
	auth := in.GetAuth()
	initCap := in.GetInitialCap()
	maxCap := in.GetMaxCap()
	idleTimeout := in.GetIdleTimeout()
	dbNum := in.GetDBNum()
	if maxCap < initCap{
		return &pb.DoResponse{Response:helper.GrpcResponseJson(helper.GrpcErrCodeTran[5001],"", 5001)}, nil
	}

	cnf := &connect.RegistCnf{
		Address: address,
		Auth: auth,
		DBNum: int(dbNum),
		InitialCap: int(initCap),
		MaxCap: int(maxCap),
		IdleTimeout: time.Duration(idleTimeout),
	}
	poolId,err := connect.Regist(cnf)
	if err != nil{
		return &pb.DoResponse{Response:helper.GrpcResponseJson(helper.GrpcErrCodeTran[5002],"", 5002)}, nil
	}


	fmt.Println(address, auth, initCap, maxCap, idleTimeout)
	return &pb.DoResponse{Response:helper.GrpcResponseJson("success", poolId, 0 )}, nil
}

// 释放连接池
func (this *RedisSrv) DoRelease(ctx context.Context, in *pb.DoReleaseRequest)(reps *pb.DoResponse, err error){
	regId := in.GetRegistId()
	poolHandler := connect.GlobalRedisHandler
	poolHandler.ReleasePool(regId)
	return &pb.DoResponse{Response:helper.GrpcResponseSuccess("success")}, nil
}

func (this *RedisSrv) DoFun(ctx context.Context, in *pb.DoRequest) (reps *pb.DoResponse, err error){
	regId := in.GetRegistId()
	doValue:=in.GetParams()
	doName := in.GetName()

	poolHandler := connect.GlobalRedisHandler
	connItem, err := poolHandler.GetConn(regId)
	defer func() {
		fmt.Println("put back a conn")
		poolHandler.PutConn(connItem)
	}()

	if err != nil{
		return &pb.DoResponse{Response:helper.GrpcResponseJson(helper.GrpcErrCodeTran[5003],"", 5003)}, nil
	}

	info := doValueInfo{}
	err = json.Unmarshal(doValue, &info)
	if err != nil{
		return &pb.DoResponse{Response:helper.GrpcResponseJson(helper.GrpcErrCodeTran[-1],"", -1)}, nil
	}
	log.Println("do redis", string(doName), len(info))
	rs, err := connItem.Conn.Do(string(doName), info...)
	if err != nil{
		return &pb.DoResponse{Response:helper.GrpcResponseJson(helper.GrpcErrCodeTran[5004], err.Error(), 5004)}, nil
	}
	showData:=this.formatShowReps(rs)


	return &pb.DoResponse{Response:helper.GrpcResponseSuccess(showData)}, nil
}

// 格式化Redis返回结果
func (this *RedisSrv) formatShowReps(rs interface{}) (showReps interface{}){
	var showData interface{}
	switch rs.(type) {
		case string:
			showData = rs.(string)
		case int32:
			showData = rs.(int32)
		case int64:
			showData = rs.(int64)
		case []interface{}:
			tmpSlice := make([]string, 0)
			for _,item := range rs.([]interface{}){
				tmpSlice = append(tmpSlice, string(item.([]byte)))
				showData = tmpSlice
			}
		default:
			showData = string(rs.([]byte))
	}

	return showData
}