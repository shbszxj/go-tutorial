package tour

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(10, 20)
	expected := 30
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
