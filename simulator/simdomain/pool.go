package simdomain

import (
	"errors"
	"sync"
)

type pool struct {
	unResolvedMessages map[string]*Message
	mu                 sync.RWMutex
}

type Pool interface {
	Push(message *Message) error
	Pop(message *Message) error
	GetMessageFromPool(messageId string) (*Message, error)
	Clear() error
	Size() int
}

func NewPool() Pool {
	return &pool{
		unResolvedMessages: make(map[string]*Message),
	}
}

func (p *pool) Push(message *Message) error {
	if !message.Verify() {
		return errors.New("invalid Message")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.unResolvedMessages[message.Id] = message
	return nil
}

func (p *pool) Pop(message *Message) error {
	if !message.Verify() {
		return errors.New("invalid message")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.unResolvedMessages, message.Id)
	return nil
}

func (p *pool) GetMessageFromPool(messageId string) (*Message, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	msg, exists := p.unResolvedMessages[messageId]
	if !exists {
		return nil, errors.New("message Not Found")
	}
	return msg, nil
}

func (p *pool) Clear() error {
	if p == nil {
		return errors.New("cannot clear a nil pool")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	clear(p.unResolvedMessages)
	return nil
}

func (p *pool) Size() int {
	return len(p.unResolvedMessages)
}
