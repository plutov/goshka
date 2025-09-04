package is

import (
	"errors"
	"reflect"
	"testing"
)

func Equal[T comparable](t testing.TB, want, got T) {
	t.Helper()
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func DeepEqual[T any](t testing.TB, want, got T) {
	t.Helper()
	if !reflect.DeepEqual(&want, &got) {
		t.Fatalf("reflect.DeepEqual(%#v, %#v) == false", want, got)
	}
}

func NilErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got: %v", err)
	}
}

func ErrorIs(t testing.TB, want, got error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Fatalf("got errors.Is(%v, %v) == false", got, want)
	}
}

func ErrorAs[T error](t testing.TB, want *T, got error) {
	t.Helper()
	if !errors.As(got, want) {
		t.Fatalf("got errors.As(%v, %T) == false", got, want)
	}
}
