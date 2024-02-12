package workflow

import "testing"

func TestMetricsShape(t *testing.T) {
	m := Metrics{Processed: 1}
	if m.Processed != 1 {
		t.Fatal("bad metric")
	}
}
