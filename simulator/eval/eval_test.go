package eval

import (
	"errors"
	"testing"

	"Regret/domain"
	"Regret/simulator/simdomain"
)

type mockPool struct {
	pushFn func(message *simdomain.Message) error
}

func (m *mockPool) Push(message *simdomain.Message) error {
	if m.pushFn != nil {
		return m.pushFn(message)
	}
	return nil
}
func (m *mockPool) Pop(message *simdomain.Message) error                            { return nil }
func (m *mockPool) GetMessageFromPool(messageId string) (*simdomain.Message, error) { return nil, nil }
func (m *mockPool) Clear() error                                                    { return nil }
func (m *mockPool) Size() int                                                       { return 0 }

type mockSimulatorClient struct{}

func (m *mockSimulatorClient) SendMessage(message *domain.Message) *domain.Carrier {
	return &domain.Carrier{}
}

func TestNewEvaluator(t *testing.T) {
	mp := &mockPool{}
	mc := &mockSimulatorClient{}

	ev := NewEvaluator(mp, mc)
	if ev == nil {
		t.Fatal("Expected NewEvaluator to return an instance, got nil")
	}
}

func TestEvaluator_EvaluateMessage(t *testing.T) {
	ev := NewEvaluator(&mockPool{}, &mockSimulatorClient{})
	expected := 2.3
	actual := ev.EvaluateMessage("test-msg-id", domain.Carrier{})

	if actual != expected {
		t.Errorf("EvaluateMessage() = %v, expected %v", actual, expected)
	}
}

func TestEvaluator_BestScore(t *testing.T) {
	ev := NewEvaluator(&mockPool{}, &mockSimulatorClient{})

	if score := ev.GetBestScore(); score != 0 {
		t.Errorf("Expected initial BestScore to be 0, got %v", score)
	}
	expectedScore := int64(150)
	ev.UpdateBestScore(expectedScore)

	if score := ev.GetBestScore(); score != expectedScore {
		t.Errorf("GetBestScore() = %v, expected %v", score, expectedScore)
	}
}

func TestEvaluator_PoolMessage(t *testing.T) {
	t.Run("Successfully pools message", func(t *testing.T) {
		mp := &mockPool{
			pushFn: func(message *simdomain.Message) error {
				return nil
			},
		}
		ev := NewEvaluator(mp, &mockSimulatorClient{})

		msg := &simdomain.Message{Message: domain.Message{Id: "msg-123"}}
		err := ev.PoolMessage(msg)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("Fails to pool message", func(t *testing.T) {
		expectedErr := errors.New("pool is full or invalid message")
		mp := &mockPool{
			pushFn: func(message *simdomain.Message) error {
				return expectedErr
			},
		}
		ev := NewEvaluator(mp, &mockSimulatorClient{})

		msg := &simdomain.Message{Message: domain.Message{Id: "msg-123"}}
		err := ev.PoolMessage(msg)

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("Expected error '%v', got '%v'", expectedErr, err)
		}
	})
}
