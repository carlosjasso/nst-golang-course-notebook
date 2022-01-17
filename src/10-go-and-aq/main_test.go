package main

import (
	"fmt"
	"testing"
)

func TestTemperature_Format(t *testing.T) {
	var currTemp Temperature = 2.75
	expected := fmt.Sprintf("%.2f °%s", currTemp, "C")
	actual := currTemp.Format(Celciius)
	if actual != expected {
		t.Fatalf("currTemp.Format(Celciius) = %q, want %q", actual, expected)
	}

}

func BenchmarkTemperature_format(b *testing.B) {
	var currTemp Temperature = 4.35
	for i := 0; i < b.N; i++ {
		_ = currTemp.Format(Celciius)
	}
}
