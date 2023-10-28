package plugin

import (
	"errors"
	"testing"
)

func TestVoid(t *testing.T) {
	twoReturns := func(err error) (int, error) {
		return 0, err
	}
	threeReturns := func(err error) (int, int, error) {
		return 0, 0, err
	}

	testErr := errors.New("test error")

	if err := Void(twoReturns(nil)); err != nil {
		t.Errorf("Void(twoReturns(nil)) = %v; want <nil>", err)
	}
	if err := Void(twoReturns(testErr)); !errors.Is(err, testErr) {
		t.Errorf("Void(twoReturns(testErr)) = %v; want %v", err, testErr)
	}

	if err := Void(threeReturns(nil)); err != nil {
		t.Errorf("Void(threeReturns(nil)) = %v; want <nil>", err)
	}
	if err := Void(threeReturns(testErr)); !errors.Is(err, testErr) {
		t.Errorf("Void(threeReturns(testErr)) = %v; want %v", err, testErr)
	}
}
