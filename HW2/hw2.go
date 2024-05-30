package main

import (
	"fmt"
	"math/rand"
)

type Animal struct {
	Name   string
	Caught bool
}

func (a *Animal) Catch() {
    for !a.Caught {
        if rand.Intn(2) == 0 { 
            a.Caught = true
            fmt.Printf("Caught %s\n", a.Name)
        } else {
            fmt.Printf("Failed to catch %s\n", a.Name)
        }
    }
}

type Cage struct {
	Number int
	Animal *Animal
}

func (c *Cage) PlaceAnimal(animal *Animal) {
	c.Animal = animal
	fmt.Printf("%s is placed in cage %d\n", animal.Name, c.Number)
}

type Zookeeper struct {
	Name    string
	Animals []*Animal
	Cages   []*Cage
}

func (z *Zookeeper) GatherAnimals() {
    fmt.Printf("Zookeeper %s starts gathering animals\n", z.Name)
    for i := 0; i < len(z.Animals); i++ {
        animal := z.Animals[i]
        animal.Catch()
        for j := 0; j < len(z.Cages); j++ {
            cage := z.Cages[j]
            if cage.Animal == nil {
                cage.PlaceAnimal(animal)
                break
            }
        }
    }
}

func main() {
	animals := []*Animal{
		{Name: "Lion"},
		{Name: "Tiger"},
		{Name: "Zebra"},
		{Name: "Giraffe"},
		{Name: "Elephant"},
	}

	cages := []*Cage{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
	}

	zookeeper := &Zookeeper{
		Name:    "John",
		Animals: animals,
		Cages:   cages,
	}

	zookeeper.GatherAnimals()
}