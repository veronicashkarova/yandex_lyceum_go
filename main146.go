package main

import "testing"


func TestGetUTFLength (t *testing.T) {
	got := GetUTFLength ('h','e','l','l','o')
	expected := 5 
	if got != expected {
		t.Fatalf(`GetUTFLength ('h','e','l','l','o') = %q, want %q`, got, expected)
	}
}

