package main

import "testing"

func TestAdd(t *testing.T) {
	t.Log("hello testing")
	out := Add(10, 20)
	expt := 30

	if out != expt {
		t.Error("error in add function")
	}
}

func TestSub(t *testing.T) {
	t.Log("hello testing sub")
	out := Sub(30, 20)
	expt := 10

	if out != expt {
		t.Error("error in add function")
	}
}
