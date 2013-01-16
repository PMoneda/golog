package term

import . "fmt"
import . "regexp"

var anonCounter <-chan int64
func init() {
    // goroutine providing a counter for anonymous variables
    c := make(chan int64)
    var i int64 = 1000
    go func() {
        for {
            c <- i
            i++
        }
    }()
    anonCounter = c
}

// See §6.1.2(a)
type Variable struct {
    Name    string
}

func NewVar(name string) Term {
    // sanity check the variable name's syntax
    isCapitalized, err := MatchString(`^[A-Z_]`, name)
    maybePanic(err)
    if !isCapitalized {
        panic("Variable names must start with a capital letter or underscore")
    }

    // make sure anonymous variables are unique
    if name == "_" {
        i := <-anonCounter
        name = Sprintf("_A%d", i)
    }
    return &Variable{
        Name:   name,
    }
}

func (self *Variable) Functor() string {
    panic("Variables have no Functor()")
}

func (self *Variable) Arity() int {
    panic("Variables have no Arity()")
}

func (self *Variable) Arguments() []Term {
    panic("Variables have no Arguments()")
}

func (self *Variable) Body() Term {
    panic("Variables have no Body()")
}

func (self *Variable) Head() Term {
    panic("Variables have no Head()")
}

func (self *Variable) IsClause() bool {
    return false
}

func (self *Variable) String() string {
    return self.Name
}

func (self *Variable) Indicator() string {
    return Sprintf("%s", self.Name)
}

func (self *Variable) Error() error {
    panic("Can't call Error() on a Variable")
}