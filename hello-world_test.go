package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if "Hello World" != hello() {
		t.Fatal("failed test")
	}
}
