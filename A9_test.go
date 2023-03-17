package main

import (
	"testing"
)

/*
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected Add(2, 3) to be 5, but got %d", result)
	}
}
*/

func TestInterpNum(t *testing.T) {
	
	result := interp(NumC{N : 2}, []Binding{})
	if (result != NumV{N : 2}){
		t.Errorf("Expected NumV(2)")
	}

}

func TestInterpId(t *testing.T) {
	
	result := interp(IdC{S : "x"}, []Binding{Binding{Name: "y", Val: NumV{N : 4}},
											 Binding{Name: "x", Val: NumV{N : 2}}})
	if (result != NumV{N : 2}){
		t.Errorf("Expected NumV(2)")
	}

}

func TestInterpString(t *testing.T) {
	
	result := interp(StrC{S : "hello"}, []Binding{})
	if (result != StrV{S: "hello"}){
		t.Errorf("Expected NumV(2)")
	}

}


