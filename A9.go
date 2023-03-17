package main

import (
	"errors"
	"fmt"
)

type ExprC interface{}

type NumC struct {
	N int
}

type IdC struct {
	S string
}

type StrC struct {
	S string
}

type AppC struct {
	Fun  ExprC
	Args []ExprC
}

type LamC struct {
	Args []string
	Body ExprC
}

type IfC struct {
	If   ExprC
	Then ExprC
	Else ExprC
}

type Value interface{}

type NumV struct{
    N int
}

type BoolV struct{
    Bool bool
}

type StrV struct{
    S string
}

type Closer interface{}

type ClosV struct {
	Args []string
	Body ExprC
	Env  []Binding
}

type Binding struct {
	Name string
	Val  Value
}

func getBinop(op string, l int, r int) (interface{}, error) {
	switch op {
	case "equal?":
		return l == r, nil
	case "+":
		return l + r, nil
	case "-":
		return l - r, nil
	case "*":
		return l * r, nil
	case "/":
		if r == 0 {
			return nil, errors.New("division by zero")
		}
		return l / r, nil
	case "<=":
		return l <= r, nil
	default:
		return nil, errors.New("invalid binop syntax")
	}
}

func getEnv(s []string, l []Value, env []Binding) []Binding {
	if len(s) == 0 {
		return env
	}
	return append([]Binding{Binding{s[0], l[0]}}, getEnv(s[1:], l[1:], env)...)
}

func lookup(forVal string, env []Binding) Value {
	if len(env) == 0 {
		panic("user-error No value match!")
	} else {
		if forVal == env[0].Name {
			return env[0].Val
		} else {
			return lookup(forVal, env[1:])
		}
	}
}

func interp(e ExprC, env []Binding) Value {
	switch e := e.(type) {
	case NumC:
		return NumV{N : e.N}
	case StrC:
		return StrV{S : e.S}
	case IdC:
		return lookup(e.S, env)
	case LamC:
		return ClosV{e.Args, e.Body, env}
	case IfC:
		ifVal := interp(e.If, env)
		if (ifVal == BoolV{Bool : true}) {
			return interp(e.Then, env)
		} else if (ifVal == BoolV{Bool : false}) {
			return interp(e.Else, env)
		} else {
			// Would be helpful to print out invalid value
			panic("if condition should be boolean value")
		}
	case AppC:
		fd := interp(e.Fun, env)
		fArgs := make([]Value, len(e.Args))
		for i, arg := range e.Args {
			fArgs[i] = interp(arg, env)
		}
		switch fd := fd.(type) {
		case func(Value, Value) Value:
			if len(fArgs) != 2 {
				panic("invalid number of arguments")
			}
			return fd(fArgs[0], fArgs[1])
		case ClosV:
			fEnv := getEnv(fd.Args, fArgs, fd.Env)
			return interp(fd.Body, fEnv)
		default:
			panic("invalid application")
		}
	default:
		panic("invalid expression")
	}
}



func main() {
	fmt.Println("hello world")
    //testExprC := AppC{Fun: IdC{S: "+"}, Args: []ExprC{NumC{N: 3}, NumC{N: 4}}}
    /*
    topEnv := []Binding{
       
        Binding{Name: "+", Val: createPrimFunc("+")},
        Binding{Name: "/", Val: createPrimFunc("/")},
        Binding{Name: "*", Val: createPrimFunc("*")},
        Binding{Name: "-", Val: createPrimFunc("-")},
  
    }
    */
    fmt.Println(interp(NumC{N : 2}, []Binding{}))
    fmt.Println(NumV{N : 2})
}
