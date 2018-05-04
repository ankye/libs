package logs

// Adapter 日志引擎适配器 console、file、conn、smtp、es、multifile
// level 日志等级
// LevelEmergency = 0
// LevelAlert = 1
// LevelCritical = 2
// LevelError = 3
// LevelWarning = 4
// LevelNotice = 5
// LevelInformational = 6
// LevelDebug = 7
// //LogFileConfig 日志文件配置
//  LogFileConfig
// 	filename" 保存的文件名
//  maxlines 每个文件保存的最大行数，默认值 1000000
// 	maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
// 	daily 是否按照每天 logrotate，默认是 true
// 	maxdays 文件最多保存多少天，默认保存 7 天
// 	rotate 是否开启 logrotate，默认是 true

// //LogMultiFileConfig 多日志文件配置 包含 LogFileConfig的属性，同时多了以下属性
// separate 需要单独写入文件的日志级别["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],设置后命名类似 test.error.log

// //LogConnConfig 网络日志配置
// reconnectOnMsg 是否每次链接都重新打开链接，默认是 false
// reconnect 是否自动重新链接地址，默认是 false
// net 发送网络链接的方式，可以使用 tcp、unix、udp 等
// addr 网络链接的地址

// //LogSMTPConfig 邮件日志配置
// username smtp 验证的用户名
// password smtp 验证密码
// host 发送的邮箱地址
// sendTos 邮件需要发送的人，支持多个
// subject 发送邮件的标题，默认是 Diagnostic message from server

//LogConfig 日志配置信息
type LogConfig struct {
	Adapter string                 `yaml:"adapter"`
	Params  map[string]interface{} `yaml:"params"`
}

//NewLogConfig 初始化日志配置
func NewLogConfig() *LogConfig {
	c := new(LogConfig)
	c.Params = make(map[string]interface{})
	return c
}
