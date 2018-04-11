package logs

//LogFileConfig 日志文件配置
type LogFileConfig struct {
	Filename string `yaml:"filename"` //保存的文件名
	Maxlines int    `yaml:"maxlines"` //每个文件保存的最大行数，默认值 1000000
	Maxsize  int    `yaml:"maxsize"`  //每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily    bool   `yaml:"daily"`    //是否按照每天 logrotate，默认是 true
	Maxdays  int    `yaml:"maxdays"`  //文件最多保存多少天，默认保存 7 天
	Rotate   bool   `yaml:"rotate"`   //是否开启 logrotate，默认是 true
}

//LogMultiFileConfig 多日志文件配置
type LogMultiFileConfig struct {
	Separate string `yaml:"separate"` //需要单独写入文件的日志级别,设置后命名类似 test.error.log
}

//LogConnConfig 网络日志配置
type LogConnConfig struct {
	ReconnectOnMsg bool   `yaml:"reconnectOnMsg"` //是否每次链接都重新打开链接，默认是 false
	Reconnect      bool   `yaml:"reconnect"`      // 是否自动重新链接地址，默认是 false
	Net            string `yaml:"net"`            //发开网络链接的方式，可以使用 tcp、unix、udp 等
	Addr           string `yaml:"addr"`           // 网络链接的地址
}

//LogSMTPConfig 邮件日志配置
type LogSMTPConfig struct {
	Username string `yaml:"username"` //smtp 验证的用户名
	Password string `yaml:"password"` //smtp 验证密码
	Host     string `yaml:"host"`     //发送的邮箱地址
	SendTos  string `yaml:"sendTos"`  //邮件需要发送的人，支持多个
	Subject  string `yaml:"subject"`  //发送邮件的标题，默认是 Diagnostic message from server
}

//LogConfig 日志配置信息
type LogConfig struct {
	Adapter   string              `yaml:"adapter"`   // console file multifile conn smtp
	Level     int                 `yaml:"level"`     //日志保存的时候的级别，默认是 Trace 级别
	File      *LogFileConfig      `yaml:"file"`      //file 配置节点
	Multifile *LogMultiFileConfig `yaml:"multifile"` //multifile 配置节点
	Conn      *LogConnConfig      `yaml:"conn"`      //conn 配置节点
	Smtp      *LogSMTPConfig      `yaml:"smtp"`      //smtp 配置节点
}

//DefaultLogConfig 初始化默认配置文件
func (c *LogConfig) DefaultLogConfig() {
	c.Adapter = AdapterFile
	c.Level = 7
	c.File.Filename = "server.log"
	c.File.Daily = true
	c.File.Maxdays = 7
	c.File.Maxlines = 1000000
	c.File.Maxsize = 1024 * 1024 * 256
	c.File.Rotate = true
	c.Multifile.Separate = "server.error.log"
	c.Conn.ReconnectOnMsg = true
	c.Conn.Reconnect = true
	c.Conn.Net = "tcp"
	c.Conn.Addr = ""
	c.Smtp.Username = ""
	c.Smtp.Host = ""
	c.Smtp.SendTos = ""
	c.Smtp.Subject = ""

}

//NewLogConfig 初始化日志配置
func NewLogConfig() *LogConfig {
	c := new(LogConfig)
	c.File = new(LogFileConfig)
	c.Multifile = new(LogMultiFileConfig)
	c.Conn = new(LogConnConfig)
	c.Smtp = new(LogSMTPConfig)
	c.DefaultLogConfig()
	return c
}
