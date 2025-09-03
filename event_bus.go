package scada

import (
	"context"
	"sync"
)

// Event represents a generic event of type T.
// Payload can be any data associated with the event.
type Event[T comparable] struct {
	Type    T
	Payload any
}

// NewEvent creates a new Event of type T with the given payload.
func NewEvent[T comparable](eventType T, payload any) Event[T] {
	return Event[T]{eventType, payload}
}

// Subscriber holds a channel to receive events and default actions
// to execute when the channel is full.
type Subscriber[EventType comparable] struct {
	ch             chan Event[EventType]
	defaultActions []func(Event[EventType])
}

// EventBus manages events and subscribers, ensuring concurrency safety
// via RWMutex and lifecycle management with context.
type EventBus[EventType comparable] struct {
	sync.RWMutex
	subscribers map[EventType][]*Subscriber[EventType]
}

// NewEventBus creates and returns a new, empty EventBus.
func NewEventBus[EventType comparable]() *EventBus[EventType] {
	subscribers := make(map[EventType][]*Subscriber[EventType])
	return &EventBus[EventType]{subscribers: subscribers}
}

// Pub publishes an event to all registered subscribers.
// If a subscriber's channel is full, its defaultActions are executed.
func (eb *EventBus[T]) Pub(event Event[T]) {
	eb.Lock()
	subs, ok := eb.subscribers[event.Type]
	eb.Unlock()

	if !ok {
		return
	}

	for _, s := range subs {
		select {
		case s.ch <- event:
		default:
			for _, action := range s.defaultActions {
				action(event)
			}
		}
	}
}

// Sub registers a subscriber for a specific event type.
// ctx controls the subscriber's lifecycle, automatically closing
// the channel when the context is canceled.
// chanSize specifies the channel buffer size.
// defaultActions are executed if the channel is full.
func (eb *EventBus[T]) Sub(
	ctx context.Context,
	eventType T,
	chanSize int,
	defaultActions ...func(Event[T]),
) <-chan Event[T] {
	sub := &Subscriber[T]{
		ch:             make(chan Event[T], chanSize),
		defaultActions: defaultActions,
	}

	go func() {
		<-ctx.Done()
		eb.unsub(eventType, sub)
	}()

	eb.Lock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], sub)
	eb.Unlock()

	return sub.ch
}

// unsub removes the subscriber from the EventBus and closes its channel.
// This function is intended for internal use only.
func (eb *EventBus[T]) unsub(eventType T, sub *Subscriber[T]) {
	eb.Lock()
	defer eb.Unlock()

	subs := eb.subscribers[eventType]
	for i, s := range subs {
		if s == sub {
			eb.subscribers[eventType] = append(subs[:i], subs[i+1:]...)
			close(s.ch)
			break
		}
	}
}
