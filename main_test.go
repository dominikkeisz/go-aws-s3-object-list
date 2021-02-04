package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestGetBucketNameFromArg(t *testing.T) {
	var want = "superbucketname"
	got := GetBucketNameFromArg([]string{"filename", "superbucketname"})
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetBucketNameFromArgExitCode(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		GetBucketNameFromArg([]string{"filename", "firstbucketname", "secondbucketname"})
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestGetBucketNameFromArgFail")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("Process ran with err %v, wanted exit code 1", err)
}
