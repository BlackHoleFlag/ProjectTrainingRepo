package main

import(
	"fmt"
	"cryptorandomizer"
	"math"
)

 

type managerStruct struct {
	name string
	money float64
	strategyMeat	float64
	strategyGreens	float64
	strategyVeggies float64
	strategySauce float64
	specialModifier int
	rating float64

}

type ingredientStruct struct {
	name string
	basePrice float64
}

type customerStruct struct {
	name string
	likes []string
	dislike []string
	readyToPay float64
}

func market(managerShoping managerStruct) (map[int]map[string]int, float64) {
													//ALL PRICES based on how much needed for 1 order of salad
	//All possible meats on market
	meatListAll := [4] ingredientStruct {{"grilled chicken", 1.85} , {"chopped steak", 4.25} , {"crispy shrimps", 2.50} , {"bacon", 2.05}}
	//All possible dressings on market
	sauceListAll := [5] ingredientStruct {{"salt and pepper only",0.01} , {"mayo", 0.15} , {"ranch", 0.29} , {"olive oil", 0.33} , {"chipottle", 0.25}}
	//All possible veggies
	vegeterianListAll := [5]ingredientStruct {{"raddish", 0.15} , {"minced onions", 0.10} , {"cherry tomatoes", 0.35} , {"cucumbers", 0.20} , {"black olives", 0.18}}
	//All possible greens
	greensListAll := [3] ingredientStruct {{"spinach", 0.33} , {"lettuce", 0.15} , {"arugula", 0.25}}

	//Lets find out how much money manager willing to spend on types of produce
	managerShopingMeat := managerShoping.strategyMeat * managerShoping.money
	managerShopingGreens := managerShoping.strategyGreens * managerShoping.money
	managerShopingVeggies := managerShoping.strategyVeggies * managerShoping.money
	managerShopingSauce := managerShoping.strategySauce * managerShoping.money
	//Creating our containers with key=name of ingredient, value - quantity
	meatContainerMap := make(map[string]int)
	greensContainerMap := make(map[string]int)
	veggiesContainerMap := make(map[string]int)
	sauceContainerMap := make(map[string]int)

	//lets start shopping
	for managerShopingMeat >= 15 {
		picked := meatListAll[cryptorandomizer.Num(4)]
		managerShopingMeat = managerShopingMeat - picked.basePrice
		meatContainerMap[picked.name]++
	}

	for managerShopingGreens >= 15 {
		picked := greensListAll[cryptorandomizer.Num(3)]
		managerShopingGreens = managerShopingGreens - picked.basePrice
		greensContainerMap[picked.name]++
	}

	for managerShopingVeggies >= 15 {
		picked := vegeterianListAll[cryptorandomizer.Num(5)]
		managerShopingVeggies = managerShopingVeggies - picked.basePrice
		veggiesContainerMap[picked.name]++
	}

	for managerShopingSauce >= 15 {
		picked := sauceListAll[cryptorandomizer.Num(5)]
		managerShopingSauce = managerShopingSauce - picked.basePrice
		sauceContainerMap[picked.name]++
	}

	//load containers in to car and going to restaurant
	var carLoaded= map[int]map[string]int{}
	carLoaded[1] = meatContainerMap
	carLoaded[2] = greensContainerMap
	carLoaded[3] = veggiesContainerMap
	carLoaded[4] = sauceContainerMap
	moneyLeft:= roundUp(managerShopingMeat + managerShopingGreens + managerShopingVeggies + managerShopingSauce)
	fmt.Printf("Manager %v spend $%v and bought:\n", managerShoping.name, managerShoping.money - moneyLeft, )
	fmt.Println(carLoaded)
	return carLoaded, moneyLeft
}

func makeCustomerLine(rating float64) []customerStruct {

	//How many customers would come today, rating dependent
	numberOfCustomers :=  math.Round(10 * rating) + math.Round(float64(cryptorandomizer.Num(9)+1) * rating)
	customerLine := []customerStruct{}
	
	for numberOfCustomers >0 {
		customerLine = append(customerLine, makeCustomerType())
		numberOfCustomers--
	}
	return customerLine
}

func makeCustomerType() customerStruct{
	//our customers names
	potentialNames := [] string {"Billi Bob", "Jeffrey Lebowski", "Sandra Cassandra", "Alf", "Maki Makinson", "Abby Woodsman", "Roy Foreman", "Liz Benedict"}
	//our customers likes/dislikes
	potentialLikeDis := [] string{"bacon", "chopped steak", "crispy shrimps", "grilled chicken", "arugula", "lettuce", "spinach", "black olives", "cherry tomatoes", "cucumbers",
	 "minced onions", "raddish", "chipottle", "mayo", "olive oil", "ranch","salt and pepper only"}
	

	customer := customerStruct{}
	//lets find out how many likes/dislikes our potential customer would have
	howManyLikes := 1 + cryptorandomizer.Num(2)
	howManyDis := 1 + cryptorandomizer.Num(2)

	//generating likes
		for howManyLikes > 0 {
			genLike := potentialLikeDis[cryptorandomizer.Num(17)]
	
			if len(customer.likes) >  0 {
				for _, r := range customer.likes {
				
				if genLike != r {
					customer.likes = append(customer.likes, genLike)	
				}
				}
			} else {
				customer.likes = append(customer.likes, genLike)
			}
			howManyLikes--
		}
	//generating dislikes
	for howManyDis > 0 {
		genDis := potentialLikeDis[cryptorandomizer.Num(17)]

		if len(customer.dislike) >  0 {
			for _, r := range customer.dislike {
			
			if genDis != r {
				customer.dislike = append(customer.dislike, genDis)	
			}
			}
		} else {
			customer.dislike = append(customer.dislike, genDis)
		}
		howManyDis--
	}
	customer.name = potentialNames[cryptorandomizer.Num(8)]
	customer.readyToPay = 5.0 + float64(cryptorandomizer.Num(10))
	

	return customer
}

func main() {
	//for now lets have one playable manager, if you want more, buy DLC
	managerMeatLover := managerStruct {
		           name : "Francissko",
		          money : 200,
		   strategyMeat : 0.60,
		 strategyGreens : 0.15,
		strategyVeggies : 0.15,
		  strategySauce : 0.10,
		specialModifier	: 1,
		         rating : 1.0,
	}

	//Lets make potential customers and return slice of them
	customerLine := makeCustomerLine(managerMeatLover.rating)
	fmt.Println(customerLine)
	//sending manager to the market, returning produce and money left
	produceAll, moneyLeft := market(managerMeatLover)
	managerMeatLover.money = moneyLeft
	//Now lets make customer line for today
	
	




	fmt.Println(produceAll)
	fmt.Println(managerMeatLover.money)
}



func roundUp(input float64)float64{
	var round float64
	pow:=math.Pow(10, float64(2))
	digit:= pow * input
	round = math.Ceil(digit)
	newVal:= round/pow
	return newVal
}