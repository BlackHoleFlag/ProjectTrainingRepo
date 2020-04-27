package main 

import "fmt"

type bentleyCarStr struct {
	car
}
type porsheCarStr struct {
	car
}
type bentleyBusStr struct {
	bus
}
type porsheBusStr struct{
	bus
}
type iAutoFactory interface {
	makeCar() iCar
	makeBus() iBus
}
type car struct {
	logo string
	year int
}
type bus struct {
	logo string
	year int
}
type bentley struct {}
type porshe struct {}



func (a *bentley) makeCar() iCar {
	return &bentleyCarStr{
		car: car{
			logo: "bentley",
			year: 2003,
		},
	}
}
func(a *bentley) makeBus() iBus {
	return &bentleyBusStr{
		bus: bus {
			logo: "bentley",
			year: 2002,
		},
	}
}
func (a *porshe) makeCar() iCar {
	return &porsheCarStr{
		car: car{
			logo: "porshe",
			year: 2013,
		},
	}
}
func(a *porshe) makeBus() iBus {
	return &porsheBusStr{
		bus: bus {
			logo: "porshe",
			year: 2012,
		},
	}
}

type iCar interface {
	setLogo (logo string)
	setYear (year int)
	getLogo() string
	getYear() int
}
type iBus interface {
	setLogo (logo string)
	setYear (year int)
	getLogo() string
	getYear() int
}
//cars
func (a *car) setLogo(logo string) {
	a.logo = logo
}
func(a *car) getLogo() string {
	return a.logo
}
func (a *car) setYear(year int) {
	a.year = year
}
func (a *car) getYear() int {
	return a.year
}
//busses
func (a *bus) setLogo(logo string) {
	a.logo = logo
}
func(a *bus) getLogo() string {
	return a.logo
}
func (a *bus) setYear(year int) {
	a.year = year
}
func (a *bus) getYear() int{
	return a.year
}

func getAutoFactory (brand string) (iAutoFactory, error) {
	if brand == "bentley" {
		return &bentley{}, nil
	}
	if brand == "porshe" {
		return &porshe{}, nil
	}
	return nil, fmt.Errorf("brand dont have licence")
}

func main() {
	bentleyFactory, _ := getAutoFactory("bentley")
	porsheFactory, _ := getAutoFactory("porshe")

	bentleyCar := bentleyFactory.makeCar()
	porsheCar := porsheFactory.makeCar()

	bentleyBus := bentleyFactory.makeBus()
	porsheBus := porsheFactory.makeBus()

	printCarDetails(bentleyCar)
	printBusDetails(bentleyBus)
	printCarDetails(porsheCar)
	printBusDetails(porsheBus)
}

func printCarDetails(c iCar) {
	fmt.Printf("Car logo: %s", c.getLogo())
	fmt.Println()
	fmt.Printf("Year: %d", c.getYear())
	fmt.Println()
}
func printBusDetails(b iBus) {
	fmt.Printf("Bus logo %s\n", b.getLogo())
	fmt.Printf("Year: %d\n", b.getYear())
}