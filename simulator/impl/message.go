package impl

type MessageResult bool

const (
	MessageResultSuccess MessageResult = true
	MessageResultFailure MessageResult = false
)

type Message struct {
	Id      string
	Text    string
	OutCome MessageResult
}

func (m *Message) Verify() bool {
	return m != nil && m.Id != ""
}
