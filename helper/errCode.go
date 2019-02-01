package helper
// GRPC 服务接口错误代码对应中文含义
var GrpcErrCodeTran = map[int]string{
	-1 : "系统出小差",
	5001:"初始化参数不符合要求",
	5002:"连接池创建失败",
	5003:"未找到连接池",
	5004:"redis执行错误",
}
