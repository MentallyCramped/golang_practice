package mascot_test

import (
	"testing"

	"github.com/MentallyCramped/golang_practice/playground/mascot_example/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Boop" {
		t.Fatal("Boop HAS to be the best mascot!")
	}
}
