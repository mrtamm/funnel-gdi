package tes

import (
	"reflect"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	task := &Task{
		Id: "task1",
		Executors: []*Executor{
			{
				Image:   "alpine",
				Command: []string{"echo", "hello world"},
			},
		},
	}

	// Protojson may (and may not) add extra whitespace between labels and values.
	// Therefore there are two possible scenarios when the test checks against specific BASE64 value(s).
	// More info: https://github.com/golang/protobuf/issues/1082
	expected1 := "ewogICJpZCI6ICJ0YXNrMSIsCiAgImV4ZWN1dG9ycyI6IFsKICAgIHsKICAgICAgImltYWdlIjogImFscGluZSIsCiAgICAgICJjb21tYW5kIjogWwogICAgICAgICJlY2hvIiwKICAgICAgICAiaGVsbG8gd29ybGQiCiAgICAgIF0KICAgIH0KICBdCn0="
	expected2 := "ewogICJpZCI6ICAidGFzazEiLAogICJleGVjdXRvcnMiOiAgWwogICAgewogICAgICAiaW1hZ2UiOiAgImFscGluZSIsCiAgICAgICJjb21tYW5kIjogIFsKICAgICAgICAiZWNobyIsCiAgICAgICAgImhlbGxvIHdvcmxkIgogICAgICBdCiAgICB9CiAgXQp9"

	encoded, err := Base64Encode(task)
	if err != nil {
		t.Fatal(err)
	}

	if encoded != expected1 && encoded != expected2 {
		str, err := MarshalToString(task)
		t.Logf("Source JSON-string: %+v (error: %+v)", str, err)
		t.Logf("actual:    %+v", encoded)
		t.Logf("expected1: %+v", expected1)
		t.Logf("expected2: %+v", expected2)
		t.Fatal("unexpected value returned from Base64Encode")
	}

	decoded, err := Base64Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(decoded, task) {
		t.Logf("expected: %+v", task)
		t.Logf("actual: %+v", decoded)
		t.Fatal("incorrect decoded task from Base64Decode")
	}
}
