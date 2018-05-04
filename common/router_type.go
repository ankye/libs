package common

//ServiceServer service server info
type ServiceServer struct {
	ServerID   int    `yaml:"serverID"`
	ServerAddr string `yaml:"serverAddr"`
}

//RouterServer 服务器结构定义
type RouterServer struct {
	DestType int           `yaml:"destType"`
	Master   ServiceServer `yaml:"master"`
	Slaver   ServiceServer `yaml:"slaver"`
}
