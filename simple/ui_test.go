package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestGreeting(t *testing.T) {
	output, entry, btn := myApp.makeUI()

	if output.Text != "Hello World" {
		t.Error("Incorrect output text")
	}

	if btn.Text != "Enter" {
		t.Error("Incorrect button text")
	}

	test.Type(entry, "Foo")
	if output.Text != "Foo" {
		t.Error("Incorrect output entry text")
	}
}
