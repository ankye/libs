package common

//Error Coce ID define
//Server ERROR Code ID 划分区域 0x5000 -- 0xFFFF
const (
	//服务器错误，一般是一些传参，校验错误等
	SSErrorCodeServerError = 0x5001
	//系统错误，一般表示系统未知错误
	SSErrorCodeSystemError = 0x5002
	//请求超时
	SSErrorCodeTimeout = 0x5003
)
