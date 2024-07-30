package util

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func TestGetContent(t *testing.T) {
	expected := "FOO=BAR\n"
	expectedLength := len(expected)

	content, _ := GetFileContent(basepath + "/../../testdata/assets/env_file")
	length := len(content)

	if length != expectedLength {
		t.Fatalf("envs length = %d, expected %d", length, expectedLength)
	}

	if content != expected {
		t.Fatalf("Found %v, expected %v", content, expected)
	}
}
