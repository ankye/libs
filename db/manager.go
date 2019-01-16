package db

import (
	"errors"
	"sync"

	"github.com/gonethopper/libs/ds/queue"
	"github.com/gonethopper/libs/grpc/pb"
)

var (
	onceDB     sync.Once
	instanceDB *Manager
)

//GetDBManagerInstance db manager singleton
func GetDBManagerInstance() *Manager {
	onceDB.Do(func() {
		instanceDB = NewDBManager()
	})
	return instanceDB
}

//NewDBManager create DBManager
func NewDBManager() *Manager {
	s := new(Manager)

	return s
}

//Manager Manager 链接管理
type Manager struct {
	Master    *SQLConnection
	Slaver    *SQLConnection
	Consumer  Consumer //消费处理
	QueueSize uint32   //队列大小
	Queue     *queue.Queue
}

//Register 注册消费者，用于消息回调
func (m *Manager) Register(c Consumer, queueSize uint32) {
	m.Consumer = c
	m.QueueSize = queueSize
	m.Queue = queue.NewQueue(int64(1), queueSize, false, m)
}

//Consume 数据从队列里面获取数据发送
func (m *Manager) Consume(queueID, lowerSequence, upperSequence int64) {
	if m.Queue == nil {
		return
	}
	for sequence := lowerSequence; sequence <= upperSequence; sequence++ {
		message := m.Queue.Pop(sequence) // see performance note on producer sample above
		//fmt.Println("get message:", message)
		if message != nil {

		}
		//	time.Sleep(1 * time.Second)
		// handle the incoming message with your application code
	}
}

//Produce 发送消息交给Stream处理
func (m *Manager) Produce(message *pb.SSMessage) error {
	if m.Queue == nil {
		return errors.New("queue is not init")
	}
	return m.Queue.Push(message)
}

//AddConnection 添加db连接
func (m *Manager) AddConnection(isMaster bool, driver, dsn, timezone string) error {

	if driver == "" || dsn == "" || timezone == "" {
		return errors.New("sql config is empty")
	}
	conn := NewSQLConnection(driver, dsn, timezone)
	if err := conn.Ping(); err != nil {
		return err
	}
	if isMaster {
		m.Master = conn
	} else {
		m.Slaver = conn
	}
	return nil
}

//Connect 检测链接
func (m *Manager) Connect() error {

	if m.Master != nil {
		if err := m.Master.Ping(); err != nil {
			return err
		}
	}

	if m.Slaver != nil {
		if err := m.Slaver.Ping(); err != nil {
			return err
		}
	}

	return nil
}

//GetMaster 获取Master链接
func (m *Manager) GetMaster() *SQLConnection {
	return m.Master
}

//GetSlaver 获取Slaver链接
func (m *Manager) GetSlaver() *SQLConnection {
	return m.Slaver
}
