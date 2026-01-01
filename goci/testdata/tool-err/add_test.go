package add

import "testing"

func TestAdd(t *testing.T) {
	exp := 3
	a := add(1, 2)

	if a != exp {
		t.Errorf("Expected %d, got %d", exp, a)
	}
}
