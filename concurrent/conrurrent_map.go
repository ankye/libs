package concurrent

import (
	"io"
	"sync"

	log "github.com/gonethopper/libs/logs"
)

type ConcurrentMap struct {
	sync.RWMutex
	disposeFlag bool
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
	Items       map[interface{}]interface{}
}

func NewCocurrentMap() *ConcurrentMap {
	m := &ConcurrentMap{}
	m.Items = make(map[interface{}]interface{})
	m.disposeFlag = false
	return m
}

func (m *ConcurrentMap) Dispose() {
	m.disposeOnce.Do(func() {
		m.disposeFlag = true
		m.Lock()
		for key, value := range m.Items {
			delete(m.Items, key)
			//删除一个减少一个
			m.disposeWait.Done()
			var err error
			switch value.(type) {
			case io.Closer:
				closer := value.(io.Closer)
				err = closer.Close()
			default:
			}
			if err != nil {
				log.Error("concurrent map :dispose map error:%d", key)
			}
		}
		m.Unlock()
		m.disposeWait.Wait()
	})
}

func (m *ConcurrentMap) Get(key interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	item, ok := m.Items[key]
	if !ok {
		return nil
	}
	return item
}

// 插入的时候，不允许插入空的value
func (m *ConcurrentMap) Set(key, value interface{}) {

	if value == nil {

		log.Error("concurrent map :set map nil value key[%v]", key)
		return
	}

	m.Lock()
	defer m.Unlock()

	m.Items[key] = value
	m.disposeWait.Add(1)
}

func (m *ConcurrentMap) Del(key interface{}) {
	item := m.Get(key)
	if item == nil {
		return
	}
	m.Lock()
	delete(m.Items, key)
	m.disposeWait.Done()
	m.Unlock()
}
