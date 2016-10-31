package event

import (
	"testing"

	"github.com/CyCoreSystems/gami"
)

func TestChannelEventLog(t *testing.T) {
	fixture := map[string]string{
		"Eventname": "EventName",
	}

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
