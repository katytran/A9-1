package main


import "fmt"


type ExprC interface{}

type NumC struct {
    Num int
}

type IfC struct {
    Test ExprC
    Then ExprC
    Else ExprC
}

type IdC struct {
    S string
}

type AppC struct {
    Fun ExprC
    Arg []ExprC
}

type LamC struct {
    Args []string
    Body ExprC
}

type StrC struct {
    S string
}





func main() {
    fmt.Println("Hello, World!")
    str := StrC{S: "hello world"}
    testExprC := AppC{Fun: IdC{S: "+"}, Arg: []ExprC{NumC{Num: 3}, NumC{Num: 4}}}
    fmt.Println(str)
    fmt.Println(testExprC)
}

func Add(a, b int) int {
    return a + b
}