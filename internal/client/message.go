package client

import (
	"container/list"
	"sync"
)

type MessageHash struct {
	mx   sync.RWMutex
	Hash map[string]*list.Element
	List *list.List
}

func NewMessageHash() *MessageHash {
	return &MessageHash{
		mx:   sync.RWMutex{},
		Hash: make(map[string]*list.Element),
		List: list.New(),
	}
}

type Node struct {
	P       int32
	Message []byte
}

func (m *MessageHash) Add(id string, n Node) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.Hash[id] = m.List.PushBack(n)
}

func (m *MessageHash) Get(id string) (*list.Element, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	node, ok := m.Hash[id]

	return node, ok
}

func (m *MessageHash) Delete(id string) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.List.Remove(m.Hash[id])
	delete(m.Hash, id)
	return
}
