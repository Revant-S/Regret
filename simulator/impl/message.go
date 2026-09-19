package impl

type MessageResult bool

const (
	MessageResultSuccess MessageResult = true
	MessageResultFailure MessageResult = false
)

type Message struct {
	Id      *string
	Text    *string
	OutCome MessageResult
}
