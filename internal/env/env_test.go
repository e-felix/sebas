package env

import (
	"github.com/e-felix/sebas/internal/util"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)
	envFilePath = basepath + "/../../testdata/assets/env_file"
)

func TestReadFile(t *testing.T) {
	expected := make([]Env, 0)
	expected = append(expected, Env{Key: "FOO", Value: "BAR"})
	expectedLength := len(expected)

	envs, _ := ReadFile(envFilePath)
	length := len(envs)

	if length != expectedLength {
		t.Fatalf("envs length = %d, expected %d", length, expectedLength)
	}

	if !reflect.DeepEqual(envs, expected) {
		t.Fatalf("Found %v, expected %v", envs, expected)
	}
}

func TestGetTokens(t *testing.T) {
	expected := make(map[string]string)
	expected["FOO"] = "BAR"

	content, _ := util.GetFileContent(envFilePath)

	if len(content) == 0 {
		t.Fatalf("content is empty")
	}

	tokens, _ := GetTokens(content)

	if !reflect.DeepEqual(tokens, expected) {
		t.Fatalf("Found %v, expected %v", tokens, expected)
	}
}

func TestConvertToEnv(t *testing.T) {
	key := "FOO"
	value := "BAR"
	expected := &Env{Key: key, Value: value}

	env := ConvertToEnv(key, value)

	if !reflect.DeepEqual(env, expected) {
		t.Fatalf("Found %v, expected %v", env, expected)
	}
}
