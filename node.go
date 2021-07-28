package dropnode

import (
	"sync"

	core "github.com/markoxley/dropcore"
)

type Node struct {
	name      string
	central   NodeAddress
	nodes     map[string][]NodeAddress
	connected bool
	inLock    sync.Mutex
	outLock   sync.Mutex
	inQueue   *core.ThreadSafeRingBuffer
	outQueue  *core.ThreadSafeRingBuffer
}

func (n *Node) Start(name string, central NodeAddress, work func(int)) error {
	n.inQueue = core.NewTSRingBuffer(256, true)
	n.outQueue = core.NewTSRingBuffer(256, true)
	go n.mainThread(work)
	return nil
}

func (n *Node) Send(msg *core.Message) {
	n.outQueue.Push(msg)
}

func (n *Node) addIncoming(msg *core.Message) {
	n.inQueue.Push(msg)
}

func (n *Node) popIncoming() (*core.Message, bool) {
	msg, ok := n.inQueue.Pop()
	if !ok {
		return nil, false
	}
	if gmsg, ok := msg.(*messages.Message); ok {
		return gmsg, true
	}
	return nil, false
}

func (n *Node) popOutgoing() (*core.Message, bool) {
	msg, ok := n.outQueue.Pop()
	if !ok {
		return nil, false
	}
	if gmsg, ok := msg.(*core.Message); ok {
		return gmsg, true
	}
	return nil, false
}

func (n *Node) mainThread(work func(int)) {

}
