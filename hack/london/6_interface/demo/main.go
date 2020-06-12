package main

import "fmt"

type Ducker interface {
	DuckWalk()
	DuckBark()
}

type ComicActor struct{}

func (c *ComicActor) DuckWalk() {
	fmt.Println("ComicActor_Walk()")
}

func (c *ComicActor) DuckBark() {
	fmt.Println("ComicActor_Bark()")
}

type UnknownAnimal struct{}

func (u *UnknownAnimal) DuckWalk() {
	fmt.Println("UnknownAnimal_Walk()")
}

func (u *UnknownAnimal) DuckBark() {
	fmt.Println("UnknownAnimal_Walk()")
}

func main() {
	var duckInterface Ducker
	duckInterface = &ComicActor{}
	duckInterface.DuckWalk()
	duckInterface.DuckBark()

	duckInterface = &UnknownAnimal{}
	duckInterface.DuckWalk()
	duckInterface.DuckBark()
}
