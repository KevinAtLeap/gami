package event

import (
	"testing"

	"github.com/bit4bit/gami"
)

func TestChannelEventLog(t *testing.T) {
	fixture := map[string]string{}

	ev := gami.AMIEvent{
		ID:     "CEL",
		Params: fixture,
	}

	evtype := New(&ev)
	if _, ok := (evtype).(ChannelEventLog); !ok {
		t.Fatal("ChannelEventLog assertion")
	}

	testEvent(t, fixture, evtype)
}
