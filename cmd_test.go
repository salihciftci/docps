package main

import (
	"testing"
)

func TestDockerCmd(t *testing.T) {
	cmd := []string{
		"run",
		"--name",
		"test",
		"--rm",
		"hello-world",
	}

	_, err := dockerCmd(cmd)
	if err != nil {
		t.Fatalf("Docker not running")
	}

}
