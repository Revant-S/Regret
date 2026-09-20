package impl

// Evaluator owns the pool
type Evaluator struct {
	MessagePool *Pool
	Score       float64
}

type EvaluatorInterface interface {
	// EvaluateMessageScore returns a score for the message based on the carrier
	//chosen and pulls the message out of the pool
	EvaluateMessageScore(messageId *string, carrier *CarrierResponse) float64
}
