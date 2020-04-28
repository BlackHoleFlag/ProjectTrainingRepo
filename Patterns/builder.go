package main 

import "fmt"

type finishedProductStruct struct {
	productName string
	wood int
	metal int
	price int
}

type buildProcessInterface interface {
	setProductName() buildProcessInterface
	setWood() buildProcessInterface
	setMetal() buildProcessInterface
	setPrice() buildProcessInterface
	getItem() finishedProductStruct
}

type manufacturingDirStruct struct {
	builder buildProcessInterface
}

func (m *manufacturingDirStruct) setBuilder(b buildProcessInterface) {
	m.builder = b
}

func (m *manufacturingDirStruct) construct() finishedProductStruct {
	m.builder.setProductName().setWood().setMetal().setPrice()
	return m.builder.getItem()
}

func (m *manufacturingDirStruct) printProduct() {
	item := m.builder.getItem()
	fmt.Printf("Item: %s \n", item.productName)
	fmt.Printf("Used wood: %d \n", item.wood)
	fmt.Printf("Used metals: %d \n", item.metal)
	fmt.Printf("Price is: %d\n", item.price)
	fmt.Printf("===============\n")

}

type pickAxeStruct struct {
	finishedProduct finishedProductStruct
}

func (p *pickAxeStruct) setProductName() buildProcessInterface {
	p.finishedProduct.productName = "pickaxe"
	return p
}
func (p *pickAxeStruct) setWood() buildProcessInterface {
	p.finishedProduct.wood = 1
	return p
}
func (p *pickAxeStruct) setMetal() buildProcessInterface {
	p.finishedProduct.metal = 2
	return p
}
func (p *pickAxeStruct) setPrice() buildProcessInterface{
	p.finishedProduct.price = p.finishedProduct.wood * 2 + p.finishedProduct.metal * 3
	return p
}
func (p *pickAxeStruct) getItem() finishedProductStruct {
	return p.finishedProduct
}

type spearStruct struct {
	finishedProduct finishedProductStruct
}

func (p *spearStruct) setProductName() buildProcessInterface {
	p.finishedProduct.productName = "simple spear"
	return p
}
func (p *spearStruct) setWood() buildProcessInterface {
	p.finishedProduct.wood = 2
	return p
}
func (p *spearStruct) setMetal() buildProcessInterface {
	p.finishedProduct.metal = 1
	return p
}
func (p *spearStruct) setPrice() buildProcessInterface{
	p.finishedProduct.price = p.finishedProduct.metal * 3 + p.finishedProduct.wood * 2
	return p
}
func (p *spearStruct) getItem() finishedProductStruct {
	return p.finishedProduct
}

func main() {
	manuf := manufacturingDirStruct{}
	pickAxe := &pickAxeStruct{}
	manuf.setBuilder(pickAxe)
	manuf.construct()
	manuf.printProduct()
	spear := &spearStruct{}
	manuf.setBuilder(spear)
	manuf.construct()
	manuf.printProduct()
}