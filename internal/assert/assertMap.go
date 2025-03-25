package internal

import "testing"

// AssertMap compares two map[string]string instances for equality.
// It fails the test if the maps differ in keys or values.
func AssertMap(t *testing.T, got, want map[string]string, msgAndArgs ...interface{}) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("map length mismatch: got %d keys, want %d keys", len(got), len(want))
		return
	}

	for key, wantValue := range want {
		if gotValue, ok := got[key]; !ok {
			t.Errorf("missing key %q in got map", key)
		} else if gotValue != wantValue {
			t.Errorf("value mismatch for key %q: got %q, want %q", key, gotValue, wantValue)
		}
	}

	for key := range got {
		if _, ok := want[key]; !ok {
			t.Errorf("unexpected key %q in got map", key)
		}
	}
}
