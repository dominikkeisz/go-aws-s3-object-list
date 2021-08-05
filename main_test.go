package main

import (
	"testing"
)

func TestGetBucketNameFromArg(t *testing.T) {
	var want = "superbucketname"
	got := GetBucketNameFromArg([]string{"filename", "superbucketname"})
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
