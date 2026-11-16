package main

import (
	"errors"
	"testing"
)

func newStore(t *testing.T, names ...string) *Store {
	t.Helper()
	store, err := OpenStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { store.Close() })
	for _, n := range names {
		if _, err := store.AddUser(n, n+"@example.test"); err != nil {
			t.Fatal(err)
		}
	}
	return store
}

func TestFindUserByName(t *testing.T) {
	store := newStore(t, "Alice", "Bob")
	u, err := store.FindUserByName(" BOB ")
	if err != nil {
		t.Fatal(err)
	}
	if u.Name != "bob" {
		t.Fatalf("found %q, want bob", u.Name)
	}
	if _, err := store.FindUserByName("nobody"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestFindUserByNameTreatsAQuoteAsText(t *testing.T) {
	store := newStore(t, "Alice", "Bob")
	u, err := store.FindUserByName("nobody' OR '1'='1")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("a name holding a quote matched user %q (err %v); want ErrNotFound", u.Name, err)
	}
}

func TestFindUserByNameWithAnApostrophe(t *testing.T) {
	store := newStore(t, "O'Brien")
	u, err := store.FindUserByName("o'brien")
	if err != nil {
		t.Fatalf("err = %v, want the user o'brien", err)
	}
	if u.Name != "o'brien" {
		t.Fatalf("found %q, want o'brien", u.Name)
	}
}
