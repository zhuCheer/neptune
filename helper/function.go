package helper

import (
	"crypto/md5"
	"fmt"
	"encoding/json"
)

type baseJsonResultFormat struct{
	Info string `json:"info"`
	Data interface{} `json:"data"`
	Code int `json:"code"`
}

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

// grpc 返回 json 通用格式
func GrpcResponseJson(info string, data interface{}, code int) ([]byte){
	result := baseJsonResultFormat{info, data, code}
	jsonStr, err :=json.Marshal(result)
	if err != nil{
		return []byte("{code:-1,info:'json decode error'}")
	}
	return jsonStr

}

// grpc 返回 成功 json 通用格式
func GrpcResponseSuccess(data interface{}) ([]byte){
	result := baseJsonResultFormat{"success", data, 0}
	jsonStr, err :=json.Marshal(result)
	if err != nil{
		return []byte("{code:-1,info:'json decode error'}")
	}
	return jsonStr

}
