package scada

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestEventType string

func TestEventBus_PublishAndSubscribe(t *testing.T) {
	bus := NewEventBus[TestEventType]()
	ctx := t.Context()

	var receivedPayload any
	ch := bus.Sub(ctx, "test", 1, func(e Event[TestEventType]) {
		receivedPayload = e.Payload
	})

	// Publish event when channel is empty
	event := NewEvent[TestEventType]("test", 42)
	bus.Pub(event)

	select {
	case e := <-ch:
		assert.Equal(t, event.Type, e.Type)
		assert.Equal(t, event.Payload, e.Payload)
	case <-time.After(time.Millisecond * 50):
		t.Fatal("expected to receive event")
	}

	// Fill channel to trigger default action
	event2 := NewEvent[TestEventType]("test", 99)
	bus.Pub(event2) // channel has size 1 and was just consumed

	// Publish another event immediately, should trigger default action
	event3 := NewEvent[TestEventType]("test", "default")
	bus.Pub(event3)

	assert.Equal(t, "default", receivedPayload)
}

func TestEventBus_UnsubscribeOnContextCancel(t *testing.T) {
	bus := NewEventBus[TestEventType]()
	ctx, cancel := context.WithCancel(context.Background())

	ch := bus.Sub(ctx, "cancelTest", 1)

	// Publish event and make sure subscriber receives it
	event := NewEvent[TestEventType]("cancelTest", 123)
	bus.Pub(event)
	select {
	case e := <-ch:
		assert.Equal(t, 123, e.Payload)
	case <-time.After(time.Millisecond * 50):
		t.Fatal("expected to receive event before cancel")
	}

	// Cancel context, subscriber channel should close
	cancel()
	time.Sleep(time.Millisecond * 50) // give goroutine time to run

	_, ok := <-ch
	assert.False(t, ok, "expected channel to be closed after context cancel")
}

func TestEventBus_NoSubscribers(t *testing.T) {
	bus := NewEventBus[TestEventType]()
	// Publish to an event type with no subscribers; should not panic
	bus.Pub(NewEvent[TestEventType]("nosubs", 0))
	assert.True(t, true, "publish with no subscribers should not panic")
}

func TestNewEvent(t *testing.T) {
	e := NewEvent[TestEventType]("type1", 123)
	assert.Equal(t, TestEventType("type1"), e.Type)
	assert.Equal(t, 123, e.Payload)
}
