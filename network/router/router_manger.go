package router

import (
	"sync"
	"time"

	"github.com/gonethopper/libs/common"
	"github.com/gonethopper/libs/grpc/pb"
	"github.com/gonethopper/libs/logs"
)

var (
	onceRM     sync.Once
	instanceRM *RouterManager
)

//GetRouterManagerInstance router manager singleton
func GetRouterManagerInstance() *RouterManager {
	onceRM.Do(func() {
		instanceRM = NewRouterManager()
	})
	return instanceRM
}

//NewRouterManager create RouterManager
func NewRouterManager() *RouterManager {
	s := new(RouterManager)
	s.Connections = make(map[int]*RouterConnection)
	s.Masters = make([]int, common.MaxEntityTypeNumber)
	s.Slavers = make([]int, common.MaxEntityTypeNumber)

	return s
}

//RouterManager router clients manager
type RouterManager struct {
	//存放router connection连接，以server id为key，多个类型路由到同一台服务器保持一个链接
	Connections map[int]*RouterConnection
	//存放主Router链接的映射关系, 第一层 server type 为key，第二层server ID 为key
	Masters []int
	//存放从Router链接的映射关系, 第一层 server type 为key，第二层server ID 为key
	Slavers   []int
	Consumer  Consumer //消费处理
	QueueSize uint32   //队列大小
}

//Register 注册消费者，用于消息回调
func (r *RouterManager) Register(c Consumer, queueSize uint32) {
	r.Consumer = c
	r.QueueSize = queueSize
}

//Consume 消费服务器回来的消息
func (r *RouterManager) Consume(m *pb.SSMessage) {
	if r.Consumer != nil {
		r.Consumer.ConsumerRouterMessage(m)
	}
}

//GetLocalServerID 获取本地服务器ID
func (r *RouterManager) GetLocalServerID() int {
	if r.Consumer != nil {
		return r.Consumer.GetServerID()
	}
	return -1
}

//addConnection 添加一个连接
func (r *RouterManager) addConnection(serverID int, address string, queueSize uint32) error {
	conn := r.Connections[serverID]
	if conn == nil {
		conn = NewRouterConnection(serverID, address, queueSize)
		r.Connections[serverID] = conn
		return nil

	}
	return logs.ErrorR("RouterManager already exist connection id[%d] address[%s]", serverID, address)
}

//AddMaster 添加主Router节点
func (r *RouterManager) AddMaster(serverID int, serverType int, address string) error {

	if err := r.addConnection(serverID, address, r.QueueSize); err != nil {
		logs.Error(err)
	}

	if r.Masters[serverType] == 0 {
		r.Masters[serverType] = serverID
	} else {

		return logs.ErrorR("RouterManager addMaster type[%d] ID[%#x] addr[%s] already exist,please check config", serverType, serverID, address)
	}

	logs.Info("RouterManager addMaster type[%d] id[%d] addr[%s] success.", serverType, serverID, address)
	return nil
}

//TryConnect 尝试连接
func (r *RouterManager) TryConnect() {
	for serverID, conn := range r.Connections {
		if !conn.IsConnected() {
			if err := conn.Connect(); err != nil {
				logs.Error("Try Connect failed ID[%#x] Address[%s], try again", serverID, conn.Address)
			} else {
				logs.Info("Try Connect success ID[%#x] Address[%s].", serverID, conn.Address)
			}

		}
	}
}

//CheckConnections 检测链接存活状态
func (r *RouterManager) CheckConnections() {

	for _ = range time.Tick(60 * time.Second) {
		r.TryConnect()
	}

}

//AddSlaver 添加从Router节点
func (r *RouterManager) AddSlaver(serverID int, serverType int, address string) error {
	if err := r.addConnection(serverID, address, r.QueueSize); err != nil {
		logs.Error(err)
	}
	if r.Slavers[serverType] == 0 {
		r.Slavers[serverType] = serverID
	} else {

		return logs.ErrorR("RouterManager addSlaver type[%d] id[%d] addr[%s] already exist,please check config", serverType, serverID, address)
	}

	logs.Info("RouterManager addSlaver type[%d] id[%d] addr[%s] success.", serverType, serverID, address)
	return nil
}

//GetRouterConnection 获取Router connection
func (r *RouterManager) GetRouterConnection(serverType int) *RouterConnection {
	if serverType >= common.MaxEntityTypeNumber {
		logs.Error("server type is invalid. %d", serverType)
		return nil
	}
	serverID := r.Masters[serverType]
	if serverID > 0 {
		conn := r.Connections[serverID]
		if conn.Stream != nil {
			return conn
		}
	}
	serverID = r.Slavers[serverType]
	if serverID > 0 {
		conn := r.Connections[serverID]
		if conn.Stream != nil {
			return conn
		}
	}
	return nil
}
