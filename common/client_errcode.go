package common

//Error Coce ID define
//Client ERROR Code ID 划分区域 0x0001 -- 0x4FFF
const (
	//客户端错误，一般是一些传参，校验错误等
	CSErrorCodeClientError = 0x0001
	//系统错误，一般表示系统未知错误，或者不希望用户了解具体情形的错误
	CSErrorCodeSystemError = 0x0002
	//系统超时
	CSErrorCodeTimeout = 0x0003
)
