package simdomain

import "Regret/domain"

type Message struct {
	domain.Message
}

func (m *Message) Verify() bool {
	return m != nil && m.Id != ""
}
