package main

import "testing"

func TestNormalizeNameIgnoresCaseAndSpaces(t *testing.T) {
	cases := map[string]string{
		"  Alice ": "alice",
		"BOB":      "bob",
		"carol":    "carol",
	}
	for in, want := range cases {
		if got := NormalizeName(in); got != want {
			t.Errorf("NormalizeName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestAddUserStoresTheNormalizedName(t *testing.T) {
	store, err := OpenStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	if _, err := store.AddUser("Alice", "alice@example.test"); err != nil {
		t.Fatal(err)
	}
	if _, err := store.AddUser("ALICE", "other@example.test"); err == nil {
		t.Fatal("a second Alice in other case was stored; want a unique-name error")
	}
}
