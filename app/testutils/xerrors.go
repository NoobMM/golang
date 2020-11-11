package testutils

import (
	"strings"
	"testing"
)

// AssertError is a function for asserting if actual error and want error are equal
func AssertError(t *testing.T, prefix string, actualErr error, wantErr error) bool {
	// true = test ok
	// false = unexpected assertion
	if actualErr == nil {
		if wantErr == nil {
			return true
		}

		t.Errorf("%s error = %v, wantErr %v", prefix, actualErr, wantErr)
		return false
	}
	if !strings.Contains(actualErr.Error(), wantErr.Error()) {
		t.Errorf("%s error = %v, wantErr %v", prefix, actualErr, wantErr)
		return false
	}
	return true
}
