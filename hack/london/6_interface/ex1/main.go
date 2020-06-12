package main

import "fmt"

type IBase interface {
	FuncBase()
}

type IDerivative interface {
	IBase
	FuncDerivative()
}

type Subject struct{} // 用来限制

func (s Subject) FuncDerivative() {
	fmt.Println("FuncDerivative()")
}

func main() {
	//var refDerivative IDerivative = &Subject{} // 限定
}
