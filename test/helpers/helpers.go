package helpers

import (
	"reflect"
	"testing"
)

func FailAssertion(t *testing.T, name string, got, want interface{}) {
	t.Helper()
	t.Errorf("assertion failed: %s\n\n\t\tgot:\n\t\t\t%v\n\n\t\twant:\n\t\t\t%v", name, got, want)
}

func AssertErrorEquals(t *testing.T, got, want error) {
	t.Helper()
	if got == nil || got.Error() != want.Error() {
		FailAssertion(t, "assertErrorEquals", got, want)
	}
}

func AssertNil(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		FailAssertion(t, "assertNil", got, nil)
	}
}

func AssertDeepEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		FailAssertion(t, "deepEqual", got, want)
	}
}
