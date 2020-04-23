package main

import ("fmt"
	"sync")

	//what kind of facinating creatures they would be and what qualities they posess
type animal struct {
	nickname string
	animalType string
	weight float32
	behavior bool
	note string
	likes []string
}
var wg sync.WaitGroup
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
		likes      : []string{"mice","apples","eggs"}}

	//And the rest of our company
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
	for _, pickedAnimal := range animalList {								//for each picked animal in range of slice of animals
		animalMap[animalList[index].animalType] = animalList[index] 		//with that, we create keys for our map = type of animal
		fmt.Println(pickedAnimal)
		index++
	}
	fmt.Println(animalMap["horse"])											//finding all "horse" type on a map

	//lets make new map with only good animals
	index = 0																//so we can reuse it again
	goodAnimalMap := make(map[string]animal)
	for _, i:= range animalList {

		switch i.behavior{
	case true: goodAnimalMap[animalList[index].animalType] = animalList[index] }
		index++
		}
	//Lets write what anymal types was good
	
	index = 0
	for _, i:= range goodAnimalMap {
		fmt.Println(i.animalType, "was very good")
		index++
	}

	//lets say we got new animal arriving
	petSeven := animal {"Fransise","polar bear" , 554.45, true, "He was on commercials many time", [] string{"ice","meat","sunglasses"}}

	//now we need add him on a map, but witch one? both?
	switch petSeven.behavior{
	case true: 
		goodAnimalMap[petSeven.animalType] = petSeven
		animalMap[petSeven.animalType] = petSeven
	default: animalMap[petSeven.animalType] = petSeven
	}
	fmt.Println("This animals a good: ")
	fmt.Println(goodAnimalMap)

	//Creating new map containing only apple lovers
	appleLoversMap := make(map[string]animal)
	numberOfLikedThings:= 0
	
	for _, pickedAnimal:= range animalMap {

		index = 0
		numberOfLikedThings = len(pickedAnimal.likes)

		for index < numberOfLikedThings {
		
			if pickedAnimal.likes[index] == "apples" {
				appleLoversMap[pickedAnimal.animalType] = pickedAnimal
			}
			
		index++	
		}
		
	}
	fmt.Println("This animals like apples: ")
	fmt.Println(appleLoversMap)
}

