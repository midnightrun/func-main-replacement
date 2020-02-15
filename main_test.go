package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunNoNamesProvided(t *testing.T) {
	args := []string{"greeter"}

	var stdout bytes.Buffer

	err := run(args, &stdout)

	if err == nil {
		t.Fatal("expected to receive an error but got nil")
	}
}

func TestRunNamesProvided(t *testing.T) {
	args := []string{"greeter", "Tom", "Jerry"}
	var stdout bytes.Buffer

	err := run(args, &stdout)

	if err != nil {
		t.Fatalf("expected to receive no error but got %v", err)
	}

	out := stdout.String()

	if !strings.Contains(out, "Hi Tom") {
		t.Fatal("Expected to get a greet message for Tom but didn't")
	}

	if !strings.Contains(out, "Hi Jerry") {
		t.Fatal("Expected to get a greet message for Jerry but didn't")
	}
}

func TestRunFormatChangeProvided(t *testing.T) {
	args := []string{"greeter", "-f \"Hola %s\"", "Tom", "Jerry"}
	var stdout bytes.Buffer

	err := run(args, &stdout)

	if err != nil {
		t.Fatalf("expected to receive no error but got %v", err)
	}

	out := stdout.String()

	if !strings.Contains(out, "Hi Tom") {
		t.Fatal("Expected to get a greet message for Tom but didn't")
	}

	if !strings.Contains(out, "Hi Jerry") {
		t.Fatal("Expected to get a greet message for Jerry but didn't")
	}
}
