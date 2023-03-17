package main

import (
	"testing"
	"reflect"
)


func TestInterpNum(t *testing.T) {
	
	result := interp(NumC{N : 2}, []Binding{})
	if (result != NumV{N : 2}){
		t.Errorf("Expected NumV(2)")
	}

}

func TestInterpId(t *testing.T) {
	
	result := interp(IdC{S : "x"}, []Binding{{Name: "y", Val: NumV{N : 4}},
											 {Name: "x", Val: NumV{N : 2}}})
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

func TestLamC(t *testing.T){
	result := interp(LamC{Args : []string{"x"}, Body : AppC{Fun: IdC{S: "+"}, Args: []ExprC{IdC{S: "x"}, NumC{N: 4}}}}, []Binding{})
	if !reflect.DeepEqual(result, ClosV{Args : []string{"x"}, Body : AppC{Fun: IdC{S: "+"}, Args: []ExprC{IdC{S: "x"}, NumC{N: 4}}}, Env : []Binding{}}){
	
		t.Error("LamC error")
	}
}


func TestLookUpBool(t *testing.T){
	result := interp(IdC{S : "true"}, []Binding{{Name : "true", Val : BoolV{Bool : true}}})
	if (result!= BoolV{Bool : true}){
		t.Error("BoolV error")
	}

}