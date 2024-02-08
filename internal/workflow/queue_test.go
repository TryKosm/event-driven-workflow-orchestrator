package workflow

import "testing"

func TestQueuePushPop(t *testing.T) {
	q := &Queue{}
	q.Push(Event{ID: "1"})
	_, ok := q.Pop()
	if !ok {
		t.Fatal("expected event")
	}
}
