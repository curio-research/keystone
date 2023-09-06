package server

import (
	"testing"
)

func TestTickTime(t *testing.T) {
	tickRate := 80
	tick := 13

	shouldTick := ShouldTriggerTick(tick, tickRate, 1000)

	if !shouldTick {
		t.Errorf("Should have ticked")
	}
}
