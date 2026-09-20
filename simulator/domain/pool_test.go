package domain

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPool_PushAndGet(t *testing.T) {

	pool := NewPool()
	message := &Message{
		Id:      "msg123",
		Text:    "This is a Test Message",
		OutCome: MessageResultFailure,
	}
	err := pool.Push(message)
	if err != nil {
		t.Fatalf("expected no error on push got : %v", err)
	}

	if pool.Size() != 1 {
		t.Fatalf("expected size of pool to be 1 got : %v", pool.Size())
	}
	msg, err2 := pool.GetMessageFromPool("msg123")
	if err2 != nil {
		t.Fatalf("expected no error on GetMessageFromPool function got: %v", err2)
	}

	if !cmp.Equal(msg, message) {
		t.Fatalf("expected message not recieved from the pool : %v", cmp.Diff(msg, message))
	}

}

func TestPool_Pop(t *testing.T) {
	pool := NewPool()
	message := &Message{
		Id:      "msg123",
		Text:    "This is a Test Message for Pop",
		OutCome: MessageResultSuccess,
	}
	err := pool.Push(message)
	if err != nil {
		t.Fatalf("expected no error on push got : %v", err)
	}

	errPop := pool.Pop(message)
	if errPop != nil {
		t.Fatalf("expected no error on pop got: %v", errPop)
	}
	msg, errGet := pool.GetMessageFromPool("msg123")
	if errGet == nil {
		t.Fatalf("expected error on GetMessageFromPool after pop, but got nil. Message still in pool: %v", msg)
	}
}

func TestPool_ValidationErrors(t *testing.T) {
	pool := NewPool()
	t.Run("Push missing ID: ", func(t *testing.T) {
		msg := &Message{Id: ""}
		if err := pool.Push(msg); err == nil {
			t.Fatalf("expected error on pushing empty Id, but got nil")
		}
	})

	t.Run("Push nil Message: ", func(t *testing.T) {
		if err := pool.Push(nil); err == nil {
			t.Fatalf("expected error on pushing nil message, but got nil")
		}
	})
}

func TestPool_Clear(t *testing.T) {
	pool := NewPool()
	message := &Message{
		Id:      "msg123",
		Text:    "This is a Test Message for Pop",
		OutCome: MessageResultSuccess,
	}
	err := pool.Push(message)
	if err != nil {
		t.Fatalf("expected no error on push got : %v", err)
	}
	err = pool.Clear()
	if err != nil {
		t.Fatalf("expected no error while clearing pool got: %v", err)
	}
	if pool.Size() != 0 {
		t.Fatalf("expected size of pool after clear to be 0 got : %v", pool.Size())
	}
}
