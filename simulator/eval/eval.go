package eval

import (
	"Regret/domain"
	"Regret/simulator/client"
	"Regret/simulator/simdomain"
)

type evaluator struct {
	pool      simdomain.Pool
	client    client.SimulatorClient
	bestScore int64
}
type Evaluator interface {
	EvaluateMessage(messageID string, selectedCarrier domain.Carrier) float64
	GetBestScore() int64
	UpdateBestScore(score int64)
	PoolMessage(message *simdomain.Message) error
}

func (e *evaluator) EvaluateMessage(messageID string, selectedCarrier domain.Carrier) float64 {
	return 2.3
}

func (e *evaluator) GetBestScore() int64 {
	return e.bestScore
}
func (e *evaluator) UpdateBestScore(score int64) {
	e.bestScore = score
}

func (e *evaluator) PoolMessage(message *simdomain.Message) error {
	err := e.pool.Push(message)
	if err != nil {
		return err
	}
	return nil
}

func NewEvaluator(pool simdomain.Pool, simulatorClient client.SimulatorClient) Evaluator {
	return &evaluator{
		pool:   pool,
		client: simulatorClient,
	}
}
