package main

import "fmt"

type everHydra struct {
	[]hydraHead								//whole hydra have set oh heads
	body string								//whole body status in text
	bodyHp float64							//how much hp
}

type hydraHead struct{
	name string								//each hydra head have own name
	headHp float64							//separate hp
	atackStruckt							//
	isCutOf bool							//status of head. true = is cut of, false=exist
}

type atackStruct struct {
	abilityName string						//each head using own abilities
	abilityDmg float64						//so have different damage
}

func main() {
	//creating abilities for each head
	abilityFire := atackStruct{"fireblast", 50}				
	abilityCold := atackStruct{"frozen lance", 40}
	abilityAcid := atackStruct("toxic spit", 30)


	thisIsMap := make map
}