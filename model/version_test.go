package model_test

import (
	"testing"

	"github.com/ubogdan/gopanel/model"
)

func TestVersion(t *testing.T) {
	if got, want := model.Version().String(), "0.1.0"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}
}
