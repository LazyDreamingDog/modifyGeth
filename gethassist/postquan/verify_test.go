package postquan

import (
	"testing"

	"github.com/ethereum/go-ethereum/event"
)

type testEvent int

func TestSub(t *testing.T) {
	mux := new(event.TypeMux)
	defer mux.Stop()

	// Subscribe
	sub := mux.Subscribe(testEvent(0))
	go func() {
		if err := mux.Post(testEvent(5)); err != nil {
			t.Errorf("Post returned unexpected error: %v", err)
		}
	}()

	// Bind
	ev := <-sub.Chan()

	if ev.Data.(testEvent) != testEvent(5) {
		t.Errorf("Got %v (%T), expected event %v (%T)",
			ev, ev, testEvent(5), testEvent(5))
	}
}
