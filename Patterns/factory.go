package main

import "fmt"



	type animalStruct struct {
		species string
		age int
	}
	type animalSectionStruct struct {
		name string
		part int
	}
	type animalFactoryStruct struct {
		species string
		section string
	}

	func (af animalFactoryStruct) newAnimal(age int) animalStruct {
		return animalStruct {
			species: af.species,
			age: age,
		}
	}

	func(af animalFactoryStruct) newSection(part int) animalSectionStruct {
		return animalSectionStruct {
			name: af.section,
			part: part,
		}
	}

	func main(){
	frogFactory := animalFactoryStruct{
		species: "frog",
		section: "amphibian",
	}
	frog := frogFactory.newAnimal(2)
	amphibian := frogFactory.newSection(1)

	fmt.Printf("species: %v %v years old\n", frog.species, frog.age)
	fmt.Printf("section: %v, part %v\n", amphibian.name, amphibian.part)

	dogFactory:= animalFactoryStruct{
		species: "dog",
		section: "mammal",
	}
	dog := dogFactory.newAnimal(5)
	mammal := dogFactory.newSection(2)

	fmt.Println(dog)
	fmt.Println(mammal)

	catFactory:= animalFactoryStruct{
		species: "cat",
		section: "mammal",
	}
	cat:= catFactory.newAnimal(3)
	mammal= catFactory.newSection(1)

	fmt.Println(cat)
	fmt.Println(catFactory.section, mammal.part)
	

}