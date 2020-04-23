package main

import "fmt"
	//what kind of facinating creatures they would be and what qualities they posess
type animal struct {
	nickname string
	animalType string
	weight float32
	behavior bool
	note string
	likes []string
}
	//Mission:PetStore
func main(){
	//Making animal example
	var petOne animal
	petOne = animal{
		nickname   : "Shares",
		animalType : "snake",
		weight     : 1.3,
		behavior   : true,
		note       : "Cuddly snake, likes apples",
		likes      : []string{"apples","mice","eggs"}}

	//And the rest of company
	petTwo := animal{"Spike", "horse", 679.99, true, "Trustworthy stallion, likes apples", []string{"grass","apples","vodka"}}
	petThree:= animal{"Ogrim", "red-ass monkey", 36.6, false, "Like to throw #$*# at humans", []string{"bananas","bananas","bananas"}}
	petFour:= animal{"Swinky", "parrot", 0.23, false, "Iam not a pet, live me alone", []string{"human meat","seeds","bugs"}}
	petFive:= animal{"Formikloid", "bug", 0.01, true, "This strange species like to eat #$*#", []string{"#$*#","grass","leftovers"}}
	petSix := animal{"Lasy", "leopard", 112.5, false, "Sometimes frindly, sometimes not", []string{"meat","itallian food","sunshine"}}
	//Lets put them all on a list
	animalList := []animal{petOne, petTwo, petThree, petFour, petFive, petSix}
	//Lets put whole list on a map
	animalMap  := make(map[string]animal)									//declaring empty map
	index:=0																//index for counting elements
	for _, pickedAnimal := range animalList {									//for each picked animal in range of slice of animals
		animalMap[animalList[index].animalType] = animalList[index] 		//note that it auto sorted by name
		fmt.Println(pickedAnimal)
		index++
	}
	fmt.Println(animalMap["horse"])
	//lets make new map with only good animals
	index2:=0
	goodAnimalMap := make(map[string]animal)
	for _, i:= range animalMap {
		index2++
		switch i.behavior{
	case false: goodAnimalMap[animalList[index2].animalType] = animalList[index2]}
	
	}
	index = 0
	for _, i:= range goodAnimalMap {
		fmt.Println(i.animalType, "was very good")
		index++
	}
	

}